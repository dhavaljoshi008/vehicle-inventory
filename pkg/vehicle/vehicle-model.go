package vehicle

type Vehicle struct {
	ID    int    `json:"id"`
	Make  string `json:"make" binding:"required"`
	Model string `json:"model" binding:"required"`
	Year  int    `json:"year" binding:"required"`
}
