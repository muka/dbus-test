package main

import (
	"fmt"
	"os"

	"create-net.org/lcapra/dbus-test/dbusdaemon"
	"create-net.org/lcapra/dbus-test/log"
	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "dbus-http-proxy"
	app.Usage = "make an explosive entrance"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "server, s",
			Usage: "start server",
		},
		cli.BoolFlag{
			Name:  "client, c",
			Usage: "start client",
		},
		cli.BoolFlag{
			Name:  "call, a",
			Usage: "call Foo",
		},
	}

	app.Action = func(c *cli.Context) {

		if c.Bool("server") {
			dbusdaemon.CreateInterfaces()
			select {}
		}

		if c.Bool("call") {
			dbusdaemon.Call("Foo")
		}

		if c.Bool("client") {
			result, err := dbusdaemon.Introspect(dbusdaemon.DbusInterface, dbusdaemon.DbusObjectPath)
			if err != nil {
				log.Critical(err)
				panic(err)
			}
			fmt.Println(result)
		}

	}

	app.Run(os.Args)

}
