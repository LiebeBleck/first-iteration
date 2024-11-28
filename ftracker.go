package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию (в километрах), которую преодолел пользователь за время тренировки.
func distance(action int) float64 {
	return float64(action) * lenStep / float64(mInKm)
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	dist := distance(action)
	return dist / duration
}

// Константы для расчета калорий при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18
	runningCaloriesMeanSpeedShift      = 1.79
)

// RunningSpentCalories возвращает количество потраченных калорий при беге.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	speed := meanSpeed(action, duration)
	return ((runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift) * weight / float64(mInKm)) * duration * float64(minInH)
}

// Константы для расчета калорий при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035
	walkingSpeedHeightMultiplier    = 0.029
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	speed := meanSpeed(action, duration) * kmhInMsec
	return ((walkingCaloriesWeightMultiplier * weight) +
		(math.Pow(speed, 2)/(height/float64(cmInM)))*walkingSpeedHeightMultiplier*weight) * duration * float64(minInH)
}

// Константы для расчета калорий при плавании.
const (
	swimmingCaloriesMeanSpeedShift   = 1.1
	swimmingCaloriesWeightMultiplier = 2
)

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	// Дистанция в километрах
	distance := float64(lengthPool*countPool) / float64(mInKm)
	// Средняя скорость = дистанция / время
	return distance / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	// Средняя скорость
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	// Расход калорий = (Средняя скорость + поправка) * множитель веса * вес * длительность
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	var dist, speed, calories float64

	switch trainingType {
	case "Бег":
		dist = distance(action)
		speed = meanSpeed(action, duration)
		calories = RunningSpentCalories(action, weight, duration)
	case "Ходьба":
		dist = distance(action)
		speed = meanSpeed(action, duration)
		calories = WalkingSpentCalories(action, duration, weight, height)
	case "Плавание":
		dist = distance(action)
		speed = swimmingMeanSpeed(lengthPool, countPool, duration)
		calories = SwimmingSpentCalories(lengthPool, countPool, duration, weight)
	default:
		return "неизвестный тип тренировки"
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		trainingType, duration, dist, speed, calories)
}
