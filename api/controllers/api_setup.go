package controllers

import (
	"changeme/api/models"
	"changeme/api/utils"
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/adrg/xdg"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	Name string
	Db  *gorm.DB
	User models.User
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
	InitLogger()
	a.Db = InitSqliteDB(a.Name, a.Name+".db")
	go a.startTimerEmit()
	err := a.Migration()
	if err != nil {
		slog.Error("Migrating DB", "err", err.Error())
	}
}

func (a *App) Migration() error {
	err := a.Db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}

func InitSqliteDB(dirName string, dbName string) *gorm.DB {
	var err error
	dir := xdg.DataHome + "/" + dirName
	err = utils.CreateDir(dir)
	if err != nil {
		slog.Error("Creating Directory", "err", err.Error())
		return nil
	}
	db, err := gorm.Open(sqlite.Open(dir + "/" + dbName), &gorm.Config{})
	if err != nil {
		slog.Error("Opening SqliteDB", "err", err.Error())
		return nil
	}
	slog.Info("SqliteDB Opened", "dir", dir, "db", dbName)
	return db
}

func InitPostgresDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=myhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil {
		slog.Error("Opening PostgresDB", "err", err.Error())
		return nil
	}
	return db
}

func InitLogger() {
	w:= os.Stderr
	// Windows is annoying per usual
	if runtime.GOOS == "windows" {
		slog.SetDefault(slog.New(
			tint.NewHandler(colorable.NewColorable(w), nil),
		),
	)}
	// Every other OS
	slog.SetDefault(slog.New( 
		tint.NewHandler(w, &tint.Options{
			Level: slog.LevelDebug,
			TimeFormat: time.Kitchen,
			NoColor: !isatty.IsTerminal(w.Fd()),
		}),
	))
	slog.Info("Starting Logger...")
}
