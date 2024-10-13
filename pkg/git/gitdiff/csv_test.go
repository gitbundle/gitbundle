// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package gitdiff_test

import (
	"encoding/csv"
	"strings"
	"testing"

	csv_module "github.com/gitbundle/modules/csv"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/git/gitdiff"

	"github.com/stretchr/testify/assert"
)

func TestCSVDiff(t *testing.T) {
	cases := []struct {
		diff  string
		base  string
		head  string
		cells [][]gitdiff.TableDiffCellType
	}{
		// case 0 - initial commit of a csv
		{
			diff: `diff --git a/unittest.csv b/unittest.csv
--- a/unittest.csv
+++ b/unittest.csv
@@ -0,0 +1,2 @@
+col1,col2
+a,a`,
			base: "",
			head: `col1,col2
a,a`,
			cells: [][]gitdiff.TableDiffCellType{
				{gitdiff.TableDiffCellAdd, gitdiff.TableDiffCellAdd},
				{gitdiff.TableDiffCellAdd, gitdiff.TableDiffCellAdd},
			},
		},
		// case 1 - adding 1 row at end
		{
			diff: `diff --git a/unittest.csv b/unittest.csv
--- a/unittest.csv
+++ b/unittest.csv
@@ -1,2 +1,3 @@
 col1,col2
-a,a
+a,a
+b,b`,
			base: `col1,col2
a,a`,
			head: `col1,col2
a,a
b,b`,
			cells: [][]gitdiff.TableDiffCellType{
				{gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellUnchanged},
				{gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellUnchanged},
				{gitdiff.TableDiffCellAdd, gitdiff.TableDiffCellAdd},
			},
		},
		// case 2 - row deleted
		{
			diff: `diff --git a/unittest.csv b/unittest.csv
--- a/unittest.csv
+++ b/unittest.csv
@@ -1,3 +1,2 @@
 col1,col2
-a,a
 b,b`,
			base: `col1,col2
a,a
b,b`,
			head: `col1,col2
b,b`,
			cells: [][]gitdiff.TableDiffCellType{
				{gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellUnchanged},
				{gitdiff.TableDiffCellDel, gitdiff.TableDiffCellDel},
				{gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellUnchanged},
			},
		},
		// case 3 - row changed
		{
			diff: `diff --git a/unittest.csv b/unittest.csv
--- a/unittest.csv
+++ b/unittest.csv
@@ -1,2 +1,2 @@
 col1,col2
-b,b
+b,c`,
			base: `col1,col2
b,b`,
			head: `col1,col2
b,c`,
			cells: [][]gitdiff.TableDiffCellType{
				{gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellUnchanged},
				{gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellChanged},
			},
		},
		// case 4 - all deleted
		{
			diff: `diff --git a/unittest.csv b/unittest.csv
--- a/unittest.csv
+++ b/unittest.csv
@@ -1,2 +0,0 @@
-col1,col2
-b,c`,
			base: `col1,col2
b,c`,
			head: "",
			cells: [][]gitdiff.TableDiffCellType{
				{gitdiff.TableDiffCellDel, gitdiff.TableDiffCellDel},
				{gitdiff.TableDiffCellDel, gitdiff.TableDiffCellDel},
			},
		},
		// case 5 - renames first column
		{
			diff: `diff --git a/unittest.csv b/unittest.csv
--- a/unittest.csv
+++ b/unittest.csv
@@ -1,3 +1,3 @@
-col1,col2,col3
+cola,col2,col3
 a,b,c`,
			base: `col1,col2,col3
a,b,c`,
			head: `cola,col2,col3
a,b,c`,
			cells: [][]gitdiff.TableDiffCellType{
				{gitdiff.TableDiffCellDel, gitdiff.TableDiffCellAdd, gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellUnchanged},
				{gitdiff.TableDiffCellDel, gitdiff.TableDiffCellAdd, gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellUnchanged},
			},
		},
		// case 6 - inserts a column after first, deletes last column
		{
			diff: `diff --git a/unittest.csv b/unittest.csv
--- a/unittest.csv
+++ b/unittest.csv
@@ -1,2 +1,2 @@
-col1,col2,col3
-a,b,c
+col1,col1a,col2
+a,d,b`,
			base: `col1,col2,col3
a,b,c`,
			head: `col1,col1a,col2
a,d,b`,
			cells: [][]gitdiff.TableDiffCellType{
				{gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellAdd, gitdiff.TableDiffCellDel, gitdiff.TableDiffCellMovedUnchanged},
				{gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellAdd, gitdiff.TableDiffCellDel, gitdiff.TableDiffCellMovedUnchanged},
			},
		},
		// case 7 - deletes first column, inserts column after last
		{
			diff: `diff --git a/unittest.csv b/unittest.csv
--- a/unittest.csv
+++ b/unittest.csv
@@ -1,2 +1,2 @@
-col1,col2,col3
-a,b,c
+col2,col3,col4
+b,c,d`,
			base: `col1,col2,col3
a,b,c`,
			head: `col2,col3,col4
b,c,d`,
			cells: [][]gitdiff.TableDiffCellType{
				{gitdiff.TableDiffCellDel, gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellAdd},
				{gitdiff.TableDiffCellDel, gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellUnchanged, gitdiff.TableDiffCellAdd},
			},
		},
		// case 8 - two columns deleted, 2 added
		{
			diff: `diff --git a/unittest.csv b/unittest.csv
--- a/unittest.csv
+++ b/unittest.csv
@@ -1,2 +1,2 @@
-col1,col2,col
-a,b,c
+col3,col4,col5
+c,d,e`,
			base: `col1,col2,col3
a,b,c`,
			head: `col3,col4,col5
c,d,e`,
			cells: [][]gitdiff.TableDiffCellType{
				{gitdiff.TableDiffCellDel, gitdiff.TableDiffCellMovedUnchanged, gitdiff.TableDiffCellDel, gitdiff.TableDiffCellAdd, gitdiff.TableDiffCellAdd},
				{gitdiff.TableDiffCellDel, gitdiff.TableDiffCellMovedUnchanged, gitdiff.TableDiffCellDel, gitdiff.TableDiffCellAdd, gitdiff.TableDiffCellAdd},
			},
		},
	}

	for n, c := range cases {
		diff, err := gitdiff.ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(c.diff), "")
		if err != nil {
			t.Errorf("ParsePatch failed: %s", err)
		}

		var baseReader *csv.Reader
		if len(c.base) > 0 {
			baseReader, err = csv_module.CreateReaderAndDetermineDelimiter(nil, strings.NewReader(c.base))
			if err != nil {
				t.Errorf("CreateReaderAndDetermineDelimiter failed: %s", err)
			}
		}
		var headReader *csv.Reader
		if len(c.head) > 0 {
			headReader, err = csv_module.CreateReaderAndDetermineDelimiter(nil, strings.NewReader(c.head))
			if err != nil {
				t.Errorf("CreateReaderAndDetermineDelimiter failed: %s", err)
			}
		}

		result, err := gitdiff.CreateCsvDiff(diff.Files[0], baseReader, headReader)
		assert.NoError(t, err)
		assert.Len(t, result, 1, "case %d: should be one section", n)

		section := result[0]
		assert.Len(t, section.Rows, len(c.cells), "case %d: should be %d rows", n, len(c.cells))

		for i, row := range section.Rows {
			assert.Len(t, row.Cells, len(c.cells[i]), "case %d: row %d should have %d cells", n, i, len(c.cells[i]))
			for j, cell := range row.Cells {
				assert.Equal(t, c.cells[i][j], cell.Type, "case %d: row %d cell %d should be equal", n, i, j)
			}
		}
	}
}
