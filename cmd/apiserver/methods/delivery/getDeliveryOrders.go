package delivery

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDeliveryOrders(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []Order
		rows, err := db.Query(`SELECT * FROM delivery`)

		if err != nil {
			panic(err.Error())
		}

		var id, oddMoney, restaurantId, countPerson, discountCard, orderId, addressId, status, paymentStatus, totalPrice, deliveryPrice, userId int
		var timeDelivery, createAt int64
		var surname, phone, email, paymentType, deliveryType, comment string
		var ruleAgree bool
		for rows.Next() {
			err := rows.Scan(&id, &surname, &phone, &email, &paymentType,
				&deliveryType, &oddMoney, &restaurantId, &timeDelivery, &countPerson,
				&discountCard, &comment, &orderId, &addressId, &createAt, &status, &paymentStatus, &totalPrice, &deliveryPrice, &ruleAgree, &userId)

			if err != nil {
				panic(err.Error())
			}

			address := Address{}
			err = db.QueryRow(`SELECT city_id, street, house, flat, floor, intercom FROM addresses WHERE id = $1`, addressId).Scan(
				&address.City,
				&address.Street,
				&address.House,
				&address.Flat,
				&address.Floor,
				&address.Intercom)
			if err != nil {
				panic(err.Error())
			}

			dishes := []Dish{}

			rowsDishes, err := db.Query(`SELECT id, count, price FROM delivery_items WHERE order_id = $1`, orderId)
			if err != nil {
				panic(err.Error())
			}

			var dishId, dishCount, dishPrice int
			for rowsDishes.Next() {
				err = rowsDishes.Scan(&dishId, &dishCount, &dishPrice)
				if err != nil {
					panic(err.Error())
				}

				dish := Dish{
					Id:    dishId,
					Count: dishCount,
					Price: dishPrice,
				}
				dishes = append(dishes, dish)
			}

			order := Order{
				Id:            id,
				Surname:       surname,
				Phone:         phone,
				Email:         email,
				PaymentType:   paymentType,
				DeliveryType:  deliveryType,
				OddMoney:      oddMoney,
				Address:       address,
				RestaurantId:  restaurantId,
				TimeDelivery:  timeDelivery,
				CountPerson:   countPerson,
				DiscountCard:  discountCard,
				Comment:       comment,
				Delivery:      Delivery{
					Order:         dishes,
					TotalPrice:    totalPrice,
					DeliveryPrice: deliveryPrice,
				},
				CreateAt:      createAt,
				Status:        status,
				PaymentStatus: paymentStatus,
				RuleAgree: ruleAgree,
				UserId: userId,
			}
			orders = append(orders, order)
		}
		c.JSON(http.StatusOK, orders)
	}
}
