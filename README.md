# BirdAI_D7057E



## Project structure
| Module   | Destination      |
| -------- |------------------|
| Backend  | `src/internal`   |
| Frontend | `src/frontend`   |
|    AI    | `src/AI`         |

# Deploy
## Prerequisites
* If local: add empty `fullchain.pem` & `privkey.pem` to root. **Already exists if in deploy server**
* Add a `/secret/.env` with JWT key in root. *Found in Backend folder on drive.*
* Add `src/AI/classification_model/mobilenet_model.keras` & `src/AI/classification_model/labels.json` *Found in AI folder on drive.*

Run `docker-compose up -d` and visit [local](localhost:443) if you are running local and [deploy](https://birdai.duckdns.org/) if you are running on deploy server.

The deployment docker-compose will run the latest stable tag **unless** you change image to `bird_ai` and the same in each of the AI deploy folders.


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
