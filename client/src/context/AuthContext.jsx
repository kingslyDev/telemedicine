import React, { createContext, useContext, useState } from 'react';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(!!localStorage.getItem('access_token'));
  const [user, setUser] = useState(JSON.parse(localStorage.getItem('user'))); // Menyimpan informasi pengguna

  const login = (token, userData) => {
    localStorage.setItem('access_token', token);
    localStorage.setItem('user', JSON.stringify(userData)); // Simpan informasi pengguna di localStorage
    setIsAuthenticated(true);
    setUser(userData); // Perbarui state user
  };

  const logout = () => {
    localStorage.removeItem('access_token');
    localStorage.removeItem('user'); // Hapus informasi pengguna saat logout
    setIsAuthenticated(false);
    setUser(null); // Reset state user
  };

  return <AuthContext.Provider value={{ isAuthenticated, user, login, logout }}>{children}</AuthContext.Provider>;
};

export const useAuth = () => useContext(AuthContext);
