package main

import "fmt"

func main() {
	mainMenu()
}

func mainMenu() {
	bookmark := make(map[string]string)

	for {
		fmt.Println("1.Посмотреть закладки")
		fmt.Println("2.Добавить закладки")
		fmt.Println("3.Удалить закладки")
		fmt.Println("4.Выход")

		operation := 0
		operation = getOperation(operation)

		if operation == 4 {
			break
		} else {
			switch operation {
			case 0:

			case 1:
				bookmarksLookup(bookmark)
			case 2:
				bookmark = addBookmark(bookmark)
			case 3:
				bookmark = removeBookmark(bookmark)

			}
		}
	}
}

func getOperation(operation int) int {
	_, err := fmt.Scanln(&operation)
	if err != nil {
		return 0
	}
	return operation
}

func bookmarksLookup(bookmark map[string]string) {
	for key, value := range bookmark {
		fmt.Printf("\n%s: %s\n", key, value)
	}
}

func addBookmark(bookmark map[string]string) map[string]string {
	var key, value string
	fmt.Println("Введите название ссылки:")
	_, err := fmt.Scanln(&key)
	if err != nil {
		return nil
	}
	fmt.Println("Введите ссылку:")

	_, err = fmt.Scanln(&value)
	if err != nil {
		return nil
	}

	bookmark[key] = value
	return bookmark
}

func removeBookmark(bookmark map[string]string) map[string]string {
	var key string
	fmt.Println("Введите название ссылки для удаления:")
	_, err := fmt.Scanln(&key)
	if err != nil {
		return nil
	}
	delete(bookmark, key)
	return bookmark
}
