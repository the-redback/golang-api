# golang-api

Prerequisite: deploy postgres,

```bash
docker run -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=pass -v $HOME/pg-data:/var/lib/postgresql/data -d postgres
```

To run,

```bash
$ make
```

To build and run,

```bash
$ make build run
```

Push Docker image,

```bash
$ export REGISTRY=<docker-registry-name>
$ make build push
```
