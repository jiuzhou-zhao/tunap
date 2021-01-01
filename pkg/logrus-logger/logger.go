package logrus_logger

import "github.com/sirupsen/logrus"

type Logger struct {
	entry *logrus.Entry
}

func NewLogger(logger *logrus.Logger) *Logger {
	if logger == nil {
		logger = logrus.New()
		logger.SetFormatter(&logrus.TextFormatter{})
		logger.SetLevel(logrus.DebugLevel)
	}
	logrus.SetLevel(logrus.DebugLevel)
	return &Logger{
		entry: logrus.NewEntry(logger),
	}
}

func (logger *Logger) WithFields(fields map[string]interface{}) {
	logger.entry = logger.entry.WithFields(fields)
}

func (logger *Logger) Debug(v ...interface{}) {
	logger.entry.Debug(v...)
}

func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.entry.Debugf(format, v...)
}

func (logger *Logger) Info(v ...interface{}) {
	logger.entry.Info(v...)
}

func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.entry.Infof(format, v...)
}

func (logger *Logger) Warn(v ...interface{}) {
	logger.entry.Warn(v...)
}

func (logger *Logger) Warnf(format string, v ...interface{}) {
	logger.entry.Warnf(format, v...)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.entry.Error(v...)
}

func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.entry.Errorf(format, v...)
}

func (logger *Logger) Fatal(v ...interface{}) {
	logger.entry.Fatal(v...)
}

func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.entry.Fatalf(format, v...)
}
