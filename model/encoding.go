package model
import (
	"time"
)

type Encoding struct {
	// Name of the resource. Can be freely chosen by the user. (required)
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	CloudRegion CloudRegion `json:"cloudRegion,omitempty"`
	// Specify a list of regions which are used in case the preferred region is down. Currently there are several restrictions. - The region has to be specific or AUTO - The region has to be for the same cloud provider as the default one - You can only configure at most 3 fallback regions 
	FallbackCloudRegions []CloudRegion `json:"fallbackCloudRegions,omitempty"`
	// Version of the encoder
	EncoderVersion string `json:"encoderVersion,omitempty"`
	// Define an external infrastructure to run the encoding on. Note If you set this value, the `cloudRegion` must be 'EXTERNAL'.
	InfrastructureId string `json:"infrastructureId,omitempty"`
	Infrastructure *InfrastructureSettings `json:"infrastructure,omitempty"`
	// Will be set to the encoder version that was actually used for the encoding. This is especially useful when starting an encoding with a version tag like STABLE or BETA.
	SelectedEncoderVersion string `json:"selectedEncoderVersion,omitempty"`
	// Will be set to the encoding mode that was actually used for the encoding. This is especially useful when starting an encoding with encoding mode STANDARD.
	SelectedEncodingMode EncodingMode `json:"selectedEncodingMode,omitempty"`
	// Contains the region which was selected when cloudregion:AUTO was specified
	SelectedCloudRegion CloudRegion `json:"selectedCloudRegion,omitempty"`
	// The current status of the encoding.
	Status Status `json:"status,omitempty"`
	// You may pass a list of groups associated with this encoding. This will enable you to group results in the statistics resource
	Labels []string `json:"labels,omitempty"`
}

