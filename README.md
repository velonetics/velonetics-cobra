Pucora Cobra
====

An adapter of the [cobra](http://github.com/spf13/cobra) lib for the [Pucora](http://pucora.in) framework

Package cmd defines the cobra command structs and an execution method for adding an improved CLI to
Pucora based api gateways

## Basic example

```
package main

import (
	"os"

	"github.com/pucora/pucora-cobra/v2"
	"github.com/pucora/pucora-viper/v2"
	"github.com/pucora/lura/v2/config"
	"github.com/pucora/lura/v2/logging"
	"github.com/pucora/lura/v2/proxy"
	pucoragin "github.com/pucora/lura/v2/router/gin"
)

func main() {

	cmd.Execute(viper.New(), func(serviceConfig config.ServiceConfig) {
		logger, _ := logging.NewLogger("DEBUG", os.Stdout, "")
		pucoragin.DefaultFactory(proxy.DefaultFactory(logger), logger).New().Run(serviceConfig)
	})

}
```

## Available commands

The `cmd` package includes four commands: `check`, `check-plugin`, `help` and `run`.

1. *check* validates the received config file.
2. *check-plugin* validates the dependencies shared between the binary and a plugin.
3. *help* displays details about any command.
4. *run* executes the passed executor once the received flags overwrite the parsed config.

```
$ ./pucora
 ╓▄█                          ▄▄▌                               ╓██████▄µ
▐███  ▄███╨▐███▄██H╗██████▄  ║██▌ ,▄███╨ ▄██████▄  ▓██▌█████▄  ███▀╙╙▀▀███╕
▐███▄███▀  ▐█████▀"╙▀▀"╙▀███ ║███▄███┘  ███▀""▀███ ████▀╙▀███H ███     ╙███
▐██████▌   ▐███⌐  ,▄████████M║██████▄  ║██████████M███▌   ███H ███     ,███
▐███╨▀███µ ▐███   ███▌  ,███M║███╙▀███  ███▄```▄▄` ███▌   ███H ███,,,╓▄███▀
▐███  ╙███▄▐███   ╙█████████M║██▌  ╙███▄`▀███████╨ ███▌   ███H █████████▀
                     ``                     `'`
Version: undefined

The API Gateway builder

Usage:
  pucora [command]

Available Commands:
  check        Validates that the configuration file is valid.
  check-plugin Checks your plugin dependencies are compatible.
  help         Help about any command
  run          Runs the Pucora server.

Flags:
  -h, --help   help for pucora

Use "pucora [command] --help" for more information about a command.

```
