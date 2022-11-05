package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/simpleittools/gormjwtapi/internal/database"
	"github.com/simpleittools/gormjwtapi/internal/models"
	"log"
	"net/http"
	"os"
	"strconv"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	user     models.User
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	var cfg config
	cfg.port, err = strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		fmt.Println(err)
	}
	database.Conn()

	//dsn := os.Getenv("DSN")
	//environment := os.Getenv("ENV")
	//db, err := driver.ConnectPostgres(dsn)
	//
	//if err != nil {
	//	log.Fatal("Cannot connect to database")
	//}
	//defer db.SQL.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		//models:      data.New(db.SQL),
		//environment: environment,
	}

	err = app.serve()

	if err != nil {
		log.Fatal(err)
	}

}

// serve starts the API web server
func (app *application) serve() error {

	app.infoLog.Println(fmt.Sprintf("API Listening at http://localhost:%d", app.config.port))

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port),
		//Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
