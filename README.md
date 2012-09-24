golog
=====

golog provides a convenient API for logging with golang.

You can use log level in your program with golog.


Useage:
=====

```go
package main

import (
	"github.com/xing4git/golog"
	"fmt"
	"os"
)

var log *golog.Logger

func init() {
	file, err := os.OpenFile("/home/xing/log/golog.log", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0664)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		os.Exit(1)
	}
	// if you want to change LOG LEVEL, you just need to change here.
	log = golog.NewLogger(file, "golog", golog.LOGLEVEL_DEBUG, golog.FLAG_LstdFlags | golog.FLAG_Lshortfile)
}

func main() {
	for i := 0; i < 5; i++ {
		log.Debug("This is ", i, "th debug info")
	}
	for i := 0; i < 5; i++ {
		log.Infof("This is %dth debug info", i)
	}
}

```

If log level is LOGLEVEL_DEBUG, all the log string will be logged. 

If log level is LOGLEVEL_FATAL, only the fatal log string will be logged. And fatal logs also cause program to exit.


LICENSE
=====

golog was just a component in uniqush(https://github.com/uniqush/uniqush-push).

Atfer forking from uniqush, I modify a lot code for convenient usage.

Uniqush's LICENSE is Apache License, Version 2.0. It's also my selection. See:

	http://www.apache.org/licenses/LICENSE-2.0