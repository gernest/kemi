// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kemi

import (
	"os"
	"testing"
)

func TestTar(t *testing.T) {
	var (
		sourceFile = "testdata/compress/hello.tar"
		destDir    = "testdata/dest/tar-only"
	)
	tarUnpacker := &Tar{}

	os.RemoveAll(destDir)
	if tarUnpacker.Support() != "tar" {
		t.Errorf("expected tar got %s", tarUnpacker.Support())
	}

	tarCopier, err := tarUnpacker.Open(sourceFile)
	if err != nil {
		t.Fatal(err)
	}
	defer tarCopier.Close()
	err = tarCopier.CopyTo(destDir)
	if err != nil {
		t.Error(err)
	}
}
