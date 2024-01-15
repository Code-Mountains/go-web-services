$ go run ./cmd/api/
```
2024/01/14 17:48:52 starting dev server on port :8000
```

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

$ curl localhost:8000/v1/books | jq
```
[
  {
    "ID": 1,
    "CreatedAt": "2024-01-14T18:18:25.784315425-06:00",
    "Title": "The Darkening of Tristram",
    "Published": 1998,
    "Pages": 310,
    "Genres": [
      "Fiction",
      "Thriller"
    ],
    "Rating": 4.8,
    "Version": 1
  },
  {
    "ID": 2,
    "CreatedAt": "2024-01-14T18:18:25.784315475-06:00",
    "Title": "The Legacy of Deckard Cain",
    "Published": 2007,
    "Pages": 432,
    "Genres": [
      "Fiction",
      "Adventure"
    ],
    "Rating": 4.4,
    "Version": 1
  }
]
```