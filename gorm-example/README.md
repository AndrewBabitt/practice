# GORM Example with Gin
- GORM -> https://gorm.io/
- Gin -> https://gin-gonic.com/

## Run 
1. first start local DB container:
- `docker run --name temp-postgres -e POSTGRES_PASSWORD=secretpassword -p 5432:5432 -d postgres`
  - this will pull the postgres image from docker and creat a docker image with the a postgres DB that has a password of `secretpassword`. Test you are able to connect to it in the terminal by running `psql -h localhost -U postgres` and entering the password

2. run cmd: `go run main.go` 
    - this will connect to the DB and create the user table if it doesn't exist 


## Simple Curl cmds:
- create user: 
```
curl --location --request POST 'localhost:8080/user/new' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "MJ1996",
    "email": "FlyKicks@mail.com"
}'
```
- find created user: `curl --location --request GET 'http://localhost:8080/user/1'`