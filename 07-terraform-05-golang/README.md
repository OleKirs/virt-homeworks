# Домашнее задание к занятию "7.5. Основы golang"

С `golang` в рамках курса, мы будем работать не много, поэтому можно использовать любой IDE. 
Но рекомендуем ознакомиться с [GoLand](https://www.jetbrains.com/ru-ru/go/).  

## Задача 1. Установите golang.
1. Воспользуйтесь инструкций с официального сайта: [https://golang.org/](https://golang.org/).
2. Так же для тестирования кода можно использовать песочницу: [https://play.golang.org/](https://play.golang.org/).

```commandline
root@deb11-test50:~/golang/3# go version
go version go1.15.15 linux/amd64
```

## Задача 2. Знакомство с gotour.
У Golang есть обучающая интерактивная консоль [https://tour.golang.org/](https://tour.golang.org/). 
Рекомендуется изучить максимальное количество примеров. В консоли уже написан необходимый код, 
осталось только с ним ознакомиться и поэкспериментировать как написано в инструкции в левой части экрана.  

## Задача 3. Написание кода. 
Цель этого задания закрепить знания о базовом синтаксисе языка. Можно использовать редактор кода 
на своем компьютере, либо использовать песочницу: [https://play.golang.org/](https://play.golang.org/).

1. Напишите программу для перевода метров в футы (1 фут = 0.3048 метр). Можно запросить исходные данные 
у пользователя, а можно статически задать в коде.
    Для взаимодействия с пользователем можно использовать функцию `Scanf`:

<details>
<summary>Пример кода</summary>
    ```
    package main
    
    import "fmt"
    
    func main() {
        fmt.Print("Enter a number: ")
        var input float64
        fmt.Scanf("%f", &input)
    
        output := input * 2
    
        fmt.Println(output)    
    }
    ```
 </details>

### Ответ: 
<details>
<summary>Текст программы</summary>

```commandline
// Convert size from SI `meters` to Imperial `feets`.
// Run programm, then input size in meters and that convert size to feets.
package main

import (
        "fmt"
        "os"
)

func convert_meters_to_feets(size_in_meters float32) (size_in_feets float32) {

        const meters_to_feet float32 = 0.3048 // describe how mach meters is in 1 feet

        size_in_feets = size_in_meters / meters_to_feet

        return size_in_feets
}

func main() {

        fmt.Println("Input sise (in meters):")

        var input_from_stdin float32
        fmt.Fscan(os.Stdin, &input_from_stdin)

//      Uncomment to debug
//      fmt.Println("Input is: %\n", input_from_stdin)

        fmt.Println("This size is: ", convert_meters_to_feets(input_from_stdin), "feet(s)")
}

```

 </details>


2. Напишите программу, которая найдет наименьший элемент в любом заданном списке, например:
    ```
    x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
    ```

### Ответ: 
<details>
<summary>Текст программы</summary>

```commandline
// Find minimal value from predefined slice with integer values (`arr` slice)
package main

import "fmt"

func main() {
	
	//set slice values
	arr := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	
	//Call func to find min value and output `min` values in stdout
	fmt.Println("Minimal value is:", Min(arr))    
}
func Min(arr []int) int {

	//set start value to 'min' variable
	min := arr[0]

	//Use `for-range` loop to compare values each to other
	for _, value := range arr {
		if value < min {      // if current value less that current `min` variable
			min = value       // then replace `min` on current value
		}
	}

	return min
}

```

 </details>


3. Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. То есть `(3, 6, 9, …)`.

### Ответ: 
<details>
<summary>Текст программы</summary>

> Можно итерировать по одному и пытаться делить "нацело", но тут решил воспользоваться тем, что целочисленное деление обратно умножению на целые числа и подобно итеративному сложению делителя самого с собой.  
> В первом варианте тело цикла будет выглядеть как-то так:
> ```commandline
> for i := 1; i <= 100; i ++ {           // Перебираем все i от 1 до 100 с инкрементом `+1`
>		if i % 3 == 0 {                // Если остаток от целочисленного деления `100`/`i` равен нулю, тогда
>		  devide_wo_remains = append(devide_wo_remains, i)  // дописать в слайс текущее значение `i`.
>		}
> ```

> Во втором так:

```commandline
// Find all integer values to limit value, that may divide on divider without remains
package main

import "fmt"

var set_divider int = 3
var set_limit int = 100

func FilterList() (devide_wo_remains []int) {
	for i := set_divider; i <= set_limit; i += set_divider {
		devide_wo_remains = append(devide_wo_remains, i)
	}
	return
}

func main() {
	//toPrint := FilterList()
	fmt.Printf("Numbers from 1 to `limit` that may divide on `divider` without a remains: \n")
	fmt.Printf("%v", FilterList())
}

```

 </details>


В виде решения ссылку на код или сам код. 

## Задача 4. Протестировать код (не обязательно).

Создайте тесты для функций из предыдущего задания. 

### Ответ:  
Тест 1  
```commandline
package main

import "testing"

func TestMain(t *testing.T) {
        var test_res float32
        test_res = convert_meters_to_feets(77)
        if test_res != 252.62466 {
                t.Error("Value must be `252.62466`, but here is:", test_res)
        }
}

```

Тест 2  
```commandline
package main

import "testing"

func TestMain(t *testing.T) {
        var test_res int
        test_res = Min([]int{17,78,87,3,18})
        if test_res != 3 {
                t.Error("Minimal value must be `3`, but here is:", test_res)
        }
}

```

Тест 3  
```commandline
package main

import "fmt"
import "testing"

func TestMain(t *testing.T) {
	var tr []int
	tr = FilterList()
	if tr[1] != 6 || tr[2] != 9 || tr[9] != 30 {
		s := fmt.Sprintf("Value must be `6, 9, 30`, but here is: %v, %v and %v", tr[1], tr[2], tr[9])
		t.Error(s)
	}
}

```

Тесты прохлодят проверку с результатом подобным этому:

```commandline
root@deb11-test50:~/golang/3# go test
PASS
ok      _/root/golang/3 0.002s

```

---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---

