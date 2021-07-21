package json

type TheaterPrefecture struct {
	Name    string    `json:"name"`
	Sub     string    `json:"sub"`
	Theater []Theater `json:"theater"`
}
