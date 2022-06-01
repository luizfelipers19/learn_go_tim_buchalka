package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

type Animal interface {
	Says() string
	HowManyLegs() int
}

type Canidae struct {
	Name         string
	Race         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

type Feline struct {
	Name         string
	Race         string
	Sound        string
	NumberOfLegs int
}

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

// receiver function
func (f *Feline) Says() string {
	return f.Sound
}

func (c *Canidae) Says() string {
	return c.Sound
}

func (f *Feline) HowManyLegs() int {
	return f.NumberOfLegs
}

func (c *Canidae) HowManyLegs() int {
	return c.NumberOfLegs
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf("I'm an animal that has %d legs, and my sound is %s. What am I?", a.HowManyLegs(), a.Says())
	fmt.Println(riddle)
}

var keyPressChan chan rune

type Vehicle struct {
	NumberOfPassengers int
	NumberOfSeats      int
	NumberOfWheels     int
}

type Car struct {
	Make       string
	Year       int
	Model      string
	IsHybrid   bool
	IsElectric bool
	Vehicle    Vehicle
}

func (v *Vehicle) ShowDetails() {
	fmt.Println("The number of passengers is", v.NumberOfPassengers)
	fmt.Println("The number of seats is", v.NumberOfSeats)
	fmt.Println("The number of wheels is", v.NumberOfWheels)
}

func (c *Car) Show() {
	fmt.Println("This car make is", c.Make)
	fmt.Println("This car model is", c.Model)
	fmt.Println("This car year is", c.Year)
	fmt.Println("It's", c.IsHybrid, "that the car is hybrid")
	fmt.Println("It's", c.IsElectric, "that the car is electric")
	c.Vehicle.ShowDetails()

}

func main() {

	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfSeats:      5,
		NumberOfPassengers: 5,
	}

	hilux := Car{
		Make:       "Toyota",
		Model:      "Hilux",
		Year:       2022,
		IsElectric: false,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	teslaX := Car{
		Make:       "Tesla",
		Model:      "X",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	hilux.Show()
	teslaX.Show()

	// Functions
	z := sumTwoInts(2, 3)
	fmt.Println(z)
	o := sumManyInts(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(o)

	var kiara Feline
	kiara.Name = "Kiara"
	kiara.Sound = "Me-ooow"
	kiara.Race = "Cat"
	kiara.NumberOfLegs = 4
	kiara.Says()

	ariel := Feline{
		Name:         "Ariel",
		Sound:        "Mewn",
		NumberOfLegs: 4,
	}
	ariel.Says()

	simba := Feline{
		Name:         "Simba",
		Sound:        "Roaarrr",
		NumberOfLegs: 4,
	}
	simba.Says()

	// ----------------
	// Channels
	keyPressChan = make(chan rune)
	go listenForKeyPress()
	fmt.Println("Press any key, or Q to quit.")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}
		keyPressChan <- char
	}

	sagu := Canidae{
		Name:         "Sagu",
		Race:         "dog",
		NumberOfLegs: 4,
		HasTail:      true,
		Sound:        "Auauvagabundoauau"}

	Riddle(&sagu)
	Riddle(&kiara)

	//luiz := Employee{
	//	Name:     "Luiz Felipe",
	//	Age:      26,
	//	Salary:   100000,
	//	FullTime: true,
	//}
	//
	//guizito := Employee{
	//	Name:     "Guilherme",
	//	Age:      28,
	//	Salary:   60000,
	//	FullTime: true,
	//}

	//	var employees []Employee
	//	employees = append(employees, luiz)
	//	employees = append(employees, guizito)
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}

func sumTwoInts(x int, y int) int {
	return x + y
}

func sumManyInts(nums ...int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	return total
}
