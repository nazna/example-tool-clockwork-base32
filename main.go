package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unsafe"

	"github.com/szktty/go-clockwork-base32"
	"github.com/urfave/cli/v2"
)

var Version = "unset"

//lint:ignore U1000 used by wasm
var buf [1024]byte

//lint:ignore U1000 used by wasm
//export getBuffer
func getBuffer() *byte {
	return &buf[0]
}

//lint:ignore U1000 used by wasm
func _cwbase32(input string, shouldDecode bool) (uint32, uint32) {
	if shouldDecode {
		decoded, _ := clockwork.Decode([]byte(input))
		trimmed := bytes.TrimSuffix(decoded, []byte{'\n'})
		p := &trimmed[0]
		up := uintptr(unsafe.Pointer(p))
		return uint32(up), uint32(len(trimmed))
	} else {
		bytes := clockwork.Encode([]byte(input))
		p := &bytes[0]
		up := uintptr(unsafe.Pointer(p))
		return uint32(up), uint32(len(bytes))
	}
}

//lint:ignore U1000 used by wasm
//export cwbase32
func cwbase32(input string, shouldDecode bool) uint32 {
	p, _ := _cwbase32(input, shouldDecode)

	return p
}

//lint:ignore U1000 used by wasm
//export cwbase32len
func cwbase32len(input string, shouldDecode bool) uint32 {
	_, len := _cwbase32(input, shouldDecode)

	return len
}

func preprocess(filename string) (string, error) {
	var r io.Reader

	switch filename {
	case "", "-":
		r = os.Stdin
	default:
		f, err := os.Open(filename)

		if err != nil {
			return "", err
		}
		defer f.Close()

		r = f
	}

	bytes, err := io.ReadAll(r)

	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(bytes), "\n"), nil
}

func cwbase32cli(input string, shouldDecode bool) (string, error) {
	if shouldDecode {
		decoded, err := clockwork.Decode([]byte(input))

		if err != nil {
			return "", err
		}

		return string(bytes.TrimSuffix(decoded, []byte{'\n'})), nil
	} else {
		encoded := clockwork.Encode([]byte(input))
		return string(encoded), nil
	}
}

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "output version information and exit",
	}

	app := &cli.App{
		Name:                 "cwbase32",
		Usage:                "Clockwork Base32 encode or decode FILE, or standard input, to standard output.",
		UsageText:            "cwbase32 [OPTION]... [FILE]",
		Version:              Version,
		EnableBashCompletion: true,
		HideHelpCommand:      true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "decode",
				Aliases: []string{"d"},
				Usage:   "decode data",
			},
		},
		Action: func(ctx *cli.Context) error {
			filename := ctx.Args().First()

			input, err := preprocess(filename)

			if err != nil {
				return err
			}

			result, err := cwbase32cli(input, ctx.Bool("decode"))

			if err != nil {
				return err
			}

			fmt.Println(result)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
