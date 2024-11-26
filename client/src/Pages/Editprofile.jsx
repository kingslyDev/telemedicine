// src/components/EditProfile/EditProfile.jsx
import React, { useState, useEffect } from 'react';
import { APIWithPrefix } from '../services/authService'; // Pastikan path ini benar
import { useAuth } from '../context/AuthContext';
import { useNavigate } from 'react-router-dom';

const EditProfile = () => {
  const { user, logout } = useAuth();
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    email: '',
    username: '',
    first_name: '',
    last_name: '',
    date_of_birth: '',
    gender: '',
    address: '',
    blood_type: '',
    emergency_contact: '',
    specialization: '',
    license_number: '',
    years_of_experience: 0,
    bio: '',
    privileges: '',
  });
  const [loading, setLoading] = useState(true);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  // Fungsi untuk mendapatkan peran pengguna dari context
  const getUserRole = () => {
    return user ? user.role : null;
  };

  // Mengambil data profil pengguna saat komponen dimuat
  useEffect(() => {
    const fetchProfile = async () => {
      const role = getUserRole();
      if (!role) {
        setError('Peran pengguna tidak ditemukan. Silakan login kembali.');
        setLoading(false);
        return;
      }

      let endpoint = '';
      switch (role) {
        case 'patient':
          endpoint = '/patient/profile';
          break;
        case 'doctor':
          endpoint = '/doctor/profile';
          break;
        case 'admin':
          endpoint = '/admin/profile';
          break;
        default:
          setError('Peran pengguna tidak valid.');
          setLoading(false);
          return;
      }

      try {
        const response = await APIWithPrefix.get(endpoint);
        const data = response.data;

        // Sesuaikan pengambilan data berdasarkan peran
        if (role === 'patient') {
          setFormData({
            email: data.patient.user.email || '',
            username: data.patient.user.username || '',
            first_name: data.patient.first_name || '',
            last_name: data.patient.last_name || '',
            date_of_birth: data.patient.date_of_birth ? data.patient.date_of_birth.split('T')[0] : '',
            gender: data.patient.gender || '',
            address: data.patient.address || '',
            blood_type: data.patient.blood_type || '',
            emergency_contact: data.patient.emergency_contact || '',
            specialization: '',
            license_number: '',
            years_of_experience: 0,
            bio: '',
            privileges: '',
          });
        } else if (role === 'doctor') {
          setFormData({
            email: data.doctor.user.email || '',
            username: data.doctor.user.username || '',
            first_name: data.doctor.first_name || '',
            last_name: data.doctor.last_name || '',
            specialization: data.doctor.specialization || '',
            license_number: data.doctor.license_number || '',
            years_of_experience: data.doctor.years_of_experience || 0,
            bio: data.doctor.bio || '',
            privileges: '',
          });
        } else if (role === 'admin') {
          setFormData({
            email: data.admin.user.email || '',
            username: data.admin.user.username || '',
            first_name: data.admin.first_name || '',
            last_name: data.admin.last_name || '',
            privileges: data.admin.privileges || '{}',
            specialization: '',
            license_number: '',
            years_of_experience: 0,
            bio: '',
            address: '',
            blood_type: '',
            emergency_contact: '',
            date_of_birth: '',
            gender: '',
          });
        }

        setLoading(false);
      } catch (err) {
        console.error('Gagal mengambil profil:', err);
        if (err.response?.status === 401) {
          setError('Token tidak valid atau telah kedaluwarsa. Silakan login kembali.');
          logout(); // Hapus token dan user dari localStorage
          navigate('/login'); // Redirect ke halaman login
        } else {
          setError(err.response?.data?.error || 'Gagal mengambil profil. Silakan coba lagi.');
        }
        setLoading(false);
      }
    };

    fetchProfile();
  }, [user, logout, navigate]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage('');
    setError('');

    const role = getUserRole();
    if (!role) {
      setError('Peran pengguna tidak ditemukan. Silakan login kembali.');
      setLoading(false);
      return;
    }

    let endpoint = '';
    switch (role) {
      case 'patient':
        endpoint = '/patient/profile';
        break;
      case 'doctor':
        endpoint = '/doctor/profile';
        break;
      case 'admin':
        endpoint = '/admin/profile';
        break;
      default:
        setError('Peran pengguna tidak valid.');
        setLoading(false);
        return;
    }

    // Buat data yang akan dikirim, sesuaikan dengan peran
    let dataToSend = {};
    if (role === 'patient') {
      dataToSend = {
        first_name: formData.first_name,
        last_name: formData.last_name,
        date_of_birth: formData.date_of_birth,
        gender: formData.gender,
        address: formData.address,
        blood_type: formData.blood_type,
        emergency_contact: formData.emergency_contact,
      };
    } else if (role === 'doctor') {
      dataToSend = {
        first_name: formData.first_name,
        last_name: formData.last_name,
        specialization: formData.specialization,
        license_number: formData.license_number,
        years_of_experience: formData.years_of_experience,
        bio: formData.bio,
        // Tambahkan field lain sesuai kebutuhan dokter
      };
    } else if (role === 'admin') {
      dataToSend = {
        first_name: formData.first_name,
        last_name: formData.last_name,
        privileges: formData.privileges,
        // Tambahkan field lain sesuai kebutuhan admin
      };
    }

    try {
      const response = await APIWithPrefix.put(endpoint, dataToSend);
      setMessage(response.data.message || 'Profil berhasil diperbarui!');
    } catch (err) {
      console.error('Gagal memperbarui profil:', err);
      if (err.response?.status === 401) {
        setError('Token tidak valid atau telah kedaluwarsa. Silakan login kembali.');
        logout(); // Hapus token dan user dari localStorage
        navigate('/login'); // Redirect ke halaman login
      } else {
        setError(err.response?.data?.error || 'Gagal memperbarui profil. Silakan coba lagi.');
      }
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="max-w-3xl mx-auto p-6 bg-white shadow-md rounded-lg">
      <h2 className="text-2xl font-bold mb-6">Edit Profile</h2>
      {message && <div className="mb-4 text-green-600">{message}</div>}
      {error && <div className="mb-4 text-red-600">{error}</div>}
      <form onSubmit={handleSubmit} className="space-y-4">
        {/* Email */}
        <div>
          <label className="block text-sm font-medium text-gray-700">Email</label>
          <input type="email" name="email" value={formData.email} disabled className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-100 text-gray-500 cursor-not-allowed" />
        </div>

        {/* Username */}
        <div>
          <label className="block text-sm font-medium text-gray-700">Username</label>
          <input type="text" name="username" value={formData.username} disabled className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-100 text-gray-500 cursor-not-allowed" />
        </div>

        {/* First Name */}
        <div>
          <label className="block text-sm font-medium text-gray-700">First Name</label>
          <input type="text" name="first_name" value={formData.first_name} onChange={handleChange} className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500" />
        </div>

        {/* Last Name */}
        <div>
          <label className="block text-sm font-medium text-gray-700">Last Name</label>
          <input type="text" name="last_name" value={formData.last_name} onChange={handleChange} className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500" />
        </div>

        {/* Date of Birth (Hanya untuk pasien) */}
        {getUserRole() === 'patient' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">Date of Birth</label>
            <input type="date" name="date_of_birth" value={formData.date_of_birth} onChange={handleChange} className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500" />
          </div>
        )}

        {/* Gender (Hanya untuk pasien) */}
        {getUserRole() === 'patient' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">Gender</label>
            <select name="gender" value={formData.gender} onChange={handleChange} className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500">
              <option value="">Select Gender</option>
              <option value="Pria">Pria</option>
              <option value="Wanita">Wanita</option>
            </select>
          </div>
        )}

        {/* Address (Hanya untuk pasien) */}
        {getUserRole() === 'patient' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">Address</label>
            <input type="text" name="address" value={formData.address} onChange={handleChange} className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500" />
          </div>
        )}

        {/* Blood Type (Hanya untuk pasien) */}
        {getUserRole() === 'patient' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">Blood Type</label>
            <input type="text" name="blood_type" value={formData.blood_type} onChange={handleChange} className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500" />
          </div>
        )}

        {/* Emergency Contact (Hanya untuk pasien) */}
        {getUserRole() === 'patient' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">Emergency Contact</label>
            <input
              type="text"
              name="emergency_contact"
              value={formData.emergency_contact}
              onChange={handleChange}
              className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        )}

        {/* Specialization (Hanya untuk dokter) */}
        {getUserRole() === 'doctor' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">Specialization</label>
            <input type="text" name="specialization" value={formData.specialization} onChange={handleChange} className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500" />
          </div>
        )}

        {/* License Number (Hanya untuk dokter) */}
        {getUserRole() === 'doctor' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">License Number</label>
            <input type="text" name="license_number" value={formData.license_number} onChange={handleChange} className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500" />
          </div>
        )}

        {/* Years of Experience (Hanya untuk dokter) */}
        {getUserRole() === 'doctor' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">Years of Experience</label>
            <input
              type="number"
              name="years_of_experience"
              value={formData.years_of_experience}
              onChange={handleChange}
              min="0"
              className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
        )}

        {/* Bio (Hanya untuk dokter) */}
        {getUserRole() === 'doctor' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">Bio</label>
            <textarea name="bio" value={formData.bio} onChange={handleChange} className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"></textarea>
          </div>
        )}

        {/* Privileges (Hanya untuk admin) */}
        {getUserRole() === 'admin' && (
          <div>
            <label className="block text-sm font-medium text-gray-700">Privileges (JSON)</label>
            <textarea
              name="privileges"
              value={formData.privileges}
              onChange={handleChange}
              placeholder='e.g., {"can_manage_users": true, "can_edit_settings": false}'
              className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
            ></textarea>
          </div>
        )}

        {/* Submit Button */}
        <div>
          <button type="submit" className="w-full py-2 px-4 bg-blue-500 text-white font-semibold rounded-md shadow-sm hover:bg-blue-600 focus:ring-2 focus:ring-blue-500 focus:ring-offset-1">
            Save Changes
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditProfile;
