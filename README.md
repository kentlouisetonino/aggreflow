### Description
> - A simple RSS aggregator application.

<br />
<br />



### Development
> - Bash scripts.
```bash
# Build the app.
chmod +x build
./build

# Run the app.
chmod +x app
./app
```

> - Helpful guides for module management.
```bash
# Installing package.
go get <github/*>

# Make vendored copy of dependencies.
go mod vendor

# Add missing and remove unused modules.
go mod tidy
```

> - Help docker commands.
```bash
# "docker exec": Used to execute a command in running docker container.
# "-i": Allows you to interact with the container, providing an open standard input (stdin).
# "-t": Allocates a pseudo-TTY, for running interactive commands like a PostgreSQL REPL.
# "psql": PostgreSQL command-line utility to interact with PostgreSQL database.
# "-U": Specifies the username to use when connecting to the databse, which is postgres.
docker exec -it <container-name> psql -U postgres
```
