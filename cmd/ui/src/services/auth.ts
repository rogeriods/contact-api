import api from '../api/client'

export const register = async (username: string, password: string) => {
  const res = await api.post('/register', { username, password })
  return res.data
}

export const login = async (username: string, password: string) => {
  const res = await api.post('/login', { username, password })

  const token = res.data.token
  localStorage.setItem('token', token)

  return res.data
}

export const logout = () => {
  localStorage.removeItem('token')
}

export const isAuthenticated = () => {
  return !!localStorage.getItem('token')
}
