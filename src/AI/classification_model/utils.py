import json
import os
import shutil


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
