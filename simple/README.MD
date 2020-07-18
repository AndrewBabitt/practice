# GO Notes:

### Initial Project:
- `go mod init <repo url>`
    - ref: https://golang.org/doc/code.html#Command

### Running:
- You can run go with two commands:
    - Specify the module you want to run by using the module specifically in the go.mod (this module is github.com/AndrewBabitt/practice/golang/simple)
        - IE: `go run github.com/AndrewBabitt/practice/golang/simple`
    - Run the specific go file:
        - `go run main.go`
    - To use run to specificall run files:
        - `go run main.go constandiota.go collections.go primsandconst.go` Continue adding file names 
        - or run module command that will run everything for you

