package database

func (db *Database) GetAllRecipes() []Recipe {
	var recipes []Recipe
	db.db.Find(&recipes)
	return recipes
}

func (db *Database) GetAllIngredients() []Ingredient {
	var ingredients []Ingredient
	db.db.Find(&ingredients)
	return ingredients
}

func (db *Database) GetIngredientById(id uint) Ingredient {
	var result Ingredient
	db.db.Find(&result, &Ingredient{ID: id})
	return result
}

func (db *Database) NewIngredient(ingredient Ingredient) {
	db.db.Create(&ingredient)
}

// Ingredients map is IngredientID -> Count.
func (db *Database) NewRecipe(name string, description string, ingredients map[uint]uint) {
	recipe := Recipe{
		Name:        name,
		Description: description,
	}

	for id, _ := range ingredients {
		i := db.GetIngredientById(id)
		recipe.Ingredients = append(recipe.Ingredients, i)
	}

	tx := db.db.Create(&recipe)
	if tx.Error != nil {
		return
	}

	for id, count := range ingredients {
		db.db.Create(&RecipeIngredient{
			RecipeID:     recipe.ID,
			IngredientID: id,
			Count:        count,
		})
	}
}
