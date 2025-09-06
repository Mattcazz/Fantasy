package types

type Team struct {
	Name     string
	Code     string
	Logo_url string
}

type Player struct {
	Id                  string
	Name                string
	Points              string
	Price               float32
	Fluctuation         float32
	Fluctuation_History []float32
	Img_url             string
}
