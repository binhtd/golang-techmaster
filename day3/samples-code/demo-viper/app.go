package main

import (
	c "demo-viper/config"
	"demo-viper/util"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	//demoViper1()
	//demoViper2()
	//demoViper3()
	//demoViper4()
	//demoViper5()
	demoViper6()
}

//https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d
//https://github.com/techschool/simplebank

func demoViper1() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println("DBDriver:", config.DBDriver)
	fmt.Println("DBSource:", config.DBSource)
	fmt.Println("ServerAddress:", config.ServerAddress)
}

//https://medium.com/@jomzsg/the-easy-way-to-handle-configuration-file-in-golang-using-viper-6b3c88d2ee79
func demoViper2() {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config default \n", err)
		os.Exit(1)
	}

	env := viper.GetString("app.env")
	producerbroker := viper.GetString("app.producerbroker")
	consumerbroker := viper.GetString("app.consumerbroker")
	linetoken := viper.GetString("app.linetoken")

	fmt.Println("---------- Example ----------")
	fmt.Println("app.env :", env)
	fmt.Println("app.producerbroker :", producerbroker)
	fmt.Println("app.consumerbroker :", consumerbroker)
	fmt.Println("app.linetoken :", linetoken)
}

//https://medium.com/@bnprashanth256/reading-configuration-files-and-environment-variables-in-go-golang-c2607f912b63
func demoViper3() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error read config file, %s", err)
	}

	viper.SetDefault("database.dbname", "test_db")
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is \t", configuration.Database.DBName)
	fmt.Println("Port is\t\t", configuration.Server.Port)
	fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
	fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)

	//read variables without using model
	fmt.Println("Reading variables without using the model..")
	fmt.Println("Database is \t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))
	fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
	fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_VAR"))
}

//https://blog.gopheracademy.com/advent-2014/configuration-with-fangs/
func demoViper4() {
	viper.BindEnv("port")
	viper.BindEnv("name", "USERNAME")

	os.Setenv("PORT", "81")
	os.Setenv("USERNAME", "spf13")
	port := viper.GetInt("port")
	name := viper.GetString("name")
	fmt.Printf("port: %d\n", port)
	fmt.Printf("name: %s\n", name)

}

const (
	defaultConfigFilename = "stingoftheviper"
	envPrefix             = "STING"
)

func NewRootCommand() *cobra.Command {
	color := ""
	number := 0

	rootCmd := &cobra.Command{
		Use:   "stingoftheviper",
		Short: "Cober and Viper together at last",
		Long:  `Demonstrate how to get cobra flags to bind to viper properly`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Working with OutOrStdout/OutOrStderr allows us to unit test our command easier
			out := cmd.OutOrStdout()

			// Print the final resolved value from binding cobra flags and viper config
			fmt.Fprintln(out, "Your favorite color is:", color)
			fmt.Fprintln(out, "The magic number is:", number)
		},
	}

	// Define cobra flags, the default value has the lowest (least significant) precedence
	rootCmd.Flags().IntVarP(&number, "number", "n", 7, "What is the magic number?")
	rootCmd.Flags().StringVarP(&color, "favorite-color", "c", "red", "Should come from flag first, then env var STING_FAVORITE_COLOR then the config file, then the default last")
	return rootCmd
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()
	v.SetConfigName(defaultConfigFilename)
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		// It's okay if there isn't a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	bindFlags(cmd, v)

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

//https://github.com/carolynvs/stingoftheviper
func demoViper5() {
	cmd := NewRootCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func demoViper6() {
	flag.Int("flagname", 1234, "help message for flagname")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	fmt.Printf("flagname %d", viper.GetInt("flagname"))
}
