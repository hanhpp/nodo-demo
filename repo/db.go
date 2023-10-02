package repo

import (
	"fmt"
	"log"
	"net/url"
	"stock-api/global"
	"stock-api/util"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
	DbSSLMode  string
	DbTimeZone string
	LogLevel   logger.LogLevel // "gorm.io/gorm/logger"

	once sync.Once
	db   *gorm.DB
}

// Db returns the gorm db connection.
func (d *Database) Db() *gorm.DB {
	d.once.Do(d.Connect)

	if d.db == nil {
		log.Fatal("entity: database not connected")
	}

	return d.db
}

func (d *Database) SetDB(db *gorm.DB) {
	d.db = db
}

// Connect creates a new gorm db connection.
func (d *Database) Connect() {
	if d == nil {
		log.Fatal("db is nil")
	}

	ourDB, err := gorm.Open(postgres.Open(d.GetDns()), &gorm.Config{
		Logger: logger.Default.LogMode(d.LogLevel),
	})

	if err != nil || ourDB == nil {
		log.Fatal(err)
	} else {
		fmt.Println("Yay! " + d.DbName + " Database Connected!")
		fmt.Println("Database Host: " + d.DbHost)
	}
	d.SetDB(ourDB)
}

// Close closes the gorm db connection.
func (g *Database) Close() {
	if g.db != nil {
		dbInstance, _ := g.db.DB()
		if err := dbInstance.Close(); err != nil {
			fmt.Println("Error while closing the database:", err)
		}

		g.db = nil
	}
}

func (db *Database) GetDns() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
		dnsProcess(db.DbHost),
		dnsProcess(db.DbPort),
		dnsProcess(db.DbUsername),
		dnsProcess(db.DbName),
		db.DbPassword,
		dnsProcess(db.DbSSLMode),
		dnsProcess(db.DbTimeZone),
	)
}

// Setup set up our db with configs from env
func (db *Database) Setup(conf *global.VecConfig) {
	if conf == nil {
		log.Fatal("config is nil")
	}

	// copy information
	db.DbName = conf.DbName
	db.DbHost = conf.DbHost
	db.DbPort = conf.DbPort
	db.DbUsername = conf.DbUsername
	db.DbPassword = conf.DbPassword
	db.DbSSLMode = conf.DbSSLMode
	db.DbTimeZone = conf.DbTimeZone
	db.LogLevel = logger.Silent
	db.once = sync.Once{}

}

func (d *Database) DB() *gorm.DB {
	if d.db == nil {
		d.Connect()
	}

	return d.db
}

func dnsProcess(str string) string {
	return url.QueryEscape(util.TrimSpaceToLower(str))
}

// Init initiate our repositories and database
func Init() {

	// get configs from env
	// set configs to db
	config := global.Config
	// fmt.Printf("%+v\n", config)
	DB.Setup(config)
	// connect to db
	DB.Connect()

	// Init repositories
	InitRepositories(DB)
}

func InitRepositories(db *Database) {
	// Init Repositories
	StockRepoInstance := NewStockRepo(db)

	// Init Server
	Server = NewServer(StockRepoInstance)
}

func DoMigration() {
	DBAutoMigration()
}

func DBAutoMigration() {
	_ = DB.DB().AutoMigrate(
		&Stock{},
	)
}

var DB = &Database{}
var Server = &server{}
