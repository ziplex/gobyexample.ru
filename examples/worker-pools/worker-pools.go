// В этом примере реализован _набор обработчиков
// (worker pool)_ с помощью горутин и каналов.

package main

import "fmt"
import "time"

// С помощью этого обработчика запустим несколько
// одновременых экземпляров. Эти обработчки будут получать
// задачи из канала `jobs` и отправлять соответствующие
// результаты в `results`. Далее сделаем паузу в секунду на
// каждую задачу для имитации обработки.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	// Для использования набора обработчиков нужно отправлять
	// ему задачи и собирать результаты. Сделаем для этого
	// 2 канала.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Здесь запускается 3 обработчика, первоначально
	// заблокированные, поскольку пока нет задач.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Отправляем 9 `jobs` и закрываем канал, чтобы
	// показать, что больше задач нет.
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// В конце собираем все результаты обработки.
	for a := 1; a <= 9; a++ {
		<-results
	}
}
