#!/bin/bash
# echo -e 'curl localhost:8000/v1/healthcheck\n'
# curl localhost:8000/v1/healthcheck

# echo -e '\ncurl localhost:8000/v1/books/1\n'
# curl localhost:8000/v1/books/1

# echo -e '\ncurl localhost:8000/v1/books\n'
# curl localhost:8000/v1/books

BODY='{"title":"The Black Soulstone","published":2001,"pages":207,"genres":["Fiction","Mystery"],"rating":3.5}'
echo -e '\ncurl -i -d "$BODY" localhost:8000/v1/books\n'
curl -i -d "$BODY" localhost:8000/v1/books

BODY='{"title":"Echoes in the Darkness","published":2019,"pages":310,"genres":["Fiction","Thriller"],"rating":4.1}'
echo -e '\ncurl -X PUT -d "$BODY" localhost:8000/v1/books/2\n'
curl -X PUT -d "$BODY" localhost:8000/v1/books/2