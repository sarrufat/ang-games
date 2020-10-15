package appconf

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var (
	ChessMaxResults     int
	CacheCleaningPeriod time.Duration
	CacheTimeOut        time.Duration
)

func AppConfiguration() {
	viper.SetDefault("chessMaxResults", "999")
	viper.SetDefault("cacheCleaningPeriod", "30s")
	viper.SetDefault("cacheTimeOut", "120s")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./appconf")
	viper.AddConfigPath("./etc/ang-games/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	loarParameters()
}

func loarParameters() {
	ChessMaxResults = viper.GetInt("chessMaxResults")
	CacheCleaningPeriod = viper.GetDuration("cacheCleaningPeriod")
	CacheTimeOut = viper.GetDuration("cacheTimeOut")
}
