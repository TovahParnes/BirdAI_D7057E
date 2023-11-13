import networking
import utils

if __name__ == '__main__':

    # Example usage of check_prerequisites:
    if utils.check_prerequisites():
        print("Running Sound Classification Model")
        print("All prerequisites are met.")
        networking.listen()
    else:
        print("Prerequisites are not met. Please make sure the required files are in the current directory.")
