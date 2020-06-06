## Simple analytics tracking-like server

1. Build `$ go build -i -v -o build/server cmd/server/main.go`
2. Run `$ ./build/server -h localhost -p 8080`
3. Embed tracking snippet

    ```html
        <img src="http://localhost:8080/hello" border="0" width="1" height="1" >
    ```