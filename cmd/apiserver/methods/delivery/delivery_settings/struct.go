package delivery_settings

type Settings = struct {
	CityId int `json:"city_id"`
	PriceForDelivery int `json:"price_for_delivery"`
	FreeDelivery int `json:"free_delivery"`
}
