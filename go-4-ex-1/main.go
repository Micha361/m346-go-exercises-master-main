package main

func computeGrade(gotPoints, maxPoints float64) float64 {
	if maxPoints <= 0 {
		panic("maxPoints must be greater than zero")
	}
	note := 1.0 + 5.0*(gotPoints/maxPoints)
	if note > 6.0 {
		return 6.0
	} else if note < 1.0 {
		return 1.0
	}
	return note
}

func main() {
	Println(computeGrade(17.5, 28.0)) 
	Println(computeGrade(25.0, 30.0)) 
	Println(computeGrade(10.0, 40.0)) 
}
