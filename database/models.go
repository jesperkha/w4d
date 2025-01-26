package database

type Recipe struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredient"`
}

type Ingredient struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type RecipeIngredient struct {
	RecipeID     uint `gorm:"primaryKey"`
	IngredientID uint `gorm:"primaryKey"`
}
