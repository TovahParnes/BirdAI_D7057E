FROM golang:1.21.1

# Set destination for COPY
WORKDIR /app

ADD . /app

RUN go mod download

RUN go env -w GO111MODULE=on

RUN go build -o /BirdAI_D7057E

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/BirdAI_D7057E"]
