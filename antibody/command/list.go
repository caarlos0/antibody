package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/akatrevorjay/antibody"
	"github.com/akatrevorjay/antibody/bundle"
)

// List all downloaded bundles
var List = cli.Command{
	Name:  "list",
	Usage: "list all currently downloaded bundles",
	Action: func(ctx *cli.Context) {
		for _, b := range bundle.List(antibody.Home()) {
			fmt.Println(b.Name())
		}
	},
}
