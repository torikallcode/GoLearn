import { request } from './client'

export const api = {
  register: (payload) => request('/auth/register', { method: 'POST', body: payload, auth: false }),
  login: (payload) => request('/auth/login', { method: 'POST', body: payload, auth: false }),
  logout: (refresh_token) =>
    request('/auth/logout', { method: 'POST', body: { refresh_token }, auth: false }),
  me: () => request('/auth/me'),
  deleteAccount: () => request('/auth/me', { method: 'DELETE' }),

  paths: () => request('/paths'),
  path: (slug) => request(`/paths/${slug}`),
  lesson: (slug) => request(`/lessons/${slug}`),
  exercises: (slug) => request(`/lessons/${slug}/exercises`),
  quizzes: (slug) => request(`/lessons/${slug}/quizzes`),
  search: (q) => request(`/search?q=${encodeURIComponent(q)}`),

  run: (code) => request('/playground/run', { method: 'POST', body: { code } }),
  submitExercise: (id, code) =>
    request(`/exercises/${id}/submit`, { method: 'POST', body: { code } }),
  answerQuiz: (id, payload) =>
    request(`/quizzes/${id}/answer`, { method: 'POST', body: payload }),
  completeLesson: (slug) => request(`/lessons/${slug}/complete`, { method: 'POST' }),

  dashboard: () => request('/me/dashboard'),
  progress: () => request('/me/progress'),
  stats: () => request('/me/stats'),
  reviewToday: () => request('/me/review/today'),
  reviewCard: (id, quality) =>
    request(`/me/review/${id}`, { method: 'POST', body: { quality } }),

  notes: (lessonId) =>
    request(`/me/notes${lessonId ? `?lesson_id=${lessonId}` : ''}`),
  saveNote: (lesson_id, content) =>
    request('/me/notes', { method: 'POST', body: { lesson_id, content } }),
  updateNote: (id, content) =>
    request(`/me/notes/${id}`, { method: 'PUT', body: { content } }),
  deleteNote: (id) => request(`/me/notes/${id}`, { method: 'DELETE' }),

  bookmarks: () => request('/me/bookmarks'),
  addBookmark: (lesson_id) => request('/me/bookmarks', { method: 'POST', body: { lesson_id } }),
  removeBookmark: (lesson_id) => request(`/me/bookmarks/${lesson_id}`, { method: 'DELETE' }),

  projects: () => request('/projects'),
  project: (slug) => request(`/projects/${slug}`),
  completeProjectStep: (id) =>
    request(`/projects/steps/${id}/complete`, { method: 'POST' }),
}
