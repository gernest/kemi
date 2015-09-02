// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kemi

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Gz implements interfaces for unpacking gzip archives
type Gz struct {
	file io.Reader
	src  string
}

// Support returns gz
func (g *Gz) Support() string {
	return "gz"
}

// Open reads from src and returns a Copier which can unpack the source using gzip format.
func (g *Gz) Open(src string) (Copier, error) {
	file, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, err
	}
	b := bytes.NewReader(file)
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}
	g.file = r
	g.src = src
	return g, nil
}

// OpenFromReader reads from src returns a Copier which can unpack the source using gzip format,
// if it fails to read then it panics.
func (g *Gz) OpenFromReader(src io.Reader) Copier {
	g.file = src
	return g
}

// CopyTo copies ungzipped contents to dest
func (g *Gz) CopyTo(dest string) error {
	os.MkdirAll(filepath.Dir(dest), directoryPerm)
	out, err := ioutil.ReadAll(g.file)
	if err != nil {
		return err
	}

	info, err := os.Stat(g.src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dest, out, info.Mode())
}

// Reader returns an io.Reader that reads from un gzipped contents.
func (g *Gz) Reader() io.Reader {
	return g.file
}

// Close closes the underlying io.Reader
func (g *Gz) Close() {
	switch g.file.(type) {
	case io.ReadCloser:
		g.file.(io.ReadCloser).Close()
	}
}
