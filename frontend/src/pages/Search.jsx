import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { api } from '../api'
import Spinner from '../components/Spinner'

export default function Search() {
  const [input, setInput] = useState('')
  const [term, setTerm] = useState('')

  useEffect(() => {
    const t = setTimeout(() => setTerm(input.trim()), 350)
    return () => clearTimeout(t)
  }, [input])

  const { data, isFetching, isError, error } = useQuery({
    queryKey: ['search', term],
    queryFn: () => api.search(term),
    enabled: term.length >= 2,
  })

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-slate-900 dark:text-white">Pencarian</h1>
        <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
          Cari materi, istilah, dan snippet di seluruh pelajaran.
        </p>
      </div>

      <input
        className="input max-w-xl"
        placeholder="Kata kunci, mis. goroutine, interface, slice…"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        autoFocus
      />

      {term.length < 2 ? (
        <p className="text-sm text-slate-500 dark:text-slate-400">
          Ketik minimal 2 karakter untuk mulai mencari.
        </p>
      ) : isFetching ? (
        <Spinner label="Mencari…" />
      ) : isError ? (
        <p className="text-sm text-red-600">{error.message}</p>
      ) : (data || []).length === 0 ? (
        <p className="text-sm text-slate-500 dark:text-slate-400">
          Tidak ada hasil untuk “{term}”.
        </p>
      ) : (
        <ul className="space-y-3">
          {data.map((hit) => (
            <li key={hit.slug}>
              <Link to={`/app/lessons/${hit.slug}`} className="card block p-4 transition hover:shadow-md">
                <span className="chip bg-slate-100 text-slate-600 dark:bg-slate-800 dark:text-slate-300">
                  {hit.path_title}
                </span>
                <h2 className="mt-2 font-semibold text-slate-900 dark:text-white">{hit.title}</h2>
                <p className="mt-1 text-sm text-slate-600 dark:text-slate-400">{hit.summary}</p>
                {hit.snippet ? (
                  <p className="mt-2 line-clamp-2 font-mono text-xs text-slate-500 dark:text-slate-400">
                    …{hit.snippet}…
                  </p>
                ) : null}
              </Link>
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}
