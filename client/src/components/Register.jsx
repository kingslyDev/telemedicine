import React, { useState } from 'react';
import { register } from '@/services/authService';
import { Button } from '@/components/ui/button';

const Register = () => {
  // State untuk form
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const [phonenumber, setPhonenumber] = useState('');
  const [role, setRole] = useState('patient'); // Default role
  const [message, setMessage] = useState('');

  const [agreed, setAgreed] = useState(false);

  // Fungsi submit form
  const handleSubmit = async (e) => {
    e.preventDefault(); // Mencegah reload halaman

    // Validasi sederhana
    if (!username || !password || !email) {
      setMessage('Please fill in all fields.');
      return;
    }

    try {
      const response = await register(username, password, email, phonenumber, role); // Panggil service register
      setMessage('Registration successful! 🎉');
    } catch (error) {
      console.error('Registration failed:', error.response?.data || error.message);
      setMessage('Registration failed. Please try again.');
    }

    // Reset form (opsional)
    setUsername('');
    setPassword('');
    setEmail('');
    setPhonenumber('');
    setRole('patient');
  };

  return (
    <div className="relative min-h-screen flex justify-center items-center">
      <div className="absolute inset-0 flex">
        <div className="w-1/2 bg-[#54b8f5]"></div>
        <div className="w-1/2 bg-white"></div>
      </div>
      {/* IMAGE */}
      <img src="https://res.cloudinary.com/dwgwb5vro/image/upload/v1732437483/hheg9dpaoaaf4j7o2gdl.png" alt="Logo" className="absolute top-[5px] left-[106px] h-auto w-auto " />
      <img src="https://res.cloudinary.com/dwgwb5vro/image/upload/v1732436950/piukbpeqz4mbqgqouicq.png" alt="Icon 2" style={{ position: 'absolute', bottom: '10px', right: '60px', width: '450px', height: '350px' }} />
      {/* Card Register */}
      <div className="relative bg-white rounded-[40px] scale-[0.85] shadow-lg p-8 w-full max-w-md z-10">
        <h2 className="text-3xl font-semibold text-center mb-4">Get Started Now</h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label htmlFor="name" className="block text-sm font-medium text-gray-700">
              Name
            </label>
            <input
              type="text"
              name="name"
              id="name"
              className="mt-1 block w-full px-3 py-2 border border-customGreen rounded-md shadow-sm focus:ring-orange-500 focus:border-orange-500 sm:text-sm"
              placeholder="Enter your name"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div>
            <label htmlFor="email" className="block text-sm font-medium text-gray-700">
              Email Address
            </label>
            <input
              name="email"
              type="email"
              id="email"
              className="mt-1 block w-full px-3 py-2 border border-customGreen rounded-md shadow-sm focus:ring-orange-500 focus:border-orange-500 sm:text-sm"
              placeholder="Enter your email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div>
            <label htmlFor="email" className="block text-sm font-medium text-gray-700">
              Phone Number
            </label>
            <input
              name="phonenumber"
              type="text"
              id="phonenumber"
              className="mt-1 block w-full px-3 py-2 border border-customGreen rounded-md shadow-sm focus:ring-orange-500 focus:border-orange-500 sm:text-sm"
              placeholder="Enter your phonenumber"
              value={phonenumber}
              onChange={(e) => setPhonenumber(e.target.value)}
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
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">
              Password
            </label>
            <select value={role} onChange={(e) => setRole(e.target.value)} className="mt-1 block w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring focus:ring-blue-300">
              <option value="patient">Patient</option>
            </select>
          </div>
          <div className="flex items-center">
            <input type="checkbox" id="terms" className="h-4 w-4 text-blue-600 focus:ring-blue-500 rounded" checked={agreed} onChange={(e) => setAgreed(e.target.checked)} required />
            <label htmlFor="terms" className="ml-2 block text-sm text-gray-900">
              I agree to the terms & policy
            </label>
          </div>
          <div>
            <Button
              className="w-full mt-10 flex justify-center py-3 px-4 rounded-md shadow-sm text-sm font-medium text-white bg-[#60addc] focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              variant="primary"
              size="large"
              type="submit"
            >
              Register
            </Button>
          </div>
        </form>
        <div className="mt-6 text-center">
          <a href="/login" className="font-medium text-black hover:text-orange-500">
            have an account? <span className="text-custom-blue">Sign In</span>
          </a>
        </div>
      </div>
    </div>
  );
};

export default Register;
