package log

import "github.com/op/go-logging"

// var logs = make(map[string]*logging.Logger)

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// var backend = logging.NewLogBackend(os.Stdout, "", 0)
// var backendFormatter = logging.NewBackendFormatter(backend, format)

// Get return an instance of a logger
// func Get(name string) *logging.Logger {
// 	logging.SetBackend(backend, backendFormatter)
// 	if _, ok := logs[name]; !ok {
// 		var log = logging.MustGetLogger(name)
// 		logs[name] = log
// 	}
// 	return logs[name]
// }

// Get return an instance of a logger
func Get(name string) *logging.Logger {
	// logging.SetBackend(backend, backendFormatter)
	logging.SetFormatter(format)
	return logging.MustGetLogger(name)
}

// Debug message
func Debug(args ...interface{}) {
	Get("default").Debug(args...)
}

// Info message
func Info(args ...interface{}) {
	Get("default").Info(args...)
}

// Notice message
func Notice(args ...interface{}) {
	Get("default").Notice(args...)
}

// Warning message
func Warning(args ...interface{}) {
	Get("default").Warning(args...)
}

// Error message
func Error(args ...interface{}) {
	Get("default").Error(args...)
}

// Critical message
func Critical(args ...interface{}) {
	Get("default").Critical(args...)
}
