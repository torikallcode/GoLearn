import { create } from 'zustand'
import { api } from '../api'
import { tokens } from '../api/tokens'

export const useAuth = create((set, get) => ({
  user: null,
  stats: null,
  status: tokens.access ? 'loading' : 'guest',

  async boot() {
    if (!tokens.access) {
      set({ status: 'guest' })
      return
    }
    try {
      const data = await api.me()
      set({ user: data.user, stats: data.stats, status: 'authed' })
    } catch {
      tokens.clear()
      set({ user: null, stats: null, status: 'guest' })
    }
  },

  async login(email, password) {
    const data = await api.login({ email, password })
    tokens.set(data.tokens)
    set({ user: data.user, stats: data.stats, status: 'authed' })
    return data
  },

  async register(name, email, password) {
    const data = await api.register({ name, email, password })
    tokens.set(data.tokens)
    set({ user: data.user, stats: data.stats, status: 'authed' })
    return data
  },

  async logout() {
    const refresh = tokens.refresh
    try {
      if (refresh) await api.logout(refresh)
    } catch {
      /* ignore */
    }
    tokens.clear()
    set({ user: null, stats: null, status: 'guest' })
  },

  hardLogout() {
    tokens.clear()
    set({ user: null, stats: null, status: 'guest' })
  },

  patchStats(stats) {
    if (stats) set({ stats })
  },

  get isAuthed() {
    return get().status === 'authed'
  },
}))
