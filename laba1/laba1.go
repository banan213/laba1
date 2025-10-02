package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	mali := "qwertyuiopasdfghjklzxcvbnm"
	veliki := "QWERTYUIOPASDFGHJKLZXCVBNM"
	spec := "!@#$%^&*()-_=+[]{}<>?/|"
	num := "1234567890"
	pas := ""

	var IFmali string
	var IFveliki string
	var IFspec string
	var IFnum string
	var pul string
	var vs int

	fmt.Print("Введіть довжину пароля: ")
	var lenStr string
	fmt.Scanln(&lenStr)
	lan, err := strconv.Atoi(lenStr)
	if err != nil {
		fmt.Println("Введено не число")
		return
	}

	if lan <= 4 || lan >= 128 {
		fmt.Println("некоректна довжина пароля")
		return
	}

	for {

		fmt.Print("Використовувати малі літери? (y/n):")
		fmt.Scanln(&IFmali)
		if IFmali == "y" {
			idx := rand.Intn(len(mali))
			s := string(mali[idx])
			pas = pas + s
			pul = pul + mali
			vs++
		}

		fmt.Print("Використовувати великі літери? (y/n):")
		fmt.Scanln(&IFveliki)
		if IFveliki == "y" {
			idx := rand.Intn(len(veliki))
			s := string(veliki[idx])
			pas = pas + s
			pul = pul + veliki
			vs++
		}

		fmt.Print("Використовувати спец. символи? (y/n):")
		fmt.Scanln(&IFspec)
		if IFspec == "y" {
			idx := rand.Intn(len(spec))
			s := string(spec[idx])
			pas = pas + s
			pul = pul + spec
			vs++
		}

		fmt.Print("Використовувати цифри? (y/n):")
		fmt.Scanln(&IFnum)
		if IFnum == "y" {
			idx := rand.Intn(len(num))
			s := string(num[idx])
			pas = pas + s
			pul = pul + num
			vs++
		}

		if vs == 0 {
			fmt.Println("ви не обрали жодну категорію")
		} else {
			break
		}
	}

	for i := 0; i < lan-vs; i++ {
		idx := rand.Intn(len(pul))
		s := string(pul[idx])
		pas = pas + s
	}
	fmt.Println("згенорований пароль:", pas)

	runes := []rune(pas)
	n := len(runes)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	pas = string(runes)
	fmt.Println("згенорований пароль по алгоритму Fisher--Yates:", pas)
}
