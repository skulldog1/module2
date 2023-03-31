package main

import (
	"bufio"
	"os"
	"strings"
)

var transMap = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "yo",
	'ж': "zh", 'з': "z", 'и': "i", 'й': "y", 'к': "k", 'л': "l", 'м': "m",
	'н': "n", 'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u",
	'ф': "f", 'х': "kh", 'ц': "ts", 'ч': "ch", 'ш': "sh", 'щ': "shch",
	'ъ': "", 'ы': "y", 'ь': "", 'э': "e", 'ю': "yu", 'я': "ya",
	'А': "A", 'Б': "B", 'В': "V", 'Г': "G", 'Д': "D", 'Е': "E", 'Ё': "Yo",
	'Ж': "Zh", 'З': "Z", 'И': "I", 'Й': "Y", 'К': "K", 'Л': "L", 'М': "M",
	'Н': "N", 'О': "O", 'П': "P", 'Р': "R", 'С': "S", 'Т': "T", 'У': "U",
	'Ф': "F", 'Х': "Kh", 'Ц': "Ts", 'Ч': "Ch", 'Ш': "Sh", 'Щ': "Shch",
	'Ъ': "", 'Ы': "Y", 'Ь': "", 'Э': "E", 'Ю': "Yu", 'Я': "Ya",
}

func main() {
	inputFile := "example.txt"
	outputFile := "example.processed.txt"

	inFile, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)

	writer := bufio.NewWriter(outFile)

	for scanner.Scan() {
		line := scanner.Text()
		processedLine := transliterate(line)
		writer.WriteString(processedLine + "\n")
	}
	writer.Flush()

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func transliterate(s string) string {
	var sb strings.Builder

	for _, c := range s {
		if val, ok := transMap[c]; ok {
			sb.WriteString(val)
		} else {
			sb.WriteRune(c)
		}
	}

	return sb.String()
}
