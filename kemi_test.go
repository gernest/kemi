// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kemi

import (
	"os"
	"testing"
)

func TestUnpack(t *testing.T) {

	archives := []struct {
		src, dest string
	}{
		{
			"testdata/compress/hello.tar", "testdata/dest/tar",
		},
		{
			"testdata/compress/hello.tar.gz", "testdata/dest/tar-gz",
		},
		{
			"testdata/compress/hello.zip", "testdata/dest/zip",
		},
	}

	// remove evrything in the destintion
	os.RemoveAll("testdata/dest")

	for _, a := range archives {
		err := Unpack(a.src, a.dest)
		if err != nil {
			t.Errorf("unpacking %s %v \n", a.src, err)
		}
	}

}

func TestGetUnpackers(t *testing.T) {
	src := []struct {
		src string
		rst []Unpacker
	}{
		{
			"sample.tar.gz",
			[]Unpacker{&Gz{}, &Tar{}},
		},
		{
			"sample.mambo.tar.gz",
			[]Unpacker{&Gz{}, &Tar{}},
		},
		{
			"sample.mambo.sux.tar.gz",
			[]Unpacker{&Gz{}, &Tar{}},
		},
	}
	for _, v := range src {
		p, err := listUnpackers(v.src)
		if err != nil {
			t.Error(err)
		}

		for k := range p {
			if p[k].Support() != v.rst[k].Support() {
				t.Errorf("expecetd %s got %s", v.rst[k].Support(), p[k].Support())
			}
		}
	}

}
