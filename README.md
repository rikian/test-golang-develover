## Setup

GO version 1.18

First
```
git clone git@github.com:rikian/test-golang-developer.git

```
```
go mod tidy
```

# Create .env file
see file .env.eample
```
# config postgres
DB_HOST=localhost
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=db_main
DB_PORT=5432

#JWT SECRET
JWT_SECRET=S4n94t_R4h4S14_BRO...

#DEBUG_MODE
DEBUG=true

# grpc address WITH PORT
GRPC_ADDRESS=127.0.0.1:12345

# migration
DB_MIGRATION=false

# status
# STATUS=DEVELOPMENT
STATUS=PRODUCTION
```

make sure, in the folder config -> config.go. you need to change projectDirName to root your working directory name
```
const projectDirName = "your root working directory name"
```
see --> https://github.com/joho/godotenv/issues/43#issuecomment-503183127

Make migration
```
make migration
```

Run with out running unit test
```
make r
```
run with running test
```
make run
```