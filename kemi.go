// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package kemi helps you win at unpacking archive files. Kemi supports .tar, .tar.gz
.zip out of the box, with option to extend and add more support with your own or your
neighbor's  implementations.
*/
package kemi

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var registeredUnpackers = map[string]Unpacker{
	"tar": &Tar{},
	"gz":  &Gz{},
	"zip": &Zip{},
}

var (
	directoryPerm os.FileMode = 0755

	// ErrNotSupported is the error used for files whose unpackers are not registered
	ErrNotSupported = errors.New("kemi: File not supported")
)

// Unpacker is an interface that wraps methods for initializing Copier implementations.
type Unpacker interface {

	// Support returns the file extension supported
	// by this unpacker without the  dot  e.g tar
	Support() string

	// Open  psses the archive filepath as a string
	Open(string) (Copier, error)

	// uses the io.Reader to return the Copier
	OpenFromReader(io.Reader) Copier
}

// Copier is the interface that wraps methods for copying extracted archive files.
type Copier interface {

	// CopyTo copies the unpacked archive to dest, it depends on what
	CopyTo(dest string) error

	// Reader returns an io.Reader that can be used by other unpackers.
	Reader() io.Reader

	// Close closes the underlying reader. It is up to the implementation to make sure
	// when reading is done to close the reader.
	Close()
}

// Unpack unpacks contents of src into dest. Where src if the path to an archive file.
func Unpack(src, dest string) error {
	unPackers, err := listUnpackers(src)
	if err != nil {
		return err
	}
	var copier Copier

	for k, v := range unPackers {
		if k == 0 {
			cp, err := v.Open(src)
			if err != nil {
				return err
			}
			copier = cp
		}
		copier = v.OpenFromReader(copier.Reader())
	}
	defer copier.Close()
	err = copier.CopyTo(dest)
	if err != nil {
		return err
	}
	return nil
}

// Optain the list of packers by observing the file extensions.
// this means, .tar.gz will have a gz and tar unpackers in the order.
func listUnpackers(src string) (rst []Unpacker, err error) {
	base := filepath.Base(src)
	c := strings.Split(base, ".")
	cLen := len(c)
	switch {
	case cLen == 2:
		outer, ok := registeredUnpackers[c[cLen-1]]
		if !ok {
			err = ErrNotSupported
			return
		}
		rst = append(rst, outer)
	case cLen > 2:
		outer, ok := registeredUnpackers[c[cLen-1]]
		if !ok {
			err = ErrNotSupported
			return
		}
		rst = append(rst, outer)
		if len(c[cLen-2]) <= 3 {
			inner, ok := registeredUnpackers[c[cLen-2]]
			if !ok {
				//err = ErrNotSupported
				return
			}
			rst = append(rst, inner)
		}
	}
	return
}

// Creates a new file name, and copies what is read from data.
func writeFile(name string, data io.Reader, perm os.FileMode) error {
	out, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, cerr := io.Copy(out, data); cerr != nil {
		return cerr
	}
	return nil
}

//Register registers u as  Unpacker for whatever format u supports.
func Register(u Unpacker) {
	registeredUnpackers[u.Support()] = u
}
