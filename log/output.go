package log

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var DelayTime = 3 * time.Minute

type Output struct {
	io.WriteCloser
	Logger     *Logger
	deleteTime time.Time
	logFile    *os.File
	fileLock   *sync.Mutex
}

func NewOutput(filePath string, logger *Logger) (*Output, error) {

	ret := &Output{
		Logger:     logger,
		deleteTime: time.Now().Add(DelayTime),
		fileLock:   &sync.Mutex{},
	}

	ret.fileLock.Lock()

	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	logFile, err := filepath.Abs(workingDir + string(os.PathSeparator) + filePath)
	if err != nil {
		return nil, err
	}

	logDir, _ := filepath.Split(logFile)

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, err
	}

	var readErr error
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		ret.logFile, readErr = os.Create(logFile)
	} else {
		ret.logFile, readErr = os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, 0660)
	}

	fmt.Println(logFile)

	if readErr != nil {
		return nil, readErr
	}

	ret.fileLock.Unlock()

	return ret, nil

}

func (o *Output) delete() bool {
	return time.Now().After(o.deleteTime)
}

func (o *Output) Write(p []byte) (n int, err error) {
	if o.logFile == nil {
		err = errors.New("logFile is nil, closed possibly")
		return
	}
	o.fileLock.Lock()
	n, err = o.logFile.Write(p)
	o.fileLock.Unlock()
	return
}

func (o *Output) Close() (e error) {
	o.fileLock.Lock()
	e = o.logFile.Close()
	o.fileLock.Unlock()
	return
}
