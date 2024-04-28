# base image
FROM golang:1.22.2-alpine3.19

# container listens on port 9000
EXPOSE 9000

# update packages in Alpine Linux container, --no-cache means package index won't be cached
RUN apk update && apk add --no-cache mysql-client build-base

# create directory 'app' and make it working directory
RUN mkdir /app
WORKDIR /app

# copy go.mod and go.sum file in 'app' directory
COPY go.mod .
COPY go.sum .

# download dependencies in go.mod and go.sum
RUN go mod download

# copy all content into 'app'
COPY . .

# copy grpc_entrypoint.sh to usr/local/bin and run it
COPY ./grpc_entrypoint.sh /usr/local/bin/grpc_entrypoint.sh
RUN /bin/chmod +x /usr/local/bin/grpc_entrypoint.sh

# build go application
RUN go build cmd/main.go

# move compiled binary to usr/local/bin
RUN mv main /usr/local/bin/

# used to run the container
CMD ["main"]

# entrypoint of container, when the container starts it starts grpc_entrypoint.sh first
ENTRYPOINT ["grpc_entrypoint.sh"]