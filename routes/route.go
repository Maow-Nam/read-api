package routes

import (
	"github.com/Maow-Nam/read-api/api"
	"github.com/labstack/echo"
)

// Init initialize api routes and set up a connection.
func Init(e *echo.Echo) {
	// Database connection.
	db, err := api.NewMongoDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	a := &api.MongoDB{
		Conn: db.Conn,
		BCol: db.BCol,
	}

	e.POST("/postbook", a.PostBook)

}
