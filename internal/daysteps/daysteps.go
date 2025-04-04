package daysteps

import (
	"fmt"
	"sprint5_tracker/internal/personaldata"
	"sprint5_tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

const (
	StepLength = 0.65
)

// DaySteps содержит данные о дневных прогулках.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
// Parse парсит строку с данными о прогулке.
func (ds *DaySteps) Parse(datastring string) error {
	// Разделяем строку на части
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("некорректный формат данных\n")
	}

	// Преобразуем количество шагов в int
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %v\n", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше 0\n")
	}
	ds.Steps = steps

	// Преобразуем длительность в time.Duration
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %v\n", err)
	}
	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	// Проверяем, что продолжительность больше 0
	if ds.Duration <= 0 {
		return "", fmt.Errorf("длительность прогулки должна быть больше 0")
	}

	// Вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps)

	// Вычисляем количество калорий
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Формируем строку результата
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
