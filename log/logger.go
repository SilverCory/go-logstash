package log

import (
	"sync"
	"time"
)

type Logger struct {
	outputs     map[string]Output
	outputsLock *sync.Mutex
}

func New() (l *Logger) {
	l = &Logger{
		outputs:     make(map[string]Output),
		outputsLock: &sync.Mutex{},
	}

	go func() {
		time.Sleep(1 * time.Minute)
		for k, v := range l.outputs {
			if v.delete() {
				l.removeOutput(k)
			}
		}
	}()

	return
}

func (l *Logger) Log(filePath string, contents []byte) (err error) {
	l.outputsLock.Lock()
	output, ok := l.outputs[filePath]
	if !ok {
		output, err := NewOutput(filePath, l)
		if err != nil {
			l.outputsLock.Unlock()
			return err
		} else {
			l.outputs[filePath] = *output
		}
	}

	l.outputsLock.Unlock()
	contents = append(contents, '\n')
	_, err = output.Write(contents)
	return

}

func (l *Logger) removeOutput(name string) {
	l.outputsLock.Lock()
	output, ok := l.outputs[name]
	if ok {
		output.Close()
	}
	delete(l.outputs, name)
	l.outputsLock.Unlock()
}
