// src/routers/PrivateRoute.jsx
import React from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

const PrivateRoute = ({ children }) => {
  const { isAuthenticated, isLoading } = useAuth();
  const location = useLocation();

  if (isLoading) {
    return <div>Loading...</div>; // Indikator loading sementara
  }

  if (!isAuthenticated) {
    return <Navigate to="/Login" replace state={{ from: location }} />;
  }

  return children;
};

export default PrivateRoute; // Pastikan ini adalah ekspor default
