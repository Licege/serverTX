package delivery

type Order = struct {
	Id int `json:"id"`
	Surname string `json:"surname"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	PaymentType string `json:"payment_type"`
	DeliveryType string `json:"delivery_type"`
	OddMoney int `json:"odd_money"`
	Address Address `json:"address"`
	RestaurantId int `json:"restaurant_id"`
	TimeDelivery int64 `json:"time_delivery"`
	CountPerson int `json:"count_person"`
	DiscountCard int `json:"discount_card"`
	Comment string `json:"comment"`
	Delivery Delivery `json:"delivery"`
	CreateAt int64 `json:"create_at"`
	Status int `json:"status"`
	PaymentStatus int `json:"payment_status"`
	RuleAgree bool `json:"rule_agree"`
	UserId int `json:"user_id"`
}

type Delivery = struct {
	Order []Dish `json:"order"`
	TotalPrice int `json:"total_price"`
	DeliveryPrice int `json:"delivery_price"`
}

type Dish = struct {
	Id int `json:"id"`
	Count int `json:"count"`
	Price int `json:"price"`
}

type Address = struct {
	City int `json:"city"`
	Street string `json:"street"`
	House string `json:"house"`
	Flat string `json:"apartment"`
	Floor string `json:"floor"`
	Intercom string `json:"intercom"`
}

type ShortDish = struct {
	Id int `json:"id"`
	Price int `json:"price"`
}

type Settings = struct {
	PriceForDelivery int `json:"price_for_delivery"`
	FreeDelivery int `json:"free_delivery"`
	IsDelivery bool `json:"is_delivery"`
}
