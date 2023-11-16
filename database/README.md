### Folder Purposes
```bash
# Directories for CRUD specific queries.
database/queries/*

# Schema for the tables that are migrated to the database.
database/schema/*

# Auto generated files using the sqlc.
database/sqlc/*
```

<br />
<br />
<br />



### Goose Guides
> - Managing migrations.

```plaintext
# Go the schema.
cd database/schema

# Adding migration in database.
goose postgres postgresql://postgres:postgres@localhost:5432/aggreflowdb up

# Dropping migration in database.
goose postgres postgresql://postgres:postgres@localhost:5432/aggreflowdb down
```

<br />
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

