import { create } from 'zustand'

const KEY = 'golearn.theme'

function initial() {
  const saved = localStorage.getItem(KEY)
  if (saved === 'light' || saved === 'dark') return saved
  return window.matchMedia?.('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
}

function apply(theme) {
  const root = document.documentElement
  root.classList.toggle('dark', theme === 'dark')
  root.style.colorScheme = theme
}

export const useTheme = create((set) => ({
  theme: initial(),
  init() {
    apply(initial())
  },
  toggle() {
    set((s) => {
      const next = s.theme === 'dark' ? 'light' : 'dark'
      localStorage.setItem(KEY, next)
      apply(next)
      return { theme: next }
    })
  },
}))
