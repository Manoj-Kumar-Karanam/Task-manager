import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import api from '../api'
import { Link } from 'react-router-dom'

function Login() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const navigate = useNavigate()

  const handleLogin = async () => {
    try {
      const res = await api.post('http://localhost:8080/auth/login', { email, password })
      localStorage.setItem('token', res.data.token)
      navigate('/dashboard')
    } catch (err) {
      alert('Login failed!')
    }
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg--500">
      <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-sm">
        <h2 className="text-2xl font-bold mb-6 text-center">Login</h2>

        {/* Email Input */}
        <input
          className="w-full px-4 py-2 mb-4 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
          placeholder="Email"
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        {/* Password Input */}
        <input
          className="w-full px-4 py-2 mb-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400"
          placeholder="Password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        {/* Forgot Password */}
        <p className="text-sm text-blue-500 text-right mb-4 cursor-pointer hover:underline">
          Forgot password?
        </p>

        {/* Login Button */}
        <button
          onClick={handleLogin}
          className="w-full bg-blue-600 text-white py-2 rounded-md font-semibold hover:bg-blue-700 transition-colors"
        >
          Login
        </button>

        {/* Signup Link */}

          <p className="text-sm text-center mt-4">
            Don’t have an account?{' '}
            <Link 
              to="/register" 
              className="text-blue-500 cursor-pointer hover:underline"
            >
              Register
            </Link>
          </p>


        
      </div>
    </div>
  )
}

export default Login
