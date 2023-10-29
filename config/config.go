package config

import (
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type envVars struct {
	DbHost string `required:"true" envconfig:"DB_HOST"`
	DbPort string `required:"true" envconfig:"DB_PORT"`
	DbName string `required:"true" envconfig:"DB_NAME"`
	DbUser string `required:"true" envconfig:"DB_USER"`
	DbPass string `required:"true" envconfig:"DB_PASS"`
}

var (
	Tpl *template.Template
	Env envVars
)

func Configure() {

	WrkDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Failed !! Error While Getting Working Directory: ", err)
	}
	log.Println("Success!! Got Working Directory")

	Tpl = template.Must(template.ParseGlob(WrkDir + "/views/*.html"))

	err = godotenv.Load(WrkDir + "/./.env")
	if err != nil {
		log.Fatalln("Failed !! Error While Loading .env File: ", err)
	}
	log.Println("Success!! Loaded .env file")

	err = envconfig.Process("", &Env)
	if err != nil {
		log.Fatalln("Failed !! Error While Processing Environment Variables: ", err)
	}
	log.Println("Success!! Processed Environment Variables")

}
