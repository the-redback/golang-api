# golang-api

Postgres deploy,

```bash
docker run -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=pass -v $HOME/pg-data:/var/lib/postgresql/data -d postgres
```

Connect to posrgres via cli,

```bash
psql --host 0.0.0.0 --port 5432 postgres -U postgres -W 
```

Create table,

```bash
CREATE TABLE conways (
    id bigint NOT NULL,
    x_axis integer,
    y_axis integer,
    grid text
);
```


