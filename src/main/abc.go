package main

import (
	"fmt"
	"time"
	"math/rand"
	//"html/template"
	"log"
	"runtime"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	fmt.Println("1111111111111111111")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(12))

	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])

	//t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")

	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println("XXXXXXXXXXXXXXXXX=> ", runtime.GOOS)

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	//struct
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)


	// Array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices Array
	primesSlices := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primesSlices[1:4]
	fmt.Println(s) // 3 5 7

	log.Fatal("ffffffffffffffffffffffffffffffffffffffffffff")

	//defer trì hoãn  ....
}
