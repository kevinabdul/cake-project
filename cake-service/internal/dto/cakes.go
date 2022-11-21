package dto

// usage of intermediate struct is done to ensure we only decode the fields
// that we want and omit the others
type CakeInsert struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
}

// to handle partial update
type CakeUpdate struct {
	ID          int      `json:"id"`
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Rating      *float64 `json:"rating"`
	Image       *string  `json:"image"`
}
