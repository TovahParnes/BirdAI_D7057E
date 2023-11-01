import re
import base64
from io import BytesIO
from PIL import Image


def decode_base64(encoded_data):
    """
    Decode a base64-encoded string and remove the data URI header.
    """

    try:
        decoded_data = re.sub('^data:image/.+;base64,', '', encoded_data)
        return decoded_data
    except Exception as e:
        # Handle any unexpected errors
        print(f"An unexpected error occurred while decoding base64: {e}")
        return None


def create_pil_image_from_base64(encoded_data):
    """
    Create a PIL image from base64-encoded data.
    """

    try:
        decoded_data = decode_base64(encoded_data)
        image_bytes = base64.b64decode(decoded_data, validate=True)
        return Image.open(BytesIO(image_bytes))
    except (base64.binascii.Error, OSError) as e:
        # Handle base64 decoding or image creation errors
        print(f"Error while creating PIL image: {e}")
        return None
    except Exception as e:
        # Handle any other unexpected errors
        print(f"An unexpected error occurred: {e}")
        return None
