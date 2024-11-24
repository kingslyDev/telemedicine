// src/context/AuthContext.jsx
import React, { createContext, useContext, useState } from 'react';
import { login as loginService } from '../services/authService';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(!!localStorage.getItem('access_token'));

  const [user, setUser] = useState(() => {
    const storedUser = localStorage.getItem('user');
    try {
      return storedUser && storedUser !== 'undefined' ? JSON.parse(storedUser) : null;
    } catch (error) {
      console.error('Error parsing user from localStorage:', error);
      return null;
    }
  });

  const [isLoading, setIsLoading] = useState(false); // Opsional: untuk status loading

  const login = async (email, password) => {
    setIsLoading(true);
    try {
      const response = await loginService(email, password);
      console.log('Login response:', response); // Debug log
      if (response && response.token) {
        localStorage.setItem('access_token', response.token);
        if (response.user) {
          localStorage.setItem('user', JSON.stringify(response.user));
          setUser(response.user);
        }
        setIsAuthenticated(true);
      } else {
        throw new Error('Invalid login response');
      }
      return response;
    } catch (error) {
      console.error('Login failed:', error.response?.data || error.message);
      throw error;
    } finally {
      setIsLoading(false);
    }
  };

  const logout = () => {
    localStorage.removeItem('access_token');
    localStorage.removeItem('user'); // Hapus informasi pengguna saat logout
    setIsAuthenticated(false);
    setUser(null); // Reset state user
  };

  return <AuthContext.Provider value={{ isAuthenticated, user, login, logout, isLoading }}>{children}</AuthContext.Provider>;
};

export const useAuth = () => useContext(AuthContext);
