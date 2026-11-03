import { Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { api } from '../api'
import { useAuth } from '../store/auth'
import { useTheme } from '../store/theme'
import ProgressBar from '../components/ProgressBar'

const features = [
  {
    icon: '🧩',
    title: 'Materi bite-sized',
    desc: 'Satu pelajaran satu konsep, 10–15 menit. Fokus, tanpa bertele-tele.',
  },
  {
    icon: '⌨️',
    title: 'Editor kode langsung',
    desc: 'Tulis dan jalankan Go di browser tanpa instalasi.',
  },
  {
    icon: '✅',
    title: 'Latihan & kuis instan',
    desc: 'Umpan balik otomatis plus penjelasan di setiap jawaban.',
  },
  {
    icon: '🔁',
    title: 'Spaced repetition',
    desc: 'Flashcard dijadwalkan agar konsep lama tidak mudah lupa.',
  },
  {
    icon: '🛠️',
    title: 'Proyek nyata',
    desc: 'Bangun portofolio dari CLI todo sampai API backend.',
  },
  {
    icon: '📈',
    title: 'Progres terukur',
    desc: 'XP, streak, dan checklist skill menuju siap kerja.',
  },
]

export default function Landing() {
  const status = useAuth((s) => s.status)
  const theme = useTheme((s) => s.theme)
  const toggleTheme = useTheme((s) => s.toggle)
  const { data: paths } = useQuery({ queryKey: ['paths'], queryFn: api.paths })

  const ctaTo = status === 'authed' ? '/app' : '/register'

  return (
    <div className="min-h-screen bg-gradient-to-b from-white to-slate-50 dark:from-slate-950 dark:to-slate-900">
      <header className="mx-auto flex max-w-6xl items-center justify-between px-4 py-5 sm:px-6">
        <div className="flex items-center gap-2">
          <span className="grid h-9 w-9 place-items-center rounded-xl bg-brand-600 font-mono text-sm font-bold text-white">
            Go
          </span>
          <span className="text-lg font-bold tracking-tight text-slate-900 dark:text-white">
            GoLearn
          </span>
        </div>
        <div className="flex items-center gap-2">
          <button className="btn-ghost" onClick={toggleTheme} aria-label="Ganti tema">
            {theme === 'dark' ? '☀️' : '🌙'}
          </button>
          {status === 'authed' ? (
            <Link className="btn-primary" to="/app">
              Buka Dashboard
            </Link>
          ) : (
            <>
              <Link className="btn-ghost" to="/login">
                Masuk
              </Link>
              <Link className="btn-primary" to="/register">
                Daftar gratis
              </Link>
            </>
          )}
        </div>
      </header>

      <section className="mx-auto max-w-6xl px-4 pb-16 pt-10 sm:px-6 lg:pt-16">
        <div className="grid items-center gap-10 lg:grid-cols-2">
          <div>
            <span className="chip bg-brand-50 text-brand-700 dark:bg-brand-900/30 dark:text-brand-300">
              Belajar · Praktik · Siap kerja
            </span>
            <h1 className="mt-4 text-4xl font-extrabold leading-tight tracking-tight text-slate-900 dark:text-white sm:text-5xl">
              Jadi backend engineer Go,{' '}
              <span className="text-brand-600 dark:text-brand-400">dari nol sampai siap kerja.</span>
            </h1>
            <p className="mt-5 max-w-xl text-lg text-slate-600 dark:text-slate-300">
              GoLearn menggabungkan materi ringkas, editor kode interaktif, latihan, kuis, dan
              pengulangan terjadwal — dirancang agar kamu benar-benar paham, bukan sekadar membaca.
            </p>
            <div className="mt-8 flex flex-wrap gap-3">
              <Link className="btn-primary px-6 py-3 text-base" to={ctaTo}>
                Mulai belajar sekarang
              </Link>
              <a className="btn-secondary px-6 py-3 text-base" href="#kurikulum">
                Lihat kurikulum
              </a>
            </div>
            <dl className="mt-10 grid max-w-md grid-cols-3 gap-4">
              {[
                ['11', 'Pelajaran Fondasi'],
                ['100%', 'Praktik langsung'],
                ['Gratis', 'Tahap awal'],
              ].map(([v, l]) => (
                <div key={l}>
                  <dt className="text-2xl font-bold text-slate-900 dark:text-white">{v}</dt>
                  <dd className="text-xs text-slate-500 dark:text-slate-400">{l}</dd>
                </div>
              ))}
            </dl>
          </div>

          <div className="card overflow-hidden p-0">
            <div className="flex items-center gap-2 border-b border-slate-200 bg-slate-100 px-4 py-2.5 dark:border-slate-800 dark:bg-slate-800">
              <span className="h-3 w-3 rounded-full bg-red-400" />
              <span className="h-3 w-3 rounded-full bg-yellow-400" />
              <span className="h-3 w-3 rounded-full bg-green-400" />
              <span className="ml-2 font-mono text-xs text-slate-500 dark:text-slate-400">
                main.go
              </span>
            </div>
            <pre className="overflow-x-auto bg-slate-950 p-5 font-mono text-sm text-slate-100">
{`package main

import (
    "fmt"
    "strings"
)

func main() {
    kalimat := "belajar go itu seru"
    kata := strings.Fields(kalimat)

    for i, k := range kata {
        fmt.Printf("%d: %s\n", i+1, strings.ToUpper(k))
    }
}`}
            </pre>
          </div>
        </div>
      </section>

      <section className="mx-auto max-w-6xl px-4 py-14 sm:px-6">
        <h2 className="text-center text-2xl font-bold text-slate-900 dark:text-white sm:text-3xl">
          Belajar dengan cara yang benar-benar efektif
        </h2>
        <div className="mt-10 grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
          {features.map((f) => (
            <div key={f.title} className="card p-6">
              <span className="text-2xl" aria-hidden>
                {f.icon}
              </span>
              <h3 className="mt-3 font-semibold text-slate-900 dark:text-white">{f.title}</h3>
              <p className="mt-1 text-sm text-slate-600 dark:text-slate-400">{f.desc}</p>
            </div>
          ))}
        </div>
      </section>

      <section id="kurikulum" className="mx-auto max-w-6xl px-4 py-14 sm:px-6">
        <h2 className="text-2xl font-bold text-slate-900 dark:text-white sm:text-3xl">
          Peta belajar
        </h2>
        <p className="mt-2 text-slate-600 dark:text-slate-400">
          Mulai dari fondasi, lalu naik ke backend engineering dan level profesional.
        </p>
        <div className="mt-8 grid gap-5 sm:grid-cols-2">
          {(paths || []).map((p, i) => (
            <div key={p.slug} className="card p-6">
              <div className="flex items-center justify-between">
                <span className="chip bg-brand-50 text-brand-700 dark:bg-brand-900/30 dark:text-brand-300">
                  Path {i + 1}
                </span>
                <span className="chip bg-slate-100 text-slate-600 dark:bg-slate-800 dark:text-slate-300">
                  {p.level}
                </span>
              </div>
              <h3 className="mt-3 text-lg font-bold text-slate-900 dark:text-white">{p.title}</h3>
              <p className="mt-1 text-sm text-slate-600 dark:text-slate-400">{p.description}</p>
              <div className="mt-4">
                <ProgressBar value={p.lesson_count ? 100 : 0} />
                <p className="mt-1 text-xs text-slate-500 dark:text-slate-400">
                  {p.lesson_count} pelajaran
                </p>
              </div>
            </div>
          ))}
        </div>
      </section>

      <footer className="border-t border-slate-200 bg-white py-8 dark:border-slate-800 dark:bg-slate-950">
        <div className="mx-auto flex max-w-6xl flex-col items-center justify-between gap-3 px-4 text-sm text-slate-500 dark:text-slate-400 sm:flex-row sm:px-6">
          <p>© {new Date().getFullYear()} GoLearn. Dibuat untuk belajar Go.</p>
          <Link className="hover:text-brand-600" to={ctaTo}>
            Mulai belajar →
          </Link>
        </div>
      </footer>
    </div>
  )
}
