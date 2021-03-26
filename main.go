package main

import "fmt"

const x string = "Hello world"

var y string = "It's a variable"
var num int

func test(numbers []float64) float64 {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Not Implemented")
}

func average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func getNewXY(x int, y int) (int, int) {
	x += 34
	y += 12
	return x, y
}

func mull(args ...int) int {
	total := 0
	for _, num := range args {
		total *= num
	}
	return total
}

func square(x *float64) {
	*x = *x * *x
}

func main() {
	fmt.Println(x)
	fmt.Println(y)
	var num int
	fmt.Scanf("%d", &num)
	fmt.Println("Your number is ", num)

	// for
	i := 1
	for i <= 10 {
		fmt.Print(i, " , ")
		i++
	}

	fmt.Println()

	// for inline
	for i := 1; i <= 10; i++ {
		// if
		if i%2 == 0 {
			fmt.Print("even ")
		} else {
			fmt.Print("odd ")
		}

		// switch
		switch i {
		case 1:
			fmt.Print(" - One ")
		case 2:
			fmt.Print(" - Two ")
		case 3:
			fmt.Print(" - Three ")
		case 4:
			fmt.Print(" - Four ")
		case 5:
			fmt.Print(" - Five ")
		default:
			fmt.Print(" - Number ")
		}
	}

	fmt.Println()

	// arrays
	var m [5]float64
	m[3] = 100.3
	fmt.Println(m)

	fmt.Println()

	for i := 0; i < len(m); i++ {
		fmt.Print(int(m[i]), " * ") // + convert types
	}

	// use range
	var total float64
	for _, value := range m {
		total += value
	}
	fmt.Println("Total: ", total)

	// short array declare
	n := [5]string{
		"Alex",
		"Olga",
		"Vika",
		"Vera",
		"Alexandra",
	}
	for _, val := range n {
		fmt.Println("-=- ", val, " -=-")
	}

	// slices - срезы - это часть массива
	var zeroSlice []string
	fmt.Println(zeroSlice)

	p := make([]int, 4)
	fmt.Println(p)

	girls := n[1:5]
	fmt.Println(girls)

	// append
	extendFamily := append(n[0:5], "Cat", "Dog")
	fmt.Println(extendFamily)

	// copy
	men := make([]string, 1)
	copy(men, n[0:5])
	fmt.Println(men)

	// MAP
	keys := make(map[string]int)
	keys["start"] = 10
	keys["end"] = 100
	fmt.Println(keys)

	keywords := map[string]string{
		"U": "Up",
		"D": "Down",
		"L": "Left",
		"R": "Right",
	}
	fmt.Println(keywords["U"])

	// multi map
	persons := map[string]map[string]string{
		"Alex": map[string]string{
			"fullname": "Alexandr Ivanov",
			"age":      "34",
		},
		"Oleg": map[string]string{
			"fullname": "Oleg Orlov",
			"age":      "54",
		},
	}
	fmt.Println(persons)

	if person, ok := persons["Alex"]; ok {
		fmt.Println(person)
	}

	// functions
	test([]float64{4, 4})

	fmt.Println("Average: ", average([]float64{4, 5, 6, 7, 8, 22}))

	x, y := getNewXY(5, 5)
	fmt.Println("New x, y : ", x, y)

	fmt.Println(mull(1, 2, 3, 4, 5))

	w := 4.0
	square(&w)
	fmt.Println("W = ", w)
}
