# task-manager-be

#### Add TODO
`curl -d "title=task1" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/api/v1/todos/
`
#### Get all TODOs
`curl -X GET http://localhost:8080/api/v1/todos/
`

#### Get one TODO
`curl -X GET http://localhost:8080/api/v1/todos/{title}
`

#### Remove TODO
`curl -X DELETE http://localhost:8080/api/v1/todos/{title}
`

#### Remove TODO
`curl -d "updatedTitle=task1" -H "Content-Type: application/x-www-form-urlencoded" -X PUT http://localhost:8080/api/v1/todos/{title}
`
