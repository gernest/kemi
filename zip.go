// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kemi

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

//Zip implements interfaces for unpacking zip archives
type Zip struct {
	file io.Reader
	src  string
}

// Support returns zip
func (z *Zip) Support() string {
	return "zip"
}

// Open opens the file and keep the *os.File forlater processing
func (z *Zip) Open(src string) (Copier, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	z.file = file
	z.src = src
	return z, nil
}

// OpenFromReader takes the src and keep if for further processing.
//
// NOTE due to the nature of the zip package from the standard library, it requires
// an io.ReaderAt as an input intead of io.Reader, we use the bytes.Reader which implements
// io.ReaderAt.
//
// TODO alternative way?
func (z *Zip) OpenFromReader(src io.Reader) Copier {
	b := &bytes.Buffer{}
	io.Copy(b, src)
	z.file = bytes.NewReader(b.Bytes())
	return z
}

// CopyTo copies the contents of the archive file to the dest path, dest should be a diretory
// its weird to unpack a zip file to a file.
func (z *Zip) CopyTo(dest string) error {
	var reader *zip.Reader
	switch z.file.(type) {
	case *os.File:
		f := z.file.(*os.File)
		info, err := f.Stat()
		if err != nil {
			return err
		}
		rd, err := zip.NewReader(f, info.Size())
		if err != nil {
			return err
		}
		reader = rd
	case *bytes.Reader:
		f := z.file.(*bytes.Reader)
		rd, err := zip.NewReader(f, int64(f.Len()))
		if err != nil {
			return err
		}
		reader = rd
	}
	defer z.Close()

	for _, f := range reader.File {
		info := f.FileInfo()
		if info.IsDir() {
			continue
		}
		out := filepath.Join(dest, f.Name)
		os.MkdirAll(filepath.Dir(out), directoryPerm)
		err := z.witeFile(out, f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *Zip) witeFile(name string, file *zip.File) error {
	if info := file.FileInfo(); info.IsDir() {
		return nil
	}
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	return writeFile(name, rc, file.Mode())
}

// Reader returns the underlying io.Reader
func (z *Zip) Reader() io.Reader {
	return z.file
}

// Close stop reading from the source io.Reader.
func (z *Zip) Close() {
	switch z.file.(type) {
	case io.ReadCloser:
		z.file.(io.ReadCloser).Close()
	}
}
