### Description
> - A RSS aggregator backend application.

> - The technologies are REST APIs, Go-Chi, Go, PostgreSQL, and Docker.

> - Current `in progress`.

<br />
<br />


### Folder Structuring Guides
```bash
# Directories for CRUD specific queries.
./database/queries/*

# Schema for the tables that are migrated to the database.
./database/schema/*

# Auto generated files using the sqlc.
./database/sqlc/*

# Main code source of the application
./src/*

# Any docker related files to consider when pulling up the image.
./docker/*

# Modules used in this project.
./vendor/*
```

<br />
<br />



### Development Guides
> - Pushing changes to all remote repository.

```bash
# Change mode the permission and make it executable.
chmod +x git-push-all

# Push all changes to remote repositories.
./git-push-all
```

> - Create a `.env` file with the following values below.

```bash
# For server connection.
PORT=13000

# For database connection.
DB_URL=postgresql://username:password@localhost:database_port/database_name?sslmode=disable
```

> - Bash scripts.

```bash
# Build the app.
chmod +x build
./build

# Run the app.
chmod +x app
./app
```

<br />
<br />



### Module Management Guides

```bash
# Installing package.
go get <github/*>

# Make vendored copy of dependencies.
go mod vendor

# Add missing and remove unused modules.
go mod tidy
```

<br />
<br />



### Docker Guides

```bash
# Pull the postgres image.
# "up": Used to start/create containers as defined in the compose file.
# "--build": Build or rebuild docker images that are defined in the compose file.
# "-d": Detaches the container from the terminal.
docker-compose up --build -d

# "docker exec": Used to execute a command in running docker container.
# "-i": Allows you to interact with the container, providing an open standard input (stdin).
# "-t": Allocates a pseudo-TTY, for running interactive commands like a PostgreSQL REPL.
# "psql": PostgreSQL command-line utility to interact with PostgreSQL database.
# "-U": Specifies the username to use when connecting to the databse, which is postgres.
docker exec -it aggreflowdb_container psql -U postgres
```

<br />
<br />



### Database Guides

```bash
# Connect to aggreflowdb using any database client.
Host: localhost
Port: 5432
Database: aggreflowdb
Username: postgres
Password: postgres
```

<br />
<br />



### Goose Guides
> - Managing migrations.

```bash
# Go the schema.
cd database/schema

# Adding migration in database.
goose postgres postgresql://postgres:postgres@localhost:5432/aggreflowdb up

# Dropping migration in database.
goose postgres postgresql://postgres:postgres@localhost:5432/aggreflowdb down
```

<br />
<br />



### SQLC Guides

```bash
# Generating sqlc files.
# Either below.
cd <root-project>
sqlc generate -f "./database/sqlc.yaml"

cd database
sqlc generate
```

