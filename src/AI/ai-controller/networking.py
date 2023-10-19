import os
from flask import Flask, request, jsonify
import requests

app = Flask(__name__)
_receiver_port = 80


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)


def send_image(_ip, _port):
    # Specify the URL of the Flask server
    url = 'http://' + _ip + ':' + _port + '/upload'

    # Open and read the image file
    with open('image.jpg', 'rb') as file:
        files = {'image': ('image.jpg', file, 'image/jpeg')}
        response = requests.post(url, files=files)

    # Check the response from the server
    if response.status_code == 200:
        print("Image received.")
        return response
    else:
        print(f'Error: {response.status_code}')
    return None


def send_to_classification(_ip, _port, _image_data):
    url = 'http://' + _ip + ':' + _port + '/upload'

    print(_image_data)

    response = requests.post(url, json=_image_data)

    # Check the response status code
    if response.status_code == 200:
        print('Image to classification data back is:', response.json())
    else:
        print(f'Image to classification data, Message request failed - {response.status_code}.')


def send_message(_ip, _port, _route):
    url = "http://" + _ip + ":" + _port + "/" + _route
    payload = {"message": "Hello, receiver!"}
    response = requests.post(url, json=payload)

    # Check the response status code
    if response.status_code == 200:
        print('Send message response:', response.json())
    else:
        print(f'Message request failed - {response.status_code}.')


# Define a route that responds with a string
@app.route('/test_response')
def test_response():
    return "Hello from AI Controller!"


@app.route('/upload', methods=['POST'])
def upload_image():
    if 'image' not in request.files:
        return jsonify({'error': 'No image provided'})

    file = request.files['image']
    if file.filename == '':
        return jsonify({'error': 'No selected file'})

    if file:
        # Save the uploaded image to a folder (e.g., 'uploads')
        file.save(os.path.join('uploads', file.filename))
        return jsonify({'message': 'Image successfully uploaded'})
