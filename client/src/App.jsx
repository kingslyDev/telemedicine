// src/App.js
import React from 'react';
import { AuthProvider } from './context/AuthContext';
import { RouterProvider } from 'react-router-dom'; // Pastikan ini diimpor
import router from './routers/AppRouter'; // Pastikan rutenya diimpor

const App = () => {
  return (
    <AuthProvider>
      <RouterProvider router={router} />
    </AuthProvider>
  );
};

export default App;
