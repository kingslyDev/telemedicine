# app.py

import os
import re
import base64
from flask import Flask, request, jsonify
from flask_cors import CORS
from tensorflow.keras.models import load_model
from tensorflow.keras.preprocessing import image
import numpy as np
from PIL import Image
from io import BytesIO
import logging
import threading

# Inisialisasi Flask app
app = Flask(__name__)
CORS(app)

# Konfigurasi logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Muat model yang telah dilatih
MODEL_PATH = 'models/pneumonia_model.h5'
try:
    model = load_model(MODEL_PATH)
    logger.info('Model telah dimuat.')
except Exception as e:
    logger.error(f"Gagal memuat model: {e}")
    model = None

# Definisikan parameter gambar
IMG_HEIGHT = 150
IMG_WIDTH = 150

# Buat lock untuk thread safety
model_lock = threading.Lock()

def base64_to_pil(img_base64):
    """
    Konversi data gambar base64 ke gambar PIL
    """
    try:
        image_data = re.sub('^data:image/.+;base64,', '', img_base64)
        pil_image = Image.open(BytesIO(base64.b64decode(image_data)))
        return pil_image
    except Exception as e:
        logger.error(f"Gagal mengonversi base64 ke PIL: {e}")
        return None

def model_predict(img, model):
    """
    Preprocessing gambar dan prediksi menggunakan model yang dimuat
    """
    try:
        # Ubah ukuran gambar
        img = img.resize((IMG_WIDTH, IMG_HEIGHT))
        # Konversi ke array
        x = image.img_to_array(img)
        # Expand dimensi untuk mencocokkan input model
        x = np.expand_dims(x, axis=0)
        # Normalisasi gambar
        x /= 255.0
        with model_lock:
            # Lakukan prediksi
            preds = model.predict(x)
        return preds
    except Exception as e:
        logger.error(f"Gagal melakukan prediksi: {e}")
        return None

@app.route('/predict', methods=['POST'])
def predict():
    try:
        # Dapatkan gambar dari permintaan
        img_data = request.json.get('image')
        if not img_data:
            return jsonify(error="Tidak ada gambar yang dikirim"), 400

        img = base64_to_pil(img_data)
        if img is None:
            return jsonify(error="Gagal memproses gambar"), 400

        if model is None:
            return jsonify(error="Model tidak tersedia"), 500

        # Prediksi kelas
        preds = model_predict(img, model)
        if preds is None:
            return jsonify(error="Gagal melakukan prediksi"), 500

        result = preds[0][0]
        if result > 0.5:
            prediction = "PNEUMONIA"
        else:
            prediction = "NORMAL"

        logger.info(f"Prediksi: {prediction}")
        return jsonify(result=prediction)

    except Exception as e:
        logger.error(f"Kesalahan saat memproses permintaan: {e}")
        return jsonify(error=str(e)), 500

if __name__ == '__main__':
    app.run(port=5002, debug=True)
