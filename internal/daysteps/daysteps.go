package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит данные о дневной активности
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse парсит строку с данными активности
func (ds *DaySteps) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("error: data format: 2 parts expected, received %d", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("error: parsing the number of steps: %w", err)
	}
	ds.Steps = steps

	if steps <= 0 {
		return fmt.Errorf("error: incorrect number of steps\n")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("error: parsing the duration: %w", err)
	}
	ds.Duration = duration

	if duration <= 0 {
		return fmt.Errorf("error: incorrect step format\n")
	}

	return nil
}

// ActionInfo формирует информацию о прогулке
func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("error: in counting calories: %w", err)
	}

	result := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n", ds.Steps, distance, calories)

	return result, nil
}
