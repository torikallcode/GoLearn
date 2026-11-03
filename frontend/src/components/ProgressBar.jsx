export default function ProgressBar({ value = 0, className = '', showLabel = false }) {
  const pct = Math.max(0, Math.min(100, Math.round(value)))
  return (
    <div className={className}>
      <div className="h-2 w-full overflow-hidden rounded-full bg-slate-200 dark:bg-slate-800">
        <div
          className="h-full rounded-full bg-brand-500 transition-all"
          style={{ width: `${pct}%` }}
        />
      </div>
      {showLabel ? (
        <p className="mt-1 text-xs text-slate-500 dark:text-slate-400">{pct}% selesai</p>
      ) : null}
    </div>
  )
}
