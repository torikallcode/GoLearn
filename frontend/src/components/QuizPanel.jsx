import { useState } from 'react'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { api } from '../api'
import { useAuth } from '../store/auth'
import Markdown from './Markdown'

function QuizCard({ quiz }) {
  const [optionId, setOptionId] = useState(null)
  const [answer, setAnswer] = useState('')
  const [result, setResult] = useState(null)
  const patchStats = useAuth((s) => s.patchStats)
  const queryClient = useQueryClient()

  const submit = useMutation({
    mutationFn: () => {
      const payload = quiz.type === 'short' ? { answer } : { option_id: optionId }
      return api.answerQuiz(quiz.id, payload)
    },
    onSuccess: (data) => {
      setResult(data)
      patchStats(data.stats)
      queryClient.invalidateQueries({ queryKey: ['dashboard'] })
    },
  })

  const isShort = quiz.type === 'short'

  return (
    <div className="rounded-xl border border-slate-200 p-4 dark:border-slate-800">
      <p className="font-medium text-slate-900 dark:text-slate-100">{quiz.question}</p>

      {isShort ? (
        <input
          className="input mt-3"
          placeholder="Tulis jawaban singkat…"
          value={answer}
          onChange={(e) => setAnswer(e.target.value)}
          disabled={!!result}
        />
      ) : (
        <div className="mt-3 space-y-2">
          {quiz.options.map((o) => {
            const isCorrectOption = result?.correct_option_id === o.id
            const chosen = optionId === o.id
            let cls =
              'border-slate-200 dark:border-slate-700 hover:border-brand-400'
            if (result) {
              if (isCorrectOption) cls = 'border-emerald-500 bg-emerald-50 dark:bg-emerald-900/20'
              else if (chosen && !isCorrectOption)
                cls = 'border-red-400 bg-red-50 dark:bg-red-900/20'
            } else if (chosen) {
              cls = 'border-brand-500 bg-brand-50 dark:bg-brand-900/20'
            }
            return (
              <label
                key={o.id}
                className={`flex cursor-pointer items-center gap-3 rounded-lg border px-3 py-2 text-sm ${cls}`}
              >
                <input
                  type="radio"
                  name={`quiz-${quiz.id}`}
                  className="accent-brand-600"
                  checked={chosen}
                  onChange={() => setOptionId(o.id)}
                  disabled={!!result}
                />
                <span>{o.text}</span>
              </label>
            )
          })}
        </div>
      )}

      {result ? (
        <div className="mt-3 space-y-2">
          <p className={`text-sm font-semibold ${result.correct ? 'text-emerald-600' : 'text-red-600'}`}>
            {result.correct ? '✅ Benar!' : '❌ Belum tepat.'}
            {!result.correct && result.correct_answer ? ` Jawaban: ${result.correct_answer}` : ''}
            {result.xp_earned > 0 ? ` (+${result.xp_earned} XP)` : ''}
          </p>
          <div className="rounded-lg bg-slate-50 p-3 text-sm text-slate-600 dark:bg-slate-800/60 dark:text-slate-300">
            <Markdown className="text-sm">{result.explanation}</Markdown>
          </div>
        </div>
      ) : (
        <button
          className="btn-primary mt-3 text-xs"
          onClick={() => submit.mutate()}
          disabled={submit.isPending || (isShort ? !answer.trim() : !optionId)}
        >
          {submit.isPending ? 'Memeriksa…' : 'Jawab'}
        </button>
      )}
    </div>
  )
}

export default function QuizPanel({ quizzes }) {
  if (!quizzes?.length) {
    return <p className="text-sm text-slate-500 dark:text-slate-400">Belum ada kuis untuk pelajaran ini.</p>
  }
  return (
    <div className="space-y-4">
      {quizzes.map((q) => (
        <QuizCard key={q.id} quiz={q} />
      ))}
    </div>
  )
}
