import React, { useState } from 'react';
import { Button } from './ui/button';
import { Link } from 'react-router-dom';
import { useAuth } from '@/context/AuthContext';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { AlertDialog, AlertDialogTrigger, AlertDialogContent, AlertDialogTitle, AlertDialogDescription, AlertDialogAction, AlertDialogCancel } from '@/components/ui/alert-dialog';

const Navbar = () => {
  const { user, logout } = useAuth();
  const [menuOpen, setMenuOpen] = useState(false);

  const handleLogout = () => {
    logout();
    setMenuOpen(false); // Close dropdown if open
  };

  const navLinks = [
    { name: 'Home', path: '#home' },
    { name: 'About Us', path: '#about' },
    { name: 'Services', path: '#services' },
  ];

  return (
    <nav className="fixed top-0 left-0 w-full z-50 bg-white shadow-md">
      <div className="container mx-auto flex items-center justify-between py-4 px-6 md:px-12">
        {/* Logo */}
        <Link to="/" className="text-lg md:text-2xl font-bold text-gray-800">
          HaloGhal
        </Link>

        {/* Desktop Navigation */}
        <div className="hidden md:flex items-center gap-8">
          {navLinks.map((link) => (
            <a key={link.name} href={link.path} className="text-sm md:text-base font-medium text-gray-800 hover:text-green-700 transition duration-200">
              {link.name}
            </a>
          ))}

          {/* Authenticated User Menu */}
          {user ? (
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <button className="text-sm md:text-base font-medium text-gray-800 hover:text-green-700 transition duration-200">{user.username}</button>
              </DropdownMenuTrigger>
              <DropdownMenuContent className="bg-white rounded shadow-md">
                <DropdownMenuLabel className="text-gray-600 font-semibold">My Account</DropdownMenuLabel>
                <DropdownMenuSeparator />
                <Link to="/profile">
                  <DropdownMenuItem>Profile</DropdownMenuItem>
                </Link>
                <Link to="/result">
                  <DropdownMenuItem>Lihat Hasil</DropdownMenuItem>
                </Link>
                <DropdownMenuItem asChild>
                  <AlertDialog>
                    <AlertDialogTrigger asChild>
                      <button className="text-red-600">Logout</button>
                    </AlertDialogTrigger>
                    <AlertDialogContent>
                      <AlertDialogTitle>Log Out</AlertDialogTitle>
                      <AlertDialogDescription>Are you sure you want to log out? This action will end your session.</AlertDialogDescription>
                      <div className="flex justify-end space-x-2 mt-4">
                        <AlertDialogCancel className="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300 transition">Cancel</AlertDialogCancel>
                        <AlertDialogAction onClick={handleLogout} className="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition">
                          Yes, Logout
                        </AlertDialogAction>
                      </div>
                    </AlertDialogContent>
                  </AlertDialog>
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          ) : (
            <Link to="/login">
              <Button className="bg-green-700 hover:bg-green-900 text-white font-medium py-2 px-4 rounded-full">Login</Button>
            </Link>
          )}
        </div>

        {/* Mobile Menu Button */}
        <button onClick={() => setMenuOpen(!menuOpen)} className="text-2xl md:hidden text-gray-800 focus:outline-none">
          <i className={`fa ${menuOpen ? 'fa-times' : 'fa-bars'}`} />
        </button>
      </div>

      {/* Mobile Menu */}
      <div className={`md:hidden bg-white shadow-md ${menuOpen ? 'block' : 'hidden'}`}>
        <ul className="flex flex-col items-start gap-4 py-4 px-6">
          {navLinks.map((link) => (
            <li key={link.name}>
              <a href={link.path} className="text-sm font-medium text-gray-800 hover:text-green-700 transition duration-200">
                {link.name}
              </a>
            </li>
          ))}
          {user ? (
            <>
              <li>
                <Link to="/profile" className="text-sm font-medium text-gray-800 hover:text-green-700 transition duration-200">
                  Profile
                </Link>
              </li>
              <li>
                <Link to="/result" className="text-sm font-medium text-gray-800 hover:text-green-700 transition duration-200">
                  Lihat Hasil
                </Link>
              </li>
              <li>
                <AlertDialog>
                  <AlertDialogTrigger asChild>
                    <button className="text-sm font-medium text-red-600">Logout</button>
                  </AlertDialogTrigger>
                  <AlertDialogContent>
                    <AlertDialogTitle>Log Out</AlertDialogTitle>
                    <AlertDialogDescription>Are you sure you want to log out? This action will end your session.</AlertDialogDescription>
                    <div className="flex justify-end space-x-2 mt-4">
                      <AlertDialogCancel className="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300 transition">Cancel</AlertDialogCancel>
                      <AlertDialogAction onClick={handleLogout} className="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition">
                        Yes, Logout
                      </AlertDialogAction>
                    </div>
                  </AlertDialogContent>
                </AlertDialog>
              </li>
            </>
          ) : (
            <li>
              <Link to="/login" className="bg-green-700 hover:bg-green-900 text-white font-medium py-2 px-4 rounded-full">
                Login
              </Link>
            </li>
          )}
        </ul>
      </div>
    </nav>
  );
};

export default Navbar;
