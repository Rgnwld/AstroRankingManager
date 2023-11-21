package AstroTypes

type UserTimeObj struct {
	Id                string `json:"id" binding:"required"`
	UserId            string `json:"userId" binding:"required"`
	TimeInMiliSeconds int    `json:"timeInMiliSeconds" binding:"required"`
	MapId             int    `json:"mapId" binding:"required"`
}

type UserNameTimeObj struct {
	Id                string `json:"id" binding:"required"`
	Username          string `json:"username" binding:"required"`
	TimeInMiliSeconds int    `json:"timeInMiliSeconds" binding:"required"`
	MapId             int    `json:"mapId" binding:"required"`
}

/*
	NOTE: Remove TimeInSeconds && MapId from UserTimeObj and replace it with TimeObj struct

	type UserTimeObj struct {
	Id            string `json:"id"   binding:"required"`
	Username      string `json:"userId"   binding:"required"`
	Time (Or something like this; Need to figure out how to throw it into a json and bind it)
}
*/

// type UserTimeObjV2 struct {
// 	Id       string  `json:"id"   binding:"required"`
// 	UserId   string  `json:"userId"   binding:"required"`
// 	TimeInfo TimeObj `json:"timeInfo" binding:"required"`
// }

type TimeObj struct {
	TimeInMiliSeconds int `json:"timeInMiliSeconds" binding:"required"`
	MapId             int `json:"mapId" binding:"required"`
}
