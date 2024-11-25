package config

import (
	"os"

	"github.com/aiteung/atdb"
)

var IteungIPAddress string = os.Getenv("ITEUNGBEV1")

var MongoString string = os.Getenv("MONGOSTRING")

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "data_mahasiswa",
}

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)

