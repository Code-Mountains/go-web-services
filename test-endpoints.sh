#!/bin/bash
echo -e 'curl localhost:8000/v1/healthcheck\n'
curl localhost:8000/v1/healthcheck

echo -e '\ncurl localhost:8000/v1/books/1\n'
curl localhost:8000/v1/books/1

echo -e '\ncurl localhost:8000/v1/books\n'
curl localhost:8000/v1/books
