package delivery_global_settings

type Settings = struct {
	IsDeliveryWorking bool `json:"is_delivery_working"`
	PhoneForSms string `json:"phone_for_sms"`
}
