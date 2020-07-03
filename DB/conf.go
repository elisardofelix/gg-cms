package DB

type conf struct {
	connectionString string
	DataBase string
	TestDataBase string
}

var Conf = conf{
	connectionString: "mongodb://192.168.50.15:27017",
	DataBase: "test",
	TestDataBase: "test2",
}