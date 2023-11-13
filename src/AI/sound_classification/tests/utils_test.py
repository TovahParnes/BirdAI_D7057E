import unittest
import os
import shutil
import tempfile
from utils import load_json_file, create_folder, delete_folder, folder_exists, generate_random_name, check_prerequisites, check_directory_for_file


class TestYourFunctions(unittest.TestCase):

    def setUp(self):
        # Create a temporary directory for testing folder-related functions
        self.temp_dir = tempfile.mkdtemp()

    def tearDown(self):
        # Clean up the temporary directory after testing
        shutil.rmtree(self.temp_dir)

    def test_load_json_file(self):
        # Create a temporary JSON file for testing
        test_json_file = os.path.join(self.temp_dir, "test_labels.json")
        with open(test_json_file, 'w') as json_file:
            json_file.write('{"1": "Label1", "2": "Label2"}')

        result = load_json_file(test_json_file)
        self.assertEqual(result, {"1": "Label1", "2": "Label2"})

    def test_create_folder(self):
        test_folder = os.path.join(self.temp_dir, "test_folder")
        create_folder(test_folder)
        self.assertTrue(os.path.exists(test_folder))

        # Test the case where the folder already exists
        create_folder(test_folder)  # This should raise a FileExistsError
        self.assertTrue(os.path.exists(test_folder))

    def test_delete_folder(self):
        test_folder = os.path.join(self.temp_dir, "test_folder")
        os.mkdir(test_folder)

        delete_folder(test_folder)
        self.assertFalse(os.path.exists(test_folder))

        # Test the case where the folder does not exist
        non_existent_folder = os.path.join(self.temp_dir, "non_existent_folder")
        delete_folder(non_existent_folder)  # This should not raise an error

    def test_folder_exists(self):
        test_folder = os.path.join(self.temp_dir, "test_folder")
        os.mkdir(test_folder)
        self.assertTrue(folder_exists(test_folder))

        non_existent_folder = os.path.join(self.temp_dir, "non_existent_folder")
        self.assertFalse(folder_exists(non_existent_folder))

    def test_generate_random_name_length(self):
        random_name = generate_random_name()
        self.assertEqual(len(random_name), 32)  # Check if the length is as expected

    def test_generate_random_name(self):
        random_name_01 = generate_random_name()
        random_name_02 = generate_random_name()

        self.assertNotEqual(random_name_01, random_name_02)  # Check if the length is as expected

    def test_check_prerequisites(self):
        # Create a temporary "images" folder for testing
        test_folder = "test_images_folder"
        self.assertFalse(os.path.exists(test_folder))
        create_folder(test_folder)

        # Test when both model and labels files exist
        with open("mobilenet_model.keras", "w") as model_file, open("labels.json", "w") as labels_file:
            self.assertTrue(check_prerequisites())

        # Test when model file is missing
        os.remove("mobilenet_model.keras")
        self.assertFalse(check_prerequisites())

        # Test when labels file is missing
        os.remove("labels.json")
        self.assertFalse(check_prerequisites())

        # Clean up the temporary "images" folder
        os.rmdir(test_folder)

    def test_check_directory_for_file(self):
        # Create a temporary directory for testing
        test_dir = "test_dir_" + generate_random_name()
        self.assertFalse(os.path.exists(test_dir))
        create_folder(test_dir)
        self.assertTrue(os.path.exists(test_dir))

        # Clean up the temporary directory
        os.rmdir(test_dir)


if __name__ == '__main__':
    unittest.main()
