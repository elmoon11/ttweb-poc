package entities

type Tag struct {
	ID   uint64
	Name string `json:"name" binding:"required"`
}
