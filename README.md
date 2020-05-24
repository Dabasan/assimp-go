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

## Run

1. Build an executable with `go build`
2. Copy the library file(s) (Windows .dll or Linux .so) to the same directory as the executable
3. Run the executable

Native libraries are located in [bin](./assimp/C/bin).

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

This project is released under the MIT license.

See also: [License for the original project](https://github.com/tbogdala/assimp-go/blob/master/LICENSE)