export default function ErrorState({ error, onRetry }) {
  return (
    <div className="card flex flex-col items-center gap-3 p-8 text-center">
      <span className="text-3xl">⚠️</span>
      <p className="font-semibold text-slate-800 dark:text-slate-100">Terjadi kesalahan</p>
      <p className="max-w-md text-sm text-slate-500 dark:text-slate-400">
        {error?.message || 'Tidak dapat memuat data. Coba lagi.'}
      </p>
      {onRetry ? (
        <button className="btn-secondary" onClick={onRetry}>
          Coba lagi
        </button>
      ) : null}
    </div>
  )
}
