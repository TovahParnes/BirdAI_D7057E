import os

import numpy as np
from flask import Flask, request, jsonify, json
import requests
from PIL import Image
import PIL

app = Flask(__name__)
_receiver_port = 3500

detection_model_ip = "172.18.0.2"
detection_model_port = "3501"

classification_model_ip = "172.18.0.3"
classification_model_port = "3502"


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)


@app.route('/evaluate_image')
def evaluate_image():

    data = request.get_json()
    images = data.get('media')

    # TODO: Convert to PIL Image.

    im1 = Image.open(r"image.jpg")

    _result_image = send_image_to_detection(im1, detection_model_ip, detection_model_port)
    _result = send_image_to_classification(_result_image, classification_model_ip, classification_model_port)

    print(_result)

    birds = {
        "bird1": {"name": "test-Sparrow", "accuracy": 0.95},
        "bird2": {"name": "test-Robin", "accuracy": 0.92},
        "bird3": {"name": "test-Eagle", "accuracy": 0.98},
        "bird4": {"name": "test-Owl", "accuracy": 0.89}
    }

    bird = {
        "name": "test-Sparrow",
        "accuracy": 0.95
    }

    # Determine the next bird number
    next_bird_number = len(birds) + 1
    new_bird_key = f"bird{next_bird_number}"

    # Add a new bird
    birds[new_bird_key] = {"name": "New-Bird", "accuracy": 0.91}

    json_structure = json.dumps({"birds": birds}, indent=4)

    # print(json_structure)
    return json_structure


def send_image_to_classification(_img, _ip, _port):     # Returns a PIL image.
    url = f'http://{_ip}:{_port}/process_image'

    # Send a POST request with the image data
    response = requests.post(url, json=convertPILToJSON(_img))

    _result = response.json()['pos']

    return _result


def send_image_to_detection(_img, _ip, _port):     # Returns a PIL image.
    url = f'http://{_ip}:{_port}/process_image'

    # Send a POST request with the image data
    response = requests.post(url, json=convertPILToJSON(_img))

    _image_data = response.json()['image_data']

    # Create a byte array from the pixel data
    pixel_bytes = bytes([value for pixel in _image_data for value in pixel])

    # Create a PIL image from the pixel data
    image = Image.frombytes('RGB', (224, 224), pixel_bytes)

    return image


def convertPILToJSON(_image):

    # Convert the image data to a list of pixel values
    pixel_data = list(_image.getdata())

    # Create a dictionary with image data
    image_info = {
        "width": _image.width,
        "height": _image.height,
        "mode": _image.mode,
        "pixels": pixel_data
    }

    return image_info
