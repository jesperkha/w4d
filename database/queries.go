package database

func (db *Database) GetAllRecipes() []Recipe {
	var recipes []Recipe
	db.db.Find(&recipes)
	return recipes
}
