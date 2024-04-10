package pkg

type Env struct {
	AppEnv                 string
	PORT                   string
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPass                 string
	DBName                 string
	ContextTimeout         int
	AccessTokenSecret      string
	RefreshTokenSecret     string
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
}

func NewEnv() *Env {
	env := Env{
		AppEnv:                "development",
		PORT:                   "8080",
		DBHost:                 "localhost",
		DBPort:                 "5432",
		DBUser:                 "postgres",
		DBPass:                 "12345",
		DBName:                 "forum",
		ContextTimeout: 30,
		AccessTokenExpiryHour: 24,
		RefreshTokenExpiryHour: 720,
		AccessTokenSecret:      "mysecret",
		RefreshTokenSecret:     "myrefreshsecret",
	}


	return &env
}



// package pkg

// import (
// 	"log"

// 	"github.com/spf13/viper"
// )

// type Env struct {
// 	// ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
// 	AppEnv                 string `mapstructure:"APP_ENV"`
// 	PORT                   string `mapstructure:"PORT"`
// 	DBHost                 string `mapstructure:"DB_HOST"`
// 	DBPort                 string `mapstructure:"DB_PORT"`
// 	DBUser                 string `mapstructure:"DB_USER"`
// 	DBPass                 string `mapstructure:"DB_PASSWORD"`
// 	DBName                 string `mapstructure:"DB_NAME"`
// 	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
// 	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
// 	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
// 	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
// 	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
// }

// func NewEnv() *Env {
// 	env := Env{}
// 	viper.SetConfigFile(".env")

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatal("Can't find the file .env: ", err)
// 	}

// 	err = viper.Unmarshal(&env)
// 	if err != nil {
// 		log.Fatal("Environment can't be loaded: ", err)
// 	}

// 	if env.AppEnv == "development" {
// 		log.Println("The app is running in development env")
// 	}

// 	return &env

// }
