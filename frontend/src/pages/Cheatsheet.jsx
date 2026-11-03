import { useMemo, useState } from 'react'

const sections = [
  {
    title: 'Dasar & Variabel',
    items: [
      {
        title: 'Deklarasi variabel',
        code: `var nama string = "Budi"\numur := 20            // short declaration\nconst Pi = 3.14`,
      },
      {
        title: 'Tipe data utama',
        code: `bool, string\nint, int8/16/32/64, uint\nfloat32, float64\nbyte (alias uint8), rune (alias int32)`,
      },
      {
        title: 'Zero value',
        code: `var i int     // 0\nvar s string  // ""\nvar b bool    // false\nvar p *int    // nil`,
      },
      {
        title: 'Control flow',
        code: `if x > 0 { /* ... */ } else { /* ... */ }\n\nswitch hari {\ncase "Senin", "Selasa":\n    fmt.Println("awal minggu")\ndefault:\n    fmt.Println("hari lain")\n}\n\nfor i := 0; i < 5; i++ {}\nfor _, v := range data {}`,
      },
    ],
  },
  {
    title: 'Fungsi & Koleksi',
    items: [
      {
        title: 'Fungsi multiple return',
        code: `func bagi(a, b float64) (float64, error) {\n    if b == 0 {\n        return 0, errors.New("pembagi nol")\n    }\n    return a / b, nil\n}`,
      },
      {
        title: 'Variadic & closure',
        code: `func jumlah(nums ...int) int {\n    total := 0\n    for _, n := range nums { total += n }\n    return total\n}\n\ncounter := func() func() int {\n    n := 0\n    return func() int { n++; return n }\n}()`,
      },
      {
        title: 'Slice & map',
        code: `s := []int{1, 2, 3}\ns = append(s, 4)\ns[0] = 10\n\nm := map[string]int{"a": 1}\nm["b"] = 2\nif v, ok := m["a"]; ok { fmt.Println(v) }\ndelete(m, "a")`,
      },
      {
        title: 'String & konversi',
        code: `strings.ToUpper("go")\nstrings.Split("a,b,c", ",")\nstrings.TrimSpace(" x ")\nn, err := strconv.Atoi("42")\ns := strconv.Itoa(42)`,
      },
    ],
  },
  {
    title: 'Struct, Pointer, Interface',
    items: [
      {
        title: 'Struct & method',
        code: `type User struct {\n    ID   int\n    Name string\n}\n\nfunc (u User) Label() string {\n    return fmt.Sprintf("%d:%s", u.ID, u.Name)\n}\n\nu := User{ID: 1, Name: "Ana"}`,
      },
      {
        title: 'Pointer',
        code: `x := 10\np := &x\n*p = 20 // x sekarang 20\n\nfunc tambah(n *int) { *n++ }`,
      },
      {
        title: 'Interface',
        code: `type Shape interface {\n    Area() float64\n}\n\nfunc totalArea(shapes []Shape) float64 {\n    var t float64\n    for _, s := range shapes { t += s.Area() }\n    return t\n}`,
      },
      {
        title: 'Error handling',
        code: `if err != nil {\n    return fmt.Errorf("proses gagal: %w", err)\n}\n\nvar ErrNotFound = errors.New("not found")\nif errors.Is(err, ErrNotFound) { /* ... */ }`,
      },
    ],
  },
  {
    title: 'Concurrency',
    items: [
      {
        title: 'Goroutine & WaitGroup',
        code: `var wg sync.WaitGroup\nfor i := 0; i < 3; i++ {\n    wg.Add(1)\n    go func(id int) {\n        defer wg.Done()\n        fmt.Println("worker", id)\n    }(i)\n}\nwg.Wait()`,
      },
      {
        title: 'Channel',
        code: `ch := make(chan int, 2) // buffered\nch <- 1\nclose(ch)\n\nfor v := range ch {\n    fmt.Println(v)\n}`,
      },
      {
        title: 'Select & context',
        code: `ctx, cancel := context.WithTimeout(context.Background(), time.Second)\ndefer cancel()\n\nselect {\ncase v := <-ch:\n    fmt.Println(v)\ncase <-ctx.Done():\n    fmt.Println("timeout")\n}`,
      },
      {
        title: 'Mutex',
        code: `var mu sync.Mutex\nmu.Lock()\ncounter++\nmu.Unlock()`,
      },
    ],
  },
  {
    title: 'Backend & HTTP',
    items: [
      {
        title: 'Handler net/http',
        code: `http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {\n    w.Header().Set("Content-Type", "application/json")\n    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})\n})\nhttp.ListenAndServe(":8080", nil)`,
      },
      {
        title: 'Gin dasar',
        code: `r := gin.Default()\n\nr.GET("/users/:id", func(c *gin.Context) {\n    id := c.Param("id")\n    c.JSON(200, gin.H{"id": id})\n})\n\nr.POST("/users", func(c *gin.Context) {\n    var body struct{ Name string \`json:"name"\` }\n    if err := c.ShouldBindJSON(&body); err != nil {\n        c.JSON(400, gin.H{"error": err.Error()})\n        return\n    }\n})`,
      },
      {
        title: 'JSON encode/decode',
        code: `data, _ := json.Marshal(v)\nvar out T\nerr := json.Unmarshal(data, &out)`,
      },
      {
        title: 'Query SQL',
        code: `row := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", id)\nvar name string\nif err := row.Scan(&name); err != nil { /* ... */ }`,
      },
    ],
  },
]

function CodeBlock({ code }) {
  const [copied, setCopied] = useState(false)
  async function copy() {
    try {
      await navigator.clipboard.writeText(code)
      setCopied(true)
      setTimeout(() => setCopied(false), 1500)
    } catch {
      /* ignore */
    }
  }
  return (
    <div className="group relative">
      <pre className="overflow-x-auto rounded-xl bg-slate-950 p-4 font-mono text-xs text-slate-100">
        {code}
      </pre>
      <button
        onClick={copy}
        className="absolute right-2 top-2 rounded-md bg-slate-800 px-2 py-1 text-[10px] text-slate-300 opacity-0 transition group-hover:opacity-100"
      >
        {copied ? 'Tersalin!' : 'Salin'}
      </button>
    </div>
  )
}

export default function Cheatsheet() {
  const [query, setQuery] = useState('')

  const filtered = useMemo(() => {
    const q = query.trim().toLowerCase()
    if (!q) return sections
    return sections
      .map((sec) => ({
        ...sec,
        items: sec.items.filter(
          (it) =>
            it.title.toLowerCase().includes(q) ||
            it.code.toLowerCase().includes(q) ||
            sec.title.toLowerCase().includes(q),
        ),
      }))
      .filter((sec) => sec.items.length > 0)
  }, [query])

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-slate-900 dark:text-white">Cheatsheet Go</h1>
        <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
          Ringkasan cepat sintaks dan pola penting. Cari lalu salin.
        </p>
      </div>

      <input
        className="input max-w-md"
        placeholder="Cari sintaks, mis. channel, struct, json…"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
      />

      <div className="space-y-8">
        {filtered.map((sec) => (
          <section key={sec.title}>
            <h2 className="mb-3 text-lg font-bold text-slate-900 dark:text-white">{sec.title}</h2>
            <div className="grid gap-4 lg:grid-cols-2">
              {sec.items.map((it) => (
                <div key={it.title} className="card p-4">
                  <h3 className="mb-2 text-sm font-semibold text-slate-700 dark:text-slate-200">
                    {it.title}
                  </h3>
                  <CodeBlock code={it.code} />
                </div>
              ))}
            </div>
          </section>
        ))}
      </div>
      {filtered.length === 0 ? (
        <p className="text-sm text-slate-500 dark:text-slate-400">Tidak ada hasil untuk “{query}”.</p>
      ) : null}
    </div>
  )
}
