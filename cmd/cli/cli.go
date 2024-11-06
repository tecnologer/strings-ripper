package cli

import (
	"fmt"

	"github.com/tecnologer/string-shunker/pkg/ripper"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:  "ripper",
		Usage: "Ripper is a tool for ripping string into a concatenated string with the defined length",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Input string to be ripped",
				Required: true,
			},
			&cli.IntFlag{
				Name:    "length",
				Aliases: []string{"l"},
				Usage:   "Length of the ripped string",
				Value:   150,
			},
		},
		Action: func(context *cli.Context) error {
			fmt.Println(ripper.Do(context.String("input"), context.Int("length")))

			return nil
		},
	}
}
