package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("---------- Калькулятор индекса массы тела ----------")
	for {
		userWeight, userHeight := getUserInput()
		IMT, err := calculateIMT(userWeight, userHeight)
		if err != nil {
			//fmt.Println("Не заданы параметры для расчета")
			//continue
			panic("Не заданы параметры для расчёта")
		}
		outputResult(IMT)
		isRepeat := repeatCalculation()
		if !isRepeat {
			break
		}
	}
}

func repeatCalculation() bool {
	var userAnswer string
	fmt.Println("Хотите рассчитать индекс еще раз (Да/Нет):")
	fmt.Scan(&userAnswer)
	if userAnswer == "Да" || userAnswer == "ДА" || userAnswer == "да" {
		return true
	}
	return false
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.1f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("Сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("Дефицит массы тела")
	case imt < 25:
		fmt.Println("Масса тела в норме")
	case imt < 30:
		fmt.Println("Избыточная масса тела")
	case imt < 35:
		fmt.Println("1-я степень ожирения")
	case imt < 40:
		fmt.Println("2-я степень ожирения")
	default:
		fmt.Println("3-я степень ожирения")
	}
}

func calculateIMT(userWeight float64, userHeight float64) (float64, error) {
	if userWeight <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_SET")
	}
	IMT := userWeight / math.Pow(userHeight/100, 2)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Println("Введите ваш рост в см:")
	fmt.Scan(&userHeight)
	fmt.Println("Введите ваш вес:")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}
