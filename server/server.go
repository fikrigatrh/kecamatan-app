package server

import (
	"fmt"
	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
	"kecamatan_app/driver"
	"kecamatan_app/router"
	"os"
)

func init() {
	gotenv.Load()
}

func StartServer()  {
	port := os.Getenv("SERVICE_PORT")
	addr := driver.Config.ServiceHost + ":" + port
	fmt.Println(driver.Config.ServicePort,"<<<<<<")
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.Router(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}
