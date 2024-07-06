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

### TODOS (Fiber)

##### Create

```bash
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"title":"curl request","body":"Post request"}' \
     http://localhost:4000/api/todos
```

##### Get

```bash
curl -i "http://localhost:4000/api/todos"
```

##### Patch

```bash
curl --header "Content-Type: application/json" \
     --request PATCH \
     "http://localhost:4000/api/todos/1/done
```

### Coins Balance (Chi)

##### Get

```bash
curl --header "Content-Type: application/json" \
     --header "Authorization: 123ABC" \
     --request GET \
     "http://localhost:8000/account/coins/?username=vinicius
```
