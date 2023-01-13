package params

type (
	GetListMovieResponse struct {
		Id          int     `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
		Image       string  `json:"image"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
	}

	GetDetailMovieResponse struct {
		Id          int     `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Rating      float64 `json:"rating"`
		Image       string  `json:"image"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
	}
)
