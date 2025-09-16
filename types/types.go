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

type APIResponse struct {
	Teams []APIResponseTeam `json:"teams"`
}

type APIResponsePlayer struct {
	Name        string `json:"name"`
	Position    string `json:"position"`
	Nationality string `json:"nationality"`
}

type APIResponseTeam struct {
	Name      string              `json:"name"`
	Crest_URL string              `json:"crest"`
	Squad     []APIResponsePlayer `json:"squad"`
}
