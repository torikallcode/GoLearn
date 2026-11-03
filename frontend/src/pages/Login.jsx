import { useState } from 'react'
import { Link, useLocation, useNavigate } from 'react-router-dom'
import { useAuth } from '../store/auth'

export default function Login() {
  const login = useAuth((s) => s.login)
  const navigate = useNavigate()
  const location = useLocation()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  async function onSubmit(e) {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      await login(email.trim(), password)
      navigate(location.state?.from || '/app', { replace: true })
    } catch (err) {
      setError(err.message || 'Gagal masuk.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-slate-50 px-4 dark:bg-slate-950">
      <div className="w-full max-w-md">
        <Link to="/" className="mb-8 flex items-center justify-center gap-2">
          <span className="grid h-10 w-10 place-items-center rounded-xl bg-brand-600 font-mono text-sm font-bold text-white">
            Go
          </span>
          <span className="text-xl font-bold text-slate-900 dark:text-white">GoLearn</span>
        </Link>
        <div className="card p-8">
          <h1 className="text-2xl font-bold text-slate-900 dark:text-white">Masuk</h1>
          <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
            Lanjutkan progres belajarmu.
          </p>
          <form className="mt-6 space-y-4" onSubmit={onSubmit}>
            <div>
              <label className="label" htmlFor="email">
                Email
              </label>
              <input
                id="email"
                type="email"
                className="input"
                autoComplete="email"
                required
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </div>
            <div>
              <label className="label" htmlFor="password">
                Password
              </label>
              <input
                id="password"
                type="password"
                className="input"
                autoComplete="current-password"
                required
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            {error ? (
              <p className="rounded-lg bg-red-50 px-3 py-2 text-sm text-red-600 dark:bg-red-900/20 dark:text-red-400">
                {error}
              </p>
            ) : null}
            <button className="btn-primary w-full" disabled={loading}>
              {loading ? 'Memproses…' : 'Masuk'}
            </button>
          </form>
          <p className="mt-5 text-center text-sm text-slate-500 dark:text-slate-400">
            Belum punya akun?{' '}
            <Link to="/register" className="font-semibold text-brand-600 hover:underline">
              Daftar
            </Link>
          </p>
        </div>
      </div>
    </div>
  )
}
