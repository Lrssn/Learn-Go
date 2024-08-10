package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {

	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds) // Time as microseconds
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile) // Filename and line
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags) // Custom logger
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:") // Custom prefix on existing logger
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil) // Structured Json log output
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)
}
