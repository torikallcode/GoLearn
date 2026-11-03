import { useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { api } from '../api'
import Markdown from '../components/Markdown'
import ProgressBar from '../components/ProgressBar'
import Spinner from '../components/Spinner'
import ErrorState from '../components/ErrorState'

export default function ProjectDetail() {
  const { slug } = useParams()
  const queryClient = useQueryClient()
  const [open, setOpen] = useState(null)

  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['project', slug],
    queryFn: () => api.project(slug),
  })

  const completeStep = useMutation({
    mutationFn: (id) => api.completeProjectStep(id),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['project', slug] }),
  })

  if (isLoading) return <Spinner full />
  if (isError) return <ErrorState error={error} onRetry={refetch} />

  const steps = data.steps || []
  const doneCount = steps.filter((s) => s.is_done).length
  const pct = steps.length ? Math.round((doneCount / steps.length) * 100) : 0

  return (
    <div className="space-y-6">
      <div>
        <Link to="/app/projects" className="text-sm text-slate-500 hover:text-brand-600 dark:text-slate-400">
          ← Semua proyek
        </Link>
        <h1 className="mt-2 text-2xl font-bold text-slate-900 dark:text-white">{data.title}</h1>
        <p className="mt-1 max-w-2xl text-sm text-slate-600 dark:text-slate-400">
          {data.description}
        </p>
        <div className="mt-4 max-w-md">
          <div className="flex items-center justify-between text-xs text-slate-500 dark:text-slate-400">
            <span>{doneCount}/{steps.length} langkah selesai</span>
            <span>{pct}%</span>
          </div>
          <ProgressBar value={pct} className="mt-1.5" />
        </div>
      </div>

      <div className="space-y-3">
        {steps.map((s, i) => {
          const isOpen = open === s.id
          return (
            <div key={s.id} className="card overflow-hidden">
              <div className="flex items-center justify-between gap-3 p-4">
                <button
                  className="flex min-w-0 flex-1 items-center gap-3 text-left"
                  onClick={() => setOpen(isOpen ? null : s.id)}
                >
                  <span className="grid h-8 w-8 shrink-0 place-items-center rounded-lg bg-brand-600 text-sm font-bold text-white">
                    {i + 1}
                  </span>
                  <span className="min-w-0">
                    <span className={`block truncate font-semibold ${s.is_done ? 'text-slate-400 line-through' : 'text-slate-900 dark:text-white'}`}>
                      {s.title}
                    </span>
                  </span>
                </button>
                <div className="flex shrink-0 items-center gap-2">
                  {s.is_done ? (
                    <span className="chip bg-emerald-50 text-emerald-700 dark:bg-emerald-900/20 dark:text-emerald-300">
                      ✓ Selesai
                    </span>
                  ) : (
                    <button
                      className="btn-secondary text-xs"
                      onClick={() => completeStep.mutate(s.id)}
                      disabled={completeStep.isPending}
                    >
                      Tandai selesai
                    </button>
                  )}
                  <button className="btn-ghost text-xs" onClick={() => setOpen(isOpen ? null : s.id)}>
                    {isOpen ? '▲' : '▼'}
                  </button>
                </div>
              </div>
              {isOpen ? (
                <div className="border-t border-slate-200 p-5 dark:border-slate-800">
                  <Markdown>{s.content_md}</Markdown>
                </div>
              ) : null}
            </div>
          )
        })}
      </div>
    </div>
  )
}
