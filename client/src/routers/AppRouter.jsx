// src/routers/AppRouter.js
import { createBrowserRouter } from 'react-router-dom';
import Register from '../components/Register';
import Login from '../components/Login';
import { useAuth } from '../context/AuthContext'; // Impor useAuth

const router = createBrowserRouter([
  {
    path: '/Register',
    element: <Register />,
  },
  {
    path: '/Login',
    element: <Login />,
  },
]);

export default router;
