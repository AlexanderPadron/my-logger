package mylogger

import "log"

func logInfo(mensaje string) {
	log.Printf("INFO - %v", mensaje)
}

func logWarning(mensaje string) {
	log.Printf("WARNING - %v", mensaje)
}

func logError(mensaje string) {
	log.Printf("ERROR - %v", mensaje)
}