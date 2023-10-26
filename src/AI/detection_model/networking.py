import os
import socket
from PIL import Image
import PIL
from flask import Flask, request, jsonify
import requests
import io
import yolo

app = Flask(__name__)
_receiver_port = 3501


@app.route('/process_image', methods=['POST'])
def process_image():
    data = request.get_json()

    width = data["width"]
    height = data["height"]
    mode = data["mode"]
    pixels = [value for pixel in data["pixels"] for value in pixel]

    # Create a byte array from the pixel data
    pixel_bytes = bytes(pixels)

    # Create a PIL image from the pixel data
    _image = Image.frombytes(mode, (width, height), pixel_bytes)

    _image.show()

    _result_image = yolo.run_classification(_image)

    pixel_data = list(_result_image.getdata())

    # Create a dictionary with image data
    image_info = {
        "width": _result_image.width,
        "height": _result_image.height,
        "mode": _result_image.mode,
        "pixels": pixel_data
    }

    return jsonify({'message': 'Image successfully received and processed', 'image_data': pixel_data})


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)

