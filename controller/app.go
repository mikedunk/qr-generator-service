package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudinary/cloudinary-go"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/mikedunk/qr-generator-service/config"
	"github.com/mikedunk/qr-generator-service/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// App export
type App struct {
	Router       *mux.Router
	Config       *config.Config
	Con          *gorm.DB
	CloudService *cloudinary.Cloudinary
}

func NewApp() *App {
	return &App{}
}

func (app *App) initializeConfig() {

	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config file : ", err)
	}

	app.Config = &config

}

func (app *App) initializeDb() {

	dbURI := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True", app.Config.DBUsername, app.Config.DBPass, app.Config.DBName)
	conn, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic(" error connecting to database")
	}
	log.Info("db Connection Established")

	if err := conn.AutoMigrate(&model.Qr{}); err != nil {
		log.Fatal(err)
	}

	app.Con = conn

}
func (app *App) initializeRouter() {
	app.Router = mux.NewRouter()
	router := app.Router
	router.Use(mux.CORSMethodMiddleware(router))

	log.Info("Router initialized")
}

func (app *App) initializeCloudServices() error {

	cld, err := cloudinary.NewFromURL(app.Config.CloudUrl)

	fmt.Printf("Typle of cld: %T", cld)

	if err != nil {
		log.Info("Unable to connect to cloud service", err)
		return err
	}

	app.CloudService = cld

	fmt.Println(cld)
	log.Info("Successfu)lly Acquired cloud connection")
	//app.UploadToCloud()
	return nil
}

func (app *App) run() {

	srv := &http.Server{
		Handler:      app.Router,
		Addr:         fmt.Sprintf(":%s", app.Config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Application running on Port ", app.Config.Port)
	log.Fatal(srv.ListenAndServe())
}

func (app *App) StartApplication() {

	app.initializeConfig()
	app.initializeDb()
	app.initializeRouter()
	app.AddQrRoutes()
	app.initializeCloudServices()
	app.run()
}

func (app *App) GetDbConnection() *gorm.DB {
	return app.Con
}
