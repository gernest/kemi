// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kemi

import (
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	var (
		sourceFile = "testdata/compress/hello.zip"
		destDir    = "testdata/dest/zip-only"
	)
	zipUnpacker := &Zip{}

	os.RemoveAll(destDir)
	if zipUnpacker.Support() != "zip" {
		t.Errorf("expected tar got %s", zipUnpacker.Support())
	}

	zipCopier, err := zipUnpacker.Open(sourceFile)
	if err != nil {
		t.Fatal(err)
	}
	defer zipCopier.Close()
	err = zipCopier.CopyTo(destDir)
	if err != nil {
		t.Error(err)
	}
}
