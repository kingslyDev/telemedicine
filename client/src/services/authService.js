import axios from 'axios';

// Base URL API backend
const API = axios.create({
  baseURL: 'http://localhost:8080', // Ganti dengan URL backend Anda
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
export const login = async (username, password) => {
  return API.post('/auth/login', { username, password });
};
