import { useEffect, useState } from 'react'
import { useMutation } from '@tanstack/react-query'
import { api } from '../api'

export default function NotesPanel({ lesson }) {
  const [content, setContent] = useState(lesson.note || '')
  const [saved, setSaved] = useState(false)

  useEffect(() => {
    setContent(lesson.note || '')
    setSaved(false)
  }, [lesson.id, lesson.note])

  const save = useMutation({
    mutationFn: () => api.saveNote(lesson.id, content),
    onSuccess: () => {
      setSaved(true)
      setTimeout(() => setSaved(false), 2000)
    },
  })

  return (
    <div className="space-y-3">
      <p className="text-sm text-slate-500 dark:text-slate-400">
        Catatan pribadi untuk pelajaran ini. Hanya kamu yang bisa melihatnya.
      </p>
      <textarea
        className="input min-h-[220px] resize-y font-mono text-sm"
        placeholder="Tulis rangkuman, pertanyaan, atau hal yang ingin kamu ingat…"
        value={content}
        onChange={(e) => { setContent(e.target.value); setSaved(false) }}
      />
      <div className="flex items-center gap-3">
        <button
          className="btn-primary text-xs"
          onClick={() => save.mutate()}
          disabled={save.isPending}
        >
          {save.isPending ? 'Menyimpan…' : 'Simpan catatan'}
        </button>
        {saved ? <span className="text-xs text-emerald-600">✓ Tersimpan</span> : null}
        {save.isError ? (
          <span className="text-xs text-red-600">{save.error.message}</span>
        ) : null}
      </div>
    </div>
  )
}
