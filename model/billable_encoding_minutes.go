package model

type BillableEncodingMinutes struct {
	EncodingMode EncodingMode `json:"encodingMode,omitempty"`
	Codec CodecConfigType `json:"codec,omitempty"`
	PerTitleResultStream StatisticsPerTitleStream `json:"perTitleResultStream,omitempty"`
	PsnrMode PsnrPerStreamMode `json:"psnrMode,omitempty"`
	// Name of the preset configuration used for the codec configuration or \"CUSTOM\" if any preset values were overridden
	Preset string `json:"preset,omitempty"`
	// Indicates if the stream was part of a live encoding.
	Live *bool `json:"live,omitempty"`
	BillableMinutes *BillableEncodingMinutesDetails `json:"billableMinutes,omitempty"`
}

