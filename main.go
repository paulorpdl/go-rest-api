package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kelseyhightower/envconfig"
	"github.com/paulorpdl/go-rest-api/pkg/config"
	"github.com/paulorpdl/go-rest-api/pkg/logger"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("iniciando aplicación...")
	log.Info("cargando configuración...")

	// Leemos configuraciones de variables y las almacenamos en un struct para uso posterior
	config := &config.Server{}
	if err := envconfig.Process("SERVER", config); err != nil {
		log.Error(err)
	}

	//Habilitamos depuración en caso de estar la variable `SERVER_DEBUG` habilitada
	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}
	log.Debugf("configuracion: %+v", config)

	//Creamos router
	r := chi.NewRouter()
	r.Use(logger.NewStructuredLogger(log.New()))

	//Creamos una respuesta a la raíz
	r.Get(config.Path, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	log.Infof("iniciando escucha en %s:%s", config.Address, config.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.Address, config.Port), r)
}
