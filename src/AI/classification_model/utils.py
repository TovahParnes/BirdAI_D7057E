import json
import os
import random
import shutil
import string


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
