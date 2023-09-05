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

## Verifying Permissions

To verify the permissions work as intended, follow the instructions in
[`seed.md`](./seed.md).

Then, using the following headers:

```
x-hasura-role: employee
x-hasura-user-id: 6db66e81-aaf4-4c0d-beaa-020ec8da720d
```

run the following query:

```graphql
query GetFlight {
  flight: flightsByPk(id: "49f4cc96-5c07-462b-ac5b-b32ba89efe3d") {
    id
  }
}
```

You'll notice the data coming back as `null`, since that employee never flew on
that particular flight.

However, if we change up the user to one that was on that particular flight, we
actually get results back:

```
x-hasura-user-id: 1c5bcc1a-94fb-4521-b371-51e74c254a40
```

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

### export

Export Hasura metadata after syncing it.

```shell
hasura --project hasura \
  metadata export
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
