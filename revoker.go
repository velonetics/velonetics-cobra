package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	revoker "github.com/pucora/pucora-revoker/v2"
	"github.com/pucora/lura/v2/logging"
	"github.com/spf13/cobra"
)

func revokerFunc(cmd *cobra.Command, _ []string) {
	if cfgFile == "" {
		cmd.Println("Please, provide the path to the configuration file with --config")
		return
	}
	serviceConfig, err := parser.Parse(cfgFile)
	if err != nil {
		cmd.Printf("ERROR parsing the configuration file: %s\n", err.Error())
		os.Exit(-1)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()
	logger, _ := logging.NewLogger("INFO", os.Stdout, "[Revoker]")
	if err := revoker.Run(ctx, serviceConfig, logger); err != nil {
		cmd.Println(err.Error())
		os.Exit(-1)
	}
}
