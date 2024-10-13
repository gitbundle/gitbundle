// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"
	_ "unsafe"

	"github.com/gitbundle/modules/base"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/timeutil"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func TestMain(m *testing.M) {
	giteaRoot := base.SetupGiteaRoot()
	if giteaRoot == "" {
		fmt.Println("Environment variable $GITBUNDLE_ROOT not set")
		os.Exit(1)
	}
	giteaBinary := "gitea"
	if runtime.GOOS == "windows" {
		giteaBinary += ".exe"
	}
	setting.AppPath = path.Join(giteaRoot, giteaBinary)
	if _, err := os.Stat(setting.AppPath); err != nil {
		fmt.Printf("Could not find gitea binary at %s\n", setting.AppPath)
		os.Exit(1)
	}

	giteaConf := os.Getenv("GITBUNDLE_CONF")
	if giteaConf == "" {
		giteaConf = path.Join(filepath.Dir(setting.AppPath), "integrations/sqlite.ini")
		fmt.Printf("Environment variable $GITBUNDLE_CONF not set - defaulting to %s\n", giteaConf)
	}

	if !path.IsAbs(giteaConf) {
		setting.CustomConf = path.Join(giteaRoot, giteaConf)
	} else {
		setting.CustomConf = giteaConf
	}

	tmpDataPath, err := os.MkdirTemp("", "data")
	if err != nil {
		fmt.Printf("Unable to create temporary data path %v\n", err)
		os.Exit(1)
	}

	setting.AppDataPath = tmpDataPath

	setting.SetCustomPathAndConf("", "", "")
	setting.LoadForTest()
	if err = git.InitOnceWithSync(context.Background()); err != nil {
		fmt.Printf("Unable to InitOnceWithSync: %v\n", err)
		os.Exit(1)
	}
	git.CheckLFSVersion()
	setting.InitDBConfig()
	setting.NewLogServices(true)

	exitStatus := m.Run()

	if err := removeAllWithRetry(setting.RepoRootPath); err != nil {
		fmt.Fprintf(os.Stderr, "os.RemoveAll: %v\n", err)
	}
	if err := removeAllWithRetry(tmpDataPath); err != nil {
		fmt.Fprintf(os.Stderr, "os.RemoveAll: %v\n", err)
	}
	os.Exit(exitStatus)
}

func removeAllWithRetry(dir string) error {
	var err error
	for i := 0; i < 20; i++ {
		err = os.RemoveAll(dir)
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	return err
}

func Test_dropTableColumns(t *testing.T) {
	x, deferable := prepareTestEnv(t, 0)
	if x == nil || t.Failed() {
		defer deferable()
		return
	}
	defer deferable()

	type DropTest struct {
		ID            int64 `xorm:"pk autoincr"`
		FirstColumn   string
		ToDropColumn  string `xorm:"unique"`
		AnotherColumn int64
		CreatedUnix   timeutil.TimeStamp `xorm:"INDEX created"`
		UpdatedUnix   timeutil.TimeStamp `xorm:"INDEX updated"`
	}

	columns := []string{
		"first_column",
		"to_drop_column",
		"another_column",
		"created_unix",
		"updated_unix",
	}

	for i := range columns {
		x.SetMapper(names.GonicMapper{})
		if err := x.Sync2(new(DropTest)); err != nil {
			t.Errorf("unable to create DropTest table: %v", err)
			return
		}
		sess := x.NewSession()
		if err := sess.Begin(); err != nil {
			sess.Close()
			t.Errorf("unable to begin transaction: %v", err)
			return
		}
		if err := dropTableColumns(sess, "drop_test", columns[i:]...); err != nil {
			sess.Close()
			t.Errorf("Unable to drop columns[%d:]: %s from drop_test: %v", i, columns[i:], err)
			return
		}
		if err := sess.Commit(); err != nil {
			sess.Close()
			t.Errorf("unable to commit transaction: %v", err)
			return
		}
		sess.Close()
		if err := x.DropTables(new(DropTest)); err != nil {
			t.Errorf("unable to drop table: %v", err)
			return
		}
		for j := range columns[i+1:] {
			x.SetMapper(names.GonicMapper{})
			if err := x.Sync2(new(DropTest)); err != nil {
				t.Errorf("unable to create DropTest table: %v", err)
				return
			}
			dropcols := append([]string{columns[i]}, columns[j+i+1:]...)
			sess := x.NewSession()
			if err := sess.Begin(); err != nil {
				sess.Close()
				t.Errorf("unable to begin transaction: %v", err)
				return
			}
			if err := dropTableColumns(sess, "drop_test", dropcols...); err != nil {
				sess.Close()
				t.Errorf("Unable to drop columns: %s from drop_test: %v", dropcols, err)
				return
			}
			if err := sess.Commit(); err != nil {
				sess.Close()
				t.Errorf("unable to commit transaction: %v", err)
				return
			}
			sess.Close()
			if err := x.DropTables(new(DropTest)); err != nil {
				t.Errorf("unable to drop table: %v", err)
				return
			}
		}
	}
}

//go:linkname dropTableColumns github.com/gitbundle/server/pkg/migrations.dropTableColumns
func dropTableColumns(sess *xorm.Session, tableName string, columnNames ...string) (err error)

//go:linkname prepareTestEnv github.com/gitbundle/server/pkg/migrations.prepareTestEnv
func prepareTestEnv(t *testing.T, skip int, syncModels ...interface{}) (*xorm.Engine, func())

//go:linkname deleteDB github.com/gitbundle/server/pkg/migrations.deleteDB
func deleteDB() error
