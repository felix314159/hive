//go:build tools
// +build tools

package tools

import (
	_ "github.com/fjl/gencodec" // stop removing this from go.mod we need it for go:gen, this line can't me in main because then it gets removed because it is a 'tool' because it declares a main package. check out go 1.24 release notes for tools.go special meaning
)

// you can now run 'go build -tags tools', then run 'cd ./simulators/ethereum/engine && go generate -v ./...'
/* Go generate fails on my machine tho:
types/types.go
panic: unsupported version: 2

goroutine 1 [running]:
internal/pkgbits.NewPkgDecoder({0x353cd0?, 0x220ac?}, {0x4005802000, 0x11091})
	/usr/local/go/src/internal/pkgbits/decoder.go:87 +0x334
go/internal/gcimporter.Import(0x40000b86c0, 0x40000aaa50, {0x353cd0, 0xd}, {0x0?, 0x40000aa510?}, 0x0?)
	/usr/local/go/src/go/internal/gcimporter/gcimporter.go:239 +0x5d4
go/importer.(*gcimports).ImportFrom(0x40002ddb78?, {0x353cd0?, 0x2f0380?}, {0x0?, 0x40004eabb8?}, 0x3?)
	/usr/local/go/src/go/importer/importer.go:102 +0x4c
go/importer.(*gcimports).Import(0x0?, {0x353cd0?, 0x40004eabb8?})
	/usr/local/go/src/go/importer/importer.go:95 +0x2c
main.(*fileScope).addImport(0x4010919c20, {0x353cd0, 0xd})
	/home/user/go/pkg/mod/github.com/fjl/gencodec@v0.1.1/types_util.go:217 +0x38
main.newMarshalerType(0x40000b8680, {0x3ec860, 0x40000be108}, 0x401217d9d0)
	/home/user/go/pkg/mod/github.com/fjl/gencodec@v0.1.1/main.go:363 +0x190
main.(*Config).process(0x40002dde60)
	/home/user/go/pkg/mod/github.com/fjl/gencodec@v0.1.1/main.go:251 +0x240
main.main()
	/home/user/go/pkg/mod/github.com/fjl/gencodec@v0.1.1/main.go:204 +0x2c8
types/types.go:159: running "gencodec": exit status 2


*/