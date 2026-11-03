import { useState } from 'react'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { api } from '../api'
import { useAuth } from '../store/auth'
import CodeEditor from './CodeEditor'
import OutputPanel from './OutputPanel'
import Markdown from './Markdown'

const diffChip = {
  mudah: 'bg-emerald-50 text-emerald-700 dark:bg-emerald-900/20 dark:text-emerald-300',
  sedang: 'bg-amber-50 text-amber-700 dark:bg-amber-900/20 dark:text-amber-300',
  sulit: 'bg-red-50 text-red-700 dark:bg-red-900/20 dark:text-red-300',
}

function ExerciseCard({ exercise }) {
  const [code, setCode] = useState(exercise.starter_code || '')
  const [hintsShown, setHintsShown] = useState(0)
  const [runResult, setRunResult] = useState(null)
  const [submission, setSubmission] = useState(null)
  const [error, setError] = useState('')
  const patchStats = useAuth((s) => s.patchStats)
  const queryClient = useQueryClient()

  const run = useMutation({
    mutationFn: () => api.run(code),
    onMutate: () => { setError(''); setRunResult(null) },
    onSuccess: setRunResult,
    onError: (e) => setError(e.message),
  })

  const submit = useMutation({
    mutationFn: () => api.submitExercise(exercise.id, code),
    onMutate: () => { setError(''); setSubmission(null) },
    onSuccess: (data) => {
      setSubmission(data)
      patchStats(data.stats)
      queryClient.invalidateQueries({ queryKey: ['dashboard'] })
      queryClient.invalidateQueries({ queryKey: ['exercises'] })
    },
    onError: (e) => setError(e.message),
  })

  const hints = exercise.hints || []

  return (
    <div className="space-y-3">
      <div className="flex flex-wrap items-center gap-2">
        <span className={`chip ${diffChip[exercise.difficulty] || diffChip.mudah}`}>
          {exercise.difficulty}
        </span>
        <span className="chip bg-brand-50 text-brand-700 dark:bg-brand-900/30 dark:text-brand-300">
          +{exercise.xp_reward} XP
        </span>
        {exercise.passed ? (
          <span className="chip bg-emerald-50 text-emerald-700 dark:bg-emerald-900/20 dark:text-emerald-300">
            ✓ Selesai
          </span>
        ) : null}
      </div>

      <Markdown className="text-sm">{exercise.instructions_md}</Markdown>

      <CodeEditor value={code} onChange={setCode} height={260} />

      <div className="flex flex-wrap items-center gap-2">
        <button className="btn-primary" onClick={() => submit.mutate()} disabled={submit.isPending}>
          {submit.isPending ? 'Menguji…' : 'Submit jawaban'}
        </button>
        <button className="btn-secondary" onClick={() => run.mutate()} disabled={run.isPending}>
          ▶ Run
        </button>
        {hints.length > 0 && hintsShown < hints.length ? (
          <button className="btn-ghost text-xs" onClick={() => setHintsShown((n) => n + 1)}>
            💡 Tampilkan hint ({hintsShown}/{hints.length})
          </button>
        ) : null}
      </div>

      {hintsShown > 0 ? (
        <ul className="space-y-1 rounded-lg bg-amber-50 p-3 text-sm text-amber-800 dark:bg-amber-900/20 dark:text-amber-200">
          {hints.slice(0, hintsShown).map((h, i) => (
            <li key={i}>💡 {h}</li>
          ))}
        </ul>
      ) : null}

      {runResult ? <OutputPanel result={runResult} /> : null}
      {error ? (
        <p className="rounded-lg bg-red-50 px-3 py-2 text-sm text-red-600 dark:bg-red-900/20 dark:text-red-400">
          {error}
        </p>
      ) : null}

      {submission ? (
        <div className="space-y-3 rounded-xl border border-slate-200 p-4 dark:border-slate-800">
          <p className={`font-semibold ${submission.passed ? 'text-emerald-600' : 'text-red-600'}`}>
            {submission.passed ? '✅ ' : '❌ '}
            {submission.message}
          </p>
          <ul className="space-y-1 text-sm">
            {submission.tests.map((t) => (
              <li key={t.index} className="flex items-start gap-2">
                <span>{t.passed ? '✅' : '❌'}</span>
                <span className="text-slate-600 dark:text-slate-300">
                  Test {t.index}
                  {t.hidden ? ' (tersembunyi)' : ''}
                  {!t.passed && t.expected !== undefined ? (
                    <span className="mt-1 block font-mono text-xs text-slate-500">
                      harapan: {t.expected || '(kosong)'} · hasil: {t.actual || '(kosong)'}
                    </span>
                  ) : null}
                </span>
              </li>
            ))}
          </ul>
          {submission.xp_earned > 0 ? (
            <p className="text-sm font-semibold text-brand-600">+{submission.xp_earned} XP</p>
          ) : null}
        </div>
      ) : null}

      {exercise.passed && exercise.solution_code ? (
        <details className="rounded-xl border border-slate-200 p-3 dark:border-slate-800">
          <summary className="cursor-pointer text-sm font-medium text-slate-700 dark:text-slate-200">
            Lihat solusi referensi
          </summary>
          <pre className="mt-2 overflow-x-auto rounded-lg bg-slate-950 p-3 font-mono text-xs text-slate-100">
            {exercise.solution_code}
          </pre>
        </details>
      ) : null}
    </div>
  )
}

export default function ExercisePanel({ exercises }) {
  const [idx, setIdx] = useState(0)
  if (!exercises?.length) {
    return <p className="text-sm text-slate-500 dark:text-slate-400">Belum ada latihan untuk pelajaran ini.</p>
  }
  return (
    <div className="space-y-4">
      {exercises.length > 1 ? (
        <div className="flex flex-wrap gap-2">
          {exercises.map((ex, i) => (
            <button
              key={ex.id}
              onClick={() => setIdx(i)}
              className={`rounded-lg px-3 py-1.5 text-xs font-medium ${
                i === idx
                  ? 'bg-brand-600 text-white'
                  : 'bg-slate-100 text-slate-600 dark:bg-slate-800 dark:text-slate-300'
              }`}
            >
              {ex.passed ? '✓ ' : ''}
              {ex.title}
            </button>
          ))}
        </div>
      ) : null}
      <ExerciseCard key={exercises[idx].id} exercise={exercises[idx]} />
    </div>
  )
}
