package personaldata

import (
	"fmt"
)

// Структура Personal
type Personal struct {
	Name   string  // Имя пользователя
	Weight float64 // Вес пользователя в кг
	Height float64 // Рост пользователя в м
}

// Метод Print()
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f кг\n", p.Weight)
	fmt.Printf("Рост: %.2f м\n\n", p.Height)
}
