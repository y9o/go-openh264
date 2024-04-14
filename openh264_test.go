package openh264

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image"
	"log"
	"os"
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

func makecolorbars() *image.YCbCr {

	colors := [][]byte{
		{180, 128, 128},
		{168, 44, 136},
		{145, 147, 44},
		{133, 63, 52},
		{63, 193, 204},
		{51, 109, 212},
		{28, 212, 120},
		{16, 128, 128},
	}
	y := make([]byte, 1920*1080*2)
	u := make([]byte, 960*540*2)
	v := make([]byte, 960*540*2)
	yy := y[1920*1080:]
	uu := u[960*540:]
	vv := v[960*540:]
	for i := 0; i < 1920; i++ {
		v := colors[i/240][0]
		for j := 0; j < 1080; j++ {
			y[1920*j+i] = v
			yy[1920*j+i] = v
		}
	}
	for i := 0; i < 960; i++ {
		v1 := colors[i/120][1]
		v2 := colors[i/120][2]
		for j := 0; j < 540; j++ {
			u[960*j+i] = v1
			v[960*j+i] = v2
			uu[960*j+i] = v1
			vv[960*j+i] = v2
		}
	}

	img := &image.YCbCr{
		Y:       y,
		Cb:      u,
		Cr:      v,
		YStride: 1920,
		CStride: 960,
		Rect: image.Rect(
			0, 0,
			1920,
			1080),
		SubsampleRatio: image.YCbCrSubsampleRatio420,
	}
	return img
}
func getSystemLibrary() string {
	switch runtime.GOOS {
	case "windows":
		return "openh264-2.4.1-win64.dll"
	default:
		return "./libopenh264-2.4.1-linux64.7.so"
	}
}
func TestMain(m *testing.M) {
	err := Open(getSystemLibrary())
	if err != nil {
		log.Fatalln(err)
	}
	Close()
	Close()
	err = Open(getSystemLibrary())
	if err != nil {
		log.Fatalln(err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestEncodeDecode(t *testing.T) {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()
	img := makecolorbars()
	//Encode
	var ppEnc *ISVCEncoder
	if ret := WelsCreateSVCEncoder(&ppEnc); ret != 0 || ppEnc == nil {
		t.Fatal("failed WelsCreateEncoder:", ret, ppEnc)
	}
	defer WelsDestroySVCEncoder(ppEnc)

	encParam := SEncParamBase{
		IUsageType:     CAMERA_VIDEO_REAL_TIME,
		IPicWidth:      1920,
		IPicHeight:     1080,
		ITargetBitrate: 1_000_000,
		FMaxFrameRate:  20,
	}
	if r := ppEnc.Initialize(&encParam); r != 0 {
		t.Fatal("Initialize", r)
	}
	defer ppEnc.Uninitialize()
	encSrcPic := SSourcePicture{
		IColorFormat: VideoFormatI420,
		IStride:      [4]int32{},
		PData:        [4]*uint8{},
		IPicWidth:    1920,
		IPicHeight:   1080,
		UiTimeStamp:  0,
	}
	encInfo := SFrameBSInfo{}

	bufbyte := make([]byte, 0, 2_000_000)
	buf := bytes.NewBuffer(bufbyte)

	encSrcPic.IStride[0] = 1920
	encSrcPic.IStride[1] = 960
	encSrcPic.IStride[2] = 960
	chunk := []int{}
	loop := 960
	loop = 40
	for i := 0; i < loop; i++ {
		pinner.Pin(&img.Y[i*2])
		pinner.Pin(&img.Cb[i])
		pinner.Pin(&img.Cr[i])
		encSrcPic.PData[0] = (*uint8)(unsafe.Pointer(&img.Y[i*2]))
		encSrcPic.PData[1] = (*uint8)(unsafe.Pointer(&img.Cb[i]))
		encSrcPic.PData[2] = (*uint8)(unsafe.Pointer(&img.Cr[i]))
		if ret := ppEnc.EncodeFrame(&encSrcPic, &encInfo); ret != CmResultSuccess {
			t.Fatalf("ppEnc.EncodeFrame(%d) != CmResultSuccess(%d)", ret, CmResultSuccess)
		}
		if encInfo.EFrameType != VideoFrameTypeSkip {
			c := 0
			for iLayer := 0; iLayer < int(encInfo.ILayerNum); iLayer++ {
				pLayerBsInfo := &encInfo.SLayerInfo[iLayer]
				var iLayerSize int32
				nallens := unsafe.Slice(pLayerBsInfo.PNalLengthInByte, pLayerBsInfo.INalCount)
				for _, l := range nallens {
					iLayerSize += l
				}
				nals := unsafe.Slice(pLayerBsInfo.PBsBuf, iLayerSize)
				c += int(iLayerSize)
				buf.Write(nals)
			}
			chunk = append(chunk, c)
		}
	}
	fh := md5.New()
	fh.Write(buf.Bytes())
	h := fmt.Sprintf("%X", fh.Sum(nil))
	switch h {
	case "B1A7ABC4E84964D23B63E0D9D94C3FC8": //40
	case "D3D4F517C6C90DA4CAE66AD9EB52A259": //960
		break
	default:
		t.Fatal(h)
	}

	var ppdec *ISVCDecoder
	if ret := WelsCreateDecoder(&ppdec); ret != 0 || ppdec == nil {
		log.Fatalln("failed WelsCreateDecoder:", ret, ppdec)
	}
	defer WelsDestroyDecoder(ppdec)

	var op int = 0
	ppdec.SetOption(DECODER_OPTION_TRACE_LEVEL, &op)

	sDecParam := SDecodingParam{}
	sDecParam.SVideoProperty.Size = uint32(unsafe.Sizeof(sDecParam.SVideoProperty))
	sDecParam.EEcActiveIdc = ERROR_CON_SLICE_MV_COPY_CROSS_IDR_FREEZE_RES_CHANGE
	sDecParam.SVideoProperty.EVideoBsType = VIDEO_BITSTREAM_DEFAULT
	if r := ppdec.Initialize(&sDecParam); r != 0 {
		log.Fatalln("failed Initialize.", r)
	}
	defer ppdec.Uninitialize()
	dataoffset := 0
	src := buf.Bytes()
	fh2 := md5.New()
	for i, l := range chunk {
		data := src[dataoffset : dataoffset+l]
		dataoffset += l
		if len(data) > 0 {
			var sDstBufInfo SBufferInfo
			sDstBufInfo.UiInBsTimeStamp = uint64(i)
			var pDst [3][]byte
			if r := ppdec.DecodeFrameNoDelay(data, len(data), &pDst, &sDstBufInfo); r != 0 {
				t.Fatal("decode", r)
			}
			if pDst[0] != nil {
				fh2.Write(pDst[0])
				fh2.Write(pDst[1])
				fh2.Write(pDst[2])
			}
		}

	}

	var num_of_frames_in_buffer int
	ppdec.GetOption(DECODER_OPTION_NUM_OF_FRAMES_REMAINING_IN_BUFFER, &num_of_frames_in_buffer)
	for i := 0; i < num_of_frames_in_buffer; i++ {
		var pDst [3][]byte
		var sDstBufInfo SBufferInfo
		if r := ppdec.FlushFrame(&pDst, &sDstBufInfo); r != 0 {
			t.Fatal("err", r)
		}
		if sDstBufInfo.IBufferStatus == 1 {
			fh2.Write(pDst[0])
			fh2.Write(pDst[1])
			fh2.Write(pDst[2])
		}
	}
	h2 := fmt.Sprintf("%X", fh2.Sum(nil))
	switch h2 {
	case "2633E81CCB23D87F4588AB6DFE02E93B":
	case "E44201F28083470357F333877278AF4E": //40
		break
	default:
		t.Fatal(h2)
	}
}

func TestEncodeExtDecode(t *testing.T) {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()
	img := makecolorbars()

	var ppEnc *ISVCEncoder
	if ret := WelsCreateSVCEncoder(&ppEnc); ret != 0 || ppEnc == nil {
		t.Fatal("failed WelsCreateEncoder:", ret, ppEnc)
	}
	defer WelsDestroySVCEncoder(ppEnc)

	var encParam SEncParamExt

	ppEnc.GetDefaultParams(&encParam)
	encParam.IUsageType = CAMERA_VIDEO_REAL_TIME
	encParam.FMaxFrameRate = 20
	encParam.IPicWidth = 1920
	encParam.IPicHeight = 1080
	encParam.ITargetBitrate = 1_000_000
	encParam.IMaxBitrate = 2_000_000
	encParam.BEnableDenoise = false
	encParam.ISpatialLayerNum = 1
	encParam.IMultipleThreadIdc = 5

	encParam.SSpatialLayers[0].IVideoWidth = 1920
	encParam.SSpatialLayers[0].IVideoHeight = 1080
	encParam.SSpatialLayers[0].FFrameRate = 20
	encParam.SSpatialLayers[0].ISpatialBitrate = 1_000_000
	encParam.SSpatialLayers[0].SSliceArgument.UiSliceMode = SM_FIXEDSLCNUM_SLICE
	encParam.SSpatialLayers[0].SSliceArgument.UiSliceNum = 8

	if ret := ppEnc.InitializeExt(&encParam); ret != 0 {
		t.Fatal("InitializeExt 0 !=", ret)
	}
	defer ppEnc.Uninitialize()
	var videoFormat int = VideoFormatI420
	ppEnc.SetOption(ENCODER_OPTION_DATAFORMAT, &videoFormat)

	encSrcPic := SSourcePicture{
		IColorFormat: VideoFormatI420,
		IStride:      [4]int32{},
		PData:        [4]*uint8{},
		IPicWidth:    1920,
		IPicHeight:   1080,
		UiTimeStamp:  0,
	}
	encInfo := SFrameBSInfo{}

	bufbyte := make([]byte, 0, 1000000)
	buf := bytes.NewBuffer(bufbyte)

	encSrcPic.IStride[0] = 1920
	encSrcPic.IStride[1] = 960
	encSrcPic.IStride[2] = 960
	chunk := []int{}
	loop := 960
	loop = 40
	for i := 0; i < loop; i++ {
		pinner.Pin(&img.Y[i*2])
		pinner.Pin(&img.Cb[i])
		pinner.Pin(&img.Cr[i])
		encSrcPic.PData[0] = (*uint8)(unsafe.Pointer(&img.Y[i*2]))
		encSrcPic.PData[1] = (*uint8)(unsafe.Pointer(&img.Cb[i]))
		encSrcPic.PData[2] = (*uint8)(unsafe.Pointer(&img.Cr[i]))
		if i == 5 {
			if r := ppEnc.ForceIntraFrame(true); r != 0 {
				t.Fatal("ForceIntraFrame", r)
			}
		}
		if ret := ppEnc.EncodeFrame(&encSrcPic, &encInfo); ret != CmResultSuccess {
			t.Fatalf("ppEnc.EncodeFrame(%d) != CmResultSuccess(%d)", ret, CmResultSuccess)
		}
		encInfo.UiTimeStamp += 1000
		if encInfo.EFrameType != VideoFrameTypeSkip {
			c := 0
			for iLayer := 0; iLayer < int(encInfo.ILayerNum); iLayer++ {
				pLayerBsInfo := &encInfo.SLayerInfo[iLayer]
				var iLayerSize int32
				nallens := unsafe.Slice(pLayerBsInfo.PNalLengthInByte, pLayerBsInfo.INalCount)
				for _, l := range nallens {
					iLayerSize += l
				}
				nals := unsafe.Slice(pLayerBsInfo.PBsBuf, iLayerSize)
				c += int(iLayerSize)
				buf.Write(nals)
			}
			chunk = append(chunk, c)
		}
		if i == 10 {
			brInfo := SBitrateInfo{}
			if r := ppEnc.GetOption(ENCODER_OPTION_BITRATE, (*int)(unsafe.Pointer(&brInfo))); r != 0 {
				t.Fatal("GetOption ENCODER_OPTION_BITRATE", r)
			}
			if brInfo.IBitrate != 1_000_000 {
				t.Fatal("ENCODER_OPTION_BITRATE 1_000_000 != ", brInfo)
			}
		}
	}
	fh := md5.New()
	fh.Write(buf.Bytes())
	h := fmt.Sprintf("%X", fh.Sum(nil))

	// fh264, err := os.OpenFile("testext.264", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fh264.Write(buf.Bytes())
	// fh264.Close()
	switch h {
	case "1609B3CA097D994915CCF07CDD165330": //40
	case "A71E9C88C8541A7E2A65723B41BD8276": //960
		break
	default:
		t.Fatal(h)
	}

	//decoder

	var ppdec *ISVCDecoder
	if ret := WelsCreateDecoder(&ppdec); ret != 0 || ppdec == nil {
		log.Fatalln("failed WelsCreateDecoder:", ret, ppdec)
	}
	defer WelsDestroyDecoder(ppdec)

	var op int = 2
	ppdec.SetOption(DECODER_OPTION_TRACE_LEVEL, &op)

	sDecParam := SDecodingParam{}
	sDecParam.EEcActiveIdc = ERROR_CON_SLICE_MV_COPY_CROSS_IDR_FREEZE_RES_CHANGE
	var sDstBufInfo SBufferInfo

	if r := ppdec.Initialize(&sDecParam); r != 0 {
		log.Fatalln("failed Initialize.", r)
	}
	defer ppdec.Uninitialize()
	dataoffset := 0
	src := buf.Bytes()
	fh2 := md5.New()
	c := 0
	for _, l := range chunk {
		data := src[dataoffset : dataoffset+l]
		dataoffset += l
		if len(data) > 0 {
			var pDst [3][]byte
			if r := ppdec.DecodeFrame2(data, len(data), &pDst, &sDstBufInfo); r != 0 {
				t.Fatal("decode", r)
			}
			if pDst[0] != nil {
				c++
				fh2.Write(pDst[0])
				fh2.Write(pDst[1])
				fh2.Write(pDst[2])
			}
			if r := ppdec.DecodeFrame2(nil, len(data), &pDst, &sDstBufInfo); r != 0 {
				t.Fatal("decode", r)
			}
			if pDst[0] != nil {
				c++
				fh2.Write(pDst[0])
				fh2.Write(pDst[1])
				fh2.Write(pDst[2])
			}
		}

	}

	var num_of_frames_in_buffer int
	ppdec.GetOption(DECODER_OPTION_NUM_OF_FRAMES_REMAINING_IN_BUFFER, &num_of_frames_in_buffer)
	for i := 0; i < num_of_frames_in_buffer; i++ {
		var pDst [3][]byte
		if r := ppdec.FlushFrame(&pDst, &sDstBufInfo); r != 0 {
			t.Fatal("err", r)
		}
		if pDst[0] != nil {
			c++
			fh2.Write(pDst[0])
			fh2.Write(pDst[1])
			fh2.Write(pDst[2])
		}
	}
	h2 := fmt.Sprintf("%X", fh2.Sum(nil))
	switch h2 {
	case "D86E18760CD1829B11207EB19739AE32": //40
	case "49272CBC00A8A62EEC1FF66D093CBF3B": //960
		break
	default:
		t.Fatal(h2)
	}
}

func TestEncodeExtDecodeNoDelay(t *testing.T) {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()
	img := makecolorbars()

	var ppEnc *ISVCEncoder
	if ret := WelsCreateSVCEncoder(&ppEnc); ret != 0 || ppEnc == nil {
		t.Fatal("failed WelsCreateEncoder:", ret, ppEnc)
	}
	defer WelsDestroySVCEncoder(ppEnc)

	var encParam SEncParamExt

	ppEnc.GetDefaultParams(&encParam)
	encParam.IUsageType = CAMERA_VIDEO_REAL_TIME
	encParam.FMaxFrameRate = 20
	encParam.IPicWidth = 1920
	encParam.IPicHeight = 1080
	encParam.ITargetBitrate = 1_000_000
	encParam.IMaxBitrate = 2_000_000
	encParam.BEnableDenoise = false
	encParam.ISpatialLayerNum = 1
	encParam.IMultipleThreadIdc = 5

	encParam.SSpatialLayers[0].IVideoWidth = 1920
	encParam.SSpatialLayers[0].IVideoHeight = 1080
	encParam.SSpatialLayers[0].FFrameRate = 20
	encParam.SSpatialLayers[0].ISpatialBitrate = 1_000_000
	encParam.SSpatialLayers[0].SSliceArgument.UiSliceMode = SM_FIXEDSLCNUM_SLICE
	encParam.SSpatialLayers[0].SSliceArgument.UiSliceNum = 8

	if ret := ppEnc.InitializeExt(&encParam); ret != 0 {
		t.Fatal("InitializeExt 0 !=", ret)
	}
	defer ppEnc.Uninitialize()
	var videoFormat int = VideoFormatI420
	ppEnc.SetOption(ENCODER_OPTION_DATAFORMAT, &videoFormat)

	encSrcPic := SSourcePicture{
		IColorFormat: VideoFormatI420,
		IStride:      [4]int32{},
		PData:        [4]*uint8{},
		IPicWidth:    1920,
		IPicHeight:   1080,
		UiTimeStamp:  0,
	}
	encInfo := SFrameBSInfo{}

	bufbyte := make([]byte, 0, 1000000)
	buf := bytes.NewBuffer(bufbyte)

	encSrcPic.IStride[0] = 1920
	encSrcPic.IStride[1] = 960
	encSrcPic.IStride[2] = 960
	chunk := []int{}
	loop := 960
	loop = 40
	for i := 0; i < loop; i++ {
		pinner.Pin(&img.Y[i*2])
		pinner.Pin(&img.Cb[i])
		pinner.Pin(&img.Cr[i])
		encSrcPic.PData[0] = (*uint8)(unsafe.Pointer(&img.Y[i*2]))
		encSrcPic.PData[1] = (*uint8)(unsafe.Pointer(&img.Cb[i]))
		encSrcPic.PData[2] = (*uint8)(unsafe.Pointer(&img.Cr[i]))
		if i == 5 {
			if r := ppEnc.ForceIntraFrame(true); r != 0 {
				t.Fatal("ForceIntraFrame", r)
			}
		}
		if ret := ppEnc.EncodeFrame(&encSrcPic, &encInfo); ret != CmResultSuccess {
			t.Fatalf("ppEnc.EncodeFrame(%d) != CmResultSuccess(%d)", ret, CmResultSuccess)
		}
		encInfo.UiTimeStamp += 1000
		if encInfo.EFrameType != VideoFrameTypeSkip {
			c := 0
			for iLayer := 0; iLayer < int(encInfo.ILayerNum); iLayer++ {
				pLayerBsInfo := &encInfo.SLayerInfo[iLayer]
				var iLayerSize int32
				nallens := unsafe.Slice(pLayerBsInfo.PNalLengthInByte, pLayerBsInfo.INalCount)
				for _, l := range nallens {
					iLayerSize += l
				}
				nals := unsafe.Slice(pLayerBsInfo.PBsBuf, iLayerSize)
				c += int(iLayerSize)
				buf.Write(nals)
			}
			chunk = append(chunk, c)
		}
		if i == 10 {
			brInfo := SBitrateInfo{}
			if r := ppEnc.GetOption(ENCODER_OPTION_BITRATE, (*int)(unsafe.Pointer(&brInfo))); r != 0 {
				t.Fatal("GetOption ENCODER_OPTION_BITRATE", r)
			}
			if brInfo.IBitrate != 1_000_000 {
				t.Fatal("ENCODER_OPTION_BITRATE 1_000_000 != ", brInfo)
			}
		}
	}
	fh := md5.New()
	fh.Write(buf.Bytes())
	h := fmt.Sprintf("%X", fh.Sum(nil))

	switch h {
	case "1609B3CA097D994915CCF07CDD165330": //40
	case "A71E9C88C8541A7E2A65723B41BD8276": //960
		break
	default:
		t.Fatal(h)
	}

	//decoder

	var ppdec *ISVCDecoder
	if ret := WelsCreateDecoder(&ppdec); ret != 0 || ppdec == nil {
		log.Fatalln("failed WelsCreateDecoder:", ret, ppdec)
	}
	defer WelsDestroyDecoder(ppdec)

	var op int = 0
	ppdec.SetOption(DECODER_OPTION_TRACE_LEVEL, &op)
	//multi thread
	op = 4
	ppdec.SetOption(DECODER_OPTION_NUM_OF_THREADS, &op)

	sDecParam := SDecodingParam{}
	sDecParam.EEcActiveIdc = ERROR_CON_SLICE_MV_COPY_CROSS_IDR_FREEZE_RES_CHANGE
	sDecParam.SVideoProperty.Size = uint32(unsafe.Sizeof(sDecParam.SVideoProperty))
	sDecParam.SVideoProperty.EVideoBsType = VIDEO_BITSTREAM_DEFAULT

	if r := ppdec.Initialize(&sDecParam); r != 0 {
		log.Fatalln("failed Initialize.", r)
	}
	defer ppdec.Uninitialize()
	dataoffset := 0
	src := buf.Bytes()
	fh2 := md5.New()
	c := 0
	for i, l := range chunk {
		data := src[dataoffset : dataoffset+l]
		dataoffset += l
		if len(data) > 0 {
			var sDstBufInfo SBufferInfo
			sDstBufInfo.UiInBsTimeStamp = uint64(i)
			var pDst [3][]byte
			if r := ppdec.DecodeFrameNoDelay(data, len(data), &pDst, &sDstBufInfo); r != 0 {
				t.Fatal("decode", r)
			}
			if pDst[0] != nil {
				c++
				fh2.Write(pDst[0])
				fh2.Write(pDst[1])
				fh2.Write(pDst[2])

				// img := &image.YCbCr{
				// 	Y:       pDst[0],
				// 	Cb:      pDst[1],
				// 	Cr:      pDst[2],
				// 	YStride: int(sDstBufInfo.UsrData_sSystemBuffer().IStride[0]),
				// 	CStride: int(sDstBufInfo.UsrData_sSystemBuffer().IStride[1]),
				// 	Rect: image.Rect(
				// 		0, 0,
				// 		int(sDstBufInfo.UsrData_sSystemBuffer().IWidth),
				// 		int(sDstBufInfo.UsrData_sSystemBuffer().IHeight)),
				// 	SubsampleRatio: image.YCbCrSubsampleRatio420,
				// }
				// outname := fmt.Sprintf("img%04d.jpg", sDstBufInfo.UiOutYuvTimeStamp)
				// fh, err := os.OpenFile(outname, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
				// if err != nil {
				// 	log.Fatalln(err)
				// }
				// err = jpeg.Encode(fh, img, &jpeg.Options{Quality: 80})
				// if err != nil {
				// 	log.Fatalln(err)
				// }
				// fh.Close()
			}
		}

	}

	var num_of_frames_in_buffer int
	ppdec.GetOption(DECODER_OPTION_NUM_OF_FRAMES_REMAINING_IN_BUFFER, &num_of_frames_in_buffer)
	for i := 0; i < num_of_frames_in_buffer; i++ {
		var sDstBufInfo SBufferInfo
		var pDst [3][]byte
		if r := ppdec.FlushFrame(&pDst, &sDstBufInfo); r != 0 {
			t.Fatal("err", r)
		}
		if pDst[0] != nil {
			c++
			fh2.Write(pDst[0])
			fh2.Write(pDst[1])
			fh2.Write(pDst[2])
			// img := &image.YCbCr{
			// 	Y:       pDst[0],
			// 	Cb:      pDst[1],
			// 	Cr:      pDst[2],
			// 	YStride: int(sDstBufInfo.UsrData_sSystemBuffer().IStride[0]),
			// 	CStride: int(sDstBufInfo.UsrData_sSystemBuffer().IStride[1]),
			// 	Rect: image.Rect(
			// 		0, 0,
			// 		int(sDstBufInfo.UsrData_sSystemBuffer().IWidth),
			// 		int(sDstBufInfo.UsrData_sSystemBuffer().IHeight)),
			// 	SubsampleRatio: image.YCbCrSubsampleRatio420,
			// }
			// outname := fmt.Sprintf("img%04d.jpg", sDstBufInfo.UiOutYuvTimeStamp)
			// fh, err := os.OpenFile(outname, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
			// if err != nil {
			// 	log.Fatalln(err)
			// }
			// err = jpeg.Encode(fh, img, &jpeg.Options{Quality: 80})
			// if err != nil {
			// 	log.Fatalln(err)
			// }
			// fh.Close()
		}
	}
	h2 := fmt.Sprintf("%X", fh2.Sum(nil))
	switch h2 {
	case "D28C5C953FB92DA432EAD8731AA6DE14": //40 multi thread
	case "D86E18760CD1829B11207EB19739AE32": //40
	case "49272CBC00A8A62EEC1FF66D093CBF3B": //960
		break
	default:
		t.Fatal(h2)
	}
}

func TestEncodeDecodeParser(t *testing.T) {
	pinner := &runtime.Pinner{}
	defer pinner.Unpin()
	img := makecolorbars()
	//Encode
	var ppEnc *ISVCEncoder
	if ret := WelsCreateSVCEncoder(&ppEnc); ret != 0 || ppEnc == nil {
		t.Fatal("failed WelsCreateEncoder:", ret, ppEnc)
	}
	defer WelsDestroySVCEncoder(ppEnc)

	encParam := SEncParamBase{
		IUsageType:     CAMERA_VIDEO_REAL_TIME,
		IPicWidth:      1920,
		IPicHeight:     1080,
		ITargetBitrate: 1_000_000,
		FMaxFrameRate:  20,
	}
	if r := ppEnc.Initialize(&encParam); r != 0 {
		t.Fatal("Initialize", r)
	}
	defer ppEnc.Uninitialize()
	encSrcPic := SSourcePicture{
		IColorFormat: VideoFormatI420,
		IStride:      [4]int32{},
		PData:        [4]*uint8{},
		IPicWidth:    1920,
		IPicHeight:   1080,
		UiTimeStamp:  0,
	}
	encInfo := SFrameBSInfo{}

	bufbyte := make([]byte, 0, 2_000_000)
	buf := bytes.NewBuffer(bufbyte)

	encSrcPic.IStride[0] = 1920
	encSrcPic.IStride[1] = 960
	encSrcPic.IStride[2] = 960
	chunk := []int{}
	for i := 0; i < 4; i++ {
		pinner.Pin(&img.Y[i*2])
		pinner.Pin(&img.Cb[i])
		pinner.Pin(&img.Cr[i])
		encSrcPic.PData[0] = (*uint8)(unsafe.Pointer(&img.Y[i*2]))
		encSrcPic.PData[1] = (*uint8)(unsafe.Pointer(&img.Cb[i]))
		encSrcPic.PData[2] = (*uint8)(unsafe.Pointer(&img.Cr[i]))
		if ret := ppEnc.EncodeFrame(&encSrcPic, &encInfo); ret != CmResultSuccess {
			t.Fatalf("ppEnc.EncodeFrame(%d) != CmResultSuccess(%d)", ret, CmResultSuccess)
		}
		if encInfo.EFrameType != VideoFrameTypeSkip {
			c := 0
			for iLayer := 0; iLayer < int(encInfo.ILayerNum); iLayer++ {
				pLayerBsInfo := &encInfo.SLayerInfo[iLayer]
				var iLayerSize int32
				nallens := unsafe.Slice(pLayerBsInfo.PNalLengthInByte, pLayerBsInfo.INalCount)
				for _, l := range nallens {
					iLayerSize += l
				}
				nals := unsafe.Slice(pLayerBsInfo.PBsBuf, iLayerSize)
				c += int(iLayerSize)
				buf.Write(nals)
			}
			chunk = append(chunk, c)
		}
	}
	fh := md5.New()
	fh.Write(buf.Bytes())
	h := fmt.Sprintf("%X", fh.Sum(nil))
	if h != "7088D614072098205142FD991B08811F" {
		t.Fatal(h)
	}

	//DecodeParser

	var ppdec *ISVCDecoder
	if ret := WelsCreateDecoder(&ppdec); ret != 0 || ppdec == nil {
		log.Fatalln("failed WelsCreateDecoder:", ret, ppdec)
	}
	defer WelsDestroyDecoder(ppdec)

	var op int = 2
	ppdec.SetOption(DECODER_OPTION_TRACE_LEVEL, &op)
	op = 0
	ppdec.SetOption(DECODER_OPTION_NUM_OF_THREADS, &op)

	sDstParseInfo := SParserBsInfo{}

	sDecParam := SDecodingParam{
		BParseOnly: true,
	}

	if r := ppdec.Initialize(&sDecParam); r != 0 {
		log.Fatalln("failed Initialize.", r)
	}
	defer ppdec.Uninitialize()
	dataoffset := 0
	src := buf.Bytes()
	c := 0
	for _, l := range chunk {
		data := src[dataoffset : dataoffset+l]
		dataoffset += l
		if len(data) > 0 {
			if r := ppdec.DecodeParser(data, len(data), &sDstParseInfo); r != 0 {
				t.Fatal("DecodeParser", r)
			}
			if sDstParseInfo.INalNum > 0 {
				c++
				s := unsafe.Slice((*int32)(sDstParseInfo.PNalLenInByte), sDstParseInfo.INalNum)
				var smax int32 = 0
				for i := 0; i < int(sDstParseInfo.INalNum); i++ {
					ss := unsafe.Slice((*byte)(sDstParseInfo.PDstBuff), smax+s[i])
					ss = ss[smax:]
					smax += s[i]
					if sDstParseInfo.ISpsWidthInPixel != 1920 {
						t.Fatal("1920 !=", sDstParseInfo.ISpsWidthInPixel)
					}
					if sDstParseInfo.ISpsHeightInPixel != 1080 {
						t.Fatal("1080 !=", sDstParseInfo.ISpsHeightInPixel)
					}
					if txt := fmt.Sprintf("%X", ss[:4]); txt != "00000001" {
						t.Fatal("00000001 !=", txt)
					}
				}

			}
		}
	}
	if r := ppdec.DecodeParser(nil, 0, &sDstParseInfo); r == 0 {
		if sDstParseInfo.INalNum > 0 {
			c++
		}
	}
	if c != 4 {
		t.Fatal(c)
	}
}

func TestWelsGetCodecVersion(t *testing.T) {
	ver := OpenH264Version{
		UMajor:    OPENH264_MAJOR,
		UMinor:    OPENH264_MINOR,
		URevision: OPENH264_REVISION,
		UReserved: OPENH264_RESERVED,
	}

	if gotVer := WelsGetCodecVersion(); !reflect.DeepEqual(gotVer, ver) {
		t.Errorf("WelsGetCodecVersion() = %v, want %v", gotVer, ver)
	}
}
