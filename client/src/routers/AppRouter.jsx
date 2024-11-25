// src/routers/AppRouter.js
import { createBrowserRouter, Navigate } from 'react-router-dom'; // Tambahkan Navigate di sini
import Register from '../components/Register';
import Login from '../components/Login';
import Dashboard from '../pages/Dashboard';
import Tespneumina from '@/Pages/check/tespneumina';
import PrivateRoute from '../routers/PrivateRoute';
import EditProfile from '@/Pages/Editprofile';

const router = createBrowserRouter([
  {
    path: '/Register',
    element: <Register />,
  },
  {
    path: '/Edit',
    element: <EditProfile />,
  },
  {
    path: '/Login',
    element: <Login />,
  },
  {
    path: '/teskesehatan',
    element: (
      <PrivateRoute>
        <Tespneumina />
      </PrivateRoute>
    ),
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
