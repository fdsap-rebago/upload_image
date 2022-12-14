package main

import (
	"log"
	"upload_image/middleware"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	//----- SWAGGER -----
	swagger "github.com/arsmn/fiber-swagger/v2"

	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	_ "upload_image/docs"
)

// @title Image Upload
// @version 1.0
// @description Image Upload Swagger
// @termsOfService http://swagger.io/terms/
// @contact.name FDSAP Support
// @contact.email
// @license.name
// @license.url
// @host 192.168.0.138:4040
// @BasePath /
func main() {
	middleware.PostgreSQLConnect()
	app := fiber.New()

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,locale",
	}))

	app.Use(logger.New())
	AppRoutes(app)

	log.Fatal(app.Listen(":4040"))
}

func AppRoutes(app *fiber.App) {
	// Swagger
	app.Get("/docs/*", swagger.HandlerDefault)
	//-----

	app.Post("/upload", UploadImage)
	app.Post("/fetch", FetchByte)
	app.Get("/fetch_image_data/:img_id", GetImageData)
	app.Get("/get_um/:userInput", GetUserManagement)
}
