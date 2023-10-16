import networking
if __name__ == '__main__':
    print("Running Classification Model")

    # # Create an instance of the ImageDataGenerator for test data
    # test_datagen = ImageDataGenerator(rescale=1.0 / 255)  # You can specify other augmentation options here
    #
    # # Load and preprocess a single image
    # image_path = 'image.jpg'
    # image = Image.open(image_path)
    #
    # image = image.resize((224, 224))  # Resize the image to your target size
    # image = np.array(image)  # Convert the image to a NumPy array
    # image = image.reshape((1,) + image.shape)  # Reshape the image to have a batch dimension
    #
    # # Assuming classification.predict() takes an image object as its argument
    # prediction = classification.predict(classification.load_model(), image)
    #
    # # Map class index to category label
    # pos = int(np.argmax(prediction, axis=-1)[0])
    # print(pos)

    networking.listen()
