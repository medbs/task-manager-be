# Task-Manager-Be

Task manager example in Go

### Prerequisites
Go 1.11
Docker & Docker Compose to run it as a container

### Run as a container
`docker build -t taskm . && docker run taskm`

`docker-compose up`

### Run as exe
`go build -o app cmd/main.go && ./app`


#### Examples
* Add Task:
`curl -d '{"Title":"t1"}' -H "Content-Type: application/json" -X POST http://localhost:8099/api/v1/tm/`

* Get all Task:
`curl -X GET http://localhost:8099/api/v1/tm/
`
* Get one Task:
`curl -X GET http://localhost:8099/api/v1/tm/{title}
`
* Remove Task:
`curl -X DELETE http://localhost:8099/api/v1/tm/{title}
`
* Update Task:
`curl -d '{"Title":"t2"}' -H "Content-Type: application/json" -X PUT http://localhost:8099/api/v1/tm/{title}
`
