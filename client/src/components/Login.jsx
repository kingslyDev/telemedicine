// src/components/Login.jsx
import React, { useState } from 'react';
import { useAuth } from '../context/AuthContext'; // Pastikan path benar
import { Button } from '../components/ui/button';
import { useNavigate, useLocation } from 'react-router-dom';

const Login = () => {
  // State untuk form
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const { login } = useAuth(); // Mengambil login dari context
  const navigate = useNavigate();
  const location = useLocation();

  // Redirect setelah login
  const from = location.state?.from?.pathname || '/dashboard';

  // Fungsi submit form
  const handleSubmit = async (e) => {
    e.preventDefault(); // Prevent page reload

    // Simple validation
    if (!email || !password) {
      setMessage('Please fill in all fields.');
      return;
    }

    try {
      const response = await login(email, password); // Call login from context
      setMessage(`Login successful! 🎉 Role: ${response.user.role}`); // Debug role
      // Redirect ke halaman dashboard atau halaman yang diminta sebelumnya
      navigate(from, { replace: true });
    } catch (error) {
      console.error('Login failed:', error.response?.data || error.message);
      setMessage(error.response?.data?.error || 'Login failed. Please check your credentials.');
    }
  };

  return (
    <div className="relative min-h-screen flex justify-center items-center">
      <div className="absolute inset-0 flex">
        <div className="w-1/2 bg-[#54b8f5]"></div>
        <div className="w-1/2 bg-white"></div>
      </div>
      {/* IMAGE */}
      <img src="https://res.cloudinary.com/dwgwb5vro/image/upload/v1732437483/hheg9dpaoaaf4j7o2gdl.png" alt="Logo" className="absolute top-[5px] left-[106px] h-auto w-auto" />
      <img
        src="https://res.cloudinary.com/dwgwb5vro/image/upload/v1732436950/piukbpeqz4mbqgqouicq.png"
        alt="Icon 2"
        style={{
          position: 'absolute',
          bottom: '10px',
          right: '60px',
          width: '450px',
          height: '350px',
        }}
      />
      {/* Card Login */}
      <div className="relative bg-white rounded-[40px] scale-[0.85] shadow-lg p-8 w-full max-w-md z-10">
        <h2 className="text-3xl font-semibold text-center mb-4">Welcome Back</h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label htmlFor="email" className="block text-sm font-medium text-gray-700">
              Email Address
            </label>
            <input
              name="email"
              type="email"
              className="mt-1 block w-full px-3 py-2 border border-customGreen rounded-md shadow-sm focus:ring-orange-500 focus:border-orange-500 sm:text-sm"
              placeholder="Enter your email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div>
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">
              Password
            </label>
            <input
              name="password"
              type="password"
              id="password"
              className="mt-1 block w-full px-3 py-2 border border-customGreen rounded-md shadow-sm focus:ring-orange-500 focus:border-orange-500 sm:text-sm"
              placeholder="Enter your password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <div>
            <Button
              className="w-full mt-10 flex justify-center py-3 px-4 rounded-md shadow-sm text-sm font-medium text-white bg-[#60addc] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              variant="primary"
              size="large"
              type="submit"
            >
              Login
            </Button>
          </div>
        </form>
        <div className="mt-6 text-center">
          <a href="/Register" className="font-medium text-black hover:text-orange-500">
            Don’t have an account? <span className="text-custom-blue">Sign Up</span>
          </a>
        </div>
        {message && <p className="mt-4 text-center text-red-500">{message}</p>}
      </div>
    </div>
  );
};

export default Login;
