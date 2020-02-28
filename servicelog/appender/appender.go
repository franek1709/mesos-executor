package appender

import "github.com/franek1709/mesos-executor-fork/servicelog"

// Appender delivers service log entries to their destination.
type Appender interface {
	Append(entries <-chan servicelog.Entry)
}
