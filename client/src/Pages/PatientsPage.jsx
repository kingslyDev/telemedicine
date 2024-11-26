import React, { useState, useEffect } from "react";
import axios from "axios";

const PatientsPage = () => {
  const [patients, setPatients] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchPatients = async () => {
      try {
        console.log("Fetching patients...");
        const response = await axios.get("http://localhost:8080/api/patients");
        console.log("Response:", response.data);
        setPatients(response.data.patients || []); // Handle cases where 'patients' is undefined
        setLoading(false);
      } catch (err) {
        console.error("Error fetching patients:", err.response?.data || err.message);
        setError(err.response?.data?.error || "Failed to fetch patients");
        setLoading(false);
      }
    };

    fetchPatients();
  }, []);

  if (loading)
    return <div className="text-center mt-10 text-lg font-semibold">Loading...</div>;
  if (error)
    return <div className="text-center mt-10 text-red-500 font-semibold">Error: {error}</div>;

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6 text-center">Data Demografis Pasien</h1>
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white border border-gray-200 shadow-md rounded-lg">
          <thead className="bg-gray-100 border-b">
            <tr>
              <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">No</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">Username</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">First Name</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">Last Name</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">Date of Birth</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">Gender</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">Address</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">Blood Type</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">Emergency Contact</th>
              {/* <th className="px-6 py-4 text-left text-sm font-medium text-gray-600">Action</th> */}
            </tr>
          </thead>
          <tbody>
            {patients.map((patient, index) => (
              <tr key={patient.id} className="border-b hover:bg-gray-50">
                <td className="px-6 py-4 text-sm text-gray-700">{index + 1}</td>
                <td className="px-6 py-4 text-sm text-gray-700">
                  {patient.user?.username || "N/A"}
                </td>
                <td className="px-6 py-4 text-sm text-gray-700">{patient.first_name}</td>
                <td className="px-6 py-4 text-sm text-gray-700">{patient.last_name}</td>
                <td className="px-6 py-4 text-sm text-gray-700">
                  {patient.date_of_birth
                    ? new Date(patient.date_of_birth).toLocaleDateString()
                    : "N/A"}
                </td>
                <td className="px-6 py-4 text-sm text-gray-700">{patient.gender || "N/A"}</td>
                <td className="px-6 py-4 text-sm text-gray-700">{patient.address || "N/A"}</td>
                <td className="px-6 py-4 text-sm text-gray-700">{patient.blood_type || "N/A"}</td>
                <td className="px-6 py-4 text-sm text-gray-700">
                  {patient.emergency_contact || "N/A"}
                </td>
                {/* <td className="px-6 py-4 text-sm">
                  <button
                    className="text-white bg-blue-500 hover:bg-blue-600 px-4 py-2 rounded"
                    onClick={() => (window.location.href = `/data/${patient.id}`)}
                  >
                    View Details
                  </button>
                </td> */}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default PatientsPage;