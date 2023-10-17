# BirdAI_D7057E



## Project structure
| Module   | Destination      |
| -------- |------------------|
| Backend  | `src/internal`   |
| Frontend | `src/frontend`   |
|    Ai    | `src/AI`         |

# Creating and launching the AI Containers automatically
Create and launch all images/containers
navigate to
```
 ..\BirdAI_D7057E\src\AI
```
```
docker-compose build
docker-compose up
```
# Creating and launching the AI Containers manually
navigate to
```
 ..\BirdAI_D7057E\src\AI
```
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
