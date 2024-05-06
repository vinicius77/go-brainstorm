#### Initiate Project

```bash
go mod init github.com/vinicius77/go-brainstorm
```

#### Install Fiber V2

```bash
go get -u github.com/gofiber/fiber/v2
```

#### Install modules

```bash
go mod tidy
```

#### Run the application

```bash
go run cmd/api/main.go
```

##### CURL Examples

##### Create

```bash
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"title":"curl request","body":"Post request"}' \
     http://localhost:4000/api/todos
```

#### Get

```bash
curl -i "http://localhost:4000/api/todos"
```

#### Patch

```bash
curl --header "Content-Type: application/json" \
     --request PATCH \
     "http://localhost:4000/api/todos/1/done
```
