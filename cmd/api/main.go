package main

import (
	"fmt"
	"github.com/joho/godotenv"
	auth "github.com/korylprince/go-ad-auth/v3"
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

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		//models:      data.New(db.SQL),
		//environment: environment,
	}

	authType := os.Getenv("AUTHTYPE")
	switch authType {
	case "INTERNAL":
		app.infoLog.Println("auth is internal")
	case "LDAP":
		// todo configure the application to use LDAP, currently this is just setup a a framework
		portNumber, _ := strconv.Atoi(os.Getenv("LDAPPORT"))
		authConfig := &auth.Config{
			Server:   os.Getenv("LDAPSERVER"),
			Port:     portNumber,
			BaseDN:   os.Getenv("LDAPBASEDN"),
			Security: auth.SecurityStartTLS,
		}
		username := os.Getenv("LDAPUSERNAME")
		password := os.Getenv("LDAPPASSWORD")

		status, err := auth.Authenticate(authConfig, username, password)

		if err != nil {
			//handle err
			return
		}

		if !status {
			//handle failed authentication
			return
		}
	default:
		app.errorLog.Fatal("authentication type improperly configured. options are LDAP or INTERNAL")
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
