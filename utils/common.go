package util

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Config struct {
	Host       string `mapstructure:"host"`
	Connection string `mapstructure:"connection"`
}

const (
	TIME_LAYOUT_YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func HttpHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.URL)
		handler.ServeHTTP(w, r)
	})
}
func Redirect(w http.ResponseWriter, r *http.Request, url string, port int) {
	http.Redirect(w, r, url, port)
}
