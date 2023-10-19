package main

// // interface declaration
// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct{}
// type spanishBot struct{}

// // One function using an interface instead of two separate print function, since it's using the same logic.
// func printGreetings(b []bot) {
// 	for _, b := range b {
// 		fmt.Println(b.getGreeting())
// 	}
// }

// func (englishBot) getGreeting() string {
// 	return "Hi there!"
// }
// func (spanishBot) getGreeting() string {
// 	return "Hola!"
// }

// ASSIGNMENT #1
// type shape interface {
// 	getArea() float64
// }
// type triangle struct {
// 	height float64
// 	base   float64
// }
// type square struct {
// 	sideLength float64
// }

// func (t triangle) getArea() float64 {
// 	return 0.5 * t.base * t.height
// }

// func (s square) getArea() float64 {
// 	return math.Pow(s.sideLength, 2)
// }

// func printArea(s shape) {
// 	fmt.Println(s.getArea())
// }

func main() {
	// // eb := englishBot{}
	// // sb := spanishBot{}

	// // printGreetings([]bot{eb, sb})

	// res, err := http.Get("https://google.com")

	// fmt.Println(res, err)

	// io.Copy(os.Stdout, res.Body)

	//----------- SOLUTION FOR ASSIGNMENT #1 -----------------
	// t := triangle{
	// 	height: 3.4,
	// 	base:   5.4,
	// }

	// s := square{
	// 	sideLength: 4,
	// }

	// printArea(t)
	// printArea(s)

	// fmt.Println(os.Args[len(os.Args)-1])

	//----------- SOLUTION FOR ASSIGNMENT #2 -----------------
	// filename := os.Args[len(os.Args)-1]

	// file, err := os.Open(filename)

	// if err != nil {
	// 	fmt.Println("Error while opening the file", err)
	// }

	// io.Copy(os.Stdout, file)

}
