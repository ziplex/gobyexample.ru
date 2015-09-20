// Go имеет встроенную поддержку для форматирования строк
// сродни `printf`. Ниже нескольно общих примеров для
// выполнения задач форматирования.

package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {

	// Go предлагает несколько "глаголов" созданных
	// для форматирования общих Go значений. Например, это
	// выведет инстанс нашей `point` структуры.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// Если значение является структурой, вариант `%+v`
	// будет включать так же имена полей структуры.
	fmt.Printf("%+v\n", p)

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.

	// Вариант форматирования `%#v` выведет Go синтакс
	// представление значения, т.е. кусочек исходного кода
	// который произведет это значение.
	fmt.Printf("%#v\n", p)

	// Чтобы вывести тип значения, используйте `%T`.
	fmt.Printf("%T\n", p)

	// Форматирование булевых значений прямолинейно.
	fmt.Printf("%t\n", true)

	// Существует множество операций для форматирования
	// целочисленных значений. Используйте `%d` для
	// десятичной системы исчисления
	fmt.Printf("%d\n", 123)

	// Это выведет бинарное представление числа.
	fmt.Printf("%b\n", 14)

	// Вывод символа, соответсвующего заданному числу.
	fmt.Printf("%c\n", 33)

	// `%x` для шестнадцатеричного исчисления
	fmt.Printf("%x\n", 456)

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.

	// Так же имеется несколько опций для форматирования
	// чисел с плавающей запятой.
	fmt.Printf("%f\n", 78.9)

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.

	// `%e` и `%E` форматирует число с плавающей запятой
	// в виде (немого другой версии) научной нотации.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// Для простого вывода строк используйте `%s`.
	fmt.Printf("%s\n", "\"string\"")

	// Для двойных ковычек как в исходниках Go, используйте  `%q`
	fmt.Printf("%q\n", "\"string\"")

	// Так же как и с целочисленными ранее, `%x` отображает
	// строку в виде шестнадцатеричного исчисления, с
	// двумя символами вывода за каждый байт ввода.
	fmt.Printf("%x\n", "hex this")

	// Чтобы показать представление ссылки, используйте `%p`.
	fmt.Printf("%p\n", &p)

	// При форматировании чисел вы часто захотите
	// контролировать ширину и точность итоговой
	// цифры. Чтобы установить ширину числа, укажите
	// число после `%`. По умолчанию, результат будет
	// отображаться по правому краю с пробелами
	// слева.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// Вы так же можете указать ширифу вывода чисел
	// с плавающей точкой, хотя зачастую вы так же
	// хотели бы ограничить десятичную точность
	// с шириной используя синтакс ширина.точность
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// Для вывода по левой стороне, укажите флаг `-`.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Возможно вы бы хотели контролировать ширину
	// форматирования строк, например, для гарантии вывода
	// в виде таблицы.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// Для вывода по левой стороне, укажите флаг `-` с числами.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// До сих пор мы видели `Printf`, который выводит
	// отформатированные строки в `os.Stdout`. `Sprintf`
	// форматирует и возвращает строку без вывода ее
	// куда-либо.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Вы можете форматировать и выводить в `io.Writers`
	// вместо `os.Stdout` используя `Fprintf`
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
