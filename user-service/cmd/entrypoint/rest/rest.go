package rest

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/go-workshop/user-service/cmd/config"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/time/rate"
)

func StartRest(mongoClient *mongo.Client) {
	log.Info("calling fii endpoints")
	e := echo.New()
	e.Static("/views/static", "views/static")
	// e.Use(middleware.Static("/views/static"))

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))

	e.Renderer = newTemplate()

	setupUserApiEndpoints(e)

	port := config.Env.RestPort

	if port == "" {
		port = ":33001"
	}

	e.Start(fmt.Sprintf(":%s", port))
}

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {

	dirs := []string{
		"views/*.html",
		"views/*/*.html",
		"views/*/*/*.html",
	}

	files := []string{}
	for _, dir := range dirs {
		ff, err := filepath.Glob(dir)
		if err != nil {
			panic(err)
		}
		files = append(files, ff...)
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &Templates{
		templates: t,
		// templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func marshalToJS(name string, object any) (map[string]template.JS, error) {
	json, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	return map[string]template.JS{name: template.JS(json)}, nil
}
