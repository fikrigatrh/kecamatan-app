package server

import (
	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
	"kecamatan_app/driver"
	"kecamatan_app/router"
)

func init() {
	gotenv.Load()
}

func StartServer()  {
	addr := driver.Config.ServiceHost + ":" + driver.Config.ServicePort
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.Router(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}
