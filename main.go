package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/guidewire/fern-reporter/pkg/graph/generated"
	"github.com/guidewire/fern-reporter/pkg/graph/resolvers"
	"gorm.io/gorm"
	"html/template"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/guidewire/fern-reporter/config"
	"github.com/guidewire/fern-reporter/pkg/api/routers"
	"github.com/guidewire/fern-reporter/pkg/db"
	"time"

	"embed"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed pkg/views/test_runs.html
var testRunsTemplate embed.FS

func main() {
	initConfig()
	initDb()
	initServer()
}

func initConfig() {
	if _, err := config.LoadConfig(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func initDb() {
	db.Initialize()
}

func initServer() {
	serverConfig := config.GetServer()
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "ACCESS_TOKEN"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	funcMap := template.FuncMap{
		"CalculateDuration": CalculateDuration,
	}
	templ, err := template.New("").Funcs(funcMap).ParseFS(testRunsTemplate, "pkg/views/test_runs.html")
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}
	router.SetHTMLTemplate(templ)

	// router.LoadHTMLGlob("pkg/views/*")
	routers.RegisterRouters(router)

	router.POST("/query", GraphqlHandler(db.GetDb()))
	router.GET("/", PlaygroundHandler("/query"))
	err = router.Run(serverConfig.Port)
	if err != nil {
		log.Fatalf("error starting routes: %v", err)
	}

}

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GraphqlHandler(gormdb *gorm.DB) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{DB: gormdb}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func CalculateDuration(start, end time.Time) string {
	duration := end.Sub(start)
	return duration.String() // or format as needed
}
