package delivery

type Order = struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	PaymentType int `json:"payment_type"`
	DeliveryType int `json:"delivery_type"`
	CashChange int `json:"cash_change"`
	Address Address `json:"address"`
	RestaurantId int `json:"restaurant_id"`
	TimeDelivery int64 `json:"time_delivery"`
	CountPerson int `json:"count_person"`
	DiscountCard int `json:"discount_card"`
	Comment string `json:"comment"`
	Dishes []Dish `json:"dishes"`
}

type Dish = struct {
	Id int `json:"id"`
	Count int `json:"count"`
}

type Address = struct {
	City int `json:"city"`
	Street string `json:"street"`
	Porch string `json:"porch"`
	Apartment string `json:"apartment"`
	Floor int `json:"floor"`
	Intercom string `json:"intercom"`
}