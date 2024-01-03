package controllers

import (
	"changeme/api/models"
	"changeme/api/utils"
	"context"
	"log"

	"github.com/adrg/xdg"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	Name string
	Db  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp(name string) *App {
	return &App{
		Name: name,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.startTimerEmit()
	a.Db = InitSqliteDB(a.Name, "changeme.db")
	err := a.Migration()
	if err != nil {
		log.Println("Error Migrating DB:", err.Error())
	}
}

func (a *App) Migration() error {
	err := a.Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Error Migrating DB:", err.Error())
		return err
	}
	return nil
}

func InitSqliteDB(dirName string, dbName string) *gorm.DB {
	var err error
	dir := xdg.DataHome + "/" + dirName
	err = utils.CreateDir(dir)
	if err != nil {
		log.Println("Error Creating Directory:", err.Error())
		return nil
	}
	db, err := gorm.Open(sqlite.Open(dir + "/" + dbName), &gorm.Config{})
	if err != nil {
		log.Println("Error Opening SqliteDB:", err.Error())
		return nil
	}
	return db
}

func InitPostgresDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=myhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil {
		log.Println("Error Opening PostgresDB:", err.Error())
		return nil
	}
	return db
}
