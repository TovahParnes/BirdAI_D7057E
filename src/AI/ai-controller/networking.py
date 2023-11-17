from flask import Flask, request, jsonify, json
import requests
from PIL import Image
from requests import RequestException

import utils
import os

from dotenv import load_dotenv

# Determine the environment (e.g., 'dev' or 'prod')
env = 'prod'  # Change this to 'prod' for the production environment

# Load the environment variables from the corresponding .env file
env_file = f'.env.{env}'
if os.path.exists(env_file):
    load_dotenv(env_file)

# Access the values using the os.getenv() function
detection_model_ip = os.getenv('DETECTION_MODEL_IP')
detection_model_port = os.getenv('DETECTION_MODEL_PORT')
classification_model_ip = os.getenv('CLASSIFICATION_MODEL_IP')
classification_model_port = os.getenv('CLASSIFICATION_MODEL_PORT')
sound_classification_model_ip = os.getenv('SOUND_CLASSIFICATION_MODEL_IP')
sound_classification_model_port = os.getenv('SOUND_CLASSIFICATION_MODEL_PORT')


app = Flask(__name__)
_receiver_port = 3500


def listen():
    """
     Starts the listening service, the service will now listen for outbound requests.
    """
    app.run(host='0.0.0.0', port=_receiver_port)


@app.route('/evaluate_sound', methods=['POST'])
def evaluate_sound():
    data = request.get_json()

    _result = send_data_to_sound_classification(data, sound_classification_model_ip, sound_classification_model_port)

    if _result is not None:
        json_data = _result.json()

        if 'label' in json_data:
            _name = json_data['label']

        if 'accuracy' in json_data:
            _accuracy = json_data['accuracy']

        birds = [{"name": _name, "accuracy": _accuracy}]
        return json.dumps({"birds": birds}, indent=4)

    return jsonify({"error": "Something was wrong with the request or something went wrong during receiving data from "
                             "sound classification model"}, 400)


@app.route('/evaluate_image', methods=['POST'])
def evaluate_image():
    """
     Takes a request of an image for detection & classification and returns the prediction.
    """

    data = request.get_json()

    try:
        media_data = data["media"]
    except KeyError:
        return jsonify({"error": "Media data is missing in the request."}, 400)

    image = utils.create_pil_image_from_base64(media_data)

    if image is not None:
        _result_image = send_image_to_detection(image, detection_model_ip, detection_model_port)
    else:
        return jsonify({"error": "While trying to create a PIL image"}, 400)

    if _result_image is not None:
        _result = send_image_to_classification(_result_image, classification_model_ip, classification_model_port)
    else:
        return jsonify({"error": "Failure while sending image to detection module"}, 400)

    if _result is None:
        return jsonify({"error": "Failure while sending image to classification module"}, 400)

    birds = [{"name": _result.json()['label'], "accuracy": _result.json()['accuracy']}]

    json_structure = json.dumps({"birds": birds}, indent=4)

    return json_structure


def send_image_to_classification(_img, _ip, _port):     # Returns a PIL image.
    """
    Send an image for detection to a remote server and receive the result as a PIL image.
    """

    url = f'http://{_ip}:{_port}/process_image'

    try:
        # Send a POST request with the image data
        response = requests.post(url, json=convert_pil_to_json(_img))
        response.raise_for_status()  # Raise an exception for HTTP errors

        return response

    except requests.exceptions.RequestException as e:
        # Handle request errors, such as network issues or server unavailability
        print(f"Request error: {e}")
        return None
    except requests.exceptions.HTTPError as e:
        # Handle HTTP errors (e.g., 4xx or 5xx status codes)
        print(f"HTTP error: {e}")
        return None
    except Exception as e:
        # Handle any other unexpected errors
        print(f"An unexpected error occurred: {e}")
        return None


def send_image_to_detection(_img, _ip, _port):
    """
    Send an image for detection to a remote server and receive the result as a PIL image.
    """

    url = f'http://{_ip}:{_port}/process_image'

    try:
        # Send a POST request with the image data
        response = requests.post(url, json=convert_pil_to_json(_img))
        response.raise_for_status()  # Raise an exception for HTTP errors

        _image_data = response.json()['image_data']

        # Create a byte array from the pixel data
        pixel_bytes = bytes([value for pixel in _image_data for value in pixel])

        # Create a PIL image from the pixel data
        image = Image.frombytes('RGB', (224, 224), pixel_bytes)

        return image

    except RequestException as e:
        print(f"Request error: {e}")
        # Handle the error as needed, e.g., log it or return a default image
        return None
    except (ValueError, KeyError) as e:
        print(f"JSON parsing error: {e}")
        # Handle JSON parsing errors, e.g., log them or return a default image
        return None
    except Exception as e:
        print(f"An unexpected error occurred: {e}")
        # Handle any other unexpected errors
        return None


def send_data_to_sound_classification(_data, _ip, _port):
    url = f'http://{_ip}:{_port}/process_sound'

    try:
        response = requests.post(url, json=_data)
        return response
    except RequestException as e:
        print(f"Request error: {e}")
        return None
    except (ValueError, KeyError) as e:
        print(f"JSON parsing error: {e}")
        return None
    except Exception as e:
        print(f"An unexpected error occurred: {e}")
        return None


def convert_pil_to_json(image):
    """
    Convert a PIL image to a JSON-compatible representation.
    """

    try:
        # Convert the image data to a list of pixel values
        pixel_data = list(image.getdata())

        # Create a dictionary with image data
        image_info = {
            "width": image.width,
            "height": image.height,
            "mode": image.mode,
            "pixels": pixel_data
        }

        return image_info
    except Exception as e:
        # Handle any unexpected errors
        print(f"An unexpected error occurred while converting PIL to JSON: {e}")
        return None

