// src/routers/AppRouter.js
import { createBrowserRouter, Navigate } from 'react-router-dom'; // Tambahkan Navigate di sini
import Register from '../components/Register';
import Login from '../components/Login';
import Dashboard from '../pages/Dashboard';
import Tespneumina from '@/Pages/check/tespneumina';
import PrivateRoute from '../routers/PrivateRoute';

const router = createBrowserRouter([
  {
    path: '/Register',
    element: <Register />,
  },
  {
    path: '/Login',
    element: <Login />,
  },
  {
    path: '/teskesehatan',
    element: <Tespneumina />,
  },
  {
    path: '/dashboard',
    element: (
      <PrivateRoute>
        <Dashboard />
      </PrivateRoute>
    ),
  },
  {
    path: '*',
    element: <Navigate to="/Login" replace />, // Redirect semua rute tak dikenal ke Login
  },
]);

export default router;
