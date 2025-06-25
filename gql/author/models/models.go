package models

type Author struct {
	ID    string  `json:"id"`
	Name1 string  `json:"name1"`
	Class string  `json:"class"`
	Bio   *string `json:"bio,omitempty"`
}

type NewAuthor struct {
	Name  string  `json:"name"`
	Class string  `json:"class"`
	Bio   *string `json:"bio,omitempty"`
}
