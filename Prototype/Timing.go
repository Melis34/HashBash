package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Label struct {
	Text string
	X    float64
	Y    float64
}

func measureExecutionTime(f func()) time.Duration {
	startTime := time.Now()

	// Call the provided function
	f()

	// Calculate the elapsed time
	elapsedTime := time.Since(startTime)
	return elapsedTime
}

func main() {
	currenttime := time.Now()
	fmt.Println("current time: %d", currenttime)
	timeString := currenttime.Format("2006-01-02 15-04")
	fmt.Println(timeString)

	n := 2 //hoeveel bytes worden gecheckt
	tofile(n, "ownmethod", timeString)
	tofile(n, "traditional", timeString)
	fmt.Println("amount of bytes checked:", n)

	var valuesown plotter.Values
	var valuestrad plotter.Values

	for i := 0; i < 10000; i++ {
		startinput := hash(generateRandomBytes(10)) //Start
		durationtrad := measureExecutionTime(func() { traditional(startinput, n) })
		durationown := measureExecutionTime(func() { ownmethod(startinput, n) })
		fmt.Println("Tradtitional time", durationtrad)
		fmt.Println("Own time", durationown)

		ownmethodstring := fmt.Sprintf("%x , %d\n", startinput, durationown.Microseconds())
		tradmethodstring := fmt.Sprintf("%x , %d\n", startinput, durationtrad.Microseconds())

		ownfilename := fmt.Sprintf("%s%d%s", "ownmethod", n, timeString)
		tradfilename := fmt.Sprintf("%s%d%s", "traditional", n, timeString)
		valuesown = append(valuesown, float64(durationown.Microseconds()))
		valuestrad = append(valuestrad, float64(durationtrad.Microseconds()))
		writeLineToFile(ownfilename, ownmethodstring)
		writeLineToFile(tradfilename, tradmethodstring)

	}

	labelsown := []Label{
		{Text: fmt.Sprintf("Total Entries: %d", len(valuesown)), X: -0.5, Y: 10},
		{Text: timeString, X: 0, Y: 5},
		{Text: timeString, X: 0, Y: 0},
	}
	labelstrad := []Label{
		{Text: fmt.Sprintf("Total Entries: %d", len(valuestrad)), X: -0.5, Y: 10},
		{Text: timeString, X: 0, Y: 5},
		{Text: "Below Label 2", X: 0, Y: 0},
	}

	titleown := "own method"
	titletrad := "traditional method"
	titleown = titleown + strconv.Itoa(n) + timeString
	titletrad = titletrad + " " + strconv.Itoa(n) + " " + timeString
	histPlot(valuesown, labelsown, titleown)
	histPlot(valuestrad, labelstrad, titletrad)

}

func generateRandomBytes(length int) []byte {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Handle the error
	}
	return randomBytes
}

func traditional(input []byte, n int) { //takes a 0 increments it till it finds collision
	output := hash(input)
	binaryString := []byte{
		'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0',
		'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0',
	} //32 bytes = 256 bits
	collision := false
	for collision == false {
		checkoutput := hash(binaryString)
		if areFirstNBytesEqual(checkoutput, output, n) {
			// fmt.Println("traditional input")
			// fmt.Println("input ", input)       //original input
			// fmt.Println("first", output)       // original hash
			// fmt.Println("check", binaryString) //input that leads to collision
			// fmt.Println("found ", checkoutput) //output of said input
			collision = true
		}
		binaryString = incrementBinary(binaryString)
	}
	return
}

func incrementBinary(binaryStr []byte) []byte {
	carry := 1

	for i := len(binaryStr) - 1; i >= 0 && carry > 0; i-- {
		bit := int(binaryStr[i] - '0')
		sum := bit + carry
		binaryStr[i] = byte(sum%2 + '0')
		carry = sum / 2
	}

	return binaryStr
}

func areFirstNBytesEqual(arr1, arr2 []byte, n int) bool {
	if len(arr1) < n || len(arr2) < n {
		return false
	}

	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func ownmethod(input []byte, n int) {
	output := hash(input)
	// prevhash := output
	collision := false
	nexthash := hash(output)
	for collision == false {
		if areFirstNBytesEqual(nexthash, output, n) {
			// fmt.Println("own method")
			// fmt.Println("input ", input)    //original input
			// fmt.Println("first", output)    // original hash
			// fmt.Println("check", prevhash)  //input that leads to collision
			// fmt.Println("found ", nexthash) //output of said input
			collision = true
		} else {
			// prevhash = nexthash
			nexthash = hash(nexthash)
		}
	}
	return
}

func hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func tofile(n int, method string, datetime string) {
	title := "Results of " + strconv.Itoa(n) + " bytes, method: " + method
	filename := method + strconv.Itoa(n) + datetime

	// Create or open the file for writing. If it exists, truncate it; if not, create a new file.
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Close the file when the function exits

	// Write the title and initial lines
	fmt.Fprintln(file, title)

	fmt.Printf("Content written to %s successfully.\n", filename)
}

func writeLineToFile(file string, line string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(line); err != nil {
		panic(err)
	}

}

func histPlot(values plotter.Values, labels []Label, title string) {
	fmt.Printf("title recieved:" + title)

	p := plot.New()

	p.Title.Text = title
	hist, err := plotter.NewHist(values, 20)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	// Set X axis limits
	p.X.Min = p.X.Min - 1.0
	p.X.Max = p.X.Max + 1.0

	// Add labels
	for _, label := range labels {
		points := make(plotter.XYs, 1)
		// Adjust X coordinate relative to the minimum X value of the histogram
		points[0].X = p.X.Min + label.X
		points[0].Y = label.Y

		labelPlot, err := plotter.NewLabels(plotter.XYLabels{
			XYs:    points,
			Labels: []string{label.Text},
		})
		if err != nil {
			panic(err)
		}
		p.Add(labelPlot)
	}

	title = title + ".png"
	if err := p.Save(6*vg.Inch, 3*vg.Inch, title); err != nil {
		panic(err)
	}
}
