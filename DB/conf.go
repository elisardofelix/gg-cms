package DB

import "os"

var dbConnStr = os.Getenv("GGCMSDBString")
var dbProd = os.Getenv("GGCMSDBProd")
var dbTest = os.Getenv("GGCMSDBTest")


type conf struct {
	connectionString string
	DataBase string
	TestDataBase string
}

var Conf = conf{
	connectionString: dbConnStr,
	DataBase: dbProd,
	TestDataBase: dbTest,
}