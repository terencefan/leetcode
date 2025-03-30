package p2115

type Recipe struct {
	name        string
	ingredients []string
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {

	var inToRecipe = make(map[string][]string)

	for i, recipe := range recipes {
		for _, ingredient := range ingredients[i] {
			inToRecipe[ingredient] = append(inToRecipe[ingredient], recipe)
		}
	}

	var inDegree = make(map[string]int)

	for i, recipe := range recipes {
		inDegree[recipe] = len(ingredients[i])
	}

	q := []string{}

	for _, supply := range supplies {
		if recipes, ok := inToRecipe[supply]; ok {
			for _, recipe := range recipes {
				inDegree[recipe]--
				if inDegree[recipe] == 0 {
					q = append(q, recipe)
				}
			}
		}
	}

	created := make([]string, 0)
	for len(q) > 0 {
		supply := q[0]
		created = append(created, supply)

		if recipes, ok := inToRecipe[supply]; ok {
			for _, recipe := range recipes {
				inDegree[recipe]--
				if inDegree[recipe] == 0 {
					q = append(q, recipe)
				}
			}
		}
		q = q[1:]
	}
	return created
}
