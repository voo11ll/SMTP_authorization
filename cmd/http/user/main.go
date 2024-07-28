package main

import (
	"auth/auth_back/config"
	"auth/auth_back/pkg/helpers/httpServerHelper"
	userHttp "auth/auth_back/pkg/services/http/user"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	router := httpServerHelper.NewRouter(userHttp.Routes)

	castomCors := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "authorization", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "", "", ""},
		Debug:            true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
	})

	mainHandler := castomCors.Handler(router)

	log.Println("Auth http listening on " + viper.GetString("http.user.host") + ":" + viper.GetString("http.user.port"))
	log.Fatal(http.ListenAndServe(viper.GetString("http.user.host")+":"+viper.GetString("http.user.port"), mainHandler))
}
