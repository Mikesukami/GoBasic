package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	for {
		fmt.Println("== This is Example to learn GO == ")
		fmt.Println("Enter 12. Exit")
		fmt.Println("Enter 13. Clear Screen")
		fmt.Print("Please chosse an exam [1-13] : ")
		var choice int
		fmt.Scan(&choice)
		fmt.Print("\n")

		switch choice {
		case 1:
			num0()
		case 2:
			num1()
		case 3:
			num1_2()
		case 4:
			num2()
		case 5:
			num3()
		case 6:
			num3_1()
		case 7:
			num4()
		case 8:
			num4_1()
		case 9:
			num5()
		case 10:
			num6()
		case 11:
			num7()
		case 12:
			fmt.Println("Exit")
			return
		case 13:
			clearScreen()
		default:
			fmt.Println("Invalid choice. please enter again.")
		}
	}
}

func num0() {
	fmt.Println("This is 1 Exam")

	i := 0

	fmt.Print("Enter num 0-3 : ")
	fmt.Scan(&i)

	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Your input is invalid. Please enter again.")
	}

	fmt.Print("\n")
}

// 2. ระหว่าง 1-100 มีเลขที่หาร3ลงตัวกี่ตัว อะไรบ้าง (for if)
func num1() {
	fmt.Println("This is 2 Exam")

	count := 0

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, " ")
			count++
		}
	}

	fmt.Println("\nCount : ", count)
	fmt.Print("\n")
}

// สร้างฟังชั่นคำนวณเลขยกกำลัง เช่น num(20,2)
func num1_2() {
	fmt.Println("This is 3 Exam")
	fmt.Print("Enter 2 number [num1 | num2] : ")
	var x, y float64
	fmt.Scan(&x, &y)
	fmt.Println("Result : ", CalPower(x, y))

	fmt.Print("\n")
}

func num2() {
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}

	fmt.Println("Max In X : ", FindMax(x))
	fmt.Println("Min In X : ", FindMin(x))

}

/*
3. 1-1000 มีเลข9รวมทั้งหมดกี่ตัว จงเขียนโปรแกรม(คำตอบข้อ3 300)
guide1: แปลง int to string
guide2: strconv.Itoa()
*/
func num3() {
	fmt.Println("This is 5 Exam")

	count := 0

	for i := 1; i <= 1000; i++ {
		s := strconv.Itoa(i)
		for _, v := range s {
			if v == '9' {
				count++
			}
		}
	}

	// for i := 1; i <= 1000; i++ {
	// 	s := fmt.Sprintf("%d", i)
	// 	for _, v := range s {
	// 		if v == '9' {
	// 			count++
	// 		}
	// 	}
	// }

	fmt.Println("Count : ", count)
	fmt.Print("\n")

}

/*
ใส่ค่าตัวเลขเข้าไปให้ฟังชั่น เพื่อหาเลข9ในจำนวนเลขที่ใส่เข้าไป

	เช่น someFunc(10000) หาเลข9ตั้งแต่1-10000
*/
func num3_1() {
	fmt.Println("This is 6 Exam")
	fmt.Print("Enter number: ")
	var num int
	fmt.Scan(&num)

	count := someFunc(num)
	fmt.Println("Count:", count)

}

func num4() {
	var myWords = "AW SOME GO!"
	var result string
	for i := 0; i < len(myWords); i++ {
		fmt.Println(result)
		if myWords[i] != ' ' {
			result += string(myWords[i])
		}
	}
	fmt.Println(result)
	fmt.Print("\n")
}

func num4_1() {
	fmt.Println("This is 7 Exam")
	fmt.Print("Enter text: ")
	text := "ine t"
	fmt.Println("Result : ", cutText(text))
	fmt.Print("\n")
}

func num5() {
	peoples := map[string]map[string]string{
		"emp1": {
			"fname": "John",
			"lname": "Doe",
		},
		"emp2": {
			"fname": "Jane",
			"lname": "Smith",
		},
		"emp3": {
			"fname": "Jack",
			"lname": "White",
		},
	}

	//print all like John Doe
	for _, v := range peoples {
		fmt.Println(v["fname"], v["lname"])
	}
}

type Company struct {
	Name  string
	Tel   string
	Email string
}

func num6() {
	c_1 := Company{
		Name:  "Mamimo Ltd.",
		Tel:   "123456",
		Email: "mamimo@gomail.com",
	}

	fmt.Printf("Name: %s\nTel: %s\nEmail: %s \n", c_1.Name, c_1.Tel, c_1.Email)
}

func num7() {
	for i := 1; i <= 6; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// func for Exam
func CalPower(x, y float64) float64 {
	return math.Pow(x, y)
}

func FindMax(x []int) int {
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

func FindMin(x []int) int {
	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return min

}

func someFunc(x int) int {
	count := 0

	for i := 1; i <= x; i++ {
		s := strconv.Itoa(i)
		for _, v := range s {
			if v == '9' {
				count++
			}
		}
	}

	return count

}

func cutText(x string) string {
	result := ""
	for i := 0; i < len(x); i++ {
		if x[i] != ' ' {
			result += string(x[i])
		}
	}
	return result
}

// Clear Screen
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
