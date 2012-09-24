package golog

import (
	"os"
	"fmt"
	"testing"
)

var file *os.File

func init() {
	var err error
	file, err = os.OpenFile("/home/xing/log/golog.log", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0664)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		os.Exit(1)
	}
}

func TestDebug(t *testing.T) {
	fmt.Printf("test debug....%v\n", file.Name())
	log := NewLogger(file, "golog", LOGLEVEL_DEBUG, FLAG_LstdFlags | FLAG_Lshortfile)
	for i := 0; i < 2; i++ {
		log.Debug("This is ", i, "th debug info")
	}
	for i := 0; i < 2; i++ {
		log.Infof("This is %dth debug info", i)
	}
}

func TestWarn(t *testing.T) {
	fmt.Printf("test Warn....%v\n", file.Name())
	log := NewLogger(file, "golog", LOGLEVEL_WARN, FLAG_Ldate | FLAG_Llongfile)
	for i := 0; i < 2; i++ {
		log.Warn("This is ", i, "th debug info")
	}
	for i := 0; i < 2; i++ {
		log.Debugf("This is %dth debug info", i) // log nothing here
	}
}

func TestFatal(t *testing.T) {
	fmt.Printf("test fatal....%v\n", file.Name())
	log := NewLogger(file, "golog", LOGLEVEL_WARN, FLAG_Ltime)
	for i := 0; i < 2; i++ {
		log.Warn("This is ", i, "th debug info") // log nothing here
	}
	for i := 0; i < 2; i++ {
		log.Fatalf("This is %dth debug info", i)	// log.Fatal* will invoke os.Exit(1), so There is only one fatal info logged.
	}
}