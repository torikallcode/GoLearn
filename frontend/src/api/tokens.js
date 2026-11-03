const ACCESS_KEY = 'golearn.access_token'
const REFRESH_KEY = 'golearn.refresh_token'

let access = localStorage.getItem(ACCESS_KEY)
let refresh = localStorage.getItem(REFRESH_KEY)

export const tokens = {
  get access() {
    return access
  },
  get refresh() {
    return refresh
  },
  set(next) {
    access = next?.access_token ?? null
    refresh = next?.refresh_token ?? null
    if (access) localStorage.setItem(ACCESS_KEY, access)
    else localStorage.removeItem(ACCESS_KEY)
    if (refresh) localStorage.setItem(REFRESH_KEY, refresh)
    else localStorage.removeItem(REFRESH_KEY)
  },
  clear() {
    access = null
    refresh = null
    localStorage.removeItem(ACCESS_KEY)
    localStorage.removeItem(REFRESH_KEY)
  },
}
