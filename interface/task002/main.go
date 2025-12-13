package main

// в чем минусы текущей реализации логгера? напиши свою обертку над ним
import "os"

type Logger interface {
	Log(message string) error
	Close() error
}

type FileLogger struct {
	file *os.File
}

func NewFileLogger(fileName string) (*FileLogger, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return &FileLogger{f}, nil
}

func (f *FileLogger) Log(message string) error {
	_, err := f.file.WriteString(message + "\n")
	return err
}

func (f *FileLogger) Close() error {
	return f.file.Close()
}

// ======= КОД ВЫШЕ НЕЛЬЗЯ МЕНЯТЬ =========

type SequentialLogger struct {
	wrppedLogger Logger
}

func NewSequentialLogger(wrppedLogger Logger) SequentialLogger {
}

func (sl SequentialLogger) Log(message string) error {
}

func (sl SequentialLogger) Close() error {
}
