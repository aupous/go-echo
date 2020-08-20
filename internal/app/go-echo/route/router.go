package route

import (
	"log"

	"github.com/aupous/go-echo/internal/app/go-echo/controller"
	myMw "github.com/aupous/go-echo/internal/app/go-echo/middleware"

	// Import postgre driver
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

// Init init routes
func Init() *echo.Echo {
	e := echo.New()

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	db, err := sqlx.Connect("pgx", "postgres://postgres:postgres@localhost:5432/pgx_test?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// Set Custom MiddleWare
	e.Use(myMw.TransactionHandler(db))

	// Routes
	{
		e.POST("/users", controller.CreateUser)
		e.GET("/users", controller.GetUsers)
		e.GET("/users/:id", controller.GetUser)
	}
	return e
}
