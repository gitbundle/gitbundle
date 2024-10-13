// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package doctor

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/db"
)

// Check represents a Doctor check
type Check struct {
	Title                      string
	Name                       string
	IsDefault                  bool
	Run                        func(ctx context.Context, logger log.Logger, autofix bool) error
	AbortIfFailed              bool
	SkipDatabaseInitialization bool
	Priority                   int
}

type wrappedLevelLogger struct {
	log.LevelLogger
}

func (w *wrappedLevelLogger) Log(skip int, level log.Level, format string, v ...interface{}) error {
	return w.LevelLogger.Log(
		skip+1,
		level,
		" - %s "+format,
		append(
			[]interface{}{
				log.NewColoredValueBytes(
					fmt.Sprintf("[%s]", strings.ToUpper(level.String()[0:1])),
					level.Color()),
			}, v...)...)
}

func initDBDisableConsole(ctx context.Context, disableConsole bool) error {
	setting.LoadFromExisting()
	setting.InitDBConfig()

	setting.NewXORMLogService(disableConsole)
	if err := db.InitEngine(ctx); err != nil {
		return fmt.Errorf("db.InitEngine: %w", err)
	}
	// some doctor sub-commands need to use git command
	if err := git.InitOnceWithSync(ctx); err != nil {
		return fmt.Errorf("git.InitOnceWithSync: %w", err)
	}
	return nil
}

// Checks is the list of available commands
var Checks []*Check

// RunChecks runs the doctor checks for the provided list
func RunChecks(ctx context.Context, logger log.Logger, autofix bool, checks []*Check) error {
	wrappedLogger := log.LevelLoggerLogger{
		LevelLogger: &wrappedLevelLogger{logger},
	}

	dbIsInit := false
	for i, check := range checks {
		if !dbIsInit && !check.SkipDatabaseInitialization {
			// Only open database after the most basic configuration check
			setting.EnableXORMLog = false
			if err := initDBDisableConsole(ctx, true); err != nil {
				logger.Error("Error whilst initializing the database: %v", err)
				logger.Error(
					"Check if you are using the right config file. You can use a --config directive to specify one.",
				)
				return nil
			}
			dbIsInit = true
		}
		logger.Info("[%d] %s", log.NewColoredIDValue(i+1), check.Title)
		logger.Flush()
		if err := check.Run(ctx, &wrappedLogger, autofix); err != nil {
			if check.AbortIfFailed {
				logger.Critical("FAIL")
				return err
			}
			logger.Error("ERROR")
		} else {
			logger.Info("OK")
			logger.Flush()
		}
	}
	return nil
}

// Register registers a command with the list
func Register(command *Check) {
	Checks = append(Checks, command)
	sort.SliceStable(Checks, func(i, j int) bool {
		if Checks[i].Priority == Checks[j].Priority {
			return Checks[i].Name < Checks[j].Name
		}
		if Checks[i].Priority == 0 {
			return false
		}
		return Checks[i].Priority < Checks[j].Priority
	})
}
