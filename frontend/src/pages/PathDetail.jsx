import { Link, useParams } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { api } from '../api'
import ProgressBar from '../components/ProgressBar'
import Spinner from '../components/Spinner'
import ErrorState from '../components/ErrorState'

const statusChip = {
  completed: 'bg-emerald-50 text-emerald-700 dark:bg-emerald-900/20 dark:text-emerald-300',
  available: 'bg-brand-50 text-brand-700 dark:bg-brand-900/30 dark:text-brand-300',
  locked: 'bg-slate-100 text-slate-500 dark:bg-slate-800 dark:text-slate-400',
}
const statusLabel = { completed: 'Selesai', available: 'Aktif', locked: 'Terkunci' }

export default function PathDetail() {
  const { slug } = useParams()
  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['path', slug],
    queryFn: () => api.path(slug),
  })

  if (isLoading) return <Spinner full />
  if (isError) return <ErrorState error={error} onRetry={refetch} />

  return (
    <div className="space-y-6">
      <div>
        <Link to="/app/roadmap" className="text-sm text-slate-500 hover:text-brand-600 dark:text-slate-400">
          ← Kembali ke roadmap
        </Link>
        <div className="mt-3 flex flex-wrap items-center gap-2">
          <span className="chip bg-slate-100 text-slate-600 dark:bg-slate-800 dark:text-slate-300">
            {data.level}
          </span>
          <span className="chip bg-brand-50 text-brand-700 dark:bg-brand-900/30 dark:text-brand-300">
            {data.completed}/{data.total} pelajaran
          </span>
        </div>
        <h1 className="mt-2 text-2xl font-bold text-slate-900 dark:text-white">{data.title}</h1>
        <p className="mt-1 max-w-2xl text-sm text-slate-600 dark:text-slate-400">
          {data.description}
        </p>
        <div className="mt-4 max-w-md">
          <ProgressBar value={data.percent} showLabel />
        </div>
      </div>

      <div className="space-y-6">
        {data.modules.map((m, mi) => (
          <section key={m.id} className="card p-6">
            <div className="flex items-center gap-3">
              <span className="grid h-8 w-8 place-items-center rounded-lg bg-brand-600 text-sm font-bold text-white">
                {mi + 1}
              </span>
              <h2 className="text-lg font-bold text-slate-900 dark:text-white">{m.title}</h2>
            </div>
            <ul className="mt-4 divide-y divide-slate-100 dark:divide-slate-800">
              {m.lessons.map((l) => {
                const locked = l.status === 'locked'
                const inner = (
                  <div className="flex items-center justify-between gap-3 py-3">
                    <div className="min-w-0">
                      <div className="flex items-center gap-2">
                        <p
                          className={`truncate font-medium ${
                            locked
                              ? 'text-slate-400'
                              : 'text-slate-800 group-hover:text-brand-600 dark:text-slate-100'
                          }`}
                        >
                          {locked ? '🔒 ' : ''}
                          {l.title}
                        </p>
                      </div>
                      <p className="mt-0.5 truncate text-xs text-slate-500 dark:text-slate-400">
                        {l.summary}
                      </p>
                    </div>
                    <div className="flex shrink-0 items-center gap-2">
                      <span className="hidden text-xs text-slate-400 sm:inline">
                        {l.estimated_minutes} mnt
                      </span>
                      <span className={`chip ${statusChip[l.status] || statusChip.locked}`}>
                        {statusLabel[l.status] || l.status}
                      </span>
                    </div>
                  </div>
                )
                return (
                  <li key={l.id}>
                    {locked ? (
                      <div className="cursor-not-allowed opacity-70" title="Selesaikan pelajaran sebelumnya dulu">
                        {inner}
                      </div>
                    ) : (
                      <Link to={`/app/lessons/${l.slug}`} className="group block">
                        {inner}
                      </Link>
                    )}
                  </li>
                )
              })}
            </ul>
          </section>
        ))}
      </div>
    </div>
  )
}
