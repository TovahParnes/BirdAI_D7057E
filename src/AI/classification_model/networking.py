from flask import Flask, request, jsonify
from PIL import Image
import io
import base64
import numpy as np


import classification

app = Flask(__name__)
_receiver_port = 80


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)


@app.route('/upload', methods=['POST'])
def upload_image():

    try:
        _model = classification.load_model()
    except:
        return jsonify({'error': 'Loading model failed.'})

    try:
        data = request.get_json()
    except:
        return jsonify({'error': 'failed to get data from response.'})

    try:
        image_data = data.get('image_data')
        if image_data is None:
            raise ValueError('Failed to get image_data from JSON data.')
    except Exception as e:
        error_msg = str(e)
        return jsonify({'error': error_msg})

    if image_data is not None:
        try:
            # Decode base64 to obtain binary image data
            image_bytes = base64.b64decode(image_data)
        except:
            return jsonify({'error': 'Decoding of data failed'})

        try:
            # Create a PIL Image from the binary data
            image = Image.open(io.BytesIO(image_bytes))
        except:
            return jsonify({'error': 'Failed to convert data to Image.'})

        try:
            image = image.resize((224, 224))  # Resize the image to your target size
            image = np.array(image)  # Convert the image to a NumPy array
            image = image.reshape((1,) + image.shape)  # Reshape the image to have a batch dimension
        except:
            return jsonify({'error': 'Reshaping the image failed.'})

        try:
            # Assuming classification.predict() takes an image object as its argument
            prediction = classification.predict(classification.load_model(), image)
        except:
            return jsonify({'error': 'Failed during classification'})

        try:
            # Map class index to category label
            pos = int(np.argmax(prediction, axis=-1)[0])
            msg = "Image received and processed successfully with position: ", pos
            return jsonify({'message': msg})
        except:
            print("Failed during return message composing")
            return jsonify({'error': 'Failed during return message composing'})

    else:
        return jsonify({'error': 'No image_data found in the request'})
