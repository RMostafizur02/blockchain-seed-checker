package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	file *os.File
}

func NewLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	
	return &Logger{file: file}, nil
}

func (l *Logger) LogHit(mnemonic string, balance float64, addresses map[string]string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("[%s] HIT - Mnemonic: %s, Balance: %.8f\n", timestamp, mnemonic, balance)
	
	// Log to file
	l.file.WriteString(message)
	
	// Log to console
	log.Printf("💰 HIT FOUND: %s", mnemonic)
}

func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}

func (l *Logger) LogError(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	errorMsg := fmt.Sprintf("[%s] ERROR - %s\n", timestamp, message)
	
	l.file.WriteString(errorMsg)
	log.Printf("❌ ERROR: %s", message)
}

func (l *Logger) LogInfo(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	infoMsg := fmt.Sprintf("[%s] INFO - %s\n", timestamp, message)
	
	l.file.WriteString(infoMsg)
	log.Printf("ℹ️  INFO: %s", message)
}
