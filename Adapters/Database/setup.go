package database

import (
	"context"
	"log"

	model "github.com/Shubhacker/RTO-BE-Dev/Adapters/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

// database connection

var DbConn model.Conn
var DBConn = model.DBConn{
	Host:     "localhost",
	Port:     "5432",
	User:     "postgres",
	Password: "admin",
	DbName:   "RTO",
	SslMode:  "disable",
}

func DBInit() *model.Conn {
	// e := godotenv.Load()
	// if e != nil {
	// 	fmt.Print(e)
	// }
	dbURI := "host=" + DBConn.Host + " port=" + DBConn.Port + " user=" + DBConn.User + " password=" + DBConn.Password + " dbname= " + DBConn.DbName + " sslmode= " + DBConn.SslMode
	//nolint:gosec // configuration
	// maxOpenConnection, err := strconv.Atoi(dbConnData.MaxConn)
	// if err != nil {
	// 	log.Panic("cannot parse DB_MAX_CONN")
	// }
	//nolint:gocritic // configuration
	// maxIdleTime, err := strconv.Atoi(dbConnData.MaxIdleTime)
	// if err != nil {
	// 	log.Panic("cannot parse DB_MAX_IDLE_TIME")
	// }
	//nolint:gocritic // configuration
	// maxConnectionLifetime, err := strconv.Atoi(dbConnData.MaxLifeTime)
	// if err != nil {
	// 	log.Panic("cannot parse DB_MAX_LIFETIME")
	// }

	//nolint:gocritic // configuration
	// healthCheckPeriod, err := strconv.Atoi(dbConnData.HealthCheckPeriod)
	// if err != nil {
	// 	log.Panic("cannot parse DB_HEALTHCHECK_PERIOD")
	// }
	// log.Println(dbURI, "< --- DB URL")
	config, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		log.Println(err.Error())
		log.Panic("cannot parse database URI")
	}
	// config.MaxConns = int32(maxOpenConnection)
	// config.MaxConnLifetime = time.Duration(maxConnectionLifetime) * time.Minute
	// config.HealthCheckPeriod = time.Duration(healthCheckPeriod) * time.Minute
	// config.MaxConnIdleTime = time.Duration(maxIdleTime) * time.Minute

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Println("Failed to connet DB " + err.Error())
	} else {
		log.Println("DB Connected successfully ")
	}

	DbConn.Db = pool
	DbConn.Schema = "report"
	return &DbConn
}
