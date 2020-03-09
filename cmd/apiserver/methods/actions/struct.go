package actions

type Promo = struct {
	Id int `json:"id"`
	Type int `json:"type"`
	IsAvailable bool `json:"is_available"`
	IsReuse bool `json:"is_reuse"`
	Label string `json:"label"`
	Description string `json:"description"`
	PicturePoster PictureType `json:"picture_poster"`
	PictureAction PictureType `json:"picture_action"`
	PictureWebSlider PictureType `json:"picture_web_slider"`
	PromotionSettings PromoSettingsType `json:"promotion_settings"`
	VisitsPeriod []VisitsPeriodType `json:"visits_period"`
	PointsForPromoCode PointsForPromoCodeType `json:"points_for_promo_code"`
	GroupsUsers GroupsUsersType `json:"groups_users"`
	DisplayInMobileApp bool `json:"display_in_mobile_app"`
	PointsForQrCode PointsForQrCodeType `json:"points_for_qr_code"`
	NumberOfVisits []NumberOfVisitsType `json:"number_of_visits"`
	Survey SurveyType `json:"survey"`
}

type PictureType = struct {
	Id int `json:"id"`
	Url string `json:"url"`
}

type PromoSettingsType = struct {
	Id int `json:"id"`
	FromDisplayTime string `json:"from_display_time"`
	ToDisplayTime string `json:"to_display_time"`
	DaysOfWeek []int `json:"days_of_week"`
	IsVisible bool `json:"is_visible"`
}

type VisitsPeriodType = struct {
	Id int `json:"id"`
	NumberOfVisits int `json:"number_of_visits"`
	NumberOfDays int `json:"number_of_days"`
	NumberOfPoints int `json:"number_of_points"`
}

type PointsForPromoCodeType = struct {
	Id int `json:"id"`
	PromoCode string `json:"promo_code"`
	NumberOfPoints int `json:"number_of_points"`
	IsTemporary bool `json:"is_temporary"`
	ActionPeriod int `json:"action_period"`
}

type GroupsUsersType = struct {
	Id int `json:"id"`
	MinPoints int `json:"min_points"`
	MaxPoints int `json:"max_points"`
	MinAge int `json:"min_age"`
	MaxAge int `json:"max_age"`
	IsMan bool `json:"is_man"`
}

type PointsForQrCodeType = struct {
	Id int `json:"id"`
	NumberOfPoints int `json:"number_of_points"`
	IsTemporary bool `json:"is_temporary"`
	ActionPeriod int `json:"action_period"`
}

type NumberOfVisitsType = struct {
	Id int `json:"id"`
	CountVisits int `json:"count_visits"`
	CountsPoints int `json:"counts_points"`
}

type SurveyType = struct {
	Id int `json:"id"`
	PointsForSurvey []struct{
		Id int `json:"id"`
		Question string `json:"question"`
		Picture PictureType `json:"picture"`
		Type int `json:"type"`
		ListOfAnswers []struct{
			Id int `json:"id"`
			Answer string `json:"answer"`
		} `json:"list_of_answers"`
	} `json:"points_for_survey"`
	NumberOfPoints int `json:"number_of_points"`
}
