import unittest
from PIL import Image
from io import BytesIO
from utils import decode_base64, create_pil_image_from_base64


class TestUtils(unittest.TestCase):
    def test_decode_base64(self):
        # Test a valid base64 encoded string
        encoded_data = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA..."
        decoded_data = decode_base64(encoded_data)
        self.assertEqual(decoded_data[:11], "iVBORw0KGgo")

        # Test an empty string
        empty_encoded_data = ""
        decoded_empty_data = decode_base64(empty_encoded_data)
        self.assertEqual(decoded_empty_data, "")

        # Test a non-base64 string
        non_base64_data = "This is not a valid base64 string"
        decoded_non_base64_data = decode_base64(non_base64_data)
        self.assertEqual(decoded_non_base64_data, non_base64_data)


if __name__ == '__main__':
    unittest.main()
