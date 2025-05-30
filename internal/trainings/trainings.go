package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training хранит данные о тренировке
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

// Parse парсит строку с данными тренировки
func (t *Training) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("error: data format: 3 parts expected, received %d", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("error: parsing the number of steps: %w", err)
	}
	t.Steps = steps

	if steps <= 0 {
		return fmt.Errorf("error: incorrect number of steps\n")
	}

	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("error: parsing duration: %w", err)
	}
	t.Duration = duration

	if duration <= 0 {
		return fmt.Errorf("error: incorrect step format\n")
	}

	return nil
}

// ActionInfo формирует информацию о тренировке
func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)

	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("error: count calories: %w", err)
	}

	result := fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, t.TrainingType, t.Duration.Hours(), distance, speed, calories)

	return result, nil
}
