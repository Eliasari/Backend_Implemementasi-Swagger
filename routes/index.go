// package routes

// import (
// 	"database/sql"

// 	"github.com/gofiber/fiber/v2"
// )

// // pastikan ini ada di file routes/alumni.go & routes/pekerjaan.go
// func RegisterRoutes(app *fiber.App, db *sql.DB) {
// 	AlumniRoutes(app, db)
// 	PekerjaanRoutes(app, db)
// 	AuthRoutes(app, db) // kasih db juga
// 	UserRoutes(app, db)
// }

package routes

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gofiber/fiber/v2"
)

// @title API Alumni Management
// @version 1.0
// @description API untuk mengelola data alumni, user, dan pekerjaan alumni
// @host localhost:3000
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func RegisterRoutes(app *fiber.App, db *mongo.Database) {

    AuthRoutes(app, db)
    UserRoutes(app, db)
    AlumniRoutes(app, db)
    PekerjaanRoutes(app, db)
    FileRoute(app, db)
}

