import { useState } from 'react'
import { NavLink, Outlet, useNavigate } from 'react-router-dom'
import { useAuth } from '../store/auth'
import { useTheme } from '../store/theme'

const nav = [
  { to: '/app', label: 'Dashboard', icon: '🏠', end: true },
  { to: '/app/roadmap', label: 'Roadmap', icon: '🗺️' },
  { to: '/app/review', label: 'Review', icon: '🔁' },
  { to: '/app/projects', label: 'Proyek', icon: '🛠️' },
  { to: '/app/cheatsheet', label: 'Cheatsheet', icon: '📌' },
  { to: '/app/search', label: 'Cari', icon: '🔍' },
]

function NavItems({ onClick }) {
  return (
    <nav className="flex flex-col gap-1">
      {nav.map((item) => (
        <NavLink
          key={item.to}
          to={item.to}
          end={item.end}
          onClick={onClick}
          className={({ isActive }) =>
            `flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition ${
              isActive
                ? 'bg-brand-50 text-brand-700 dark:bg-brand-900/30 dark:text-brand-300'
                : 'text-slate-600 hover:bg-slate-100 dark:text-slate-300 dark:hover:bg-slate-800'
            }`
          }
        >
          <span aria-hidden>{item.icon}</span>
          {item.label}
        </NavLink>
      ))}
    </nav>
  )
}

function Brand() {
  return (
    <NavLink to="/app" className="flex items-center gap-2">
      <span className="grid h-9 w-9 place-items-center rounded-xl bg-brand-600 font-mono text-sm font-bold text-white">
        Go
      </span>
      <span className="text-lg font-bold tracking-tight text-slate-900 dark:text-white">
        GoLearn
      </span>
    </NavLink>
  )
}

export default function Layout() {
  const user = useAuth((s) => s.user)
  const logout = useAuth((s) => s.logout)
  const theme = useTheme((s) => s.theme)
  const toggleTheme = useTheme((s) => s.toggle)
  const [mobileOpen, setMobileOpen] = useState(false)
  const navigate = useNavigate()

  async function handleLogout() {
    await logout()
    navigate('/')
  }

  return (
    <div className="min-h-screen lg:grid lg:grid-cols-[260px_1fr]">
      {/* Sidebar desktop */}
      <aside className="hidden border-r border-slate-200 bg-white px-4 py-6 dark:border-slate-800 dark:bg-slate-950 lg:flex lg:flex-col">
        <Brand />
        <div className="mt-8 flex-1">
          <NavItems />
        </div>
        <div className="mt-4 space-y-2 border-t border-slate-200 pt-4 dark:border-slate-800">
          <NavLink
            to="/app/profile"
            className={({ isActive }) =>
              `flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium ${
                isActive
                  ? 'bg-brand-50 text-brand-700 dark:bg-brand-900/30 dark:text-brand-300'
                  : 'text-slate-600 hover:bg-slate-100 dark:text-slate-300 dark:hover:bg-slate-800'
              }`
            }
          >
            <span aria-hidden>👤</span>
            <span className="truncate">{user?.name || 'Profil'}</span>
          </NavLink>
          <div className="flex items-center gap-2">
            <button className="btn-ghost flex-1 text-xs" onClick={toggleTheme}>
              {theme === 'dark' ? '☀️ Terang' : '🌙 Gelap'}
            </button>
            <button className="btn-ghost flex-1 text-xs" onClick={handleLogout}>
              Keluar
            </button>
          </div>
        </div>
      </aside>

      <div className="flex min-h-screen flex-col">
        {/* Topbar mobile */}
        <header className="flex items-center justify-between border-b border-slate-200 bg-white px-4 py-3 dark:border-slate-800 dark:bg-slate-950 lg:hidden">
          <Brand />
          <div className="flex items-center gap-1">
            <button className="btn-ghost text-sm" onClick={toggleTheme} aria-label="Ganti tema">
              {theme === 'dark' ? '☀️' : '🌙'}
            </button>
            <button
              className="btn-ghost text-sm"
              onClick={() => setMobileOpen((v) => !v)}
              aria-label="Buka menu"
            >
              ☰
            </button>
          </div>
        </header>

        {mobileOpen ? (
          <div className="border-b border-slate-200 bg-white px-4 py-4 dark:border-slate-800 dark:bg-slate-950 lg:hidden">
            <NavItems onClick={() => setMobileOpen(false)} />
            <div className="mt-3 flex gap-2 border-t border-slate-200 pt-3 dark:border-slate-800">
              <button className="btn-secondary flex-1 text-xs" onClick={() => { setMobileOpen(false); navigate('/app/profile') }}>
                Profil
              </button>
              <button className="btn-secondary flex-1 text-xs" onClick={handleLogout}>
                Keluar
              </button>
            </div>
          </div>
        ) : null}

        <main className="flex-1 px-4 py-6 sm:px-6 lg:px-8">
          <div className="mx-auto w-full max-w-6xl">
            <Outlet />
          </div>
        </main>
      </div>
    </div>
  )
}
