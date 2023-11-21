

# Build the Angular frontend
FROM node:16-alpine AS frontend-builder
WORKDIR /app
COPY src/frontend/package.json src/frontend/package-lock.json ./
RUN npm install
COPY src/frontend/ ./
RUN npm run build

# Build the Golang backend
FROM golang:1.21.1 AS backend-builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o birdai

# Create the final image
FROM shared-requirements AS final-image
WORKDIR /app

# Copy the built Angular app from the frontend builder stage
COPY --from=frontend-builder /app/dist /app/src/frontend/dist

# Copy the built Golang binary from the backend builder stage
COPY --from=backend-builder /app/birdai /app

COPY fullchain.pem .
COPY privkey.pem .

# Copy the .env files
COPY compose/.env /app/.env
COPY secret/.env /app/secret/.env

# Expose the port that your Golang server will run on
EXPOSE 443

# AI-controller
WORKDIR /app/ai-controller
COPY src/AI/ai-controller /app/ai-controller
RUN pip install --no-cache-dir -r requirements.txt

# Classification
WORKDIR /app/classification_model
COPY src/AI/classification_model /app/classification_model
RUN apt-get update && apt-get install -y libgl1-mesa-glx && rm -rf /var/lib/apt/lists/*
RUN pip install --no-cache-dir -r requirements.txt

# Detection
WORKDIR /app/detection_model
COPY src/AI/detection_model /app/detection_model
RUN apt-get update && apt-get install -y libgl1-mesa-glx && rm -rf /var/lib/apt/lists/*
RUN pip install --no-cache-dir -r requirements.txt

# Sound
WORKDIR /app/sound_classification
COPY src/AI/sound_classification /app/sound_classification
RUN apt-get update && apt-get install -y libgl1-mesa-glx && rm -rf /var/lib/apt/lists/*
RUN pip install --no-cache-dir -r requirements.txt

# Start your Golang application
WORKDIR /app
CMD ["./birdai"]