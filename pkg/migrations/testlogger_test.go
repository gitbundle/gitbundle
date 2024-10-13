// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations_test

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
	_ "unsafe"

	"github.com/gitbundle/modules/json"
	"github.com/gitbundle/modules/log"
)

// TestLogger is a logger which will write to the testing log
type TestLogger struct {
	log.WriterLogger
}

// Printf takes a format and args and prints the string to os.Stdout
func Printf(format string, args ...interface{}) {
	if log.CanColorStdout {
		for i := 0; i < len(args); i++ {
			args[i] = log.NewColoredValue(args[i])
		}
	}
	fmt.Fprintf(os.Stdout, "\t"+format, args...)
}

// NewTestLogger creates a TestLogger as a log.LoggerProvider
func NewTestLogger() log.LoggerProvider {
	logger := &TestLogger{}
	logger.Colorize = log.CanColorStdout
	logger.Level = log.TRACE
	return logger
}

// Init inits connection writer with json config.
// json config only need key "level".
func (log *TestLogger) Init(config string) error {
	err := json.Unmarshal([]byte(config), log)
	if err != nil {
		return err
	}
	log.NewWriterLogger(writerCloser)
	return nil
}

// Content returns the content accumulated in the content provider
func (log *TestLogger) Content() (string, error) {
	return "", fmt.Errorf("not supported")
}

// Flush when log should be flushed
func (log *TestLogger) Flush() {
}

// ReleaseReopen does nothing
func (log *TestLogger) ReleaseReopen() error {
	return nil
}

// GetName returns the default name for this implementation
func (log *TestLogger) GetName() string {
	return "test"
}

func init() {
	log.Register("test", NewTestLogger)
}

var (
	//go:linkname prefix github.com/gitbundle/server/pkg/migrations.prefix
	prefix string
	//go:linkname slowTest github.com/gitbundle/server/pkg/migrations.slowTest
	slowTest time.Duration
	//go:linkname slowFlush github.com/gitbundle/server/pkg/migrations.slowFlush
	slowFlush time.Duration

	//go:linkname writerCloser github.com/gitbundle/server/pkg/migrations.writerCloser
	writerCloser *testLoggerWriterCloser
)

type testLoggerWriterCloser struct {
	sync.RWMutex
	t []*testing.TB
}

//go:linkname (*testLoggerWriterCloser).Write github.com/gitbundle/server/pkg/migrations.(*testLoggerWriterCloser).Write
func (w *testLoggerWriterCloser) Write(p []byte) (int, error)

//go:linkname (*testLoggerWriterCloser).Close github.com/gitbundle/server/pkg/migrations.(*testLoggerWriterCloser).Close
func (w *testLoggerWriterCloser) Close() error
