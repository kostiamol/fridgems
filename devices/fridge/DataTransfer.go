package fridge

import (
	"github.com/KharkivGophers/device-smart-house/config"
	"github.com/KharkivGophers/device-smart-house/tcp/connectionupdate"
	"github.com/KharkivGophers/device-smart-house/models"
	log "github.com/Sirupsen/logrus"
)

//DataTransfer func sends request as JSON to the centre
func DataTransfer(config *config.DevConfig, reqChan chan models.Request, c *models.Control) {

	// for data transfer
	transferConnParams := models.TransferConnParams{
		// HostOut: GetEnvCenter("CENTER_PORT_3030_TCP_ADDR"),
		HostOut: "0.0.0.0",
		PortOut: "3030",
		ConnTypeOut: "tcp",
	}

	conn := connectionupdate.GetDial(transferConnParams.ConnTypeOut, transferConnParams.HostOut, transferConnParams.PortOut)

	for {
		select {
		case r := <-reqChan:
			go func() {
				defer func() {
					if a := recover(); a != nil {
						log.Error(a)
						c.Close()
					}
				} ()
				connectionupdate.Send(r, conn)
			}()
		case <- c.Controller:
			log.Error("Data Transfer Failed")
			return
		}
	}
}

