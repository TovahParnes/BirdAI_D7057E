# Stage 1: Build the Angular frontend
FROM node:16-alpine AS frontend-builder
WORKDIR /app
COPY src/frontend/package.json src/frontend/package-lock.json ./
RUN npm install
COPY src/frontend/ ./
RUN npm run build

# Stage 2: Build the Golang backend
FROM golang:1.21.1 AS backend-builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o birdai

# Stage 3: Create the final image
FROM alpine:3.14
WORKDIR /app

# Copy the built Angular app from the frontend builder stage
COPY --from=frontend-builder /app/dist /app/src/frontend/dist

# Copy the built Golang binary from the backend builder stage
COPY --from=backend-builder /app/birdai /app

# Copy the .env files
COPY compose/.env /app/.env
COPY secret/.env /app/secret/.env

# Expose the port that your Golang server will run on
EXPOSE 8080

# Set environment variables if needed
# ENV SOME_ENV_VARIABLE=value

# Start your Golang application
CMD ["./birdai"]
