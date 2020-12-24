# pg-backup

Pg-backup is a simple tool to export pg dumps to different targets. The idea behind the tool is to bootstrap a simple server that periodically executes `pg_dump` and upload the content to a target destination

Currently, the tool support two main commands: `export` and `start`.

`export` allows to run a single pg dump and export the result.
`start` bootstrap a cron that basically executes the export command periodically

## Configuration

The tool is implemented using [cobra](https://github.com/spf13/cobra), therefore it is customizable through `env variables`, `flags` or using a `config file`.
You can find the help output of the tool below and the list of available flags.

In order to run the tool using `env variables`, they should be in uppercase format.

For example, in order to run the export command specifying output extension you can run:

> OUTPUT_EXT=psql pg-backup export

or using the flag:

> pg-backup export --output_ext psql

In the same way, you can run the cron with a different crontime in one of the following ways:
> CRONTIME="@daily" pg-backup start

or

> pg-backup start --crontime @daily

## Options

``` txt
Backup service to export pg dump to external storage

Usage:
  backup [command]

Available Commands:
  export      Export dump
  help        Help about any command
  start       Export dump cron
  version     backup version

Flags:
      --aws_access_key string   Sets AWS access key
      --aws_s3_bucket string    Sets AWS S3 bucket (default "bucket")
      --aws_s3_region string    Sets AWS S3 region (default "us-east-1")
      --aws_secret_key string   Sets AWS secret key
      --crontime string         Sets crontime - [default: @daily] (default "@daily")
      --debug                   Sets logging level to Debug
  -h, --help                    help for backup
      --output_ext string       Sets output extension - [default: psql] (default "psql")
      --output_prefix string    Adds prefix to output name.
      --output_time             Sets if output name should include time (default true)
      --pg_db string            Sets postgres db (default "postgres")
      --pg_host string          Sets postgres host (default "localhost")
      --pg_pass string          Sets postgres password
      --pg_port string          Sets postgres port (default "5432")
      --pg_user string          Sets postgres user (default "postgres")
      --target string           Sets export target between s3, file - [default: s3]

Use "backup [command] --help" for more information about a command.
```

## Docker

The repo includes a docker file to run the service in an alpine container. We suggest to use this in a docker-compose file like following:

```yaml
version: "3.8"

x-pg-env: &pg-env
  POSTGRES_USER: &pg-user test
  POSTGRES_PASSWORD: &pg-pass test
  POSTGRES_DB: &pg-db test

x-pg-backup-env: &pg-backup-env
  CRONTIME: "@daily"

services:
  postgres:
    image: postgres
    container_name: postgres
    restart: always
    environment:
      <<: *pg-env
    ports:
      - 5432:5432
    volumes:
      - data:/var/lib/postgresql/data

  pg-backup:
    container_name: pg-backup
    build:
      context: .
      dockerfile: ./Dockerfile
    restart: always
    environment:
      <<: [*pg-env, *pg-backup-env]
    volumes:
      - ./backups:/root/backups

volumes:
  data:
```
