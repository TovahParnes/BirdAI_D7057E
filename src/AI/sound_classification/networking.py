from flask import Flask

import utils

app = Flask(__name__)
_receiver_port = 3503


def listen():
    app.run(host='0.0.0.0', port=_receiver_port)


@app.route('/process_sound', methods=['POST'])
def process_sound():
    pass


def get_label_from_position(_position):
    # Load labels from a JSON file and retrieve the label associated with a given position.
    _labels = utils.load_json_file("labels.json")
    _entry = _labels[str(_position)]

    if _entry is not None:
        return _entry
    else:
        print("Entry", str(_position), "not found in the dictionary")
