package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Distance рассчитывает дистанцию в километрах
func Distance(steps int, height float64) float64 {

	if steps < 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient

	distanceInMeters := float64(steps) * stepLength

	return distanceInMeters / mInKm
}

// MeanSpeed рассчитывает среднюю скорость
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	durationInHours := duration.Hours()

	return distance / durationInHours
}

// RunningSpentCalories рассчитывает калории при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect parameters: steps=%d, weight=%.2f, height=%.2f, duration=%v", steps, weight, height, duration)
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

// WalkingSpentCalories рассчитывает калории при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps < 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect parameters: steps=%d, weight=%.2f, height=%.2f, duration=%v", steps, weight, height, duration)
	}

	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return calories * walkingCaloriesCoefficient, nil
}
