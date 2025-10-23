package main

import (
	"os"
	"strconv"
)

var file string = "counter.txt"

func main() {
	//Inc1()
	// readCounter()
	// writeCounter(2)
	Reset()
	Inc1()
	Inc1()
	Add(5)
	Dec1()
	Value()
}
func readCounter() (int, error) {
	val, _ := os.ReadFile(file)

	return strconv.Atoi(string(val))
}

func writeCounter(v int) error {

	return os.WriteFile(file, []byte(strconv.Itoa(v)), 0644)
}
func Inc1() error {
	val, _ := os.ReadFile(file)

	num, _ := strconv.Atoi(string(val))

	num++

	return os.WriteFile(file, []byte(strconv.Itoa(num)), 0644)
} // +1 и запись в файл
func Dec1() error {
	val, _ := os.ReadFile(file)

	num, _ := strconv.Atoi(string(val))

	num--

	return os.WriteFile(file, []byte(strconv.Itoa(num)), 0644)
} // -1 и запись в файл
func Add(n int) error {
	val, _ := os.ReadFile(file)

	num, _ := strconv.Atoi(string(val))

	num += n

	return os.WriteFile(file, []byte(strconv.Itoa(num)), 0644)
} // +n и запись в файл
func Reset() error {
	val, _ := os.ReadFile(file)

	num, _ := strconv.Atoi(string(val))

	num = 0

	return os.WriteFile(file, []byte(strconv.Itoa(num)), 0644)
} // 0 и запись в файл
func Value() (int, error) {
	val, _ := os.ReadFile(file)

	return strconv.Atoi(string(val))
} // просто вернуть текущее значение (читает из файла)
