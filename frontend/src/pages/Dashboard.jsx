import { Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { api } from '../api'
import { useAuth } from '../store/auth'
import ProgressBar from '../components/ProgressBar'
import Spinner from '../components/Spinner'
import ErrorState from '../components/ErrorState'

function StatCard({ label, value, hint, accent }) {
  return (
    <div className="card p-5">
      <p className="text-xs font-medium uppercase tracking-wide text-slate-500 dark:text-slate-400">
        {label}
      </p>
      <p className={`mt-2 text-3xl font-bold ${accent || 'text-slate-900 dark:text-white'}`}>
        {value}
      </p>
      {hint ? <p className="mt-1 text-xs text-slate-400">{hint}</p> : null}
    </div>
  )
}

export default function Dashboard() {
  const user = useAuth((s) => s.user)
  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['dashboard'],
    queryFn: api.dashboard,
  })

  if (isLoading) return <Spinner full />
  if (isError) return <ErrorState error={error} onRetry={refetch} />

  const stats = data.stats || {}
  const pct = data.total_lessons ? Math.round((data.total_completed / data.total_lessons) * 100) : 0

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-slate-900 dark:text-white">
          Halo, {user?.name?.split(' ')[0] || 'Sobat'} 👋
        </h1>
        <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
          Lanjutkan momentum belajarmu hari ini.
        </p>
      </div>

      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
        <StatCard label="XP" value={stats.xp ?? 0} hint={`Level ${stats.level ?? 1}`} accent="text-brand-600 dark:text-brand-400" />
        <StatCard label="Streak" value={`${stats.current_streak ?? 0} 🔥`} hint={`Terpanjang ${stats.longest_streak ?? 0} hari`} />
        <StatCard label="Pelajaran selesai" value={data.total_completed} hint={`dari ${data.total_lessons}`} />
        <StatCard
          label="Review hari ini"
          value={data.due_reviews}
          hint={data.due_reviews > 0 ? 'Siap diulang' : 'Semua beres'}
          accent={data.due_reviews > 0 ? 'text-amber-500' : undefined}
        />
      </div>

      <div className="grid gap-5 lg:grid-cols-3">
        <div className="card p-6 lg:col-span-2">
          <div className="flex items-center justify-between">
            <h2 className="font-semibold text-slate-900 dark:text-white">Lanjutkan belajar</h2>
            <Link to="/app/roadmap" className="text-sm font-medium text-brand-600 hover:underline">
              Roadmap
            </Link>
          </div>
          {data.continue ? (
            <div className="mt-4 flex flex-col gap-3 rounded-xl bg-brand-50 p-4 dark:bg-brand-900/20 sm:flex-row sm:items-center sm:justify-between">
              <div>
                <p className="text-xs uppercase tracking-wide text-brand-700 dark:text-brand-300">
                  Berikutnya
                </p>
                <p className="mt-1 font-semibold text-slate-900 dark:text-white">
                  {data.continue.title}
                </p>
                <p className="text-xs text-slate-500 dark:text-slate-400">
                  ± {data.continue.minutes} menit
                </p>
              </div>
              <Link className="btn-primary shrink-0" to={`/app/lessons/${data.continue.slug}`}>
                Mulai pelajaran
              </Link>
            </div>
          ) : (
            <p className="mt-4 text-sm text-slate-500 dark:text-slate-400">
              Semua pelajaran tersedia sudah selesai. Kerja bagus! 🎉
            </p>
          )}

          <div className="mt-5">
            <div className="flex items-center justify-between text-sm">
              <span className="text-slate-600 dark:text-slate-300">Progres keseluruhan</span>
              <span className="font-semibold text-slate-900 dark:text-white">{pct}%</span>
            </div>
            <ProgressBar value={pct} className="mt-2" />
          </div>

          {data.recent?.length ? (
            <div className="mt-6">
              <h3 className="text-sm font-semibold text-slate-700 dark:text-slate-200">
                Terakhir diselesaikan
              </h3>
              <ul className="mt-2 divide-y divide-slate-100 dark:divide-slate-800">
                {data.recent.map((r) => (
                  <li key={r.lesson_id} className="flex items-center justify-between py-2 text-sm">
                    <Link
                      to={`/app/lessons/${r.slug}`}
                      className="truncate pr-3 font-medium text-slate-700 hover:text-brand-600 dark:text-slate-200"
                    >
                      {r.title}
                    </Link>
                    <span className="shrink-0 text-xs text-slate-400">{r.path_title}</span>
                  </li>
                ))}
              </ul>
            </div>
          ) : null}
        </div>

        <div className="space-y-4">
          <div className="card p-6">
            <h2 className="font-semibold text-slate-900 dark:text-white">Review hari ini</h2>
            <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
              {data.due_reviews > 0
                ? `${data.due_reviews} flashcard menunggu diulang.`
                : 'Tidak ada review yang jatuh tempo.'}
            </p>
            <Link
              to="/app/review"
              className={`mt-4 w-full ${data.due_reviews > 0 ? 'btn-primary' : 'btn-secondary'}`}
            >
              {data.due_reviews > 0 ? 'Mulai review' : 'Lihat review'}
            </Link>
          </div>

          <div className="card p-6">
            <h2 className="font-semibold text-slate-900 dark:text-white">Progres per path</h2>
            <div className="mt-4 space-y-4">
              {(data.paths || []).map((p) => (
                <Link key={p.slug} to={`/app/roadmap/${p.slug}`} className="block">
                  <div className="flex items-center justify-between text-sm">
                    <span className="truncate pr-2 font-medium text-slate-700 dark:text-slate-200">
                      {p.title}
                    </span>
                    <span className="shrink-0 text-xs text-slate-500">
                      {p.completed}/{p.total}
                    </span>
                  </div>
                  <ProgressBar value={p.percent} className="mt-1.5" />
                </Link>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
