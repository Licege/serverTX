package delivery

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostDelivery(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := Order{}
		err := c.BindJSON(&requestBody)
		if err != nil {
			panic(err.Error())
		}

		if !Valid(requestBody, db) {
			c.Status(http.StatusBadRequest)
		} else {
			var addressId int
			err := db.QueryRow(`INSERT INTO addresses(city_id, street, house, flat, floor, intercom) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
				requestBody.Address.City,
				requestBody.Address.Street,
				requestBody.Address.House,
				requestBody.Address.Flat,
				requestBody.Address.Floor,
				requestBody.Address.Intercom).Scan(&addressId)

			if err != nil {
				panic(err.Error())
			}

			var orderId int
			err = db.QueryRow(`INSERT INTO delivery(surname, phone, email, payment_type, delivery_type, odd_money, restaurant_id, time_delivery, discount_card, comment, address_id, create_at, total_price, delivery_price, rule_agree, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING order_id`,
				requestBody.Surname,
				requestBody.Phone,
				requestBody.Email,
				requestBody.PaymentType,
				requestBody.DeliveryType,
				requestBody.OddMoney,
				requestBody.RestaurantId,
				requestBody.TimeDelivery,
				requestBody.DiscountCard,
				requestBody.Comment,
				addressId,
				requestBody.CreateAt,
				requestBody.Delivery.TotalPrice,
				requestBody.Delivery.DeliveryPrice,
				requestBody.RuleAgree,
				requestBody.UserId).Scan(&orderId)
			if err != nil {
				panic(err.Error())
			}

			for i := 0; i < len(requestBody.Delivery.Order); i++ {
				_, err := db.Exec(`INSERT INTO delivery_items(order_id, id, count, price) VALUES ($1, $2, $3, $4)`,
					orderId,
					requestBody.Delivery.Order[i].Id,
					requestBody.Delivery.Order[i].Count,
					requestBody.Delivery.Order[i].Price)

				if err != nil {
					panic(err.Error())
				}
			}

			c.Status(http.StatusOK)
		}
	}
}

func Valid(data Order, db *sql.DB) bool {
	/*Получаем меню*/
	var menu []ShortDish

	rows, err := db.Query(`SELECT id, price FROM dishes`)
	if err != nil {
		panic(err.Error())
	}
	var id, price int

	for rows.Next(){
		err := rows.Scan(&id, &price)
		if err != nil {
			panic(err.Error())
		}
		newDish := ShortDish{
			Id:          id,
			Price:       price,
		}
		menu = append(menu, newDish)
	}

	/*Проверяем цены*/
	count := 0
	totalPrice := 0

	for i := 0; i < len(data.Delivery.Order); i++ {
		for j := 0; j < len(menu); j++ {
			if data.Delivery.Order[i].Id == menu[j].Id {
				totalPrice += menu[j].Price * data.Delivery.Order[i].Count
				count ++

				if data.Delivery.Order[i].Price != menu[j].Price * data.Delivery.Order[i].Count {
					return false
				}
			}
		}
	}

	/*Кол-во проверок = кол-ву блюд в заказе и корректность totalPrice*/
	if len(data.Delivery.Order) != count || data.Delivery.TotalPrice != totalPrice {
		return false
	}

	/*Проверка параметров доставки*/

	if data.DeliveryType == "home" {
		settings := Settings{}
		err = db.QueryRow(`SELECT price_for_delivery, free_delivery, is_delivery FROM delivery_settings WHERE city_id = $1`, data.Address.City).Scan(
			&settings.PriceForDelivery,
			&settings.FreeDelivery,
			&settings.IsDelivery)
		if err != nil {
			panic(err.Error())
		}
		if !settings.IsDelivery {
			return false
		}
		if (data.Delivery.DeliveryPrice != 0 && data.Delivery.TotalPrice > settings.FreeDelivery) || (data.Delivery.DeliveryPrice != settings.PriceForDelivery && data.Delivery.TotalPrice < settings.FreeDelivery) {
			return false
		}
	} else if data.DeliveryType == "restaurant" {
		if data.Delivery.DeliveryPrice != 0 {
			return false
		}
	} else {
		return false
	}



	/*Пользовательское соглашение*/
	if !data.RuleAgree {
		return false
	}

	return true
}
