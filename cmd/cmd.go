package cmd

import (

	"github.com/keyeMyria/gin-wire-plate-cli/cmd/new"
	"github.com/urfave/cli"

)

// NewCommand 创建项目命令
func NewCommand() cli.Command {
	return cli.Command{
		Name:    "new",
		Aliases: []string{"n"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dir, d",
				Usage: "project directory",
			},
			&cli.StringFlag{
				Name:  "pkg, p",
				Usage: "package name",
			},
		},
		Action: func(c *cli.Context) error {
			cfg := new.Config{
				Dir:       c.String("dir"),
				PkgName:   c.String("pkg"),
			}
			return new.Exec(cfg)
		},
	}
}

