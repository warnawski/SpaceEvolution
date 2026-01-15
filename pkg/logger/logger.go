package logger

import (
	"log"
	"os"
)

func ConfigureLogger() {

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.SetOutput(os.Stdout)
	log.SetPrefix("[SPACE-EVO] ]")
}
