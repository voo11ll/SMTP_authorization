package main

import (
	"auth/auth_back/config"
	"auth/auth_back/pkg/helpers/httpServerHelper"
	organizationHttp "auth/auth_back/pkg/services/http/organization"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	router := httpServerHelper.NewRouter(organizationHttp.Routes)

	log.Println("Organization http listening on " + viper.GetString("http.organization.host") + ":" + viper.GetString("http.auorganizationth.port"))
	log.Fatal(http.ListenAndServe(viper.GetString("http.organization.host")+":"+viper.GetString("http.organization.port"), router))
}
