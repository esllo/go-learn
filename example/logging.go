package example

import (
	"io"
	"log"
	"os"
)

func standardLogging() {
	log.Println("Println")
}

var customLogger *log.Logger

func customLoggerLogging() {
	customLogger = log.New(os.Stdout, "INFO : ", log.LstdFlags)
	customLogger.Println("Custom Log")
}

func fileLogging() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	customLogger = log.New(logFile, "INFO : ", log.Ldate|log.Llongfile)
	customLogger.Println("File log")
}

func changeLogOutput() {
	logFile, err := os.OpenFile("newLog.txt", os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Change SetOutput to log file")

	log.SetOutput(os.Stdout)
	log.Println("Cahgne SetOutput to stdout")
}

func multipleLogging() {
	logFile, err := os.OpenFile("multiLog.txt", os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multipleLogger := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multipleLogger)
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("Multiple Log")

	log.SetOutput(os.Stdout)
	log.Println("Std log")
}

func ExampleLogging() {
	standardLogging()
	customLoggerLogging()
	fileLogging()
	changeLogOutput()
	multipleLogging()
}
