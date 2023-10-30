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

detection_model_ip = "172.20.0.4" # TODO THIS IS RANDOMIZED PLZ FIX
detection_model_port = "3501"

classification_model_ip = "172.20.0.3" # TODO THIS IS RANDOMIZED PLZ FIX
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

    # TODO: Convert to PIL Image.

    _result_image = send_image_to_detection(im1, detection_model_ip, detection_model_port)
    _result = send_image_to_classification(_result_image, classification_model_ip, classification_model_port)

    print(_result)

    birds = [
        {"name": str(_result), "accuracy": 1.00}
    ]

#     bird = {
#         "name": "test-Sparrow",
#         "accuracy": 0.95
#     }

    # Determine the next bird number
#     next_bird_number = len(birds) + 1
#     new_bird_key = f"bird{next_bird_number}"
#
#     # Add a new bird
#     birds[new_bird_key] = {"name": "New-Bird", "accuracy": 0.91}

    json_structure = json.dumps({"birds": birds}, indent=4)

    # print(json_structure)
    return json_structure


def send_image_to_classification(_img, _ip, _port):     # Returns a PIL image.
    url = f'http://{_ip}:{_port}/process_image'

    # Send a POST request with the image data
    response = requests.post(url, json=convertPILToJSON(_img))

    _result = response.json()['label']

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
