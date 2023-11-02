package main
import(
	"os"
)

type AppConfig struct{
	Address string
	Port string
	AppName string
}


func PopulateConfig()(*AppConfig){
	var config AppConfig
	addr,aexists := os.LookupEnv("ADDRESS")
	if !aexists{
		addr = ""
	}
	port,pexists := os.LookupEnv("PORT")
	if !pexists{
		port = "8080"
	}

	config.Address = addr
	config.Port = port
	config.AppName="go crawler"

	return &config

}

type Book struct{
	ID string
	Genre []string
	Author string
	price string
}