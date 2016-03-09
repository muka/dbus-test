package dbusdaemon

import (
	"encoding/json"

	"create-net.org/lcapra/dbus-test/log"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
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
func Call(method string) (xml string, err error) {
	log.Debug("Calling method Foo")
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Error("Error getting session bus instance")
		return xml, err
	}
	var s []string
	log.Debug("Calling " + DbusInterface + "." + method + "()")
	call := conn.Object(DbusInterface, DbusObjectPath).Call(DbusInterface+"."+method, 0)

	log.Debug(call.Body)
	log.Debug(call.Args)
	log.Debug(call.Destination)
	log.Debug(call.Err)
	log.Debug(call.Method)
	log.Debug(call.Path)
	log.Debug("*************")

	err = call.Store(&s)
	if err != nil {
		log.Error("Error while calling " + method)
		log.Error(err)
	}
	return xml, err
}
