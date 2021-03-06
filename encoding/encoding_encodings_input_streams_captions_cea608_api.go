package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsCaptionsCea608Api struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsCaptionsCea608Api(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsCaptionsCea608Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsCaptionsCea608Api{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsCaptionsCea608Api) Create(encodingId string, cea608CaptionInputStream model.Cea608CaptionInputStream) (*model.Cea608CaptionInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.Cea608CaptionInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/captions/cea608", &cea608CaptionInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsCaptionsCea608Api) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/captions/cea608/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsCaptionsCea608Api) Get(encodingId string, inputStreamId string) (*model.Cea608CaptionInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.Cea608CaptionInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/captions/cea608/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsCaptionsCea608Api) List(encodingId string, queryParams ...func(*query.Cea608CaptionInputStreamListQueryParams)) (*pagination.Cea608CaptionInputStreamsListPagination, error) {
    queryParameters := &query.Cea608CaptionInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.Cea608CaptionInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/captions/cea608", nil, &responseModel, reqParams)
    return responseModel, err
}

