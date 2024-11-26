import { createBrowserRouter, Navigate } from 'react-router-dom';
import Register from '../components/Register';
import Login from '../components/Login';
import Dashboard from '../pages/Dashboard';
import Tespneumina from '@/Pages/check/tespneumina';
import PrivateRoute from '../routers/PrivateRoute';
import EditProfile from '@/Pages/Editprofile';
import PatientsPage from '@/Pages/PatientsPage';
import PatientDetails from '@/Pages/PatientDetails';

const router = createBrowserRouter([
  // Public Routes
  {
    path: '/Register',
    element: <Register />,
  },
  {
    path: '/Login',
    element: <Login />,
  },

  // Private Routes
  {
    path: '/Edit',
    element: (
      <PrivateRoute>
        <EditProfile />
      </PrivateRoute>
    ),
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
    path: '/data',
    element: (
      <PrivateRoute>
        <PatientsPage />
      </PrivateRoute>
    ),
  },
  {
    path: '/data/:id',
    element: (
      <PrivateRoute>
        <PatientDetails />
      </PrivateRoute>
    ),
  },

  {
    path: '*',
    element: <Navigate to="/Login" replace />,
  },
]);

export default router;