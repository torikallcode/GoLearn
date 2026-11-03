import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { api } from '../api'
import { useAuth } from '../store/auth'
import Markdown from '../components/Markdown'
import Spinner from '../components/Spinner'
import ErrorState from '../components/ErrorState'
import Playground from '../components/Playground'
import ExercisePanel from '../components/ExercisePanel'
import QuizPanel from '../components/QuizPanel'
import NotesPanel from '../components/NotesPanel'

const panelTabs = [
  { id: 'playground', label: 'Playground' },
  { id: 'latihan', label: 'Latihan' },
  { id: 'kuis', label: 'Kuis' },
  { id: 'catatan', label: 'Catatan' },
]

export default function Lesson() {
  const { slug } = useParams()
  const queryClient = useQueryClient()
  const patchStats = useAuth((s) => s.patchStats)
  const [panelTab, setPanelTab] = useState('playground')
  const [mobileView, setMobileView] = useState('materi')
  const [bookmarked, setBookmarked] = useState(false)

  const lessonQuery = useQuery({ queryKey: ['lesson', slug], queryFn: () => api.lesson(slug) })
  const exercisesQuery = useQuery({
    queryKey: ['exercises', slug],
    queryFn: () => api.exercises(slug),
    enabled: !!lessonQuery.data,
  })
  const quizzesQuery = useQuery({
    queryKey: ['quizzes', slug],
    queryFn: () => api.quizzes(slug),
    enabled: !!lessonQuery.data,
  })

  const lesson = lessonQuery.data
  useEffect(() => {
    if (lesson) setBookmarked(!!lesson.bookmarked)
  }, [lesson?.id, lesson?.bookmarked])

  const complete = useMutation({
    mutationFn: () => api.completeLesson(slug),
    onSuccess: (data) => {
      patchStats(data.stats)
      queryClient.invalidateQueries({ queryKey: ['dashboard'] })
      queryClient.invalidateQueries({ queryKey: ['lesson', slug] })
      queryClient.invalidateQueries({ queryKey: ['path'] })
      queryClient.invalidateQueries({ queryKey: ['paths'] })
    },
  })

  const bookmark = useMutation({
    mutationFn: () =>
      bookmarked ? api.removeBookmark(lesson.lesson.id) : api.addBookmark(lesson.lesson.id),
    onSuccess: () => setBookmarked((v) => !v),
  })

  if (lessonQuery.isLoading) return <Spinner full />
  if (lessonQuery.isError) {
    return <ErrorState error={lessonQuery.error} onRetry={lessonQuery.refetch} />
  }

  const { lesson: l, module, path, prev, next } = lesson

  return (
    <div className="space-y-5">
      {/* Header */}
      <div>
        <div className="flex flex-wrap items-center gap-2 text-sm text-slate-500 dark:text-slate-400">
          <Link to="/app/roadmap" className="hover:text-brand-600">Roadmap</Link>
          <span>/</span>
          <Link to={`/app/roadmap/${path.slug}`} className="hover:text-brand-600">
            {path.title}
          </Link>
          <span>/</span>
          <span>{module.title}</span>
        </div>
        <div className="mt-2 flex flex-wrap items-start justify-between gap-3">
          <div>
            <h1 className="text-2xl font-bold text-slate-900 dark:text-white">{l.title}</h1>
            <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
              ± {l.estimated_minutes} menit · {l.summary}
            </p>
          </div>
          <div className="flex items-center gap-2">
            <button
              className="btn-secondary text-xs"
              onClick={() => bookmark.mutate()}
              disabled={bookmark.isPending}
            >
              {bookmarked ? '🔖 Tersimpan' : '📑 Simpan'}
            </button>
            <button
              className={`text-xs ${lesson.completed ? 'btn-secondary' : 'btn-primary'}`}
              onClick={() => complete.mutate()}
              disabled={complete.isPending || lesson.completed}
            >
              {lesson.completed ? '✓ Selesai' : complete.isPending ? 'Menyimpan…' : 'Tandai selesai'}
            </button>
          </div>
        </div>
      </div>

      {/* Mobile view switch */}
      <div className="flex rounded-lg bg-slate-100 p-1 dark:bg-slate-800 lg:hidden">
        {[
          ['materi', 'Materi'],
          ['praktik', 'Praktik'],
        ].map(([id, label]) => (
          <button
            key={id}
            onClick={() => setMobileView(id)}
            className={`flex-1 rounded-md px-3 py-1.5 text-sm font-medium ${
              mobileView === id
                ? 'bg-white text-brand-700 shadow-sm dark:bg-slate-900 dark:text-brand-300'
                : 'text-slate-500 dark:text-slate-400'
            }`}
          >
            {label}
          </button>
        ))}
      </div>

      <div className="grid gap-6 lg:grid-cols-2">
        {/* Content */}
        <div className={`${mobileView === 'materi' ? 'block' : 'hidden'} lg:block`}>
          <article className="card p-6">
            <Markdown>{l.content_md}</Markdown>
          </article>

          <div className="mt-4 flex items-center justify-between gap-3">
            {prev ? (
              <Link to={`/app/lessons/${prev.slug}`} className="btn-secondary min-w-0 text-left">
                <span className="truncate text-xs text-slate-400">← Sebelumnya</span>
                <span className="block truncate text-sm">{prev.title}</span>
              </Link>
            ) : (
              <span />
            )}
            {next ? (
              <Link to={`/app/lessons/${next.slug}`} className="btn-primary min-w-0 text-right">
                <span className="truncate text-xs text-white/70">Berikutnya →</span>
                <span className="block truncate text-sm">{next.title}</span>
              </Link>
            ) : (
              <span />
            )}
          </div>
        </div>

        {/* Practice panel */}
        <div className={`${mobileView === 'praktik' ? 'block' : 'hidden'} lg:block`}>
          <div className="card sticky top-4 p-4">
            <div className="flex flex-wrap gap-1 rounded-lg bg-slate-100 p-1 dark:bg-slate-800">
              {panelTabs.map((t) => (
                <button
                  key={t.id}
                  onClick={() => setPanelTab(t.id)}
                  className={`flex-1 rounded-md px-2 py-1.5 text-xs font-medium transition ${
                    panelTab === t.id
                      ? 'bg-white text-brand-700 shadow-sm dark:bg-slate-900 dark:text-brand-300'
                      : 'text-slate-500 hover:text-slate-700 dark:text-slate-400'
                  }`}
                >
                  {t.label}
                </button>
              ))}
            </div>

            <div className="mt-4">
              {panelTab === 'playground' ? <Playground /> : null}
              {panelTab === 'latihan' ? (
                exercisesQuery.isLoading ? (
                  <Spinner label="Memuat latihan…" />
                ) : (
                  <ExercisePanel exercises={exercisesQuery.data || []} />
                )
              ) : null}
              {panelTab === 'kuis' ? (
                quizzesQuery.isLoading ? (
                  <Spinner label="Memuat kuis…" />
                ) : (
                  <QuizPanel quizzes={quizzesQuery.data || []} />
                )
              ) : null}
              {panelTab === 'catatan' ? <NotesPanel lesson={l} /> : null}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
