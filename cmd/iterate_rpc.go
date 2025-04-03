// write go code here
package main

import (
	"log"
	"os"

	"github.com/fufonglabs/rpc-iterator/iter"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Action: RunRPC,
		Name:   "Replay-RPC",
		Usage:  "Iterate RPC requests from a given path.",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func RunRPC(ctx *cli.Context) error {
	// read file
	iter, _ := iter.NewFileReader(ctx.Context, ctx.Args().Get(0))
	defer iter.Close()

	// iterating
	if iter.Next() {
		val := iter.Value()
		log.Println(val)
	}

	return nil
}
