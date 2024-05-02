# Start from golang:1.12-alpine base image
FROM golang:1.22-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git

# Set the Current Working Directory inside the container
RUN mkdir /app
WORKDIR /app

COPY . .

RUN go get -d .

RUN go mod download

# Build the Go app
RUN go build -o /aiplus_demo

EXPOSE 8080

# Run the executable
CMD ["/aiplus_demo"]
