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

> - Helpful [docker documentation](https://github.com/kentlouisetonino/aggreflow/blob/develop/docs/docker.md).

```bash
# Connect to aggreflowdb using any database client.
Host: localhost
Port: 5432 
Database: aggreflowdb
Username: postgres
Password: postgres
```

