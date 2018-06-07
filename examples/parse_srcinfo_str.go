package main

import (
	"fmt"
	"github.com/morganamilo/go-srcinfo"
)

const str = `
pkgbase = gdc-bin
	pkgver = 6.3.0+2.068.2
	pkgrel = 1
	url = https://gdcproject.org/
	arch = i686
	arch = x86_64
	license = GPL
	source_i686 = http://gdcproject.org/downloads/binaries/6.3.0/i686-linux-gnu/gdc-6.3.0+2.068.2.tar.xz
	md5sums_i686 = cc8dcd66b189245e39296b1382d0dfcc
	source_x86_64 = http://gdcproject.org/downloads/binaries/6.3.0/x86_64-linux-gnu/gdc-6.3.0+2.068.2.tar.xz
	md5sums_x86_64 = 16d3067ebb3938dba46429a4d9f6178f

pkgname = gdc-bin
	pkgdesc = Compiler for D programming language which uses gcc backend
	depends = gdc-gcc
	depends = perl
	depends = binutils
	depends = libgphobos
	provides = d-compiler=2.068.2
	provides = gdc=6.3.0+2.068.2

pkgname = gdc-gcc
	pkgdesc = The GNU Compiler Collection - C and C++ frontends (from GDC, gdcproject.org)
	provides = gcc=6.3.0
	provides = gcc-libs=6.3.0

pkgname = libgphobos-lib32
	pkgdesc = Standard library for D programming language, GDC port
	provides = d-runtime-lib32
	provides = d-stdlib-lib32
`

func main() {
	info, err := srcinfo.Parse(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info)
}
