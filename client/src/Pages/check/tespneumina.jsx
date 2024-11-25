import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function Tespneumina() {
  const [image, setImage] = useState(null);
  const [previewUrl, setPreviewUrl] = useState(null);
  const [prediction, setPrediction] = useState(null);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  // Handle file selection
  const handleFileChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setImage(file);
      setPreviewUrl(URL.createObjectURL(file));
      setPrediction(null); // Reset prediction when a new image is selected
    }
  };

  // Handle drag and drop
  const handleDrop = (e) => {
    e.preventDefault();
    const file = e.dataTransfer.files[0];
    if (file) {
      setImage(file);
      setPreviewUrl(URL.createObjectURL(file));
      setPrediction(null); // Reset prediction when a new image is dropped
    }
  };

  // Handle clearing the image
  const handleClear = () => {
    setImage(null);
    setPreviewUrl(null);
    setPrediction(null);
  };

  // Handle submit
  const handleSubmit = () => {
    if (!image) {
      alert('Silakan pilih gambar terlebih dahulu');
      return;
    }

    setLoading(true);

    const reader = new FileReader();
    reader.onloadend = () => {
      const base64String = reader.result;
      // Send base64String to your backend
      fetch('http://127.0.0.1:5002/predict', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ image: base64String }),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.result) {
            setPrediction(data.result);
          } else if (data.error) {
            alert('Error: ' + data.error);
          }
        })
        .catch((error) => {
          console.error('Error:', error);
          alert('Terjadi kesalahan saat memproses gambar.');
        })
        .finally(() => setLoading(false));
    };
    reader.readAsDataURL(image);
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 relative">
      {/* Button Kembali */}
      <button className="absolute top-4 left-4 px-4 py-2 bg-gray-300 text-black rounded-lg hover:bg-gray-400 transition" onClick={() => navigate('/dashboard')}>
        Kembali
      </button>

      <h1 className="text-2xl font-bold text-gray-800 mb-6">Pneumonia Detection</h1>
      <div className="w-80 h-48 border-2 border-dashed border-gray-400 rounded-lg flex items-center justify-center relative bg-white" onDragOver={(e) => e.preventDefault()} onDrop={handleDrop}>
        {previewUrl ? (
          <img src={previewUrl} alt="Preview" className="max-w-full max-h-full object-contain rounded-md" />
        ) : (
          <div className="text-center text-gray-500">
            <p>Drop image here or click to select</p>
            <input type="file" accept="image/*" className="absolute w-full h-full opacity-0 cursor-pointer" onChange={handleFileChange} />
          </div>
        )}
      </div>

      <div className="mt-4 flex flex-col items-center gap-4">
        <div className="flex gap-4">
          <button onClick={handleSubmit} className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition">
            Submit
          </button>
          <button onClick={handleClear} className="px-4 py-2 bg-gray-300 text-black rounded-lg hover:bg-gray-400 transition">
            Clear
          </button>
        </div>
        {prediction && (
          <div className="mt-4">
            <span className={`text-xl font-semibold ${prediction === 'PNEUMONIA' ? 'text-red-500' : 'text-gray-800'}`}>Prediction: {prediction}</span>
          </div>
        )}
      </div>

      {/* Button Konsultasi Dokter */}
      <button className="absolute bottom-4 right-4 px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 transition" onClick={() => navigate('/dokterkonsul')}>
        Konsultasi Dokter
      </button>
    </div>
  );
}
