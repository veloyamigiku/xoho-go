package json

type TheaterRes struct {
	Header     TheaterHeader       `json:"header"`
	Prefecture []TheaterPrefecture `json:"prefecture"`
}
