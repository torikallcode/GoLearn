import { Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { api } from '../api'
import Spinner from '../components/Spinner'
import ErrorState from '../components/ErrorState'

const levelChip = {
  Pemula: 'bg-emerald-50 text-emerald-700 dark:bg-emerald-900/20 dark:text-emerald-300',
  Menengah: 'bg-amber-50 text-amber-700 dark:bg-amber-900/20 dark:text-amber-300',
  Lanjut: 'bg-red-50 text-red-700 dark:bg-red-900/20 dark:text-red-300',
}

export default function Projects() {
  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['projects'],
    queryFn: api.projects,
  })

  if (isLoading) return <Spinner full />
  if (isError) return <ErrorState error={error} onRetry={refetch} />

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-slate-900 dark:text-white">Proyek Portofolio</h1>
        <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
          Bangun proyek nyata bertahap untuk memantapkan skill backend dan melengkapi portofolio.
        </p>
      </div>

      <div className="grid gap-5 sm:grid-cols-2">
        {(data || []).map((p) => (
          <Link key={p.slug} to={`/app/projects/${p.slug}`} className="card p-6 transition hover:shadow-md">
            <span className={`chip ${levelChip[p.level] || ''}`}>{p.level}</span>
            <h2 className="mt-3 text-lg font-bold text-slate-900 dark:text-white">{p.title}</h2>
            <p className="mt-1 text-sm text-slate-600 dark:text-slate-400">{p.description}</p>
            <p className="mt-4 text-sm font-medium text-brand-600 dark:text-brand-400">
              Buka proyek →
            </p>
          </Link>
        ))}
      </div>
      {(!data || data.length === 0) ? (
        <p className="text-sm text-slate-500 dark:text-slate-400">Belum ada proyek tersedia.</p>
      ) : null}
    </div>
  )
}
