# kemi [![Build Status](https://travis-ci.org/gernest/kemi.svg)](https://travis-ci.org/gernest/kemi)[![Coverage Status](https://coveralls.io/repos/gernest/kemi/badge.svg?branch=master&service=github)](https://coveralls.io/github/gernest/kemi?branch=master)[![GoDoc](https://godoc.org/github.com/gernest/kemi?status.svg)](https://godoc.org/github.com/gernest/kemi)

kemi helps you win at unpacking archive files with Go(a.k.a golang). Kemi supports .tar, .tar.gz
.zip out of the box, with option to extend and add more support with your own or your
neighbor's  implementations.

## Current archives supported

* .tar

* .tar.gz

* .zip


## Installation

	go get github.com/gernest/kemi

## How to use

### Case you want to unpack an archive file, from your golang applications.

say you want to unpack `foo.tar.gz` to `path/to/bar` 

```go
package main

import(
	"github.com/gernest/kemi"
)

func main(){
	
	err:=kemi.Unpack("foo.tar.gz","path/to/bar")
	if err!=nil{
		// Handle your error
	}
}
```

### Case you love your own implementation

If you want to use your own unpacking implementation, please see the `Unparker` interface definition. And you can check on the `Zip` or `Gz` or `Tar` structs for how to implement.

Say you have implemented  your own `Unpacker` named `Desux` for your fancy archive format `sux`

You wan  unpack `foo.sux.gz` to `path/to/bar` like this

```go
package main

import(
	"github.com/gernest/kemi"
)

func main(){
	
	// register the sux implementation
	kemi.Register(Desux)
	
	err:=kemi.Unpack("foo.sux.gz","path/to/bar")
	if err!=nil{
		// Handle your error
	}
}
```


# Contributing

Start with clicking the star button to make the author and his neighbors happy. Then fork it and submit a pull request for whatever change you want to be added to this project.

Or Open an issue for any questions.

## Author
Geofrey Ernest <geofreyernest@live.com>

twitter  : [@gernesti](https://twitter.com/gernesti)

Facebook : [Geofrey Ernest](https://www.facebook.com/geofrey.ernest.35)



## Licence
This project is released under MIT licence see [LICENCE](LICENCE) for more details.
