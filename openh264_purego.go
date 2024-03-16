//go:build !cgoopenh264

// See the openh264 source code for how to use it.
package openh264

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

func (v *SBufferInfo) UsrData_sSystemBuffer() *SSysMEMBuffer {
	return (*SSysMEMBuffer)(unsafe.Pointer(&v.UsrData))
}

var pOpenh264Library uintptr = 0

var WelsGetCodecVersionEx func(*OpenH264Version) int

var WelsCreateDecoder func(**ISVCDecoder) int
var WelsDestroyDecoder func(*ISVCDecoder) int

var WelsCreateSVCEncoder func(**ISVCEncoder) int
var WelsDestroySVCEncoder func(*ISVCEncoder) int

// Load OpenH264 binary file (purego)
func Open(library string) error {
	if pOpenh264Library != 0 {
		return nil
	}
	libraryptr, err := openLibrary(library)
	if err != nil {
		return err
	}
	pOpenh264Library = libraryptr
	purego.RegisterLibFunc(&WelsCreateDecoder, libraryptr, "WelsCreateDecoder")
	purego.RegisterLibFunc(&WelsDestroyDecoder, libraryptr, "WelsDestroyDecoder")
	purego.RegisterLibFunc(&WelsGetCodecVersionEx, libraryptr, "WelsGetCodecVersionEx")
	purego.RegisterLibFunc(&WelsCreateSVCEncoder, libraryptr, "WelsCreateSVCEncoder")
	purego.RegisterLibFunc(&WelsDestroySVCEncoder, libraryptr, "WelsDestroySVCEncoder")
	return nil
}

// Close OpenH264 library (purego)
func Close() error {
	if pOpenh264Library != 0 {
		err := closeLibrary(pOpenh264Library)
		pOpenh264Library = 0
		return err
	}
	return nil
}

func WelsGetCodecVersion() (ver OpenH264Version) {
	if WelsGetCodecVersionEx != nil {
		WelsGetCodecVersionEx(&ver)
	}
	return
}

func (v *ISVCDecoder) Initialize(p *SDecodingParam) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.Initialize)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(p)),
	)
	return int(r1)
}

func (v *ISVCDecoder) Uninitialize() int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.Uninitialize)),
		uintptr(unsafe.Pointer(v)),
	)
	return int(r1)
}

func (v *ISVCDecoder) SetOption(eOptionId int, pOption *int) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.SetOption)),
		uintptr(unsafe.Pointer(v)),
		uintptr(eOptionId),
		uintptr(unsafe.Pointer(pOption)),
	)
	return int(r1)
}

func (v *ISVCDecoder) GetOption(eOptionId int, pOption *int) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.GetOption)),
		uintptr(unsafe.Pointer(v)),
		uintptr(eOptionId),
		uintptr(unsafe.Pointer(pOption)),
	)
	return int(r1)
}

func data2dst(data [3]*byte, ppDst *[3][]byte, pDstInfo *SBufferInfo) {
	ppDst[0] = nil
	ppDst[1] = nil
	ppDst[2] = nil
	if pDstInfo.IBufferStatus == 1 {
		height := pDstInfo.UsrData_sSystemBuffer().IHeight
		stride := pDstInfo.UsrData_sSystemBuffer().IStride
		ppDst[0] = unsafe.Slice((*byte)(data[0]), height*stride[0])
		height = height / 2
		ppDst[1] = unsafe.Slice((*byte)(data[1]), height*stride[1])
		ppDst[2] = unsafe.Slice((*byte)(data[2]), height*stride[1])
	}
}

func (v *ISVCDecoder) DecodeFrameNoDelay(pSrc []byte, iSrcLen int, ppDst *[3][]byte, pDstInfo *SBufferInfo) int {
	var data [3]*byte
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.DecodeFrameNoDelay)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&pSrc[0])),
		uintptr(iSrcLen),
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(pDstInfo)),
	)
	data2dst(data, ppDst, pDstInfo)
	return int(r1)
}
func (v *ISVCDecoder) DecodeFrame2(pSrc []byte, iSrcLen int, ppDst *[3][]byte, pDstInfo *SBufferInfo) int {
	var data [3]*byte
	var _pSrc unsafe.Pointer
	if pSrc == nil {
		_pSrc = nil
	} else {
		_pSrc = unsafe.Pointer(&pSrc[0])
	}
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.DecodeFrame2)),
		uintptr(unsafe.Pointer(v)),
		uintptr(_pSrc),
		uintptr(iSrcLen),
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(pDstInfo)),
	)
	data2dst(data, ppDst, pDstInfo)
	return int(r1)
}

func (v *ISVCDecoder) FlushFrame(ppDst *[3][]byte, pDstInfo *SBufferInfo) int {
	var data [3]*byte
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.FlushFrame)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(pDstInfo)),
	)
	data2dst(data, ppDst, pDstInfo)
	return int(r1)
}
func (v *ISVCDecoder) DecodeParser(pSrc []byte, iSrcLen int, pDstInfo *SParserBsInfo) int {
	var _pSrc unsafe.Pointer
	if pSrc == nil {
		_pSrc = nil
	} else {
		_pSrc = unsafe.Pointer(&pSrc[0])
	}
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.DecodeParser)),
		uintptr(unsafe.Pointer(v)),
		uintptr(_pSrc),
		uintptr(iSrcLen),
		uintptr(unsafe.Pointer(pDstInfo)),
	)
	return int(r1)
}

// func (v *ISVCDecoder) DecodeFrameEx(pSrc []byte, iSrcLen int, pDst []byte, iDstStride int, iDstLen, iWidth, iHeight, iColorFormat *int) int {
// 	r1, _, _ := purego.SyscallN(
// 		uintptr(unsafe.Pointer(v.vtbl.DecodeFrameEx)),
// 		uintptr(unsafe.Pointer(v)),
// 		uintptr(unsafe.Pointer(&pSrc[0])),
// 		uintptr(iSrcLen),
// 		uintptr(unsafe.Pointer(&pDst[0])),
// 		uintptr(iDstStride),
// 		uintptr(unsafe.Pointer(iDstLen)),
// 		uintptr(unsafe.Pointer(iWidth)),
// 		uintptr(unsafe.Pointer(iHeight)),
// 		uintptr(unsafe.Pointer(iColorFormat)),
// 	)
// 	return int(r1)
// }

func (v *ISVCEncoder) Initialize(p *SEncParamBase) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.Initialize)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(p)),
	)
	return int(r1)
}
func (v *ISVCEncoder) InitializeExt(p *SEncParamExt) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.InitializeExt)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(p)),
	)
	return int(r1)
}

func (v *ISVCEncoder) GetDefaultParams(p *SEncParamExt) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.GetDefaultParams)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(p)),
	)
	return int(r1)
}
func (v *ISVCEncoder) Uninitialize() int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.Uninitialize)),
		uintptr(unsafe.Pointer(v)),
	)
	return int(r1)
}

func (v *ISVCEncoder) EncodeFrame(kpSrcPic *SSourcePicture, pBsInfo *SFrameBSInfo) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.EncodeFrame)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(kpSrcPic)),
		uintptr(unsafe.Pointer(pBsInfo)),
	)
	return int(r1)
}
func (v *ISVCEncoder) EncodeParameterSets(pBsInfo *SFrameBSInfo) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.EncodeParameterSets)),
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pBsInfo)),
	)
	return int(r1)
}

func (v *ISVCEncoder) ForceIntraFrame(bIDR bool) int {
	var b uintptr = 0
	if bIDR {
		b = 1
	}
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.ForceIntraFrame)),
		uintptr(unsafe.Pointer(v)),
		b,
	)
	return int(r1)
}

func (v *ISVCEncoder) SetOption(eOptionId int, pOption *int) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.SetOption)),
		uintptr(unsafe.Pointer(v)),
		uintptr(eOptionId),
		uintptr(unsafe.Pointer(pOption)),
	)
	return int(r1)
}

func (v *ISVCEncoder) GetOption(eOptionId int, pOption *int) int {
	r1, _, _ := purego.SyscallN(
		uintptr(unsafe.Pointer(v.vtbl.GetOption)),
		uintptr(unsafe.Pointer(v)),
		uintptr(eOptionId),
		uintptr(unsafe.Pointer(pOption)),
	)
	return int(r1)
}
