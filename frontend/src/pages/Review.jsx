import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { api } from '../api'
import { useAuth } from '../store/auth'
import Spinner from '../components/Spinner'
import ErrorState from '../components/ErrorState'
import ProgressBar from '../components/ProgressBar'

const ratings = [
  { q: 1, label: 'Lupa', hint: 'Tidak ingat', cls: 'bg-red-500 hover:bg-red-600' },
  { q: 3, label: 'Sulit', hint: 'Ingat susah payah', cls: 'bg-amber-500 hover:bg-amber-600' },
  { q: 4, label: 'Bagus', hint: 'Ingat dengan baik', cls: 'bg-brand-600 hover:bg-brand-700' },
  { q: 5, label: 'Mudah', hint: 'Sangat mudah', cls: 'bg-emerald-600 hover:bg-emerald-700' },
]

export default function Review() {
  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['review-today'],
    queryFn: api.reviewToday,
  })
  const [index, setIndex] = useState(0)
  const [revealed, setRevealed] = useState(false)
  const [reviewed, setReviewed] = useState(0)
  const patchStats = useAuth((s) => s.patchStats)
  const queryClient = useQueryClient()

  useEffect(() => {
    setIndex(0)
    setRevealed(false)
    setReviewed(0)
  }, [data])

  const review = useMutation({
    mutationFn: ({ id, quality }) => api.reviewCard(id, quality),
    onSuccess: () => {
      setReviewed((n) => n + 1)
      setRevealed(false)
      setIndex((i) => i + 1)
      queryClient.invalidateQueries({ queryKey: ['dashboard'] })
      api.stats().then((s) => patchStats(s.stats)).catch(() => {})
    },
  })

  if (isLoading) return <Spinner full />
  if (isError) return <ErrorState error={error} onRetry={refetch} />

  const cards = data || []
  const card = cards[index]
  const done = index >= cards.length

  return (
    <div className="mx-auto max-w-2xl space-y-5">
      <div>
        <h1 className="text-2xl font-bold text-slate-900 dark:text-white">Review Harian</h1>
        <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
          Pengulangan terjadwal agar konsep lama tetap melekat.
        </p>
      </div>

      {cards.length > 0 ? (
        <div>
          <div className="flex items-center justify-between text-xs text-slate-500 dark:text-slate-400">
            <span>
              Kartu {Math.min(index + 1, cards.length)} dari {cards.length}
            </span>
            <span>{reviewed} diulas</span>
          </div>
          <ProgressBar value={(index / cards.length) * 100} className="mt-1.5" />
        </div>
      ) : null}

      {cards.length === 0 ? (
        <div className="card p-10 text-center">
          <p className="text-4xl">🎉</p>
          <h2 className="mt-3 text-lg font-bold text-slate-900 dark:text-white">
            Tidak ada review hari ini
          </h2>
          <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
            Selesaikan pelajaran untuk memasukkan flashcard ke antrean review.
          </p>
          <Link to="/app/roadmap" className="btn-primary mt-5">
            Ke roadmap
          </Link>
        </div>
      ) : done ? (
        <div className="card p-10 text-center">
          <p className="text-4xl">✅</p>
          <h2 className="mt-3 text-lg font-bold text-slate-900 dark:text-white">
            Review selesai!
          </h2>
          <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
            Kamu mengulas {reviewed} kartu. Sampai jumpa besok.
          </p>
          <div className="mt-5 flex justify-center gap-2">
            <Link to="/app" className="btn-primary">Ke dashboard</Link>
            <button className="btn-secondary" onClick={refetch}>Muat ulang</button>
          </div>
        </div>
      ) : (
        <div className="card p-8">
          <p className="text-xs uppercase tracking-wide text-slate-400">
            {card.lesson_title}
          </p>
          <p className="mt-3 min-h-[60px] text-lg font-medium text-slate-900 dark:text-white">
            {card.front}
          </p>

          {revealed ? (
            <div className="mt-4 rounded-xl bg-slate-50 p-4 text-slate-700 dark:bg-slate-800/60 dark:text-slate-200">
              {card.back}
            </div>
          ) : (
            <button className="btn-secondary mt-4 w-full" onClick={() => setRevealed(true)}>
              Tampilkan jawaban
            </button>
          )}

          {revealed ? (
            <div className="mt-5">
              <p className="mb-2 text-xs text-slate-500 dark:text-slate-400">
                Seberapa mudah kamu mengingatnya?
              </p>
              <div className="grid grid-cols-2 gap-2 sm:grid-cols-4">
                {ratings.map((r) => (
                  <button
                    key={r.q}
                    className={`btn text-white ${r.cls}`}
                    disabled={review.isPending}
                    onClick={() => review.mutate({ id: card.flashcard_id, quality: r.q })}
                    title={r.hint}
                  >
                    {r.label}
                  </button>
                ))}
              </div>
            </div>
          ) : null}
        </div>
      )}
    </div>
  )
}
