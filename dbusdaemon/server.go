package dbusdaemon

import (
	"errors"
	"fmt"

	"create-net.org/lcapra/dbus-test/log"

	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	"github.com/godbus/dbus/prop"
)

const (
	// DbusObjectPath default ObjectPath
	DbusObjectPath = "/org/agile/HttpProxy" // DBUS object path
	// DbusInterface default Interface
	DbusInterface = "org.agile.HttpProxy" // DBUS Interface
)

type foo string

func (f foo) Foo() (string, *dbus.Error) {
	fmt.Println(f)
	return string(f), nil
}

func (f foo) FooPlus(what string) (string, *dbus.Error) {
	r := string(f) + " plus < " + what + " >"
	fmt.Println(r)
	return r, nil
}

// PropsSpec lists the available properties of the interface
var PropsSpec = map[string]map[string]*prop.Prop{
	DbusInterface: {
		"SomeInt": {
			int32(0),
			true,
			prop.EmitTrue,
			func(c *prop.Change) *dbus.Error {
				log.Info(c.Name, "changed to", c.Value)
				return nil
			},
		},
	},
}

// CreateInterfaces create a dbus server to expose services
func CreateInterfaces() (err error) {
	log.Debug("Getting session bus conn")
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}

	log.Debug("Request name")
	reply, err := conn.RequestName(DbusInterface, dbus.NameFlagDoNotQueue)
	if err != nil {
		return err
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		log.Error("name already taken")
		return errors.New(DbusInterface + " name already taken")
	}

	f := foo("Bar!")

	log.Debug("Export Foo")
	conn.Export(f, DbusObjectPath, DbusInterface)

	props := prop.New(conn, DbusObjectPath, PropsSpec)
	n := &introspect.Node{
		Name: DbusObjectPath,
		Interfaces: []introspect.Interface{
			introspect.IntrospectData,
			prop.IntrospectData,
			{
				Name:       DbusInterface,
				Methods:    introspect.Methods(f),
				Properties: props.Introspection(DbusInterface),
			},
		},
	}

	log.Debug("Export interfaces")
	conn.Export(introspect.NewIntrospectable(n), DbusObjectPath,
		"org.freedesktop.DBus.Introspectable")

	// conn.Export(introspect.Introspectable(intro), DbusObjectPath,
	// 	"org.freedesktop.DBus.Introspectable")

	log.Debug("Listening on " + DbusInterface + " / " + DbusObjectPath)

	return nil
}
