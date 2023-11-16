import librosa
import numpy as np
from PIL import Image
from flask import Flask, jsonify, request
from keras.preprocessing.image import ImageDataGenerator
from matplotlib import pyplot as plt
import io

import classification
import utils

app = Flask(__name__)
_receiver_port = 3503


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)


def sound_to_spectrogram(sound_data):
    y, sr = librosa.load(io.BytesIO(sound_data), sr=None)

    # Generate a mel spectrogram
    spectrogram = librosa.feature.melspectrogram(y=y, sr=sr)
    spectrogram_db = librosa.power_to_db(spectrogram, ref=np.max)

    plt.figure(figsize=(10, 4))
    librosa.display.specshow(spectrogram_db, x_axis='time', y_axis='mel')
    plt.colorbar(format='%+2.0f dB')
    plt.title('Spectrogram')

    # Save the plot as an image in memory
    image_stream = io.BytesIO()
    plt.savefig(image_stream, format='png')
    image_stream.seek(0)

    # Create a PIL image from the saved image stream
    return Image.open(image_stream)


@app.route('/process_sound', methods=['POST'])
def process_sound():
    data = request.get_json()

    media_data = data["media"]
    decoded_data = utils.preprocess_sound(media_data)

    image = sound_to_spectrogram(decoded_data)

    _prefix_path = "images/"
    _folder_name = utils.generate_random_name()

    utils.create_folder(_prefix_path + _folder_name)

    _full_path = _prefix_path + _folder_name + "/bird01"
    utils.create_folder(_prefix_path + _folder_name + "/bird01")

    image.save(_full_path + "/image.png")   # Future, fix so that more birds can be saved.
    general_datagen = ImageDataGenerator(rescale=1. / 255)  # for training, validation and testing data

    testr_directory = _prefix_path + _folder_name

    testr_generator = general_datagen.flow_from_directory(
        testr_directory,
        target_size=(224, 224),
        batch_size=1
    )

    _result = classification.predict(classification.load_model(), testr_generator[0][0])

    # Get the predicted class probabilities for all classes
    predicted_probabilities = _result[0]

    # Get the predicted class probability
    predicted_class_probability = np.argsort(predicted_probabilities)[::-1]

    # Select the top three classes with the highest probabilities
    top_three_indices = predicted_class_probability[:3]

    utils.delete_folder(_prefix_path + _folder_name)

    _pos = int(np.argmax(_result, axis=-1)[0])

    _label = get_label_from_position(_pos)

    # Get the corresponding probabilities
    top_three_probabilities = [predicted_probabilities[i] for i in top_three_indices]

    return jsonify({'message': 'Image successfully received and processed', 'label': _label, 'accuracy': float(top_three_probabilities[0])})


def get_label_from_position(_position):
    # Load labels from a JSON file and retrieve the label associated with a given position.
    _labels = utils.load_json_file("labels.json")
    _entry = _labels[str(_position)]

    if _entry is not None:
        return _entry
    else:
        print("Entry", str(_position), "not found in the dictionary")
