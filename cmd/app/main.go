package main

import (
	"goth-todo/internal/config"
	"goth-todo/internal/db"
	"goth-todo/internal/logging"
	"os"

	"github.com/joho/godotenv"
)

var statusCache map[int]string

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9090
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	godotenv.Load()

	logging.ConfigureLogging()

	db.ConnectDB()
	db.Migrate()

	app := config.RunApp(db.DB)

	app.Router.Run(os.Getenv("PORT"))
}
