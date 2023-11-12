### Helpful Docker Information

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
