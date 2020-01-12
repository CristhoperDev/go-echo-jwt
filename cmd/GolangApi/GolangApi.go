package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/subosito/gotenv"
	"go-echo-jwt/internal/connection"
	"go-echo-jwt/internal/handler"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var startTime string

//Error log
var (
	publicEndpoints = make(map[string]string)
	Error           *log.Logger
)

//Init app
func Init(errorHandle io.Writer) {
	gotenv.Load()
	rand.Seed(time.Now().UTC().UnixNano())
	//Trace = log.New(traceHandle,
	//  "TRACE: ",
	//  log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func startApplication(c echo.Context) error {
	return c.String(http.StatusOK, "Application stated on PORT: 3000!")
}

func main() {

	// startup code
	startTime = time.Now().String()
	programLog, err := os.OpenFile("GolangApi/GolangApi.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer programLog.Close()
	Init(programLog)
	log.SetOutput(programLog)
	connection.InitDal()

	e := echo.New()
	e.GET("/", startApplication)
	e.POST("/register", handler.RegisterUser)
	e.POST("/login", handler.UserLogin)
	e.POST("/token", handler.RefreshToken)

	r := e.Group("/")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("users", handler.GetUsers)

	e.Logger.Fatal(e.Start(":3000"))
}