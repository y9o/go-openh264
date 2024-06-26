// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo.exe -godefs _codec.go

package openh264

type ISVCEncoderVtbl struct {
	Initialize		*[0]byte
	InitializeExt		*[0]byte
	GetDefaultParams	*[0]byte
	Uninitialize		*[0]byte
	EncodeFrame		*[0]byte
	EncodeParameterSets	*[0]byte
	ForceIntraFrame		*[0]byte
	SetOption		*[0]byte
	GetOption		*[0]byte
}
type ISVCDecoderVtbl struct {
	Initialize		*[0]byte
	Uninitialize		*[0]byte
	DecodeFrame		*[0]byte
	DecodeFrameNoDelay	*[0]byte
	DecodeFrame2		*[0]byte
	FlushFrame		*[0]byte
	DecodeParser		*[0]byte
	DecodeFrameEx		*[0]byte
	SetOption		*[0]byte
	GetOption		*[0]byte
}

type OpenH264Version struct {
	UMajor		uint32
	UMinor		uint32
	URevision	uint32
	UReserved	uint32
}

const (
	DsErrorFree		= 0x0
	DsFramePending		= 0x1
	DsRefLost		= 0x2
	DsBitstreamError	= 0x4
	DsDepLayerLost		= 0x8
	DsNoParamSets		= 0x10
	DsDataErrorConcealed	= 0x20
	DsRefListNullPtrs	= 0x40
	DsInvalidArgument	= 0x1000
	DsInitialOptExpected	= 0x2000
	DsOutOfMemory		= 0x4000
	DsDstBufNeedExpan	= 0x8000
)

const (
	ENCODER_OPTION_DATAFORMAT		= 0x0
	ENCODER_OPTION_IDR_INTERVAL		= 0x1
	ENCODER_OPTION_SVC_ENCODE_PARAM_BASE	= 0x2
	ENCODER_OPTION_SVC_ENCODE_PARAM_EXT	= 0x3
	ENCODER_OPTION_FRAME_RATE		= 0x4
	ENCODER_OPTION_BITRATE			= 0x5
	ENCODER_OPTION_MAX_BITRATE		= 0x6
	ENCODER_OPTION_INTER_SPATIAL_PRED	= 0x7
	ENCODER_OPTION_RC_MODE			= 0x8
	ENCODER_OPTION_RC_FRAME_SKIP		= 0x9
	ENCODER_PADDING_PADDING			= 0xa
	ENCODER_OPTION_PROFILE			= 0xb
	ENCODER_OPTION_LEVEL			= 0xc
	ENCODER_OPTION_NUMBER_REF		= 0xd
	ENCODER_OPTION_DELIVERY_STATUS		= 0xe
	ENCODER_LTR_RECOVERY_REQUEST		= 0xf
	ENCODER_LTR_MARKING_FEEDBACK		= 0x10
	ENCODER_LTR_MARKING_PERIOD		= 0x11
	ENCODER_OPTION_LTR			= 0x12
	ENCODER_OPTION_COMPLEXITY		= 0x13
	ENCODER_OPTION_ENABLE_SSEI		= 0x14
	ENCODER_OPTION_ENABLE_PREFIX_NAL_ADDING	= 0x15
	ENCODER_OPTION_SPS_PPS_ID_STRATEGY	= 0x16
	ENCODER_OPTION_CURRENT_PATH		= 0x17
	ENCODER_OPTION_DUMP_FILE		= 0x18
	ENCODER_OPTION_TRACE_LEVEL		= 0x19
	ENCODER_OPTION_TRACE_CALLBACK		= 0x1a
	ENCODER_OPTION_TRACE_CALLBACK_CONTEXT	= 0x1b
	ENCODER_OPTION_GET_STATISTICS		= 0x1c
	ENCODER_OPTION_STATISTICS_LOG_INTERVAL	= 0x1d
	ENCODER_OPTION_IS_LOSSLESS_LINK		= 0x1e
	ENCODER_OPTION_BITS_VARY_PERCENTAGE	= 0x1f
)

const (
	DECODER_OPTION_END_OF_STREAM				= 0x1
	DECODER_OPTION_VCL_NAL					= 0x2
	DECODER_OPTION_TEMPORAL_ID				= 0x3
	DECODER_OPTION_FRAME_NUM				= 0x4
	DECODER_OPTION_IDR_PIC_ID				= 0x5
	DECODER_OPTION_LTR_MARKING_FLAG				= 0x6
	DECODER_OPTION_LTR_MARKED_FRAME_NUM			= 0x7
	DECODER_OPTION_ERROR_CON_IDC				= 0x8
	DECODER_OPTION_TRACE_LEVEL				= 0x9
	DECODER_OPTION_TRACE_CALLBACK				= 0xa
	DECODER_OPTION_TRACE_CALLBACK_CONTEXT			= 0xb
	DECODER_OPTION_GET_STATISTICS				= 0xc
	DECODER_OPTION_GET_SAR_INFO				= 0xd
	DECODER_OPTION_PROFILE					= 0xe
	DECODER_OPTION_LEVEL					= 0xf
	DECODER_OPTION_STATISTICS_LOG_INTERVAL			= 0x10
	DECODER_OPTION_IS_REF_PIC				= 0x11
	DECODER_OPTION_NUM_OF_FRAMES_REMAINING_IN_BUFFER	= 0x12
	DECODER_OPTION_NUM_OF_THREADS				= 0x13
)

const (
	ERROR_CON_DISABLE					= 0x0
	ERROR_CON_FRAME_COPY					= 0x1
	ERROR_CON_SLICE_COPY					= 0x2
	ERROR_CON_FRAME_COPY_CROSS_IDR				= 0x3
	ERROR_CON_SLICE_COPY_CROSS_IDR				= 0x4
	ERROR_CON_SLICE_COPY_CROSS_IDR_FREEZE_RES_CHANGE	= 0x5
	ERROR_CON_SLICE_MV_COPY_CROSS_IDR			= 0x6
	ERROR_CON_SLICE_MV_COPY_CROSS_IDR_FREEZE_RES_CHANGE	= 0x7
)

const (
	FEEDBACK_NON_VCL_NAL	= 0x0
	FEEDBACK_VCL_NAL	= 0x1
	FEEDBACK_UNKNOWN_NAL	= 0x2
)

const (
	NON_VIDEO_CODING_LAYER	= 0x0
	VIDEO_CODING_LAYER	= 0x1
)

const (
	SPATIAL_LAYER_0		= 0x0
	SPATIAL_LAYER_1		= 0x1
	SPATIAL_LAYER_2		= 0x2
	SPATIAL_LAYER_3		= 0x3
	SPATIAL_LAYER_ALL	= 0x4
)

const (
	VIDEO_BITSTREAM_AVC	= 0x0
	VIDEO_BITSTREAM_SVC	= 0x1
	VIDEO_BITSTREAM_DEFAULT	= 0x1
)

const (
	NO_RECOVERY_REQUSET	= 0x0
	LTR_RECOVERY_REQUEST	= 0x1
	IDR_RECOVERY_REQUEST	= 0x2
	NO_LTR_MARKING_FEEDBACK	= 0x3
	LTR_MARKING_SUCCESS	= 0x4
	LTR_MARKING_FAILED	= 0x5
)

type SLTRRecoverRequest struct {
	UiFeedbackType		uint32
	UiIDRPicId		uint32
	ILastCorrectFrameNum	int32
	ICurrentFrameNum	int32
	ILayerId		int32
}
type SLTRMarkingFeedback struct {
	UiFeedbackType	uint32
	UiIDRPicId	uint32
	ILTRFrameNum	int32
	ILayerId	int32
}
type SLTRConfig struct {
	BEnableLongTermReference	bool
	ILTRRefNum			int32
}

const (
	RC_QUALITY_MODE			= 0x0
	RC_BITRATE_MODE			= 0x1
	RC_BUFFERBASED_MODE		= 0x2
	RC_TIMESTAMP_MODE		= 0x3
	RC_BITRATE_MODE_POST_SKIP	= 0x4
	RC_OFF_MODE			= -0x1
)

const (
	PRO_UNKNOWN		= 0x0
	PRO_BASELINE		= 0x42
	PRO_MAIN		= 0x4d
	PRO_EXTENDED		= 0x58
	PRO_HIGH		= 0x64
	PRO_HIGH10		= 0x6e
	PRO_HIGH422		= 0x7a
	PRO_HIGH444		= 0x90
	PRO_CAVLC444		= 0xf4
	PRO_SCALABLE_BASELINE	= 0x53
	PRO_SCALABLE_HIGH	= 0x56
)

const (
	LEVEL_UNKNOWN	= 0x0
	LEVEL_1_0	= 0xa
	LEVEL_1_B	= 0x9
	LEVEL_1_1	= 0xb
	LEVEL_1_2	= 0xc
	LEVEL_1_3	= 0xd
	LEVEL_2_0	= 0x14
	LEVEL_2_1	= 0x15
	LEVEL_2_2	= 0x16
	LEVEL_3_0	= 0x1e
	LEVEL_3_1	= 0x1f
	LEVEL_3_2	= 0x20
	LEVEL_4_0	= 0x28
	LEVEL_4_1	= 0x29
	LEVEL_4_2	= 0x2a
	LEVEL_5_0	= 0x32
	LEVEL_5_1	= 0x33
	LEVEL_5_2	= 0x34
)
const (
	WELS_LOG_QUIET		= 0x0
	WELS_LOG_ERROR		= 0x1
	WELS_LOG_WARNING	= 0x2
	WELS_LOG_INFO		= 0x4
	WELS_LOG_DEBUG		= 0x8
	WELS_LOG_DETAIL		= 0x10
	WELS_LOG_RESV		= 0x20
	WELS_LOG_LEVEL_COUNT	= 0x6
	WELS_LOG_DEFAULT	= 0x2
)

const (
	SM_SINGLE_SLICE		= 0x0
	SM_FIXEDSLCNUM_SLICE	= 0x1
	SM_RASTER_SLICE		= 0x2
	SM_SIZELIMITED_SLICE	= 0x3
	SM_RESERVED		= 0x4
)

type SSliceArgument struct {
	UiSliceMode		uint32
	UiSliceNum		uint32
	UiSliceMbNum		[35]uint32
	UiSliceSizeConstraint	uint32
}

const (
	VF_COMPONENT	= 0x0
	VF_PAL		= 0x1
	VF_NTSC		= 0x2
	VF_SECAM	= 0x3
	VF_MAC		= 0x4
	VF_UNDEF	= 0x5
	VF_NUM_ENUM	= 0x6
)

const (
	CP_RESERVED0	= 0x0
	CP_BT709	= 0x1
	CP_UNDEF	= 0x2
	CP_RESERVED3	= 0x3
	CP_BT470M	= 0x4
	CP_BT470BG	= 0x5
	CP_SMPTE170M	= 0x6
	CP_SMPTE240M	= 0x7
	CP_FILM		= 0x8
	CP_BT2020	= 0x9
	CP_NUM_ENUM	= 0xa
)

const (
	TRC_RESERVED0		= 0x0
	TRC_BT709		= 0x1
	TRC_UNDEF		= 0x2
	TRC_RESERVED3		= 0x3
	TRC_BT470M		= 0x4
	TRC_BT470BG		= 0x5
	TRC_SMPTE170M		= 0x6
	TRC_SMPTE240M		= 0x7
	TRC_LINEAR		= 0x8
	TRC_LOG100		= 0x9
	TRC_LOG316		= 0xa
	TRC_IEC61966_2_4	= 0xb
	TRC_BT1361E		= 0xc
	TRC_IEC61966_2_1	= 0xd
	TRC_BT2020_10		= 0xe
	TRC_BT2020_12		= 0xf
	TRC_NUM_ENUM		= 0x10
)

const (
	CM_GBR		= 0x0
	CM_BT709	= 0x1
	CM_UNDEF	= 0x2
	CM_RESERVED3	= 0x3
	CM_FCC		= 0x4
	CM_BT470BG	= 0x5
	CM_SMPTE170M	= 0x6
	CM_SMPTE240M	= 0x7
	CM_YCGCO	= 0x8
	CM_BT2020NC	= 0x9
	CM_BT2020C	= 0xa
	CM_NUM_ENUM	= 0xb
)

const (
	ASP_UNSPECIFIED	= 0x0
	ASP_1x1		= 0x1
	ASP_12x11	= 0x2
	ASP_10x11	= 0x3
	ASP_16x11	= 0x4
	ASP_40x33	= 0x5
	ASP_24x11	= 0x6
	ASP_20x11	= 0x7
	ASP_32x11	= 0x8
	ASP_80x33	= 0x9
	ASP_18x11	= 0xa
	ASP_15x11	= 0xb
	ASP_64x33	= 0xc
	ASP_160x99	= 0xd
	ASP_EXT_SAR	= 0xff
)

type SSpatialLayerConfig struct {
	IVideoWidth			int32
	IVideoHeight			int32
	FFrameRate			float32
	ISpatialBitrate			int32
	IMaxSpatialBitrate		int32
	UiProfileIdc			uint32
	UiLevelIdc			uint32
	IDLayerQp			int32
	SSliceArgument			SSliceArgument
	BVideoSignalTypePresent		bool
	UiVideoFormat			uint8
	BFullRange			bool
	BColorDescriptionPresent	bool
	UiColorPrimaries		uint8
	UiTransferCharacteristics	uint8
	UiColorMatrix			uint8
	BAspectRatioPresent		bool
	EAspectRatio			uint32
	SAspectRatioExtWidth		uint16
	SAspectRatioExtHeight		uint16
}

const (
	CAMERA_VIDEO_REAL_TIME		= 0x0
	SCREEN_CONTENT_REAL_TIME	= 0x1
	CAMERA_VIDEO_NON_REAL_TIME	= 0x2
	SCREEN_CONTENT_NON_REAL_TIME	= 0x3
	INPUT_CONTENT_TYPE_ALL		= 0x4
)

const (
	LOW_COMPLEXITY		= 0x0
	MEDIUM_COMPLEXITY	= 0x1
	HIGH_COMPLEXITY		= 0x2
)

const (
	CONSTANT_ID			= 0x0
	INCREASING_ID			= 0x1
	SPS_LISTING			= 0x2
	SPS_LISTING_AND_PPS_INCREASING	= 0x3
	SPS_PPS_LISTING			= 0x6
)

type SEncParamBase struct {
	IUsageType	uint32
	IPicWidth	int32
	IPicHeight	int32
	ITargetBitrate	int32
	IRCMode		int32
	FMaxFrameRate	float32
}
type SEncParamExt struct {
	IUsageType			uint32
	IPicWidth			int32
	IPicHeight			int32
	ITargetBitrate			int32
	IRCMode				int32
	FMaxFrameRate			float32
	ITemporalLayerNum		int32
	ISpatialLayerNum		int32
	SSpatialLayers			[4]SSpatialLayerConfig
	IComplexityMode			uint32
	UiIntraPeriod			uint32
	INumRefFrame			int32
	ESpsPpsIdStrategy		uint32
	BPrefixNalAddingCtrl		bool
	BEnableSSEI			bool
	BSimulcastAVC			bool
	IPaddingFlag			int32
	IEntropyCodingModeFlag		int32
	BEnableFrameSkip		bool
	IMaxBitrate			int32
	IMaxQp				int32
	IMinQp				int32
	UiMaxNalSize			uint32
	BEnableLongTermReference	bool
	ILTRRefNum			int32
	ILtrMarkPeriod			uint32
	IMultipleThreadIdc		uint16
	BUseLoadBalancing		bool
	ILoopFilterDisableIdc		int32
	ILoopFilterAlphaC0Offset	int32
	ILoopFilterBetaOffset		int32
	BEnableDenoise			bool
	BEnableBackgroundDetection	bool
	BEnableAdaptiveQuant		bool
	BEnableFrameCroppingFlag	bool
	BEnableSceneChangeDetect	bool
	BIsLosslessLink			bool
	BFixRCOverShoot			bool
	IIdrBitrateRatio		int32
}
type SVideoProperty struct {
	Size		uint32
	EVideoBsType	uint32
}
type SDecodingParam struct {
	PFileNameRestructed	*int8
	UiCpuLoad		uint32
	UiTargetDqLayer		uint8
	EEcActiveIdc		uint32
	BParseOnly		bool
	SVideoProperty		SVideoProperty
}
type SLayerBSInfo struct {
	UiTemporalId		uint8
	UiSpatialId		uint8
	UiQualityId		uint8
	EFrameType		uint32
	UiLayerType		uint8
	ISubSeqId		int32
	INalCount		int32
	PNalLengthInByte	*int32
	PBsBuf			*uint8
}
type SFrameBSInfo struct {
	ILayerNum		int32
	SLayerInfo		[128]SLayerBSInfo
	EFrameType		uint32
	IFrameSizeInBytes	int32
	UiTimeStamp		int64
}
type SSourcePicture struct {
	IColorFormat	int32
	IStride		[4]int32
	PData		[4]*uint8
	IPicWidth	int32
	IPicHeight	int32
	UiTimeStamp	int64
}
type SBitrateInfo struct {
	ILayer		uint32
	IBitrate	int32
}
type SDumpLayer struct {
	ILayer		int32
	PFileName	*int8
}
type SProfileInfo struct {
	ILayer		int32
	UiProfileIdc	uint32
}
type SLevelInfo struct {
	ILayer		int32
	UiLevelIdc	uint32
}
type SDeliveryStatus struct {
	BDeliveryFlag	bool
	IDropFrameType	int32
	IDropFrameSize	int32
}
type SDecoderCapability struct {
	IProfileIdc	int32
	IProfileIop	int32
	ILevelIdc	int32
	IMaxMbps	int32
	IMaxFs		int32
	IMaxCpb		int32
	IMaxDpb		int32
	IMaxBr		int32
	BRedPicCap	bool
	Pad_cgo_0	[3]byte
}
type SParserBsInfo struct {
	INalNum			int32
	PNalLenInByte		*int32
	PDstBuff		*uint8
	ISpsWidthInPixel	int32
	ISpsHeightInPixel	int32
	UiInBsTimeStamp		uint64
	UiOutBsTimeStamp	uint64
}
type SEncoderStatistics struct {
	UiWidth				uint32
	UiHeight			uint32
	FAverageFrameSpeedInMs		float32
	FAverageFrameRate		float32
	FLatestFrameRate		float32
	UiBitRate			uint32
	UiAverageFrameQP		uint32
	UiInputFrameCount		uint32
	UiSkippedFrameCount		uint32
	UiResolutionChangeTimes		uint32
	UiIDRReqNum			uint32
	UiIDRSentNum			uint32
	UiLTRSentNum			uint32
	IStatisticsTs			int64
	ITotalEncodedBytes		uint32
	ILastStatisticsBytes		uint32
	ILastStatisticsFrameCount	uint32
	Pad_cgo_0			[4]byte
}
type SDecoderStatistics struct {
	UiWidth				uint32
	UiHeight			uint32
	FAverageFrameSpeedInMs		float32
	FActualAverageFrameSpeedInMs	float32
	UiDecodedFrameCount		uint32
	UiResolutionChangeTimes		uint32
	UiIDRCorrectNum			uint32
	UiAvgEcRatio			uint32
	UiAvgEcPropRatio		uint32
	UiEcIDRNum			uint32
	UiEcFrameNum			uint32
	UiIDRLostNum			uint32
	UiFreezingIDRNum		uint32
	UiFreezingNonIDRNum		uint32
	IAvgLumaQp			int32
	ISpsReportErrorNum		int32
	ISubSpsReportErrorNum		int32
	IPpsReportErrorNum		int32
	ISpsNoExistNalNum		int32
	ISubSpsNoExistNalNum		int32
	IPpsNoExistNalNum		int32
	UiProfile			uint32
	UiLevel				uint32
	ICurrentActiveSpsId		int32
	ICurrentActivePpsId		int32
	IStatisticsLogInterval		uint32
}
type SVuiSarInfo struct {
	UiSarWidth			uint32
	UiSarHeight			uint32
	BOverscanAppropriateFlag	bool
	Pad_cgo_0			[3]byte
}

const (
	MAX_TEMPORAL_LAYER_NUM	= 4
	MAX_SPATIAL_LAYER_NUM	= 4
	MAX_QUALITY_LAYER_NUM	= 4
	MAX_LAYER_NUM_OF_FRAME	= 128
	MAX_NAL_UNITS_IN_LAYER	= 128
	MAX_RTP_PAYLOAD_LEN	= 1000
	AVERAGE_RTP_PAYLOAD_LEN	= 800
	AUTO_REF_PIC_COUNT	= -1
	UNSPECIFIED_BIT_RATE	= 0
)

const (
	VideoFormatRGB		= 0x1
	VideoFormatRGBA		= 0x2
	VideoFormatRGB555	= 0x3
	VideoFormatRGB565	= 0x4
	VideoFormatBGR		= 0x5
	VideoFormatBGRA		= 0x6
	VideoFormatABGR		= 0x7
	VideoFormatARGB		= 0x8
	VideoFormatYUY2		= 0x14
	VideoFormatYVYU		= 0x15
	VideoFormatUYVY		= 0x16
	VideoFormatI420		= 0x17
	VideoFormatYV12		= 0x18
	VideoFormatInternal	= 0x19
	VideoFormatNV12		= 0x1a
	VideoFormatVFlip	= 0x80000000
)

const (
	VideoFrameTypeInvalid	= 0x0
	VideoFrameTypeIDR	= 0x1
	VideoFrameTypeI		= 0x2
	VideoFrameTypeP		= 0x3
	VideoFrameTypeSkip	= 0x4
	VideoFrameTypeIPMixed	= 0x5
)

const (
	CmResultSuccess		= 0x0
	CmInitParaError		= 0x1
	CmUnknownReason		= 0x2
	CmMallocMemeError	= 0x3
	CmInitExpected		= 0x4
	CmUnsupportedData	= 0x5
)
const (
	DEBLOCKING_IDC_0	= 0x0
	DEBLOCKING_IDC_1	= 0x1
	DEBLOCKING_IDC_2	= 0x2
)
const (
	ET_NONE		= 0x0
	ET_IP_SCALE	= 0x1
	ET_FMO		= 0x2
	ET_IR_R1	= 0x4
	ET_IR_R2	= 0x8
	ET_IR_R3	= 0x10
	ET_FEC_HALF	= 0x20
	ET_FEC_FULL	= 0x40
	ET_RFS		= 0x80
)

type SliceInfo struct {
	PBufferOfSlices		*uint8
	ICodedSliceCount	int32
	PLengthOfSlices		*uint32
	IFecType		int32
	UiSliceIdx		uint8
	UiSliceCount		uint8
	IFrameIndex		int8
	UiNalRefIdc		uint8
	UiNalType		uint8
	UiContainingFinalNal	uint8
	Pad_cgo_0		[6]byte
}
type SRateThresholds struct {
	IWidth			int32
	IHeight			int32
	IThresholdOfInitRate	int32
	IThresholdOfMaxRate	int32
	IThresholdOfMinRate	int32
	IMinThresholdFrameRate	int32
	ISkipFrameRate		int32
	ISkipFrameStep		int32
}
type SSysMEMBuffer struct {
	IWidth	int32
	IHeight	int32
	IFormat	int32
	IStride	[2]int32
}
type SBufferInfo struct {
	IBufferStatus		int32
	UiInBsTimeStamp		uint64
	UiOutYuvTimeStamp	uint64
	UsrData			[20]byte
	PDst			[3]*uint8
}

const (
	FRAME_NUM_PARAM_SET	= -1
	FRAME_NUM_IDR		= 0
	DEBLOCKING_OFFSET	= 6
	DEBLOCKING_OFFSET_MINUS	= -6
)

const (
	OPENH264_MAJOR		= 2
	OPENH264_MINOR		= 4
	OPENH264_REVISION	= 1
	OPENH264_RESERVED	= 2401
)
