package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	data := []string{
		"there is 3pm, but im still alive to write this code snippet",
		"чистый код лучше, чем заумный код",
		"ваш код станет наследием будущих программистов",
		"задумайтесь об этом",
	}
	// Создаем буфер
	buffer := bytes.NewBuffer([]byte{})

	// Записываем данные в буфер
	for _, s := range data {
		buffer.WriteString(s + "\n")
	}

	// Создаем файл и записываем данные в него
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Читаем данные из файла в новый буфер
	newBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(newBuffer, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Выводим данные из нового буфера
	fmt.Println(newBuffer.String())
}
