import torch
import os
import shutil
from PIL import Image, ImageOps, ImageDraw
from ultralytics import YOLO
model = YOLO("yolov8m.pt")

def load_images_paths(path):
    image_paths = []
    # Get a list of image file names in the directory
    image_files = [f for f in os.listdir(path) if f.endswith((".jpg", ".png", ".jpeg"))]
    for image_file in image_files:
        image_paths.append(os.path.join(path, image_file))
    return image_paths

def predict_image(image, rotations):
    all_bird_data = []
    rotated_images = rotate_image(image, rotations, _multiple_images=False)
    for img in rotated_images:      # If rotation == 0 it won't loop
        results = model.predict(source=img, name="prediction")
        print(results[0].boxes.data)
        for index in range(0, len(results[0].boxes.data)):
            # IS IT A BIRD?
            if (results[0].boxes.data[index][5] == 14):
                # Extract and parse data
                tx = torch.ceil(results[0].boxes.data[index][0]).item()
                ty = torch.ceil(results[0].boxes.data[index][1]).item()
                tw = torch.ceil(results[0].boxes.data[index][2]).item()
                th = torch.ceil(results[0].boxes.data[index][3]).item()
                score = results[0].boxes.data[index][4].item()

                # Extracts the bird box and makes it a separate image
                boxes = results[0].boxes.cpu().numpy()
                edges = boxes[index].xyxy[0].astype(int)  # get corner points as int
                box_image = image.crop(edges)

                bird_data = [box_image, tx, ty, tw, th, score]

                all_bird_data.append(bird_data)
    return all_bird_data

def rotate_image(image, rotations, _multiple_images):
    images = []
    if rotations == 0:
        images.append(image)
        return images
    if _multiple_images:
        for rotation in range(rotations+1):
            images.append(image.rotate(rotation*90, expand=True))
    else:
        images.append(image.rotate(rotations*90, expand=True))
    return images

def crop_images_from_folder(path, target_size):
    prediction_dirs = [d for d in os.listdir(path) if os.path.isdir(os.path.join(path, d))]
    image_number = 0
    for prediction_dir in prediction_dirs:
        # Construct the path to the "crop" folder within the prediction directory
        crop_folder = os.path.join(path, prediction_dir, "crops/bird/")

        # Ensure the "crop" folder exists
        if not os.path.exists(crop_folder):
            continue  # Skip if "crop" folder is not found

        # Get a list of image files in the "crop" folder
        image_files = [f for f in os.listdir(crop_folder) if f.endswith((".jpg", ".png", ".jpeg"))]

        result_images = []

        for image_file in image_files:
            image_number += 1

            # Construct the full path to the image file
            image_path = os.path.join(crop_folder, image_file)

            os.makedirs("cropped_images", exist_ok=True)

            image = Image.open(image_path)
            paddedImage = pad_image(image, target_size)
            res_image = paddedImage.resize(target_size, Image.LANCZOS)
            result_images.append(res_image)

        return result_images

def crop_images(image, target_size):
    paddedImage = pad_image(image, target_size)
    result = paddedImage.resize(target_size, Image.LANCZOS)
    return result

def pad_image(image, _target_size):
        imageWidth, imageHeight = image.size
        currentAspectRatio = imageWidth/imageHeight
        targetAspectRatio = _target_size[0]/_target_size[1]
        if currentAspectRatio > targetAspectRatio:  # Vertical padding
            verticalPadding = int(((imageWidth/targetAspectRatio)-imageHeight)/2)
            #print("vertical padding: ", verticalPadding)
            padding = (0, verticalPadding, 0, verticalPadding)
            paddedImage = ImageOps.expand(image, padding, fill="white")

        elif currentAspectRatio < targetAspectRatio:    # Horizontal padding
            horizontalPadding = int(((targetAspectRatio*imageHeight)-imageWidth)/2)
            #print("horizontal padding: ", horizontalPadding)
            padding = (horizontalPadding, 0, horizontalPadding, 0)
            paddedImage = ImageOps.expand(image, padding, fill="white")

        else:
            paddedImage = image

        return paddedImage

def delete_images(path):
    os.makedirs("cropped_images", exist_ok=True)

    # Get a list of all files and subdirectories in the directory
    directory_contents = os.listdir(path)

    # Iterate through the contents and remove them
    for item in directory_contents:
        item_path = os.path.join(path, item)
        if os.path.isfile(item_path):
            # Remove files
            os.remove(item_path)
        elif os.path.isdir(item_path):
            # Remove directories and their contents (recursively)
            shutil.rmtree(item_path)

def run_classification(_image):
    # prediction_results[Bird index][0 = image, 5 = accuracy]
    prediction_results = predict_image(_image, rotations=0)

    # The bird with the highest accuracy is always first in the array, so we use that as input
    _res_image = crop_images(prediction_results[0][0], target_size=(224, 224))
    if _res_image:
        return _res_image
    else:
        return None