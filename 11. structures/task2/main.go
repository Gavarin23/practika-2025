package main

import "fmt"

type Engine struct {
	Horsepower int
	Type       string
}

type Car struct {
	Make   string
	Model  string
	Year   int
	Engine Engine
}

func main() {
	car := Car{
		Make:  "BMW",
		Model: "X5",
		Year:  2018,
		Engine: Engine{
			Horsepower: 181,
			Type:       "Hybrid",
		},
	}

	fmt.Printf("Автомобиль: %s %s (%d)\n", car.Make, car.Model, car.Year)
	fmt.Printf("Двигатель: %d л.с., тип: %s\n", car.Engine.Horsepower, car.Engine.Type)
}
