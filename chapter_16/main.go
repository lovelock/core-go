package main

import "fmt"

// AnimalCategory 代表动物分类学中的基本分类法。
type AnimalCategory struct {
	kingdom string // 界。
	phylum string  // 门。
	class string  // 纲。
	order string  // 目。
	family string  // 科。
	genus string  // 属。
	species string // 种。
}

func (ac AnimalCategory) String() string {
	 return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}


type Animal struct {
	specialName string
	AnimalCategory
}

func (animal Animal) String() string {
	return fmt.Sprintf("%s", animal.specialName) + animal.AnimalCategory.String()
}

func main() {
	ac := new(AnimalCategory)
	ac.class = "fuck"
	ac.species = "suck"
	animal := Animal{
		specialName:"slate",
		AnimalCategory: *ac,
	}
	str1 := animal.String()
	fmt.Println(str1)
	//ac.String()
}
