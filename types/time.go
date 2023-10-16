package AstroTypes

type UserTimeObj struct {
	Id            string `json:"id"   binding:"required"`
	Username      string `json:"username"   binding:"required"`
	TimeInSeconds int    `json:"timeInSeconds"  binding:"required"`
	Map           int    `json:"map"   binding:"required"`
}

type TimeObj struct {
	TimeInSeconds int `json:"timeInSeconds"  binding:"required"`
}
