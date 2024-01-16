
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

## json.MarshalIndent()
```
$ curl localhost:8000/v1/books
[
        {
                "id": 1,
                "title": "The Darkening of Tristram",
                "published": 1998,
                "pages": 310,
                "genres": [
                        "Fiction",
                        "Thriller"
                ],
                "rating": 4.8
        },
        {
                "id": 2,
                "title": "The Legacy of Deckard Cain",
                "published": 2007,
                "pages": 432,
                "genres": [
                        "Fiction",
                        "Adventure"
                ],
                "rating": 4.4
        }
]
```

## Testing all endpoints at once using bash script
```
$ ./test-endpoints.sh 
curl localhost:8000/v1/healthcheck

{
        "environment": "dev",
        "status": "available",
        "version": "0.0.1"
}

curl localhost:8000/v1/books/1

{
        "id": 1,
        "title": "Echoes in the Darkness",
        "published": 2019,
        "pages": 300,
        "genres": [
                "Fiction",
                "Thriller"
        ],
        "rating": 4.5
}

curl localhost:8000/v1/books

[
        {
                "id": 1,
                "title": "The Darkening of Tristram",
                "published": 1998,
                "pages": 310,
                "genres": [
                        "Fiction",
                        "Thriller"
                ],
                "rating": 4.8
        },
        {
                "id": 2,
                "title": "The Legacy of Deckard Cain",
                "published": 2007,
                "pages": 432,
                "genres": [
                        "Fiction",
                        "Adventure"
                ],
                "rating": 4.4
        }
]
```

## POST request using custom BODY
```
BODY='{"title":"The Black Soulstone","published":2001,"pages":207,"genres":["Fiction","Mystery"],"rating":3.5}'

curl -i -d "$BODY" localhost:8000/v1/books

HTTP/1.1 200 OK
Date: Tue, 16 Jan 2024 00:49:31 GMT
Content-Length: 53
Content-Type: text/plain; charset=utf-8

{The Black Soulstone 2001 207 [Fiction Mystery] 3.5}
```


# Using Postgres SQL Backend
## Create New Book Record
```
$ BODY='{"title":"The Black Soulstone","published":2001,"pages":207,"genres":["Fiction","Mystery"],"rating":3.5}'
echo -e '\ncurl -i -d "$BODY" localhost:8000/v1/books\n'
curl -i -d "$BODY" localhost:8000/v1/books

curl -i -d "$BODY" localhost:8000/v1/books

HTTP/1.1 201 Created
Content-Type: application/json
Location: v1/books/1
Date: Tue, 16 Jan 2024 22:15:02 GMT
Content-Length: 164

{
        "book": {
                "id": 1,
                "title": "The Black Soulstone",
                "published": 2001,
                "pages": "207",
                "genres": [
                        "Fiction",
                        "Mystery"
                ],
                "rating": 3.5
        }
}
```

## GET All Book Records
```
$ curl localhost:8000/v1/books
{
        "books": [
                {
                        "id": 1,
                        "title": "The Black Soulstone",
                        "published": 2001,
                        "pages": "207",
                        "genres": [
                                "Fiction",
                                "Mystery"
                        ],
                        "rating": 3.5
                }
        ]
}
```


## Update One Book Record
```
$ curl localhost:8000/v1/books
{
        "books": [
                {
                        "id": 1,
                        "title": "The Black Soulstone",
                        "published": 2001,
                        "pages": "207",
                        "genres": [
                                "Fiction",
                                "Mystery"
                        ],
                        "rating": 3.5
                },
                {
                        "id": 2,
                        "title": "Echoes in the Darkness",
                        "published": 2019,
                        "pages": "300",
                        "genres": [
                                "Fiction",
                                "Thriller"
                        ],
                        "rating": 4.3
                }
        ]
}
```