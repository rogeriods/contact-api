import { useState } from 'react'
import { register } from '../services/auth'
import { useNavigate } from 'react-router-dom'

export default function Register() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const navigate = useNavigate()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    try {
      await register(username, password)
      navigate('/')
    } catch {
      alert('Register failed')
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h2>Register</h2>

      <input placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />

      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />

      <button type="submit">Register</button>
    </form>
  )
}
