package models

type Author struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Class string  `json:"class"`
	Bio   *string `json:"bio,omitempty"`
}

type NewAuthor struct {
	Name  string  `json:"name"`
	Class string  `json:"class"`
	Bio   *string `json:"bio,omitempty"`
}
