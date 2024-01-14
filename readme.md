$ go run ./cmd/api


$ curl localhost:8000/v1/healthcheck
```
status: available
env: dev
version: 0.0.1
```

$ curl -X POST localhost:8000/v1/healthcheck
```
Method Not Allowed
```

$ curl  localhost:8000/v1/books/1
```
Display the details of book with ID: 1
```

$ curl -X PUT  localhost:8000/v1/boo
ks/1
```
Update the details of teh book with ID: 1
```

$ curl -X DELETE localhost:8000/v1/b
ooks/1
```
Delete a book with ID: 1
```

$ curl localhost:8000/v1/healthcheck
```
{"environment":"dev","status":"available","version":"0.0.1"}
```

$ curl localhost:8000/v1/healthcheck | jq
```
{
  "environment": "dev",
  "status": "available",
  "version": "0.0.1"
}
```