package api

type CalculatePacksRequest struct {
	OrderQuantity int `json:"order_quantity"`
}

type PackListRequest struct {
	Packs []int `json:"packs"`
}
