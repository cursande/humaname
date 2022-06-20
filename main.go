package main

import (
	"log"
	"os"

	"github.com/cursande/humaname/src"
	"github.com/urfave/cli/v2"
)


func printName(name string, ctx *cli.Context) {
	print(name)

	if !ctx.Bool("remove-trailing-newline") {
		print("\n")
	}

	return
}

func genderFromFlag(ctx *cli.Context) humaname.Gender {
	switch ctx.String("gender") {
	case "m":
		return humaname.Male
	case "f":
		return humaname.Female
	default:
		return humaname.All
	}

}

func main() {
	app := &cli.App{
		Name:    "humaname",
		Usage:   "For generating regular human-sounding names",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "both",
				Aliases: []string{"b"},
				Action: func(ctx *cli.Context) error {
					gender := genderFromFlag(ctx)
					printName(humaname.PickRandomName(humaname.Both, gender), ctx)
					return nil
				},
			},
			{
				Name:    "first",
				Aliases: []string{"f"},
				Action: func(ctx *cli.Context) error {
					gender := genderFromFlag(ctx)
					printName(humaname.PickRandomName(humaname.First, gender), ctx)
					return nil
				},
			},
			{
				Name:    "last",
				Aliases: []string{"l"},
				Action: func(ctx *cli.Context) error {
					gender := genderFromFlag(ctx)
					printName(humaname.PickRandomName(humaname.Last, gender), ctx)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "remove-trailing-newline",
				Aliases: []string{"r"},
				Usage:   "Remove trailing newline from printed value",
			},
			&cli.StringFlag{
				Name:    "gender",
				Aliases: []string{"g"},
				Usage:   "Return names only for given gender (default: no filtering)",
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
