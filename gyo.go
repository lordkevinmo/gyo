package gyo

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const version = "1.0.0"

type Gyo struct {
	AppName  string
	Debug    bool
	Version  string
	errorLog *log.Logger
	infoLog  *log.Logger
	rootPath string
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
	g.errorLog = errorLog
	g.infoLog = infoLog
	g.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	g.Version = version

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
