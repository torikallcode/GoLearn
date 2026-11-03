import { Link } from 'react-router-dom'

export default function NotFound() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center gap-4 bg-slate-50 px-4 text-center dark:bg-slate-950">
      <p className="font-mono text-6xl font-bold text-brand-600">404</p>
      <h1 className="text-xl font-bold text-slate-900 dark:text-white">Halaman tidak ditemukan</h1>
      <p className="text-sm text-slate-500 dark:text-slate-400">
        Sepertinya kamu tersesat. Kembali ke tempat yang aman.
      </p>
      <Link className="btn-primary" to="/">
        Ke halaman utama
      </Link>
    </div>
  )
}
