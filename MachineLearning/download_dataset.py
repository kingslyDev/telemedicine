import os
from kaggle.api.kaggle_api_extended import KaggleApi

# Set path ke file kaggle.json di dalam direktori proyek
os.environ['KAGGLE_CONFIG_DIR'] = os.path.abspath('.')

# Pastikan direktori data ada
DATA_DIR = 'data'
if not os.path.exists(DATA_DIR):
    os.makedirs(DATA_DIR)

# Inisialisasi API Kaggle
api = KaggleApi()
api.authenticate()

# Unduh dataset
api.dataset_download_files('paultimothymooney/chest-xray-pneumonia', path=DATA_DIR, unzip=True)

print("Dataset telah diunduh dan diekstrak di direktori 'data'")
