from flask import Flask, request, jsonify
from PIL import Image
import io
import numpy as np
import json

import classification

app = Flask(__name__)
_receiver_port = 3502


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)


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

    _image = np.array(_image)  # Convert the image to a NumPy array
    _image = _image.reshape((1,) + _image.shape)  # Reshape the image to have a batch dimension

    _result = classification.predict(classification.load_model(), _image)

    _pos = int(np.argmax(_result, axis=-1)[0])

    _label = get_label_from_position(_pos)
    return jsonify({'message': 'Image successfully received and processed', 'label': _label})

def get_label_from_position(_position):

    _labels = load_json_file("labels.json")
    _entry = _labels[str(_position)]

    if _entry is not None:
        return _entry
    else:
        print("Entry", str(_position), "not found in the dictionary")

def load_json_file(_filename):
    with open(_filename, 'r') as json_file:
        _labels = json.load(json_file)
        return _labels