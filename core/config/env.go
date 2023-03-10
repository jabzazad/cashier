package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var (
	// ENV decleare global environment
	ENV = &Environment{}
)

// Environment environment
type Environment struct {
	App struct {
		ProjectID   string   `mapstructure:"PROJECT_ID"`
		Env         string   `mapstructure:"ENV"`
		WebBaseURL  string   `mapstructure:"WEB_BASE_URL"`
		APIBaseURL  string   `mapstructure:"API_BASE_URL"`
		Release     bool     `mapstructure:"RELEASE"`
		Port        string   `mapstructure:"PORT"`
		Environment string   `mapstructure:"ENVIRONMENT"`
		Sources     []string `mapstructure:"SOURCES"`
		Password    string   `mapstructure:"PASSWORD"`
		Host        string   `mapstructure:"HOST"`
		API         string   `mapstructure:"API"`
	} `mapstructure:"APP"`
	Swagger struct {
		Title       string   `mapstructure:"TITLE"`
		Version     string   `mapstructure:"VERSION"`
		Host        string   `mapstructure:"HOST"`
		BaseURL     string   `mapstructure:"BASE_URL"`
		Description string   `mapstructure:"DESCRIPTION"`
		Schemes     []string `mapstructure:"SCHEMES"`
		Enable      bool     `mapstructure:"ENABLE"`
	} `mapstructure:"SWAGGER"`
	HTTPServer struct {
		ReadTimeout       time.Duration `mapstructure:"READ_TIMEOUT"`
		WriteTimeout      time.Duration `mapstructure:"WRITE_TIMEOUT"`
		ReadHeaderTimeout time.Duration `mapstructure:"READ_HEADER_TIMEOUT"`
	} `mapstructure:"HTTP_SERVER"`
	CashierDeskPath string         `mapstructure:"CASHIERDATA_PATH"`
	PostgreSQL      DatabaseConfig `mapstructure:"POSTGRE_SQL"`
}

// DatabaseConfig database config model
type DatabaseConfig struct {
	Host                string        `mapstructure:"HOST"`
	Port                int           `mapstructure:"PORT"`
	Username            string        `mapstructure:"USERNAME"`
	Password            string        `mapstructure:"PASSWORD"`
	DatabaseName        string        `mapstructure:"DATABASE_NAME"`
	DatabaseCompanyName string        `mapstructure:"DATABASE_COMPANY_NAME"`
	DriverName          string        `mapstructure:"DRIVER_NAME"`
	Timeout             string        `mapstructure:"TIMEOUT"`
	Enable              bool          `mapstructure:"ENABLE"`
	MaxIdleConns        int           `mapstructure:"MAX_IDLE_CONNS"`
	MaxOpenConns        int           `mapstructure:"MAX_OPEN_CONNS"`
	ConnMaxLifetime     time.Duration `mapstructure:"MAX_LIFE_TIME"`
}

// Read init env
func Read(path string, env string) error {
	v := viper.New()
	v.SetConfigName(fmt.Sprintf("config.%s", env))
	v.AddConfigPath(path)
	v.AutomaticEnv()
	v.SetConfigType("yml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	err := v.Unmarshal(&ENV)
	if err != nil {
		return err
	}
	return nil
}
