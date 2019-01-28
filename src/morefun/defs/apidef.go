package defs

type SalesCombo struct {
	Id int `json:"id"`
	ComboName string `json:"combo_name"`
	SiteNum int `json:"site_num"`
	PlatformNum int `json:"platform_num"`
	RoadFee int `json:"road_fee"`
	Status int `json:"status"`
	CreateTime int `json:"create_time"`
	DeleteTime int `json:"delete_time"`
}
