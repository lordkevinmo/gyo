package gyo

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/lordkevinmo/gyo/render"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

type Gyo struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Routes   *chi.Mux
	RootPath string
	Renderer *render.Render
	JetViews *jet.Set
	config   config
}

type config struct {
	port     string
	renderer string
}

func (g *Gyo) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"data",
			"handlers",
			"logs",
			"middlewares",
			"migrations",
			"public",
			"tmp",
			"views",
		},
	}

	err := g.Init(pathConfig)
	if err != nil {
		return err
	}

	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	errorLog, infoLog := g.startLoggers()
	g.ErrorLog = errorLog
	g.InfoLog = infoLog
	g.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	g.Version = version
	g.RootPath = rootPath
	g.Routes = g.routes().(*chi.Mux)

	g.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	g.JetViews = views
	g.createRenderer()

	return nil
}

func (g *Gyo) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := g.createDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gyo) ListAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     g.ErrorLog,
		Handler:      g.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	g.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	g.ErrorLog.Fatal(err)
}

func (g *Gyo) checkDotEnv(path string) error {
	err := g.createFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}
func (g *Gyo) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	return errorLog, infoLog
}

func (g *Gyo) createRenderer() {
	renderer := render.Render{
		Renderer: g.config.renderer,
		RootPath: g.RootPath,
		Port:     g.config.port,
		JetViews: g.JetViews,
	}
	g.Renderer = &renderer
}
