import { useEffect } from 'react'
import { Routes, Route, Navigate, useLocation } from 'react-router-dom'
import { useAuth } from './store/auth'
import { useTheme } from './store/theme'
import { setAuthFailureHandler } from './api/client'

import Layout from './components/Layout'
import Spinner from './components/Spinner'

import Landing from './pages/Landing'
import Login from './pages/Login'
import Register from './pages/Register'
import Dashboard from './pages/Dashboard'
import Roadmap from './pages/Roadmap'
import PathDetail from './pages/PathDetail'
import Lesson from './pages/Lesson'
import Review from './pages/Review'
import Projects from './pages/Projects'
import ProjectDetail from './pages/ProjectDetail'
import Cheatsheet from './pages/Cheatsheet'
import Search from './pages/Search'
import Profile from './pages/Profile'
import NotFound from './pages/NotFound'

function Protected({ children }) {
  const status = useAuth((s) => s.status)
  const location = useLocation()
  if (status === 'loading') return <Spinner full />
  if (status !== 'authed') {
    return <Navigate to="/login" replace state={{ from: location.pathname }} />
  }
  return children
}

function GuestOnly({ children }) {
  const status = useAuth((s) => s.status)
  if (status === 'loading') return <Spinner full />
  if (status === 'authed') return <Navigate to="/app" replace />
  return children
}

export default function App() {
  const boot = useAuth((s) => s.boot)
  const hardLogout = useAuth((s) => s.hardLogout)
  const initTheme = useTheme((s) => s.init)

  useEffect(() => {
    initTheme()
    setAuthFailureHandler(hardLogout)
    boot()
  }, [boot, hardLogout, initTheme])

  return (
    <Routes>
      <Route path="/" element={<Landing />} />
      <Route path="/login" element={<GuestOnly><Login /></GuestOnly>} />
      <Route path="/register" element={<GuestOnly><Register /></GuestOnly>} />

      <Route
        path="/app"
        element={
          <Protected>
            <Layout />
          </Protected>
        }
      >
        <Route index element={<Dashboard />} />
        <Route path="roadmap" element={<Roadmap />} />
        <Route path="roadmap/:slug" element={<PathDetail />} />
        <Route path="lessons/:slug" element={<Lesson />} />
        <Route path="review" element={<Review />} />
        <Route path="projects" element={<Projects />} />
        <Route path="projects/:slug" element={<ProjectDetail />} />
        <Route path="cheatsheet" element={<Cheatsheet />} />
        <Route path="search" element={<Search />} />
        <Route path="profile" element={<Profile />} />
      </Route>

      <Route path="*" element={<NotFound />} />
    </Routes>
  )
}
