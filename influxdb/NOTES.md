# InfluxDB Docker Compose Notes

## Existing Deployment Caution

The root `docker-compose.yml` is a good fresh-stack baseline, but do not run it blindly against an existing server deployment.

The main risk is not that Docker Compose overwrites an existing volume. The main risk is that Compose creates a new empty named volume because the Compose project name or volume names do not match the volumes already used on the server.

For example, this Compose volume declaration:

```yaml
volumes:
  influxdb-data:
```

will usually create or use a volume named like:

```text
thg_influxdb-data
```

If the server currently uses a different volume name, InfluxDB may start with an empty data directory and appear to have lost data, even though the old volume still exists.

## Before Running Compose On The Server

Check existing Docker volumes:

```sh
docker volume ls
```

Check the current InfluxDB container mounts:

```sh
docker inspect <influxdb-container-name>
```

Look for the mounted source used for:

```text
/var/lib/influxdb2
/etc/influxdb2
```

Then update `docker-compose.yml` to use the existing volume names or bind mount paths.

## Using Existing Named Volumes

If the existing server volume is named `influxdb-data`, configure it explicitly:

```yaml
volumes:
  influxdb-data:
    external: true
    name: influxdb-data
```

Do the same for the InfluxDB config volume if one exists:

```yaml
volumes:
  influxdb-config:
    external: true
    name: influxdb-config
```

## Using Existing Host Directories

If the server uses bind mounts instead of Docker named volumes, use the host paths directly:

```yaml
services:
  influxdb:
    volumes:
      - /srv/thg/influxdb/data:/var/lib/influxdb2
      - /srv/thg/influxdb/config:/etc/influxdb2
```

## InfluxDB Init Variables

Do not add `DOCKER_INFLUXDB_INIT_*` variables when attaching to an existing InfluxDB deployment.

Those variables are useful for first-run setup of a fresh InfluxDB instance, but they are unnecessary and potentially confusing for an already initialized server.

The safer existing-deployment service shape is:

```yaml
services:
  influxdb:
    image: influxdb:2.7
    restart: unless-stopped
    ports:
      - "8086:8086"
    volumes:
      - influxdb-data:/var/lib/influxdb2
      - influxdb-config:/etc/influxdb2
```

The `thgsink` service still needs a valid token, organization, and bucket:

```text
THG_INFLUXDB_WRITE_TOKEN
THG_INFLUXDB_ORGANIZATION
THG_INFLUXDB_BUCKET
```

## Validation Before Starting

Render and validate the Compose configuration without starting containers:

```sh
docker compose config
```

Once volume names and mount paths are confirmed, start the stack:

```sh
docker compose up -d --build
```

## Session Resume

OpenCode session ID:

```text
ses_1cd84b5a8ffeBRBXWB6tEAtpYq
```

Resume this conversation from the repo root with:

```sh
opencode --session ses_1cd84b5a8ffeBRBXWB6tEAtpYq
```
