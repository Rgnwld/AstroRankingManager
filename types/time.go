package AstroTypes

type TimeObj struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	TimeInSeconds int    `json:"timeInSeconds"`
	Map           int    `json:"map"`
}
