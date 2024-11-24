// src/services/authService.js
import axios from 'axios';

// Base URL API backend
const API = axios.create({
  baseURL: 'http://localhost:8080', // Ganti dengan URL backend Anda jika berbeda
});

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
