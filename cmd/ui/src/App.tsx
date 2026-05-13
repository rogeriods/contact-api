import { BrowserRouter, Routes, Route, useNavigate, useLocation } from 'react-router-dom'
import Login from './pages/Login'
import Register from './pages/Register'
import Contacts from './pages/Contacts'
import { logout } from './services/auth'
import ProtectedRoute from './components/ProtectedRoute'

function Header() {
  const navigate = useNavigate()
  const location = useLocation()

  const handleLogout = () => {
    logout()
    navigate('/')
  }

  // Only show the header/logout button when not on the login or register pages
  if (location.pathname === '/' || location.pathname === '/register') {
    return null
  }

  return (
    <header
      style={{
        display: 'flex',
        justifyContent: 'flex-end',
        padding: '1rem',
        backgroundColor: '#f8f9fa',
        borderBottom: '1px solid #e9ecef',
        marginBottom: '1rem',
      }}
    >
      <button
        onClick={handleLogout}
        style={{
          padding: '0.5rem 1rem',
          backgroundColor: '#007bff',
          color: 'white',
          border: 'none',
          borderRadius: '4px',
          cursor: 'pointer',
          fontWeight: 'bold',
          transition: 'background-color 0.2s',
        }}
      >
        Logout
      </button>
    </header>
  )
}

function App() {
  return (
    <BrowserRouter>
      <Header />
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route
          path="/contacts"
          element={
            <ProtectedRoute>
              <Contacts />
            </ProtectedRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  )
}

export default App
