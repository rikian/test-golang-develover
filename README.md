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

make sure, in the file config -> config.go. you need to change projectDirName to root your working directory
```
const projectDirName = "your root working directory"
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