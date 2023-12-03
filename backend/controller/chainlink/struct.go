package chainlinkController

type ChainlinkPriceRequest struct {
	TokenAddress     string `json:"token_address"`
	PriceFeedAddress string `json:"price_feed_address"`
}
