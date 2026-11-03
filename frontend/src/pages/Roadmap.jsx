import { Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { api } from '../api'
import { useAuth } from '../store/auth'
import Spinner from '../components/Spinner'
import ErrorState from '../components/ErrorState'

export default function Roadmap() {
  const user = useAuth((s) => s.user)
  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['paths'],
    queryFn: api.paths,
  })

  if (isLoading) return <Spinner full />
  if (isError) return <ErrorState error={error} onRetry={refetch} />
  void user

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-slate-900 dark:text-white">Roadmap Belajar</h1>
        <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
          Ikuti path secara berurutan. Setiap pelajaran terbuka setelah pelajaran sebelumnya selesai.
        </p>
      </div>

      <div className="relative space-y-4 border-l-2 border-brand-200 pl-6 dark:border-brand-800">
        {(data || []).map((p, i) => {
          const isAvailable = p.lesson_count > 0
          const body = (
            <div className="card p-6 transition hover:shadow-md">
              <div className="flex flex-wrap items-center gap-2">
                <span className="chip bg-brand-50 text-brand-700 dark:bg-brand-900/30 dark:text-brand-300">
                  Path {i + 1}
                </span>
                <span className="chip bg-slate-100 text-slate-600 dark:bg-slate-800 dark:text-slate-300">
                  {p.level}
                </span>
                {!isAvailable ? (
                  <span className="chip bg-amber-50 text-amber-700 dark:bg-amber-900/20 dark:text-amber-300">
                    Segera hadir
                  </span>
                ) : (
                  <span className="chip bg-emerald-50 text-emerald-700 dark:bg-emerald-900/20 dark:text-emerald-300">
                    {p.lesson_count} pelajaran
                  </span>
                )}
              </div>
              <h2 className="mt-3 text-lg font-bold text-slate-900 dark:text-white">{p.title}</h2>
              <p className="mt-1 text-sm text-slate-600 dark:text-slate-400">{p.description}</p>
              {isAvailable ? (
                <p className="mt-4 text-sm font-medium text-brand-600 dark:text-brand-400">
                  Buka path →
                </p>
              ) : (
                <p className="mt-4 text-sm text-slate-400">Konten sedang disiapkan.</p>
              )}
            </div>
          )
          return (
            <div key={p.slug} className="relative">
              <span className="absolute -left-[35px] top-8 grid h-5 w-5 place-items-center rounded-full border-2 border-brand-400 bg-white text-[10px] dark:bg-slate-950">
                {i + 1}
              </span>
              {isAvailable ? <Link to={`/app/roadmap/${p.slug}`}>{body}</Link> : body}
            </div>
          )
        })}
      </div>
    </div>
  )
}
