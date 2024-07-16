package testtask

type Users struct {
	Serie      int    `json:"passport_serie" db:"passport_serie"`
	Number     int    `json:"passport_number" db:"passport_number"`
	Surname    string `json:"surname" db:"surname"`
	Name       string `json:"name" db:"name" binding:"required"`
	Patronymic string `json:"patronymic,omitempty" db:"patronymic"`
	Address    string `json:"address" db:"address" binding:"required"`
}
