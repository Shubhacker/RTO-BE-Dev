package model

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// host= localhost port= 5432 user= postgres password= admin dbname= RTO sslmode= disable
type Conn struct {
	Db     *pgxpool.Pool
	Schema string
}

type DBConn struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

type CreateRequestDBModel struct {
	ReportId      string
	ImageUrl      []string
	OffenseList   []string
	ByRTO         bool
	VehicleNumber string
	Location      string
	Comment       string
	ReportedBy    string
	Social        bool
	TotalFine     int64
	RTOApproved   bool
}
