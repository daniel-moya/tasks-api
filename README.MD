### Get all
``` bash
curl http://localhost:8080
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
--data '{"title": "Ta", "done": true }'
```
