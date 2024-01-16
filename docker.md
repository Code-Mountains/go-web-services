$ docker pull postgres
```
Using default tag: latest
latest: Pulling from library/postgres
2f44b7a888fa: Pull complete 
6d49150dabe2: Pull complete 
18d6a86d0fbf: Pull complete 
4c9385c30bce: Pull complete 
550091272acc: Pull complete 
2720859ac49e: Pull complete 
b8091cf53545: Pull complete 
f3ca5fbd89cd: Pull complete 
22fbbce47a56: Pull complete 
b3b5e3b65b95: Pull complete 
917e5b76e085: Pull complete 
7f21ce9572c6: Pull complete 
4ea3941c8572: Pull complete 
848fee03c034: Pull complete 
Digest: sha256:49c276fa02e3d61bd9b8db81dfb4784fe814f50f778dce5980a03817438293e3
Status: Downloaded newer image for postgres:latest
docker.io/library/postgres:latest
```

$ docker run --name reading-list-db-container -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres
```
770ca7bd112f4f5ca1bcd4913efbbf3aa8f9f132fdbdc173404317524091b031
```

$ docker ps
```
CONTAINER ID   IMAGE      COMMAND                  CREATED          STATUS          PORTS                                       NAMES
770ca7bd112f   postgres   "docker-entrypoint.s…"   44 seconds ago   Up 43 seconds   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   reading-list-db-container
```

## PSQL Client
```
$ sudo apt-get install postgresql-client-common -y
$ sudo apt-get install -y postgresql-client
```

## Export PSQL password to .bashrc
```
export PGPASSWORD='mysecretpassword' >> .bashrc
```

$ psql -h localhost -p 5432 -U postgres
```
psql (14.10 (Ubuntu 14.10-0ubuntu0.22.04.1), server 16.1 (Debian 16.1-1.pgdg120+1))
WARNING: psql major version 14, server major version 16.
         Some psql features might not work.
Type "help" for help.

postgres=# exit
```