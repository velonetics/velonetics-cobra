package main

import (
	"os"

	cmd "github.com/pucora/pucora-cobra/v2"
	koanf "github.com/pucora/pucora-koanf"
	"github.com/pucora/lura/v2/config"
	"github.com/pucora/lura/v2/logging"
	"github.com/pucora/lura/v2/proxy"
	"github.com/pucora/lura/v2/router/gin"
)

func main() {
	cmd.Execute(koanf.New(), func(serviceConfig config.ServiceConfig) {
		logger, _ := logging.NewLogger("DEBUG", os.Stdout, "")
		gin.DefaultFactory(proxy.DefaultFactory(logger), logger).New().Run(serviceConfig)
	})
}
