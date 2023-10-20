import base64
import os
import socket
from flask import Flask, request, jsonify
import requests
import io

import yolo

app = Flask(__name__)
_receiver_port = 80


@app.route('/upload', methods=['POST'])
def upload_image():
    if 'image' not in request.files:
        return jsonify({'error': 'No image provided'})

    file = request.files['image']
    if file.filename == '':
        return jsonify({'error': 'No selected file'})

    if file:
        _result_image = yolo.run_classification(file)

        if _result_image is None:
            return jsonify({'error': 'No object detected in image.'})
        else:
            # Assuming _result_image is a PIL (Pillow) Image object
            # Convert it to base64
            img_bytes = io.BytesIO()
            _result_image.save(img_bytes, format='PNG')
            img_base64 = base64.b64encode(img_bytes.getvalue()).decode('utf-8')

            # Return the image as a base64 string in JSON response
            return jsonify({'image_data': img_base64})
    else:
        return jsonify({'error': 'Image detection failed'})


def create_upload_folder():
    _upload_folder = 'uploads'  # Define the folder where uploaded files will be stored
    os.makedirs(_upload_folder, exist_ok=True)  # Create the 'uploads' folder if it doesn't exist
    app.config['UPLOAD_FOLDER'] = _upload_folder  # Set the app's upload folder


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)


def get_local_ip():
    try:
        # Create a socket object and connect to an external server (e.g., Google's DNS server)
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        s.connect(("8.8.8.8", 80))
        local_ip = s.getsockname()[0]
        s.close()
        return local_ip
    except Exception as e:
        print(f"Error: {e}")
        return None


def send_hello_message(_ip, _port):
    url = "http://" + _ip + ":" + _port + "/receive_and_send_back_name"
    payload = {"message": "Hello, receiver!"}
    response = requests.post(url, json=payload)

    # Check the response status code
    if response.status_code == 200:
        print('Hello message, Request was successful.')
        print('Response JSON:', response.json())
    else:
        print(f'Hello message, Request failed with status code {response.status_code}.')
