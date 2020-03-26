package delivery_settings

type Settings = struct {
	Id int `json:"id"`
	CityId int `json:"city_id"`
	PriceForDelivery int `json:"price_for_delivery"`
	FreeDelivery int `json:"free_delivery"`
	IsDelivery bool `json:"is_delivery"`
}
