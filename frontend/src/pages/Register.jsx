import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../store/auth'

export default function Register() {
  const register = useAuth((s) => s.register)
  const navigate = useNavigate()
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  async function onSubmit(e) {
    e.preventDefault()
    setError('')
    if (password.length < 8) {
      setError('Password minimal 8 karakter.')
      return
    }
    setLoading(true)
    try {
      await register(name.trim(), email.trim(), password)
      navigate('/app', { replace: true })
    } catch (err) {
      setError(err.message || 'Gagal mendaftar.')
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
          <h1 className="text-2xl font-bold text-slate-900 dark:text-white">Buat akun</h1>
          <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
            Gratis. Mulai dari pelajaran pertama.
          </p>
          <form className="mt-6 space-y-4" onSubmit={onSubmit}>
            <div>
              <label className="label" htmlFor="name">
                Nama
              </label>
              <input
                id="name"
                className="input"
                autoComplete="name"
                required
                minLength={2}
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </div>
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
                autoComplete="new-password"
                required
                minLength={8}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
              <p className="mt-1 text-xs text-slate-400">Minimal 8 karakter.</p>
            </div>
            {error ? (
              <p className="rounded-lg bg-red-50 px-3 py-2 text-sm text-red-600 dark:bg-red-900/20 dark:text-red-400">
                {error}
              </p>
            ) : null}
            <button className="btn-primary w-full" disabled={loading}>
              {loading ? 'Memproses…' : 'Daftar'}
            </button>
          </form>
          <p className="mt-5 text-center text-sm text-slate-500 dark:text-slate-400">
            Sudah punya akun?{' '}
            <Link to="/login" className="font-semibold text-brand-600 hover:underline">
              Masuk
            </Link>
          </p>
        </div>
      </div>
    </div>
  )
}
