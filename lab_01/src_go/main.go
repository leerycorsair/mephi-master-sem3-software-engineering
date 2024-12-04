package main

import (
	"fmt"
	"os"
	"resourceRegistrator/manager"
	"resourceRegistrator/model"
)

func MenuPrint() {
	fmt.Println("Меню:")
	fmt.Println("\t1) Создать ресурсы")
	fmt.Println("\t2) Загрузить значения ресурсов из файла")
	fmt.Println("\t3) Добавить запрос")
	fmt.Println("\t4) Освободить ресурс")
	fmt.Println("\t5) Такт")
	fmt.Println("\t6) Информация о ресурсах")
	fmt.Println("\t7) Информация о запросах")
	fmt.Println("\t0) Выход")
}

func MenuHandle(m *manager.Manager, option int) {
	switch option {
	case 1:
		var resCnt int
		fmt.Print("\tВведите число ресурсов:")
		_, err := fmt.Scan(&resCnt)
		if err != nil {
			fmt.Println("Некорректный ввод, не число.")
			return
		}
		m.InitModel(resCnt)
	case 2:
		var filePath string
		fmt.Print("\tВведите путь до файла:")
		_, err := fmt.Scan(&filePath)
		if err != nil {
			fmt.Println("Некорректный ввод.")
			return
		}
		err = m.InitFromFile(filePath)
		if err != nil {
			fmt.Println(err.Error())
		}
	case 3:
		if m.CheckModel() {
			fmt.Println("Ресурсы не созданы")
			return
		}
		var resId string
		fmt.Print("\tВведите название ресурса:")
		_, err := fmt.Scan(&resId)
		if err != nil {
			fmt.Println("Некорректный ввод.")
			return
		}
		var reqT int
		fmt.Print("\tВведите продолжительность запроса:")
		_, err = fmt.Scan(&reqT)
		if err != nil {
			fmt.Println("Некорректный ввод, не число.")
			return
		}
		if reqT <= 0 {
			fmt.Println("Некорректное время.")
			return
		}
		err = m.AddRequest(*manager.NewRequest(resId, reqT))
		if err != nil {
			fmt.Println(err.Error())
		}
	case 4:
		if m.CheckModel() {
			fmt.Println("Ресурсы не созданы")
			return
		}
		var resId string
		fmt.Print("\tВведите название ресурса:")
		_, err := fmt.Scan(&resId)
		if err != nil {
			fmt.Println("Некорректный ввод.")
			return
		}
		err = m.FreeResource(resId)
		if err != nil {
			fmt.Println(err.Error())
		}
	case 5:
		if m.CheckModel() {
			fmt.Println("Ресурсы не созданы")
			return
		}
		m.Work()
	case 6:
		if m.CheckModel() {
			fmt.Println("Ресурсы не созданы")
			return
		}
		fmt.Println(m.ModelInfo())
	case 7:
		if m.CheckModel() {
			fmt.Println("Ресурсы не созданы")
			return
		}
		fmt.Println(m.RequestsInfo())
	case 0:
		fmt.Println("Завершение работы...")
		os.Exit(0)
	default:
		fmt.Println("Некорректный ввод, неизвестная опция.")
	}
}
func main() {
	manager := manager.NewManager(model.NewEmptyModel())
	for {
		MenuPrint()
		var option int
		fmt.Print("Выберите опцию: ")
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Некорректный ввод, не число.")
			continue
		}
		MenuHandle(manager, option)
	}
}
