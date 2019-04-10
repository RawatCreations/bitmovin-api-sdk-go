package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMachineLearningObjectDetectionApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMachineLearningObjectDetectionCustomdataApi
    Results *EncodingEncodingsMachineLearningObjectDetectionResultsApi
}

func NewEncodingEncodingsMachineLearningObjectDetectionApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMachineLearningObjectDetectionApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMachineLearningObjectDetectionApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMachineLearningObjectDetectionCustomdataApi(configs...)
    api.Customdata = customdataApi
    resultsApi, err := NewEncodingEncodingsMachineLearningObjectDetectionResultsApi(configs...)
    api.Results = resultsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMachineLearningObjectDetectionApi) List(encodingId string, queryParams ...func(*query.ObjectDetectionConfigurationListQueryParams)) (*pagination.ObjectDetectionConfigurationsListPagination, error) {
    queryParameters := &query.ObjectDetectionConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.ObjectDetectionConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMachineLearningObjectDetectionApi) Create(encodingId string, objectDetectionConfiguration model.ObjectDetectionConfiguration) (*model.ObjectDetectionConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }
    payload := model.ObjectDetectionConfiguration(objectDetectionConfiguration)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/machine-learning/object-detection", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMachineLearningObjectDetectionApi) Get(encodingId string, objectDetectionId string) (*model.ObjectDetectionConfiguration, error) {
    var resp *model.ObjectDetectionConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["object_detection_id"] = objectDetectionId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMachineLearningObjectDetectionApi) Delete(encodingId string, objectDetectionId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["object_detection_id"] = objectDetectionId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}", &resp, reqParams)
    return resp, err
}
