// src/services/authService.js
import axios from 'axios';

// Instance untuk endpoint tanpa prefix `/api`
const API = axios.create({
  baseURL: 'http://localhost:8080/', // Untuk login dan register
});

// Instance untuk endpoint dengan prefix `/api`
const APIWithPrefix = axios.create({
  baseURL: 'http://localhost:8080/api', // Untuk endpoint dengan prefix `/api`
});

// Tambahkan interceptor untuk menyertakan token
APIWithPrefix.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('access_token'); // Ambil token dari 'access_token'
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
      console.log('Authorization header set:', config.headers['Authorization']); // Debugging
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Fungsi untuk register
export const register = (username, password, email, phonenumber, role) => {
  return API.post('/auth/register', {
    username,
    password,
    email,
    phone_number: phonenumber,
    role,
  });
};

// Fungsi untuk login
export const login = async (email, password) => {
  try {
    const response = await API.post('/auth/login', { email, password });
    return response.data; // Mengembalikan data dari respons
  } catch (error) {
    console.error('Backend login error:', error.response?.data || error.message);
    throw error;
  }
};

// Default export untuk API dengan prefix `/api`
export { APIWithPrefix };
