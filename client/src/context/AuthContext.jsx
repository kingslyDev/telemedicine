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

  const login = async (email, password) => {
    const response = await loginService(email, password);
    if (response && response.token && response.user) {
      // Simpan token ke localStorage
      localStorage.setItem('access_token', response.token);

      // Simpan user ke localStorage, termasuk role
      localStorage.setItem('user', JSON.stringify(response.user));
      setUser(response.user); // Simpan user ke state

      setIsAuthenticated(true);
    } else {
      throw new Error('Invalid login response');
    }
    return response;
  };

  const logout = () => {
    localStorage.removeItem('access_token');
    localStorage.removeItem('user');
    setIsAuthenticated(false);
    setUser(null);
  };

  return <AuthContext.Provider value={{ isAuthenticated, user, login, logout }}>{children}</AuthContext.Provider>;
};

export const useAuth = () => useContext(AuthContext);
