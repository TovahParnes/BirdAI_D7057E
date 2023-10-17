# BirdAI_D7057E



## Project structure
| Module   | Destination      |
| -------- |------------------|
| Backend  | `src/internal`   |
| Frontend | `src/frontend`   |
| AI | `src/AI`   |



# Setup & Run AI

## Prerequisites
Docker Desktop
A terminal like git bash or other CLI

## Foreword
I would recommend starting the AI Controller at the last stage. It might still work, but it's not tested.
Start by building all modules, then start the AI modules and then start the AI_Controller module.

## 1. Create a docker network
Open a bash terminal or other cli, then write `docker network create my_network
` to create a new network

## 2. Setup & Run Detection model
Clone down the repository from GitHub, then open a bash terminal or other CLI, then navigate to project/AI/detection_model, then run the following command to build the model.

    docker build . -t detection_model --no-cache

Once built, then run bash terminal or other CLI with the following command while still inside the detection_model project path.

    docker run -d --network my_network --name detection_model detection_model
## 3. Setup & Run Classification model
Clone down the repository from GitHub, then open a bash terminal or other CLI, then navigate to project/AI/classification_model, then run the following command to build the model.

    docker build . -t classification_model --no-cache

Once built, then run bash terminal or other CLI with the following command while still inside the classification_model project path.

    docker run -d --network my_network --name classification_model classification_model
## 4. Setup & Run AI Controller
Clone down the repository from GitHub, then open a bash terminal or other CLI, then navigate to project/AI/ai_controller, then run the following command to build the model.

    docker build . -t ai_controller --no-cache

Once built, then run bash terminal or other CLI with the following command while still inside the ai_controller project path.

    docker run -d --network my_network --name ai_controller ai_controller
