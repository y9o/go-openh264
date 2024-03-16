//go:build cgoopenh264

// Calling the OpenH264 library from Go
package openh264

/*
#cgo CFLAGS: -Iinc_openh264
#cgo  windows LDFLAGS: openh264.dll
#cgo !windows LDFLAGS: -lopenh264
#include "codec_api.h"

//Encoder

long EncoderInitialize(ISVCEncoder*e, const SEncParamBase* pParam){
	return (*e)->Initialize(e,  pParam);
}
long EncoderInitializeExt(ISVCEncoder*e, const SEncParamExt* pParam){
	return (*e)->InitializeExt(e,  pParam);
}
long EncoderGetDefaultParams(ISVCEncoder*e, SEncParamExt* pParam){
	return (*e)->GetDefaultParams(e,  pParam);
}
long EncoderUninitialize(ISVCEncoder*e){
	return (*e)->Uninitialize(e);
}
long EncoderEncodeFrame(ISVCEncoder*e, const SSourcePicture* kpSrcPic, SFrameBSInfo* pBsInfo){
	return (*e)->EncodeFrame(e, kpSrcPic, pBsInfo);
}
long EncodeEncodeParameterSets(ISVCEncoder*e, SFrameBSInfo* pBsInfo){
	return (*e)->EncodeParameterSets(e, pBsInfo);
}
long EncodeForceIntraFrame(ISVCEncoder*e, bool bIDR){
	return (*e)->ForceIntraFrame(e, bIDR);
}
long EncodeSetOption(ISVCEncoder*e, int eOptionId, void* pOption){
	return (*e)->SetOption(e,eOptionId,pOption);
}
long EncodeGetOption(ISVCEncoder*e, int eOptionId, void* pOption){
	return (*e)->GetOption(e,eOptionId,pOption);
}

//Decoder

long DecoderInitialize(ISVCDecoder*d, const SDecodingParam* pParam){
	return (*d)->Initialize(d,  pParam);
}

long DecoderUninitialize(ISVCDecoder*d){
	return (*d)->Uninitialize(d);
}

long DecoderDecodeFrame (ISVCDecoder*d, const unsigned char* pSrc,const int iSrcLen,unsigned char** ppDst,int* pStride,int* iWidth,int* iHeight){
	return (*d)->DecodeFrame (d,  pSrc, iSrcLen, ppDst, pStride, iWidth, iHeight);
}

long DecoderDecodeFrameNoDelay (ISVCDecoder*d, const unsigned char* pSrc,const int iSrcLen,unsigned char** ppDst,SBufferInfo* pDstInfo){
	return (*d)->DecodeFrameNoDelay (d,  pSrc,iSrcLen, ppDst, pDstInfo);
}

long DecoderDecodeFrame2 (ISVCDecoder*d, const unsigned char* pSrc,const int iSrcLen,unsigned char** ppDst,SBufferInfo* pDstInfo){
	return (*d)->DecodeFrame2 (d,  pSrc, iSrcLen,  ppDst, pDstInfo);
}

long DecoderFlushFrame(ISVCDecoder*d, unsigned char** ppDst,SBufferInfo* pDstInfo){
	return (*d)->FlushFrame(d,  ppDst, pDstInfo);
}

long DecoderDecodeParser(ISVCDecoder*d, const unsigned char* pSrc,const int iSrcLen,SParserBsInfo* pDstInfo){
	return (*d)->DecodeParser(d,  pSrc, iSrcLen, pDstInfo);
}

long DecoderDecodeFrameEx(ISVCDecoder*d, const unsigned char* pSrc,const int iSrcLen,unsigned char* pDst,int iDstStride,int* iDstLen,int* iWidth,int* iHeight,int* iColorFormat){
	return (*d)->DecodeFrameEx(d, pSrc, iSrcLen, pDst, iDstStride, iDstLen, iWidth, iHeight,iColorFormat);
}

long DecoderGetOption(ISVCDecoder*d, int eOptionId, void* pOption){
	return (*d)->GetOption(d,eOptionId,pOption);
}
long DecoderSetOption(ISVCDecoder*d, int eOptionId, void* pOption){
	return (*d)->SetOption(d,eOptionId,pOption);
}
*/
import "C"
import (
	"unsafe"
)

func WelsCreateSVCEncoder(ppEncoder **ISVCEncoder) int {
	r0 := C.WelsCreateSVCEncoder((**C.ISVCEncoder)(unsafe.Pointer(ppEncoder)))
	return int(r0)
}

func WelsDestroySVCEncoder(pEncoder *ISVCEncoder) {
	C.WelsDestroySVCEncoder((*C.ISVCEncoder)(unsafe.Pointer(pEncoder)))
}

func (v *ISVCEncoder) Initialize(p *SEncParamBase) int {
	r0 := C.EncoderInitialize((*C.ISVCEncoder)(unsafe.Pointer(v)), (*C.SEncParamBase)(unsafe.Pointer(p)))
	return int(r0)
}

func (v *ISVCEncoder) InitializeExt(p *SEncParamExt) int {
	r0 := C.EncoderInitializeExt((*C.ISVCEncoder)(unsafe.Pointer(v)), (*C.SEncParamExt)(unsafe.Pointer(p)))
	return int(r0)
}

func (v *ISVCEncoder) GetDefaultParams(p *SEncParamExt) int {
	r0 := C.EncoderGetDefaultParams((*C.ISVCEncoder)(unsafe.Pointer(v)), (*C.SEncParamExt)(unsafe.Pointer(p)))
	return int(r0)
}

func (v *ISVCEncoder) Uninitialize() int {
	r0 := C.EncoderUninitialize((*C.ISVCEncoder)(unsafe.Pointer(v)))
	return int(r0)
}

func (v *ISVCEncoder) EncodeFrame(kpSrcPic *SSourcePicture, pBsInfo *SFrameBSInfo) int {
	r0 := C.EncoderEncodeFrame((*C.ISVCEncoder)(unsafe.Pointer(v)), (*C.SSourcePicture)(unsafe.Pointer(kpSrcPic)), (*C.SFrameBSInfo)(unsafe.Pointer(pBsInfo)))
	return int(r0)
}

func (v *ISVCEncoder) EncodeParameterSets(pBsInfo *SFrameBSInfo) int {
	r0 := C.EncodeEncodeParameterSets((*C.ISVCEncoder)(unsafe.Pointer(v)), (*C.SFrameBSInfo)(unsafe.Pointer(pBsInfo)))
	return int(r0)
}

func (v *ISVCEncoder) ForceIntraFrame(bIDR bool) int {
	r0 := C.EncodeForceIntraFrame((*C.ISVCEncoder)(unsafe.Pointer(v)), C.bool(bIDR))
	return int(r0)
}

func (v *ISVCEncoder) SetOption(eOptionId int, pOption *int) int {
	r0 := C.EncodeSetOption((*C.ISVCEncoder)(unsafe.Pointer(v)), C.int(eOptionId), unsafe.Pointer(pOption))
	return int(r0)
}
func (v *ISVCEncoder) GetOption(eOptionId int, pOption *int) int {
	r0 := C.EncodeGetOption((*C.ISVCEncoder)(unsafe.Pointer(v)), C.int(eOptionId), unsafe.Pointer(pOption))
	return int(r0)
}

func WelsCreateDecoder(ppDecoder **ISVCDecoder) int {
	r0 := C.WelsCreateDecoder((**C.ISVCDecoder)(unsafe.Pointer(ppDecoder)))
	return int(r0)
}

func WelsDestroyDecoder(pDecoder *ISVCDecoder) {
	C.WelsDestroyDecoder((*C.ISVCDecoder)(unsafe.Pointer(pDecoder)))
}

func WelsGetCodecVersionEx(ver *OpenH264Version) {
	C.WelsGetCodecVersionEx((*C.OpenH264Version)(unsafe.Pointer(ver)))
}

func WelsGetCodecVersion() (ver OpenH264Version) {
	WelsGetCodecVersionEx(&ver)
	return
}

// If CGO_ENABLED=1, specify the path to the library.
func Open(library string) error {
	return nil
}
func Close() error {
	return nil
}
func (v *ISVCDecoder) Initialize(p *SDecodingParam) int {
	r := C.DecoderInitialize((*C.ISVCDecoder)(unsafe.Pointer(v)), (*C.SDecodingParam)(unsafe.Pointer(p)))
	return int(r)
}
func (v *ISVCDecoder) Uninitialize() int {
	r := C.DecoderUninitialize((*C.ISVCDecoder)(unsafe.Pointer(v)))
	return int(r)
}

func (v *ISVCDecoder) DecodeFrameNoDelay(pSrc []byte, iSrcLen int, ppDst *[3][]byte, pDstInfo *SBufferInfo) int {
	var data [3]*C.uchar
	r := C.DecoderDecodeFrameNoDelay(
		(*C.ISVCDecoder)(unsafe.Pointer(v)),
		(*C.uchar)(unsafe.Pointer(&pSrc[0])),
		C.int(iSrcLen),
		(**C.uchar)(unsafe.Pointer(&data)),
		(*C.SBufferInfo)(unsafe.Pointer(pDstInfo)),
	)
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
	return int(r)
}
func (v *ISVCDecoder) DecodeFrame2(pSrc []byte, iSrcLen int, ppDst *[3][]byte, pDstInfo *SBufferInfo) int {
	var data [3]*C.uchar
	var psrc *C.uchar
	if pSrc != nil {
		psrc = (*C.uchar)(unsafe.Pointer(&pSrc[0]))
	}
	r := C.DecoderDecodeFrame2(
		(*C.ISVCDecoder)(unsafe.Pointer(v)),
		psrc,
		C.int(iSrcLen),
		(**C.uchar)(unsafe.Pointer(&data)),
		(*C.SBufferInfo)(unsafe.Pointer(pDstInfo)),
	)
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
	return int(r)
}

func (v *ISVCDecoder) FlushFrame(ppDst *[3][]byte, pDstInfo *SBufferInfo) int {
	var data [3]*C.uchar
	r := C.DecoderFlushFrame(
		(*C.ISVCDecoder)(unsafe.Pointer(v)),
		(**C.uchar)(unsafe.Pointer(&data)),
		(*C.SBufferInfo)(unsafe.Pointer(pDstInfo)),
	)
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
	return int(r)
}

func (v *ISVCDecoder) DecodeParser(pSrc []byte, iSrcLen int, pDstInfo *SParserBsInfo) int {
	var psrc *C.uchar
	if pSrc != nil {
		psrc = (*C.uchar)(unsafe.Pointer(&pSrc[0]))
	}
	r := C.DecoderDecodeParser(
		(*C.ISVCDecoder)(unsafe.Pointer(v)),
		psrc,
		C.int(iSrcLen),
		(*C.SParserBsInfo)(unsafe.Pointer(pDstInfo)),
	)
	return int(r)
}

func (v *ISVCDecoder) SetOption(eOptionId int, pOption *int) int {
	r := C.DecoderSetOption((*C.ISVCDecoder)(unsafe.Pointer(v)), C.int(eOptionId), unsafe.Pointer(pOption))
	return int(r)
}
func (v *ISVCDecoder) GetOption(eOptionId int, pOption *int) int {
	return int(C.DecoderGetOption((*C.ISVCDecoder)(unsafe.Pointer(v)), C.int(eOptionId), unsafe.Pointer(pOption)))
}

func (v *SBufferInfo) UsrData_sSystemBuffer() *SSysMEMBuffer {
	tmp := (*C.SBufferInfo)(unsafe.Pointer(v))
	return (*SSysMEMBuffer)(unsafe.Pointer(&tmp.UsrData))
}
