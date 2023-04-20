The first attempts to write a RESTful api server on golang while maintaining the principle of a clean project architecture.

run postgres db in docker container
```
    docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d postgres
```

connect to container
```
    docker exec -it todo-db /bin/bash
```

create migrations file
```
    migrate create -ext sql -dir ./migrates -seq init
```

up/down migrations
```
    migrate -path ./migrates -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up/down
```

fix dirty migration
    in table `schema_migrations` set `version` on `1` and `dirty` on `false`

