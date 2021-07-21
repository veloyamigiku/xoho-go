package json

type Theater struct {
	Name string `json:"name"`
	Sub  string `json:"sub"`
	Url  string `json:"url"`
	Type []Type `json:"type"`
}
