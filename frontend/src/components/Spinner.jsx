export default function Spinner({ full = false, label = 'Memuat…' }) {
  const content = (
    <div className="flex flex-col items-center gap-3 text-slate-500 dark:text-slate-400">
      <span className="h-7 w-7 animate-spin rounded-full border-2 border-current border-t-transparent" />
      {label ? <span className="text-sm">{label}</span> : null}
    </div>
  )
  if (!full) return content
  return <div className="flex min-h-[60vh] items-center justify-center">{content}</div>
}
