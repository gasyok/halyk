package auction

import desc "github.com/gasyok/Assessment/pkg/auction/v1"

type Controller struct {
	auctionService auctionService

	desc.UnimplementedAuctionServer
}

func NewController(auctionService auctionService) *Controller {
	return &Controller{
		auctionService: auctionService,
	}
}

func (c *Controller) WithAuctionService(auctionService auctionService) {
	c.auctionService = auctionService
}
