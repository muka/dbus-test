package dbusdaemon

import (
	"errors"
	"fmt"

	"create-net.org/lcapra/dbus-http-proxy/log"

	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
)

const (
	// DbusObjectPath default ObjectPath
	DbusObjectPath = "/org/agile/HttpProxy" // DBUS object path
	// DbusInterface default Interface
	DbusInterface = "org.agile.HttpProxy" // DBUS Interface
)

const intro = `
<node>
	<interface name="` + DbusInterface + `">
		<method name="Foo">
			<arg direction="out" type="s"/>
		</method>
	</interface>` + introspect.IntrospectDataString + `</node> `

type foo string

func (f foo) Foo() (string, *dbus.Error) {
	fmt.Println(f)
	return string(f), nil
}

// CreateInterfaces create a dbus server to expose services
func CreateInterfaces() (err error) {
	log.Debug("Init session bus")
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

	log.Debug("Export Foo")
	f := foo("Bar!")
	conn.Export(f, DbusObjectPath, DbusInterface)
	log.Debug("Export introspectable")
	conn.Export(introspect.Introspectable(intro), DbusObjectPath,
		"org.freedesktop.DBus.Introspectable")

	log.Debug("Listening on " + DbusInterface + " / " + DbusObjectPath)

	return nil
}
