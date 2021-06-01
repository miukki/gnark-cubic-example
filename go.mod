module example.com/cubicfrontend

go 1.16

replace example.com/cubic => ../cubic

require (
	example.com/cubic v1.1.0
	github.com/consensys/gnark v0.4.0
	github.com/consensys/gnark-crypto v0.4.1-0.20210428083642-6bd055b79906
	github.com/fxamacker/cbor/v2 v2.3.0 // indirect
	golang.org/x/sys v0.0.0-20210531225629-47163c9f4e4f // indirect
)
