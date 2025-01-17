package client

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/signalr/2023-02-01/signalr"
	webpubsub_v2023_02_01 "github.com/hashicorp/go-azure-sdk/resource-manager/webpubsub/2023-02-01"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	SignalRClient   *signalr.SignalRClient
	WebPubSubClient *webpubsub_v2023_02_01.Client
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	signalRClient, err := signalr.NewSignalRClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, err
	}
	o.Configure(signalRClient.Client, o.Authorizers.ResourceManager)

	webPubSubClient := webpubsub_v2023_02_01.NewClientWithBaseURI(o.ResourceManagerEndpoint, func(c *autorest.Client) {
		o.ConfigureClient(c, o.ResourceManagerAuthorizer)
	})

	return &Client{
		SignalRClient:   signalRClient,
		WebPubSubClient: &webPubSubClient,
	}, nil
}
