#This sets the base image of the container to the latest official Go image from Docker Hub.
#It provides a pre-configured environment with Go installed, so you don’t have to install Go manually in the container.
FROM golang:1.23.4 AS builder
# Sets the working directory inside the container to /app.
# All subsequent commands will run from this directory. It helps organize your files and avoids cluttering the root directory.
WORKDIR /app
# Copies go.mod and go.sum files from your local directory to the container’s current working directory (/app).
#These files list dependencies. By copying them first and running go mod download before copying the rest of the code, Docker can cache the downloaded modules and avoid repeating that step unless dependencies change (this speeds up rebuilds).
COPY go.mod go.sum ./
#Downloads the Go dependencies listed in go.mod and go.sum.
#Why: Prepares your build environment with all required packages, ensuring your app compiles with all its dependencies.
RUN go mod download
#Copies all files from your current directory (on host) into the working directory inside the container (/app).
#Transfers your source code into the container so it can be built and run.
COPY . .
# Compiles your Go application and outputs the binary named main.
#Converts your source code into an executable the container can run.
RUN go build -o main .
# Second stage — minimal runtime image
# to make alpine bianry with go binary which i making in above step
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o main . 
FROM arm64v8/alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
#Tells Docker that the container will listen on port 80.
#It's a documentation feature—it does not publish the port. You still need to use -p or --publish when running the container to map it to your local machine.
EXPOSE 80
#Defines the default executable that will run when the container starts.
ENTRYPOINT ["./main"]