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
	app.Name = "dbus-test"
	app.Usage = "manage dbus service and query for its services"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "server, s",
			Usage: "start server",
		},
		cli.BoolFlag{
			Name:  "inspect, i",
			Usage: "inspect service",
		},
		cli.BoolFlag{
			Name:  "call, c",
			Usage: "call Foo and FooPlus",
		},
		cli.BoolFlag{
			Name:  "property, p",
			Usage: "change property and receive a signal update",
		},
	}

	app.Action = func(c *cli.Context) {

		if c.Bool("server") {
			dbusdaemon.CreateInterfaces()
			select {}
		}

		if c.Bool("call") {
			log.Debug("Call Foo")
			dbusdaemon.Call("Foo")
			log.Debug("Call FooPlus")
			dbusdaemon.Call("FooPlus", "ciao ciao")
		}

		if c.Bool("property") {

			log.Debug("Register for proeprty signal SomeInt")
			err := dbusdaemon.RegisterSignal("SomeInt")
			if err != nil {
				log.Critical(err)
				panic(err)
			}

			log.Debug("Get SomeInt")
			val, err := dbusdaemon.Get("SomeInt", int32(42))
			if err != nil {
				log.Error("Error calling set")
				log.Critical(err)
				panic(err)
			}

			log.Debug("Current value", val)

			log.Debug("Set SomeInt=42")
			err = dbusdaemon.Set("SomeInt", int32(42))
			if err != nil {
				log.Error("Error calling set")
				log.Critical(err)
				panic(err)
			}
		}

		if c.Bool("inspect") {
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
