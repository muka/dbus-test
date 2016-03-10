package dbusdaemon

import (
	"encoding/json"
	"errors"
	"fmt"

	"create-net.org/lcapra/dbus-test/log"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	"github.com/godbus/dbus/prop"
)

// Introspect introspect a service
func Introspect(interfaceName string, objectPath string) (xml string, err error) {
	log.Debug("Starting introspection")
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Error("Error getting session bus instance")
		return xml, err
	}
	node, err := introspect.Call(conn.Object(DbusInterface, DbusObjectPath))
	if err != nil {
		log.Error("Error during introspection call")
		return xml, err
	}
	data, err := json.MarshalIndent(node, "", "    ")
	if err != nil {
		log.Error("Error marshalling response")
		return xml, err
	}
	xml = string(data)
	return xml, nil
}

// Call a remote function
func Call(method string, args ...interface{}) (xml string, err error) {
	return call(DbusInterface, DbusObjectPath, method, args...)
}

var props *prop.Properties

func getProperties() (*prop.Properties, error) {

	conn, err := dbus.SessionBus()
	if err != nil {
		log.Error("Error getting session bus instance")
		return nil, err
	}

	if props == nil {
		props = prop.New(conn, DbusInterface, PropsSpec)
	}
	return props, nil
}

// Set a service property
func Set(property string, value interface{}) (err error) {
	properties, err := getProperties()
	if err != nil {
		return err
	}
	dbusErr := properties.Set(DbusInterface, property, dbus.MakeVariant(value))
	if dbusErr != nil {
		log.Error("Error calling propertiesSet")
		err = errors.New(dbusErr.Error())
	}
	return err
}

// Get a service property
func Get(property string, value interface{}) (val interface{}, err error) {
	properties, err := getProperties()
	if err != nil {
		return nil, err
	}
	variant, dbusErr := properties.Get(DbusInterface, property)
	if dbusErr != nil {
		log.Error("Error calling propertiesSet")
		err = errors.New(dbusErr.Error())
	}
	return variant, err
}

func call(dbusInterface string, dbusObjectPath dbus.ObjectPath, method string, args ...interface{}) (xml string, err error) {
	log.Debug("Calling method " + method)
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Error("Error getting session bus instance")
		return xml, err
	}
	var s string
	log.Debug("Calling " + dbusInterface + "." + method + "()")
	call := conn.Object(dbusInterface, dbusObjectPath).Call(dbusInterface+"."+method, 0, args...)
	err = call.Store(&s)
	if err != nil {
		log.Error("Error while calling " + method)
		log.Error(err)
	}
	return xml, err
}

//RegisterSignal register for a propery change signal
func RegisterSignal(propertyName string) (err error) {

	conn, err := dbus.SessionBus()
	if err != nil {
		log.Error("Failed to connect to session bus")
		return err
	}

	go func() {

		conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
			"type='signal',path='/org/freedesktop/DBus',interface='org.freedesktop.DBus',sender='org.freedesktop.DBus'")

		c := make(chan *dbus.Signal, 10)
		conn.Signal(c)
		for v := range c {
			fmt.Println(v)
		}

	}()

	return nil
}
