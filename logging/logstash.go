package logging

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

const (
	red    = 31
	green  = 32
	yellow = 33
	blue   = 34
	gray   = 37
)

// Fields ..
type Fields map[string]interface{}

// inst ..
type inst struct {
	fields Fields
	msg    string
	time   string
	level  Level
}

// Entry ..
type Entry interface {
	Panic(string)
	Fatal(string)
	Error(string)
	Warn(string)
	Info(string)
	Debug(string)
	WithFields(Fields) *inst
}

// WithFields add more field
func WithFields(fields Fields) (entry Entry) {
	entry = &inst{fields: fields}
	entry.WithFields(fields)
	return
}

// WithFields ..
func (i *inst) WithFields(fields Fields) *inst {
	for k, v := range i.fields {
		i.fields[k] = v
	}
	return i
}

// Panic ..
func Panic(str string) {
	i := &inst{}
	i.Panic(str)
}

// Panic ..
func (i *inst) Panic(str string) {
	i.msg = str
	i.level = PanicLevel
	i.output()
}

// Fatal ..
func Fatal(str string) {
	i := &inst{}
	i.Fatal(str)
}

// Fatal ..
func (i *inst) Fatal(str string) {
	i.msg = str
	i.level = FatalLevel
	i.output()
}

// Error ..
func Error(str string) {
	i := &inst{}
	i.Error(str)
}

// Error ..
func (i *inst) Error(str string) {
	i.msg = str
	i.level = ErrorLevel
	i.output()
}

// Warn ..
func Warn(str string) {
	i := &inst{}
	i.Warn(str)
}

// Warn ..
func (i *inst) Warn(str string) {
	i.msg = str
	i.level = WarnLevel
	i.output()
}

// Info ..
func Info(str string) {
	i := &inst{}
	i.Info(str)
}

// Info ..
func (i *inst) Info(str string) {
	i.msg = str
	i.level = InfoLevel
	i.output()
}

// Debug ..
func Debug(str string) {
	i := &inst{}
	i.Debug(str)
}

// Debug ..
func (i *inst) Debug(str string) {
	i.msg = str
	i.level = DebugLevel
	i.output()
}

func (i *inst) output() {
	var color int
	var err error
	var waitWrite []byte
	if i.level > logLevel {
		return
	}
	switch i.level {
	case DebugLevel:
		color = gray
	case WarnLevel:
		color = yellow
	case ErrorLevel, FatalLevel, PanicLevel:
		color = red
	default:
		color = blue
	}
	levelText := strings.ToUpper(i.level.String())[0:4]
	var output string
	if i.fields == nil {
		i.fields = Fields{}
	}
	if _, file, line, ok := runtime.Caller(2); ok {
		i.fields["_file"] = filepath.Base(file)
		i.fields["_line"] = line
	}

	if _, file, line, ok := runtime.Caller(3); ok {
		i.fields["_file"] = filepath.Base(file)
		i.fields["_line"] = line
	}
	t := time.Now()
	i.time = t.Format("15:04:05.000")
	i.fields["__time"] = t.Format("01-02T15:04:05.000")

	var keys []string
	for key := range i.fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if key != "_line" && key != "_file" && key != "__time" {
			color := green
			output += fmt.Sprintf(" \x1b[%dm%s\x1b[0m=%+v", color, key, i.fields[key])
		}
	}

	for _, key := range keys {
		color := red
		if key == "_line" || key == "_file" {
			output += fmt.Sprintf(" \x1b[%dm%s\x1b[0m=%+v", color, key[1:], i.fields[key])
		}
	}

	fmt.Printf("\x1b[%dm%s\x1b[0m[%s] %-40s %s\n", color, levelText, i.time, i.msg, output)

	i.fields["level"] = i.level.String()
	i.fields["msg"] = i.msg
	if waitWrite, err = json.Marshal(i.fields); err != nil {
		Fatal("Cannot convert fields to string.")
		return
	}
	waitWrite = append(waitWrite, '\n')

	if _, err := logger.Write(waitWrite); err != nil {
		Fatal("Cannot write log to file.")
	}

	if PanicLevel == i.level {
		os.Exit(-1)
	}
}
