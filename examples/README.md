A small example to generate realistic Hasura schemas.

## Directory Structure

- `diagrams` contains digrams visually documenting the example schema.
- `hasura` contains Hasura metadata for our example.
- `policies` is where we store logic for our permissions rules.
  - The name and structure of this directory doesn't matter.
  - You can structure with whatever subdirectories you won't. The CLI will
    recursively scan through it and parse any permissions rules.
  - Only the structure of the YAML matters.

## Entity Relationship Diagram

<img src="./diagrams/erd.svg" width="500" />

## Tasks

### start

Start Hasura

```shell
docker compose up --detach
```

### console

Launch the Hasura Console UI

```shell
hasura --project hasura \
  console
```

### migrate_format

Format migrations.

```shell
sleek -n \
  --indent-spaces 2 \
  --uppercase \
  hasura/migrations/**/*.sql
```

### migrate_status

List migrations

```shell
hasura --project hasura \
  migrate \
    --database-name default \
    status
```

### migrate_squash

Squash all migrations

```shell
hasura --project hasura \
  migrate \
    --database-name default \
    squash --delete-source \
    --from 1693763548205 \
    --name init
```
