ASSIMP-GO
=========

This is a [Go][golang] library that wraps the use of the Open Asset Import Library known as [assimp][assimp-link].


Requirements
------------

This does require `cgo` which means that `gcc` should be in your path
when trying to build using this library.

Software requirements:

* [Mathgl][mgl] - for 3d math
* [Gombz][gombz-link] - used as a file format and data structure for
  the information pulled from assimp.

C code is derived from [Assimp 5.0.1](https://github.com/assimp/assimp/releases/tag/v5.0.1).

## Installation

```
go get github.com/dabasan/assimp-go
```

### Linux

The above command reports some errors because it does not include the Linux shared object.

```
# github.com/dabasan/assimp-go
/usr/lib/go-1.10/pkg/tool/linux_amd64/link: running gcc failed: exit status 1
/usr/bin/ld: cannot find -lassimp
collect2: error: ld returned 1 exit status
```

1. Download [libassimp](https://github.com/Dabasan/assimp-go/releases/download/v1.0.0/libassimp.tar.gz) and extract it
2. Move the files into $GOPATH/src/github.com/dabasan/assimp-go/assimp/C/bin

## Run

1. Build an executable with `go build`
2. Copy the library file(s) (Windows .dll or Linux .so) to the same directory as the executable
3. Run the executable

### Donwload

- [DLL for Windows](https://github.com/Dabasan/assimp-go/releases/download/v1.0.0/assimp.zip)
- [SO for Linux](https://github.com/Dabasan/assimp-go/releases/download/v1.0.0/libassimp.tar.gz)

Usage
-----

The module can be used to load files supported by [Assimp][assimp-link] and converts
them into the [Gombz][gombz-link] meshes. Once imported, you can load an file
(e.g. .OBJ or .FBX file) by using the following call:


```
srcMeshes, err := assimp.ParseFile(srcFilepath)
```

LICENSE
=======

Assimp-go is released under the BSD license. See the [LICENSE][license-link] file for more details.


[golang]: https://golang.org/
[license-link]: https://raw.githubusercontent.com/tbogdala/assimp-go/master/LICENSE
[assimp-link]: http://assimp.sourceforge.net/
[mgl]: https://github.com/go-gl/mathgl
[gombz-link]: https://github.com/tbogdala/gombz
