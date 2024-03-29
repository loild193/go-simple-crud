# Simple CRUD using Go

## Stacks:
- Go
- Gin framework
- Air for hot reloading
- Postgres database
- Adminer

## Prerequisite:
- Docker

## How to Start

- Create `.env` from `.env.example` with your own modifications:

```sh
cp .env.example .env
```
Fill the value with your own setup

- Create `.env.production.docker` from `.env.production.docker.example` with your own modifications:

```sh
cp .env.production.docker.example .env.production.docker
```
Fill the value with your own setup

- Grant permission for `docker-entrypoint.sh`:
```sh
chmod +x docker-entrypoint.sh
```

- Run command on `dev` stage:
```sh
docker compose -f docker-compose.dev.yml up -d --build 
```
After that, execute server container with:
```sh
docker compose exec go_simple_crud sh
```
& run below command to start `dev` server:
```sh
air -c .air.toml
```

- Run command on `production` stage:
```sh
docker compose up -d --build
```