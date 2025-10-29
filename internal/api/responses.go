package api

type CalculatePacksResponse struct {
	PacksUsed []PackUsage `json:"packs_used"`
}

type PackUsage struct {
	PackSize int `json:"pack_size"`
	Quantity int `json:"quantity"`
}

type PackListResponse struct {
	Packs []PackInfo `json:"packs"`
}

type PackInfo struct {
	Size int `json:"size"`
}
