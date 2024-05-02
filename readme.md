## Prerequisites
- docker installed
- postgresql installed and running on localhost:5432
- database named *aiplus* is present in postgresql
## Build and run
### Using docker
build
`sudo docker build -t aiplus_demo .`
run image
`docker run -p 8080:8080 --network="host" aiplus_demo`
### Using real machine
build
`go build`
run executable
`./aiplus_demo`
