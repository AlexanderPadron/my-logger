package mylogger

import "log"

func LogInfo(mensaje string) {
	log.Printf("INFO - %v", mensaje)
}

func LogWarning(mensaje string) {
	log.Printf("WARNING - %v", mensaje)
}

func LogError(mensaje string) {
	log.Printf("ERROR - %v", mensaje)
}