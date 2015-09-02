// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kemi

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

// Tar implements interfaces for unpackig tar archives
type Tar struct {
	src  string
	file io.Reader
}

// Support returns tar
func (t *Tar) Support() string {
	return "tar"
}

// Open opens src and keeps *os.File for further processing, src should be a tar archive
// file.
func (t *Tar) Open(src string) (Copier, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	t.file = file
	return t, nil
}

// OpenFromReader keeps the src for further processing, src should be reading from a
// tar compatible content
func (t *Tar) OpenFromReader(src io.Reader) Copier {
	t.file = src
	return t
}

// CopyTo unpacks the underlying tar archive source to dest
func (t *Tar) CopyTo(dest string) error {
	reader := tar.NewReader(t.file)

	var terr error
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		out := filepath.Join(dest, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			continue
		}
		os.MkdirAll(filepath.Dir(out), directoryPerm)

		err = writeFile(out, reader, info.Mode())
		if err != nil {
			terr = err
			break
		}

	}
	return terr
}

// Reader returns the underlying reader
func (t *Tar) Reader() io.Reader {
	return t.file
}

// Close closes the underlying tar archive file.
func (t *Tar) Close() {
	switch t.file.(type) {
	case io.ReadCloser:
		rc := t.file.(io.ReadCloser)
		rc.Close()
	}
}
