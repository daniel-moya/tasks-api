# Todo API, a simple but Fast todo api
## Get Started
``` bash
git clone https://github.com/daniel-moya/tasks-api.git
cd todo-api
```

### Install deps
``` bash
go get .
```

### Setup DB
``` bash
./setup
```

### Run server
``` bash
go run .
```

### Get all
``` bash
curl http://localhost:8080/tasks
```

### Get task by Id
``` bash
curl http://localhost:8080/tasks/:id
```

### Add new Task
``` bash
curl http://localhost:8080/tasks \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"title": "New Task", "done": true }'
```
