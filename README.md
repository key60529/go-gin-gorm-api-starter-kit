# go-gin-gorm-api-starter-kit

> A golang API which include gin and gorm packages starter kit.



## Requirements

- go: go 1.9 or newer
- dep: stable 0.4.1 or newer

#### Packages

- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/jinzhu/gorm)



## Usage

#### Database (MySQL)

```go
// app/database/db.go
...
const (
  //DbHost for example 127.0.0.1:3306
  DbHost = ""
  //DbUser username to login your mysql
  DbUser = ""
  //DbPassword mysql password
  DbPassword = ""
  //DbName database name
  DbName = ""
)
...
```



#### Serve

```sh
cd your-path/go-gin-gorm-api-starter-kit/
go run main.go
```

