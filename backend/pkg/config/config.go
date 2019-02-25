package config

import (
	"io"
	"os"

	mysqlDriver "github.com/go-sql-driver/mysql" // mysql driver for database/sql
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// values holds the possible config values
type values struct {
	RunSwagger  bool
	MySQLConfig *mysqlDriver.Config
	LogWriter   io.Writer
	Port        int
}

// Values the config values returned from viper
var Values values

func init() {
	viper.SetDefault("swagger", false)

	// Swagger
	flag.Bool("swagger", true, "run the swagger docs")
	// Log File
	flag.String("logfile", "", "access log file location (defaults to stdout)")
	// MySQLConfig
	flag.StringP("dbuser", "u", "", "mysql db username")
	flag.StringP("dbpass", "p", "", "mysql db password")
	flag.StringP("dbname", "n", "", "mysql db name")
	flag.String("dbnet", "", "mysql db net")
	flag.String("dbaddr", "", "mysql db addr")
	// app settings
	flag.Int("port", 8081, "port to run the api on")

	flag.Parse()

	viper.AutomaticEnv()
	viper.SetEnvPrefix("codecamp")
	viper.BindPFlags(flag.CommandLine)

	// Swagger
	Values.RunSwagger = viper.GetBool("swagger")

	// Log File
	if viper.GetString("logfile") != "" {
		Values.LogWriter = getFileWriter(viper.GetString("logfile"))
	} else {
		Values.LogWriter = os.Stdout
	}

	// MySQLConfig
	Values.MySQLConfig = mysqlDriver.NewConfig()
	Values.MySQLConfig.ParseTime = true // needed default for our mysql implementation
	Values.MySQLConfig.User = viper.GetString("dbuser")
	Values.MySQLConfig.Passwd = viper.GetString("dbpass")
	Values.MySQLConfig.DBName = viper.GetString("dbname")
	Values.MySQLConfig.Net = viper.GetString("dbnet")
	Values.MySQLConfig.Addr = viper.GetString("dbaddr")

	// App settings
	Values.Port = viper.GetInt("port")
}

func getFileWriter(fileName string) *os.File {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return f
}
