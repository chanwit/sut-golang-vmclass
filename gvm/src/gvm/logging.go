package gvm

import "fmt"

const (
    NONE   = iota
    INFO
    DEBUG
)

const LoggingLevel = INFO

func _info(a ...interface{}) {
    if LoggingLevel == INFO {
        fmt.Println(a)
    }
}

func _infof(s string, a ...interface{}) {
    if LoggingLevel == INFO {
        fmt.Printf(s, a)
    }
}


func _debug(a ...interface{}) {
    if(LoggingLevel == DEBUG) {
        fmt.Println(a)
    }
}

func _debugf(s string, a ...interface{}) {
    if(LoggingLevel == DEBUG) {
        fmt.Printf(s, a)
    }
}