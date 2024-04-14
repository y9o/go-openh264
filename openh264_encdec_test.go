// This example demonstrates encoding a YCbCr image.
package openh264_test

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"image"
	"io"
	"runtime"
	"strings"
	"unsafe"

	"github.com/y9o/go-openh264"
)

func Example_encode() {
	err := openh264.Open("./openh264-2.4.1-win64.dll")
	if err != nil {
		return
	}
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()

	bufbyte := make([]byte, 0, 10_000)
	out_h264 := bytes.NewBuffer(bufbyte)
	input_frame := image.NewYCbCr(image.Rect(0, 0, 1920, 1080), image.YCbCrSubsampleRatio420)

	var ppEnc *openh264.ISVCEncoder
	if ret := openh264.WelsCreateSVCEncoder(&ppEnc); ret != 0 || ppEnc == nil {
		return
	}
	defer openh264.WelsDestroySVCEncoder(ppEnc)

	encParam := openh264.SEncParamBase{
		IUsageType:     openh264.CAMERA_VIDEO_REAL_TIME,
		IPicWidth:      1920,
		IPicHeight:     1080,
		ITargetBitrate: 1_000_000,
		FMaxFrameRate:  20,
	}
	if r := ppEnc.Initialize(&encParam); r != 0 {
		return
	}
	defer ppEnc.Uninitialize()

	encSrcPic := openh264.SSourcePicture{
		IColorFormat: openh264.VideoFormatI420,
		IStride:      [4]int32{},
		PData:        [4]*uint8{},
		IPicWidth:    1920,
		IPicHeight:   1080,
		UiTimeStamp:  0,
	}
	encSrcPic.IStride[0] = 1920
	encSrcPic.IStride[1] = 960
	encSrcPic.IStride[2] = 960

	pinner.Pin(&input_frame.Y[0])
	pinner.Pin(&input_frame.Cb[0])
	pinner.Pin(&input_frame.Cr[0])
	encSrcPic.PData[0] = (*uint8)(unsafe.Pointer(&input_frame.Y[0]))
	encSrcPic.PData[1] = (*uint8)(unsafe.Pointer(&input_frame.Cb[0]))
	encSrcPic.PData[2] = (*uint8)(unsafe.Pointer(&input_frame.Cr[0]))

	encInfo := openh264.SFrameBSInfo{}
	if ret := ppEnc.EncodeFrame(&encSrcPic, &encInfo); ret != openh264.CmResultSuccess {
		return
	}
	if encInfo.EFrameType != openh264.VideoFrameTypeSkip {
		for iLayer := 0; iLayer < int(encInfo.ILayerNum); iLayer++ {
			pLayerBsInfo := &encInfo.SLayerInfo[iLayer]
			var iLayerSize int32
			nallens := unsafe.Slice(pLayerBsInfo.PNalLengthInByte, pLayerBsInfo.INalCount)
			for _, l := range nallens {
				iLayerSize += l
			}
			nals := unsafe.Slice(pLayerBsInfo.PBsBuf, iLayerSize)
			out_h264.Write(nals)
		}
	}
}

func Example_decode() {
	err := openh264.Open("./openh264-2.4.1-win64.dll")
	if err != nil {
		return
	}
	readerBase64 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data_zlib_h264))
	readerZip, _ := zlib.NewReader(readerBase64)
	in_h264, _ := io.ReadAll(readerZip)

	var ppdec *openh264.ISVCDecoder
	if ret := openh264.WelsCreateDecoder(&ppdec); ret != 0 || ppdec == nil {
		return
	}
	defer openh264.WelsDestroyDecoder(ppdec)

	sDecParam := openh264.SDecodingParam{}
	sDecParam.EEcActiveIdc = openh264.ERROR_CON_SLICE_MV_COPY_CROSS_IDR_FREEZE_RES_CHANGE
	if r := ppdec.Initialize(&sDecParam); r != 0 {
		return
	}
	defer ppdec.Uninitialize()

	for len(in_h264) > 4 {
		var sDstBufInfo openh264.SBufferInfo
		pos := bytes.Index(in_h264[4:], []byte{0, 0, 0, 1})
		length := len(in_h264)
		if pos != -1 {
			length = pos + 4
		}
		var pDst [3][]byte
		if r := ppdec.DecodeFrameNoDelay(in_h264[:length], length, &pDst, &sDstBufInfo); r != 0 {
			return
		}
		if pDst[0] != nil {
			_ = &image.YCbCr{
				Y:       pDst[0],
				Cb:      pDst[1],
				Cr:      pDst[2],
				YStride: int(sDstBufInfo.UsrData_sSystemBuffer().IStride[0]),
				CStride: int(sDstBufInfo.UsrData_sSystemBuffer().IStride[1]),
				Rect: image.Rect(
					0, 0,
					int(sDstBufInfo.UsrData_sSystemBuffer().IWidth),
					int(sDstBufInfo.UsrData_sSystemBuffer().IHeight)),
				SubsampleRatio: image.YCbCrSubsampleRatio420,
			}
		}

		if pos == -1 {
			break
		}
		in_h264 = in_h264[pos+4:]
	}

	var num_of_frames_in_buffer int
	ppdec.GetOption(openh264.DECODER_OPTION_NUM_OF_FRAMES_REMAINING_IN_BUFFER, &num_of_frames_in_buffer)
	for i := 0; i < num_of_frames_in_buffer; i++ {
		var sDstBufInfo openh264.SBufferInfo
		var pDst [3][]byte
		if r := ppdec.FlushFrame(&pDst, &sDstBufInfo); r != 0 {
			return
		}
		if pDst[0] != nil {
			_ = &image.YCbCr{
				Y:       pDst[0],
				Cb:      pDst[1],
				Cr:      pDst[2],
				YStride: int(sDstBufInfo.UsrData_sSystemBuffer().IStride[0]),
				CStride: int(sDstBufInfo.UsrData_sSystemBuffer().IStride[1]),
				Rect: image.Rect(
					0, 0,
					int(sDstBufInfo.UsrData_sSystemBuffer().IWidth),
					int(sDstBufInfo.UsrData_sSystemBuffer().IHeight)),
				SubsampleRatio: image.YCbCrSubsampleRatio420,
			}
		}
	}

}

const data_zlib_h264 = `eJxiYGBgTHc6oNHT62DDKPzpgI2S4AqQWMY5mwYQnbqDgYWBgcWiJ4BBIL6PgYl14wEGh21+AwN2vV43ikbRKBpFo2gUjaJRNIpG0SgaRSMEDbwLSEerB9wFpKNXA+6CUTSKRtEoGkWjaBQNezSJZmbP7mBgYnjA0cLAyFDA4pk5IKAAEAAA//8onmGU`
