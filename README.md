### Description
> - A simple RSS aggregator backend application. The technologies are
    REST APIs, Go-Chi, Go, PostgreSQL, and Docker. Currently in progress,

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

> - Pulling the postgres from docker.

```bash
docker-compose up --build -d
```

> - Connecting to the database.
```bash
# Connect to aggreflowdb using any database client.
Host: localhost
Port: 5432 
Database: aggreflowdb
Username: postgres
Password: postgres
```

