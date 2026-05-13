import { useState } from 'react'
import { login } from '../services/auth'
import { useNavigate } from 'react-router-dom'

export default function Login() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const navigate = useNavigate()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    try {
      await login(username, password)
      navigate('/contacts')
    } catch (err) {
      alert('Login failed')
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h2>Login</h2>

      <input placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />

      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />

      <button type="submit">Login</button>
    </form>
  )
}
