package main


import (
	c "envWithViper/config"
	"fmt"
	"github.com/spf13/viper"
)


func main() {
	viper.SetConfigName("config")

	viper.AddConfigPath("./config")

	viper.AutomaticEnv()


	viper.SetConfigType("yml")

	var configurations c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("database.dbname", "test.db")

	err := viper.Unmarshal(&configurations)

	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	fmt.Println("Reading variable using the model")
	fmt.Println("Database is\t", configurations.Database.DBName)
	fmt.Println("Port is\t\t", configurations.Server.Port)
	fmt.Println("EXAMPLE_PATH is\t", configurations.EXAMPLE_PATH)
	fmt.Println("EXAMPLE_VAR is\t", configurations.EXAMPLE_VAR)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Database is\t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))
	fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
	fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_VAR"))
}
