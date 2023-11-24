package main

import (
	"log"
	"webhook/app/controller"
	_ "webhook/docs" // load API Docs files (Swagger
	"webhook/pkg/apiserver"
	"webhook/pkg/configs"
	"webhook/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

// @title webhook API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email venkateshwarachinnasamy@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
// @in header
func main() {

	config, err := configs.GetConfig()
	if err != nil {
		log.Println("Failed to load config! Reason: ", err)
		return
	}
	c := configs.GetFiberConfig(config.Server)
	// Define a new Fiber app with config.
	app := fiber.New(c)

	h := controller.NewHandler(config)

	routes.RegisterSwaggerRoute(app)
	routes.RegisterRoutes(app, h)

	apiserver.StartFiberWithGracefulShutdown(app, config.Server)

}
