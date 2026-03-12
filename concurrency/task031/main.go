package main

// напишите функции writer, double, reader.
// writer - пишет значения от 1 до 100 в возвращаемый канал
// double - возвращает канал и читает из принимаемого канала значения и умножает их
// reader - выводит получаемые значения из канала
// По сути это реализация Pipeline

func main() {
	reader(double(writer()))
}

func writer() <-chan int {
	return nil
}

func double(in <-chan int) <-chan int {
	return nil
}

func reader(in <-chan int) {
}
