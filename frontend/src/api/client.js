import { tokens } from './tokens'

export const API_BASE = import.meta.env.VITE_API_BASE || '/api/v1'

export class ApiError extends Error {
  constructor(status, code, message) {
    super(message)
    this.name = 'ApiError'
    this.status = status
    this.code = code
  }
}

let authFailureHandler = null
let refreshPromise = null

export function setAuthFailureHandler(fn) {
  authFailureHandler = fn
}

async function parse(res) {
  let body = null
  try {
    body = await res.json()
  } catch {
    body = null
  }
  if (!res.ok || (body && body.success === false)) {
    const err = body?.error || {}
    throw new ApiError(
      res.status,
      err.code || 'error',
      err.message || `Permintaan gagal (${res.status}).`,
    )
  }
  return body?.data ?? null
}

async function tryRefresh() {
  if (!tokens.refresh) return false
  if (refreshPromise) return refreshPromise
  refreshPromise = (async () => {
    try {
      const res = await fetch(`${API_BASE}/auth/refresh`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ refresh_token: tokens.refresh }),
      })
      if (!res.ok) return false
      const body = await res.json()
      if (!body?.success || !body.data?.tokens) return false
      tokens.set(body.data.tokens)
      return true
    } catch {
      return false
    } finally {
      refreshPromise = null
    }
  })()
  return refreshPromise
}

export async function request(path, { method = 'GET', body, auth = true, retry = true } = {}) {
  const headers = {}
  if (body !== undefined) headers['Content-Type'] = 'application/json'
  if (auth && tokens.access) headers.Authorization = `Bearer ${tokens.access}`

  const res = await fetch(`${API_BASE}${path}`, {
    method,
    headers,
    body: body !== undefined ? JSON.stringify(body) : undefined,
  })

  if (res.status === 401 && auth && retry && tokens.refresh) {
    const refreshed = await tryRefresh()
    if (refreshed) return request(path, { method, body, auth, retry: false })
    tokens.clear()
    authFailureHandler?.()
  }

  return parse(res)
}
