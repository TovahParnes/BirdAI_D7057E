import os

from flask import Flask, request, jsonify, json
import requests
from PIL import Image
import PIL
import base64
import re
from io import BytesIO

app = Flask(__name__)
_receiver_port = 3500

detection_model_ip = "172.20.0.4"
detection_model_port = "3501"

classification_model_ip = "172.20.0.3"
classification_model_port = "3502"


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)


@app.route('/evaluate_image', methods=['POST'])
def evaluate_image():

    data = request.get_json()

    _loaded_base64_data = data["media"]

    _loaded_base64_without_header = re.sub('^data:image/.+;base64,', '', _loaded_base64_data)

    image = base64.b64decode(_loaded_base64_without_header, validate=True)

    # Create a byte array from the pixel data
    pixel_bytes = bytes(image)

    # Create a PIL image from the pixel data
    im1 = Image.open(BytesIO(pixel_bytes))

    _result_image = send_image_to_detection(im1, detection_model_ip, detection_model_port)
    _result = send_image_to_classification(_result_image, classification_model_ip, classification_model_port)

    birds = [
        {"name": _result.json()['label'], "accuracy": _result.json()['accuracy']}
    ]

    json_structure = json.dumps({"birds": birds}, indent=4)

    return json_structure


def send_image_to_classification(_img, _ip, _port):     # Returns a PIL image.
    url = f'http://{_ip}:{_port}/process_image'

    # Send a POST request with the image data
    response = requests.post(url, json=convertPILToJSON(_img))

    _result = response

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
