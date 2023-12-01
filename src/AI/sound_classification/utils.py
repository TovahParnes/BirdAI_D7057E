import json
import os
import random
import shutil
import string
import re
import base64

def load_json_file(_filename):
    # Load JSON data from a file and return it as a Python dictionary.
    with open(_filename, 'r') as json_file:
        _labels = json.load(json_file)
        return _labels


def create_folder(folder_name):
    # Create a folder with the specified name and print a message indicating success or existence.
    try:
        os.mkdir(folder_name)
        print(f"Folder '{folder_name}' created successfully.")
    except FileExistsError:
        print(f"Folder '{folder_name}' already exists.")


def delete_folder(folder_path):
    # Delete a folder and its contents and print a message indicating success, non-existence, or error.
    try:
        shutil.rmtree(folder_path)
        print(f"Folder '{folder_path}' and its contents have been deleted.")
    except FileNotFoundError:
        print(f"Folder '{folder_path}' does not exist.")
    except Exception as e:
        print(f"An error occurred while deleting '{folder_path}': {str(e)}")


def folder_exists(folder_path):
    # Check if a folder exists at the specified path.
    return os.path.exists(folder_path)


def generate_random_name(length=32):
    characters = string.ascii_letters + string.digits  # Use letters and digits
    random_name = ''.join(random.choice(characters) for _ in range(length))
    return random_name


def check_images_folder():
    # Check if the "images" folder exists in the parent directory. If not, create it.
    parent_directory = os.path.dirname(os.path.abspath(__file__))
    images_folder = os.path.join(parent_directory, "images")

    if not os.path.exists(images_folder):
        create_folder(images_folder)


def check_prerequisites():
    model_filename = "mobilenet_model.keras"
    labels_filename = "labels.json"

    # Check if the "images" folder exists in the parent directory. If not, create it.
    check_images_folder()

    model_exists = check_directory_for_file(model_filename)
    labels_exists = check_directory_for_file(labels_filename)

    if model_exists and labels_exists:
        return True
    else:
        missing_files = []
        if not model_exists:
            missing_files.append(model_filename)
        if not labels_exists:
            missing_files.append(labels_filename)

        print("The following prerequisite files are missing:")
        for missing_file in missing_files:
            print(f"- {missing_file}")

        return False


def check_directory_for_file(filename, path=None):
    if path is None:
        path = os.getcwd()  # Default to the current working directory

    file_path = os.path.join(path, filename)
    return os.path.exists(file_path)


def decode_base64(encoded_data):
    """
    Decode a base64-encoded string and remove the data URI header.
    """
    data_match = re.match(r'^data:(.+);base64,', encoded_data)
    if data_match:
        data_type = data_match.group(1)
        decoded_data = re.sub(f'^data:{data_type};base64,', '', encoded_data)
        return data_type, decoded_data


def preprocess_sound(encoded_data):
    return base64.b64decode(decode_base64(encoded_data)[1], validate=True)
