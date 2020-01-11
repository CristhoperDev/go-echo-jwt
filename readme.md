# Golang using echo and jwt
Sources for creating a web server in Golang using the [Echo](https://github.com/labstack/echo) Package.

## Description 
Starting to create a project with connection to database to use echo and jwt

## Structure
```
.
├── cmd                  main applications of the project
│   └── GolangApi        the API server application
├── internal             private application and library code
│   ├── connection       connection to dabase
│   ├── models           entity definitions and domain logic
│   ├── handelr          controller 
│   └── dao              Data Access Object
│   ├── accesslog        access log middleware
│   ├── graceful         graceful shutdown of HTTP server
│   ├── log              structured and context-aware logger
│   └── pagination       paginated list
└── database
```


## Installation
 Clone the project:
 ```
  git clone https://github.com/CristhoperDev/go-echo-jwt.git
 ```
 cd into it:
 ```
  cd go-echo-jwt
 ```

 cd into database:
 ```
  cd database
 ```

 upload database:
 ```
  Uplaod database on mysql
 ```
 
 generate executable:
 ```
  cd cmd/GolangApi
  go build 
 ```
 
 run the server: 
 
 ```
  go run ./GolangApi/
 ```