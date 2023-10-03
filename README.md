# BirdAI_D7057E



## Project structure
| Module   | Destination      |
| -------- |------------------|
| Backend  | `src/internal`   |
| Frontend | `src/frontend`   |
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
|    AI    | `src/AI`         |

# Creating and launching the AI Containers automatically
Create and launch all images/containers
```
 cd ..\BirdAI_D7057E\src\AI
```
```
docker-compose build
docker-compose up
```
# Creating and launching the AI Containers manually
```
 cd ..\BirdAI_D7057E\src\AI
```
Create shared requirements image
```
docker build -t shared-requirements-image .
```
Create the network
```
docker network create my_network
```
Classification Model
```
docker build . -t classification_model --no-cache
docker run -d --network my_network --name classification_model classification_model
```
Detection model
```
docker build . -t detection_model --no-cache
docker run -d --network my_network --name detection_model detection_model
```
AI Controller
```
docker build . -t ai_controller --no-cache
docker run -d --network my_network --name ai_controller ai_controller
```
=======
=======
|    Ai    | `src/AI`         |
>>>>>>> bb530a9 (Update README.md)
=======
|    AI    | `src/AI`         |
>>>>>>> eaf74ed (Update README.md)

# Creating and launching the AI Containers automatically
Create and launch all images/containers
```
 cd ..\BirdAI_D7057E\src\AI
```
```
docker-compose build
docker-compose up
```
# Creating and launching the AI Containers manually
```
 cd ..\BirdAI_D7057E\src\AI
```
Create shared requirements image
```
docker build -t shared-requirements-image .
```
Create the network
```
docker network create my_network
```
Classification Model
```
docker build . -t classification_model --no-cache
docker run -d --network my_network --name classification_model classification_model
```
Detection model
```
docker build . -t detection_model --no-cache
docker run -d --network my_network --name detection_model detection_model
```
AI Controller
```
docker build . -t ai_controller --no-cache
docker run -d --network my_network --name ai_controller ai_controller
<<<<<<< HEAD
>>>>>>> de18963 (Update README.md)
=======
```
>>>>>>> a529ec4 (Update README.md)
