# go-simple-inventory

The simple inventory API implemented in Go with Fiber.

## How To Use

There are two branches in this repository:

- `local-storage` for implementation with local storage.
- `main` for implementation with MySQL database.

1. Clone the repository.

2. If `local-storage` branch is chosen, run the application with this command.

```
go run main.go
```

3. If `main` branch is chosen, copy the `.env` file.

```
cp .env.example .env
```

4. Run the MySQL server in the local machine. For Docker user, the commands for running MySQL server can be checked in `commands.sh` file. Then, create a new database.

```sql
CREATE DATABASE inventory;
```

5. Run the application with this command.

```
go run main.go
```


## How to Build with Docker

1. Build the application's image.

```
docker compose build
```

2. Run the application.

```
docker compose up -d
```

3. Stop the application.

```
docker compose down
```