## CRUD generator

### Usage

Suppose you want to write CRUD APIs for the `us_users` table in the database, and specify the corresponding model name as `User` in the code, then proceed as follows.

First, specify the database configuration in `.env`, and then execute in the project root directory:

```bash
bin/go run ./cmd/gencode -t us_users -m User
```

Command line parameter meaning:

```
-t: database table name, example: us_users
-m: model name to generate, example: User
```