package model
import (
	"time"
)

type GcsServiceAccountInput struct {
	// Name of the resource. Can be freely chosen by the user.
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
	// GCS projectId (required)
	ServiceAccountCredentials string `json:"serviceAccountCredentials,omitempty"`
	// Name of the bucket (required)
	BucketName string `json:"bucketName,omitempty"`
	CloudRegion GoogleCloudRegion `json:"cloudRegion,omitempty"`
}
func (o GcsServiceAccountInput) InputType() InputType {
    return InputType_GCS_SERVICE_ACCOUNT
}

