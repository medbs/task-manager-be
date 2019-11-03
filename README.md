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
* Add TODO:
`curl -d "title=task1" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/api/v1/todos/
`
* Get all TODOs:
`curl -X GET http://localhost:8080/api/v1/todos/
`

* Get one TODO:
`curl -X GET http://localhost:8080/api/v1/todos/{title}
`

* Remove TODO:
`curl -X DELETE http://localhost:8080/api/v1/todos/{title}
`

* Update TODO:
`curl -d "updatedTitle=task1" -H "Content-Type: application/x-www-form-urlencoded" -X PUT http://localhost:8080/api/v1/todos/{title}
`
