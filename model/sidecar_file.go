package model
import (
	"time"
)

// A file that is added to an encoding. The size limit for a sidecar file is 10 MB
type SidecarFile struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp expressed in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp expressed in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource
	Id string `json:"id,omitempty"`
	// Id of input
	InputId string `json:"inputId,omitempty"`
	// Path to sidecar file
	InputPath string `json:"inputPath,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	ErrorMode SidecarErrorMode `json:"errorMode,omitempty"`
}

