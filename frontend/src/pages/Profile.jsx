import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { api } from '../api'
import { useAuth } from '../store/auth'
import { useTheme } from '../store/theme'
import Spinner from '../components/Spinner'
import ErrorState from '../components/ErrorState'

export default function Profile() {
  const user = useAuth((s) => s.user)
  const logout = useAuth((s) => s.logout)
  const theme = useTheme((s) => s.theme)
  const toggleTheme = useTheme((s) => s.toggle)
  const navigate = useNavigate()
  const queryClient = useQueryClient()
  const [confirmDelete, setConfirmDelete] = useState(false)
  const [busy, setBusy] = useState(false)

  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['stats'],
    queryFn: api.stats,
  })

  async function handleDelete() {
    setBusy(true)
    try {
      await api.deleteAccount()
      await logout()
      navigate('/')
    } finally {
      setBusy(false)
    }
  }

  if (isLoading) return <Spinner full />
  if (isError) return <ErrorState error={error} onRetry={refetch} />

  const stats = data.stats || {}
  const badges = data.badges || []

  return (
    <div className="mx-auto max-w-3xl space-y-6">
      <h1 className="text-2xl font-bold text-slate-900 dark:text-white">Profil & Pengaturan</h1>

      <div className="card p-6">
        <div className="flex items-center gap-4">
          <span className="grid h-14 w-14 place-items-center rounded-full bg-brand-600 text-xl font-bold text-white">
            {(user?.name || '?').charAt(0).toUpperCase()}
          </span>
          <div>
            <p className="text-lg font-bold text-slate-900 dark:text-white">{user?.name}</p>
            <p className="text-sm text-slate-500 dark:text-slate-400">{user?.email}</p>
          </div>
        </div>
      </div>

      <div className="grid grid-cols-2 gap-4 sm:grid-cols-4">
        {[
          ['XP', stats.xp ?? 0],
          ['Level', stats.level ?? 1],
          ['Streak', `${stats.current_streak ?? 0} 🔥`],
          ['Pelajaran', data.lessons_done ?? 0],
        ].map(([label, value]) => (
          <div key={label} className="card p-4 text-center">
            <p className="text-xs uppercase tracking-wide text-slate-500 dark:text-slate-400">{label}</p>
            <p className="mt-1 text-2xl font-bold text-slate-900 dark:text-white">{value}</p>
          </div>
        ))}
      </div>

      <div className="card p-6">
        <h2 className="font-semibold text-slate-900 dark:text-white">Badge</h2>
        {badges.length === 0 ? (
          <p className="mt-2 text-sm text-slate-500 dark:text-slate-400">
            Belum ada badge. Selesaikan pelajaran pertama untuk mendapatkannya.
          </p>
        ) : (
          <div className="mt-3 flex flex-wrap gap-2">
            {badges.map((b, i) => (
              <span
                key={i}
                className="chip bg-amber-50 text-amber-700 dark:bg-amber-900/20 dark:text-amber-300"
                title={b.description}
              >
                🏅 {b.name}
              </span>
            ))}
          </div>
        )}
      </div>

      <div className="card p-6">
        <h2 className="font-semibold text-slate-900 dark:text-white">Preferensi</h2>
        <div className="mt-3 flex items-center justify-between">
          <span className="text-sm text-slate-600 dark:text-slate-300">Mode tampilan</span>
          <button className="btn-secondary text-xs" onClick={toggleTheme}>
            {theme === 'dark' ? '☀️ Terang' : '🌙 Gelap'}
          </button>
        </div>
      </div>

      <div className="card border-red-200 p-6 dark:border-red-900/40">
        <h2 className="font-semibold text-red-600 dark:text-red-400">Zona berbahaya</h2>
        <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">
          Menghapus akun akan menghapus seluruh progres, catatan, dan data kamu secara permanen.
        </p>
        {confirmDelete ? (
          <div className="mt-3 flex items-center gap-2">
            <button className="btn bg-red-600 text-white hover:bg-red-700" onClick={handleDelete} disabled={busy}>
              {busy ? 'Menghapus…' : 'Ya, hapus akun saya'}
            </button>
            <button className="btn-secondary" onClick={() => setConfirmDelete(false)}>
              Batal
            </button>
          </div>
        ) : (
          <button className="btn-secondary mt-3 text-red-600" onClick={() => setConfirmDelete(true)}>
            Hapus akun
          </button>
        )}
      </div>

      <button
        className="btn-secondary w-full"
        onClick={async () => {
          await logout()
          queryClient.clear()
          navigate('/')
        }}
      >
        Keluar dari akun
      </button>
    </div>
  )
}
