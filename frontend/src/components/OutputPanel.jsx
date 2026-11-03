import { friendlyError } from '../utils/friendlyError'

export default function OutputPanel({ result, running, error }) {
  const hint = result?.stderr ? friendlyError(result.stderr) : ''
  return (
    <div className="rounded-xl border border-slate-200 bg-slate-950 p-4 font-mono text-xs dark:border-slate-800">
      <div className="mb-2 flex items-center justify-between text-slate-400">
        <span>Output</span>
        {result ? <span>{result.run_ms} ms</span> : null}
      </div>
      {running ? (
        <p className="text-slate-400">Menjalankan kode…</p>
      ) : error ? (
        <pre className="whitespace-pre-wrap text-red-400">{error}</pre>
      ) : result ? (
        <div className="space-y-2">
          {result.stdout ? (
            <pre className="whitespace-pre-wrap text-emerald-300">{result.stdout}</pre>
          ) : null}
          {result.stderr ? (
            <pre className="whitespace-pre-wrap text-red-400">{result.stderr}</pre>
          ) : null}
          {hint ? (
            <p className="rounded-lg border border-amber-500/30 bg-amber-500/10 px-3 py-2 text-amber-300">
              💡 {hint}
            </p>
          ) : null}
          {!result.stdout && !result.stderr ? (
            <p className="text-slate-500">(tidak ada output)</p>
          ) : null}
        </div>
      ) : (
        <p className="text-slate-500">Tekan Run untuk menjalankan kode.</p>
      )}
    </div>
  )
}
