import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

const PatientDetails = () => {
  const { id } = useParams();
  const [patient, setPatient] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchPatientDetails = async () => {
      try {
        console.log("Fetching patient details...");
        const response = await axios.get(`http://localhost:8080/api/patients/${id}`);
        console.log("Response:", response.data);
        setPatient(response.data.patient);
      } catch (err) {
        console.error("Error fetching patient details:", err.response || err.message);
        setError(err.response?.data?.error || "Failed to fetch patient details");
      } finally {
        setLoading(false);
      }
    };
    fetchPatientDetails();
  }, [id]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;
  if (!patient) return <div>No patient data available</div>;

  return (
    <div>
      <h1>
        {patient.first_name} {patient.last_name}'s Details
      </h1>
      <p><strong>Email:</strong> {patient.user?.email || "N/A"}</p>
      <p>
        <strong>Date of Birth:</strong>{" "}
        {patient.date_of_birth
          ? new Date(patient.date_of_birth).toLocaleDateString()
          : "N/A"}
      </p>
      <p><strong>Gender:</strong> {patient.gender || "N/A"}</p>
      <p><strong>Address:</strong> {patient.address || "N/A"}</p>
      <p><strong>Blood Type:</strong> {patient.blood_type || "N/A"}</p>
      <h2>Medical Records</h2>
      <ul>
        {patient.medical_records?.length ? (
          patient.medical_records.map((record) => (
            <li key={record.id}>
              <strong>Diagnosis:</strong> {record.diagnosis || "N/A"}<br />
              <strong>Doctor:</strong> {record.doctor?.name || "N/A"}<br />
              <strong>Treatment Plan:</strong> {record.treatment_plan || "N/A"}
            </li>
          ))
        ) : (
          <p>No medical records found</p>
        )}
      </ul>
    </div>
  );
};

export default PatientDetails;