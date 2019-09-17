package main


import (
	"log"
	"os"

	"github.com/keyeMyria/gin-wire-plate-cli/cmd"
	"github.com/urfave/cli"

)


func main() {
	app := cli.NewApp()
	app.Name = "gin-wire-plate-cli"
	app.Description = "gin-wire-plate cli"
	app.Version = "0.2.0"
	app.Commands = []cli.Command{
		cmd.NewCommand(),
		// cmd.GenerateCommand(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf(err.Error())
	}
}


