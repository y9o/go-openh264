# Go binding for OpenH264

Call the OpenH264 [binaries library](https://github.com/cisco/openh264/releases) from Go to perform encoding and decoding without using Cgo.

Since [purego](https://github.com/ebitengine/purego) is used instead of Cgo, there is no need to prepare a C compiler at build time, even in a Windows.
However, Cgo is used to generate codec_api.go.

Examples for encoding and decoding image.YCbCr to video can be found in openh264_test.go.

## Cgo

`go build -tags cgoopenh264`

