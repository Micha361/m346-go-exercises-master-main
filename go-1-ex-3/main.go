package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	augen := rand.Intn(6) + 1
	zeitpunkt := time.Now()

	fmt.Fprintln(os.Stdout, "Der Würfel zeigt", augen, "Augen")
	fmt.Fprintln(os.Stderr, "Der Würfel wurde geworfen am", zeitpunkt)

	datei1, _ := os.Create("augen.txt") 
	defer datei1.Close()
	fmt.Fprintln(datei1, "Der Würfel zeigt", augen, "Augen")

	datei2, _ := os.Create("wurf.log")
	defer datei2.Close()
	fmt.Fprintln(datei2, "Der Würfel wurde geworfen am", zeitpunkt)
}
