package notifications
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
)

type NotificationsWebhooksEncodingEncodingsApi struct {
    apiClient *common.ApiClient
    Finished *NotificationsWebhooksEncodingEncodingsFinishedApi
    Error *NotificationsWebhooksEncodingEncodingsErrorApi
    TransferError *NotificationsWebhooksEncodingEncodingsTransferErrorApi
}

func NewNotificationsWebhooksEncodingEncodingsApi(configs ...func(*common.ApiClient)) (*NotificationsWebhooksEncodingEncodingsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &NotificationsWebhooksEncodingEncodingsApi{apiClient: apiClient}

    finishedApi, err := NewNotificationsWebhooksEncodingEncodingsFinishedApi(configs...)
    api.Finished = finishedApi
    errorApi, err := NewNotificationsWebhooksEncodingEncodingsErrorApi(configs...)
    api.Error = errorApi
    transferErrorApi, err := NewNotificationsWebhooksEncodingEncodingsTransferErrorApi(configs...)
    api.TransferError = transferErrorApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

