from flask import Flask, request, jsonify
from PIL import Image
import numpy as np
from keras.preprocessing.image import ImageDataGenerator

import classification
import utils

app = Flask(__name__)
_receiver_port = 3502


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)


@app.route('/process_image', methods=['POST'])
def process_image():
    data = request.get_json()

    _prefix_path = "images/"
    _folder_name = utils.generate_random_name()

    utils.create_folder(_prefix_path + _folder_name)

    _full_path = _prefix_path + _folder_name + "/bird01"
    utils.create_folder(_prefix_path + _folder_name + "/bird01")

    width = data["width"]
    height = data["height"]
    mode = data["mode"]
    pixels = [value for pixel in data["pixels"] for value in pixel]

    _pixel_bytes = bytes(pixels)    # Create a byte array from the pixel data

    _image_from_bytes = Image.frombytes(mode, (width, height), _pixel_bytes)   # Create a PIL image from the pixel data

    _image_from_bytes.save(_full_path + "/image.jpg")   # Future, fix so that more birds can be saved.

    general_datagen = ImageDataGenerator(rescale=1. / 255)  # for training, validation and testing data

    testr_directory = _prefix_path + _folder_name

    testr_generator = general_datagen.flow_from_directory(
        testr_directory,
        target_size=(224, 224),
        batch_size=1
    )

    _result = classification.predict(classification.load_model(), testr_generator[0][0])

    _pos = int(np.argmax(_result, axis=-1)[0])

    _label = get_label_from_position(_pos)

    utils.delete_folder(_prefix_path + _folder_name)

    return jsonify({'message': 'Image successfully received and processed', 'label': _label})


def get_label_from_position(_position):
    # Load labels from a JSON file and retrieve the label associated with a given position.
    _labels = utils.load_json_file("labels.json")
    _entry = _labels[str(_position)]

    if _entry is not None:
        return _entry
    else:
        print("Entry", str(_position), "not found in the dictionary")

