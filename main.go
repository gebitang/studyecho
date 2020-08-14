package main

/**
follow:
https://mp.weixin.qq.com/s/DuEGITdOYHYOXk-v-u2V5A
https://mp.weixin.qq.com/s/vg9OSO4g0KG7iDQ7GXoUSQ
*/
import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *CustomValidator) lazyInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}

func (c *CustomValidator) Validate(i interface{}) error {
	c.lazyInit()
	return c.validate.Struct(i)
}

// demo only. use config or parameter in prod env
var cookieStore = sessions.NewCookieStore([]byte("studyecho"))

func init() {
	rand.Seed(time.Now().UnixNano())

	err := os.Mkdir("log", 0755)
	if err != nil {
		panic(err)
	}
}

func mainOld() {
	// get an instance
	e := echo.New()

	e.Validator = &CustomValidator{}

	// register router
	//e.GET("/", func(context echo.Context) error {
	//	return context.String(http.StatusOK, "Hello, World")
	//})

	//config log
	configLogger(e)

	// register static router
	e.Static("img", "img")
	e.File("/favicon.ico", "img/favicon.ico")

	//register middleware
	setMiddleware(e)

	// register router
	RegisterRoutes(e)

	// start server
	e.Logger.Fatal(e.Start(":2020"))

}

// config log
func configLogger(e *echo.Echo) {
	// set log level
	e.Logger.SetLevel(log.INFO)

	// record business logic
	echoLog, err := os.OpenFile("log/echo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	// print to std and file
	e.Logger.SetOutput(io.MultiWriter(os.Stdout, echoLog))
}

func setMiddleware(e *echo.Echo) {
	// access log
	accessLog, err := os.OpenFile("log/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	// put to std out and file
	middleware.DefaultLoggerConfig.Output = accessLog
	e.Use(middleware.Logger())

	// custom middleware
	e.Use(AutoLogin)

	e.Use(middleware.Recover())
}
