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

	Milliliters float32
	Grams       float32

	// All of these are per unit
	Calories      uint
	Carbohydrates float32
	Sugars        float32
	Protein       float32
	Fats          float32
	Fiber         float32
	Sodium        float32
}

type RecipeIngredient struct {
	RecipeID     uint `gorm:"primaryKey"`
	IngredientID uint `gorm:"primaryKey"`

	Count    uint
	Optional bool
}
