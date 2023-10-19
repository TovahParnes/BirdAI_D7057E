import keras
from tensorflow import keras


def load_model():
    return keras.models.load_model("mobilenet_model.keras")


def predict(_model, _image):
    return _model.predict(_image)
