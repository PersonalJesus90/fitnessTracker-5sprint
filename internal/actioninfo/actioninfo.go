package actioninfo

import (
	"fmt"
	"log"
)

// DataParser определяет интерфейс для работы с данными активности
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info выводит информацию о тренировках или прогулках
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("error: parsing data: %v", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error: generating information: %v", err)
			continue
		}

		fmt.Println(info)
	}
}
