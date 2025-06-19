package domain

type Attribuition struct {
	Id         int64  `json:"_id"`
	Name       string `json:"name"`
	FileName   string `json:"filename"`
	Type       string `json:"type"`
	Author     string `json:"author"`
	Link       string `json:"link"`
	Licence    string `json:"licence"`
	LicenceUrl string `json:"licenceUrl"`
}

type Type struct {
	Id   int64  `json:"_id"`
	Name string `json:"name"`
}

type Licence struct {
	Id   int64  `json:"_id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type Query struct {
	Text  string `json:"text"`
	Order string `json:"order"`
}
