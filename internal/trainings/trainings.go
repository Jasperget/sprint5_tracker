package trainings

import (
	"fmt"
	"sprint5_tracker/internal/personaldata"
	"sprint5_tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

// создайте структуру Training
type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

// Метод Parse() парсит строку с данными о тренировке
func (t *Training) Parse(datastring string) (err error) {
	// Разделяем строку на части
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("некорректный формат данных")
	}

	// Преобразуем количество шагов в int
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %v", err)
	}
	t.Steps = steps

	// Проверяем тип тренировки
	if parts[1] != "Бег" && parts[1] != "Ходьба" {
		return fmt.Errorf("неизвестный тип тренировки: %s", parts[1])
	}
	t.TrainingType = parts[1]

	// Преобразуем длительность в time.Duration
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %v", err)
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	// Проверяем, что продолжительность больше 0
	if t.Duration <= 0 {
		return "", fmt.Errorf("длительность тренировки должна быть больше 0")
	}

	// Вычисляем дистанцию
	distance := spentenergy.Distance(t.Steps)

	// Вычисляем среднюю скорость
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	// Рассчитываем калории в зависимости от типа тренировки
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	// Обрабатываем ошибку, если она возникла
	if err != nil {
		return "", err
	}

	// Формируем строку результата
	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpeed,
		calories,
	)

	return result, nil
}
