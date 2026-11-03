package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Result struct {
	Stdout string  `json:"stdout"`
	Stderr string  `json:"stderr"`
	Events []Event `json:"events,omitempty"`
	RunMS  int64   `json:"run_ms"`
	Error  string  `json:"error,omitempty"`
	Exit   int     `json:"exit_code"`
}

type Event struct {
	Message string `json:"message"`
	Kind    string `json:"kind"`
	Delay   int64  `json:"delay"`
}

type Runner interface {
	Run(ctx context.Context, code, stdin string) (Result, error)
	SupportsStdin() bool
	Name() string
}

type Playground struct {
	endpoint string
	client   *http.Client
	timeout  time.Duration
	maxCode  int
	maxOut   int
}

func NewPlayground(endpoint string, timeout time.Duration, maxCode, maxOut int) *Playground {
	return &Playground{
		endpoint: endpoint,
		client:   &http.Client{Timeout: timeout + 2*time.Second},
		timeout:  timeout,
		maxCode:  maxCode,
		maxOut:   maxOut,
	}
}

func (p *Playground) Name() string        { return "playground" }
func (p *Playground) SupportsStdin() bool { return false }

func (p *Playground) Run(ctx context.Context, code, stdin string) (Result, error) {
	if len(code) > p.maxCode {
		return Result{}, errors.New("kode terlalu panjang")
	}
	if stdin != "" {
		return Result{}, errors.New("runner playground tidak mendukung input; aktifkan RUNNER_MODE=docker untuk latihan dengan stdin")
	}
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	form := url.Values{"body": {code}, "withVet": {"true"}}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return Result{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	start := time.Now()
	resp, err := p.client.Do(req)
	if err != nil {
		return Result{}, fmt.Errorf("playground request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Result{}, fmt.Errorf("playground status %d", resp.StatusCode)
	}
	var payload struct {
		Errors string  `json:"Errors"`
		Events []Event `json:"Events"`
	}
	if err = json.NewDecoder(io.LimitReader(resp.Body, int64(p.maxOut)*8)).Decode(&payload); err != nil {
		return Result{}, fmt.Errorf("playground decode: %w", err)
	}

	var stdout, stderr bytes.Buffer
	for _, ev := range payload.Events {
		switch ev.Kind {
		case "stdout":
			stdout.WriteString(ev.Message)
		case "stderr":
			stderr.WriteString(ev.Message)
		}
	}
	exit := 0
	if strings.TrimSpace(payload.Errors) != "" {
		exit = 1
		if stderr.Len() == 0 {
			stderr.WriteString(payload.Errors)
		}
	}
	return Result{
		Stdout: truncate(stdout.String(), p.maxOut),
		Stderr: truncate(stderr.String(), p.maxOut),
		Events: payload.Events,
		RunMS:  time.Since(start).Milliseconds(),
		Exit:   exit,
	}, nil
}

type Docker struct {
	image   string
	timeout time.Duration
	maxCode int
	maxOut  int
}

func NewDocker(image string, timeout time.Duration, maxCode, maxOut int) *Docker {
	return &Docker{image: image, timeout: timeout, maxCode: maxCode, maxOut: maxOut}
}

func (d *Docker) Name() string        { return "docker" }
func (d *Docker) SupportsStdin() bool { return true }

func (d *Docker) Run(ctx context.Context, code, stdin string) (Result, error) {
	if len(code) > d.maxCode {
		return Result{}, errors.New("kode terlalu panjang")
	}
	dir, err := os.MkdirTemp("", "golearn-run-*")
	if err != nil {
		return Result{}, err
	}
	defer os.RemoveAll(dir)
	if err = os.WriteFile(filepath.Join(dir, "main.go"), []byte(code), 0o600); err != nil {
		return Result{}, err
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		return Result{}, err
	}
	ctx, cancel := context.WithTimeout(ctx, d.timeout)
	defer cancel()

	args := []string{
		"run", "--rm", "-i", "--network", "none",
		"--memory", "128m", "--cpus", "0.5", "--pids-limit", "64",
		"--read-only", "--tmpfs", "/tmp:rw,size=64m",
		"-e", "GOCACHE=/tmp/gocache", "-e", "HOME=/tmp",
		"-v", filepath.ToSlash(abs) + ":/app:ro", "-w", "/app",
		d.image, "sh", "-c", "go run main.go",
	}
	cmd := exec.CommandContext(ctx, "docker", args...)
	cmd.Stdin = strings.NewReader(stdin)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	start := time.Now()
	runErr := cmd.Run()
	res := Result{
		Stdout: truncate(stdout.String(), d.maxOut),
		Stderr: truncate(stderr.String(), d.maxOut),
		RunMS:  time.Since(start).Milliseconds(),
	}
	if ctx.Err() == context.DeadlineExceeded {
		res.Stderr = "Waktu eksekusi habis."
		res.Exit = 124
		return res, nil
	}
	if runErr != nil {
		res.Exit = 1
		return res, nil
	}
	return res, nil
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "\n... output dipotong"
}
