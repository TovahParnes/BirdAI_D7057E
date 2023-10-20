import networking as net

detection_model_ip = "172.18.0.2"
detection_model_port = "80"

classification_model_ip = "172.18.0.3"
classification_model_port = "80"


if __name__ == '__main__':
    _result = net.send_image(detection_model_ip, detection_model_port)
    # net.send_to_classification(classification_model_ip, classification_model_port, _result)

