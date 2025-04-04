package actioninfo

import (
	"fmt"
)

// Интерфейс DataParser
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Функция Info обрабатывает данные и выводит информацию об активности
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Парсим строку с данными
		err := dp.Parse(data)
		if err != nil {
			fmt.Printf("Ошибка парсинга данных: %v\n\n", err)
			continue // Пропускаем некорректные данные
		}

		// Получаем информацию о действии
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка получения информации: %v\n\n", err)
			continue // Пропускаем ошибки при получении информации
		}

		// Выводим информацию
		fmt.Println(info)
	}
}
