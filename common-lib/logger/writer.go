package logger

func (l *logger) writeLog(message string) {
	l.mu.Lock()
	l.logging.Println(message)
	l.mu.Unlock()
}
