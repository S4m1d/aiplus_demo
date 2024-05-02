## Prerequisites
- docker installed
- postgresql installed and running on localhost:5432
- database named *aiplus* is present in postgresql
- repository clonned
- IMPORTANT: you added your own .env file into root directory of this project and enriched it in respect with .env.example file
## Build and run
### Using docker
- build
`docker build -t aiplus_demo .`
- run image
`docker run -p 8080:8080 --network="host" aiplus_demo`
### Using real machine
- build
`go build`
- run executable
`./aiplus_demo`
## Test
POST /employees
 body:
```json
{
    "firstName":"John",
    "middleName":"Johnovich",
    "lastName":"Doe",
    "phoneNumber":"88005553535",
    "city":"Bishkek"
}
```
