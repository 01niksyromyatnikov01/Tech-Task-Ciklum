package main

import (
	"fmt"
	"time"
)



/* Settings part  */
const precision = "2"    // Precision for the money left output
/* End Settings  */


/* Input values structure */
type structValues struct {
	OldCarPrice int           // Old car price
	NewCarPrice int           // New/secondhand car price
	SavingPerMonth int          // Man can save during the one month
	PercentLossByMonth float64  // Percent of loss increasing
}


/* Result values structure */
type structResult struct {
	months int   // Months required for purchase
	left float64 // Money left after purchase
}


func main() {

	input := getInput()

	results := Calculate(input)       // structValues{2000, 8000, 1000, 1.5}

	fmt.Printf("\nMoney left after purchase : $%."+precision+"f \nMonths required for purchase: %d \n", results.left, results.months)


	time.Sleep(600 * time.Second) // windows console window pause
}


func getInput() *structValues {

	input := new(structValues)

	fmt.Println("Enter old car price(integer), new car price(integer), saving for one month(integer) and percent of loss increasing(float).\nExample: 2000 8000 1000 1.5\n")

	fmt.Scan(&input.OldCarPrice,&input.NewCarPrice,&input.SavingPerMonth,&input.PercentLossByMonth)


	return input
}



func Calculate(input *structValues)  *structResult {
	OldCarPrice := input.OldCarPrice
	NewCarPrice := input.NewCarPrice
	SavingPerMonth := input.SavingPerMonth
	PercentLossByMonth := input.PercentLossByMonth


	months := 0
	priceOld := float64(OldCarPrice)
	priceNew := float64(NewCarPrice)

	for ; priceOld + float64(months * SavingPerMonth) < priceNew; months++ {
		if months % 2 == 1  {     // check if it is the end of every two months
			PercentLossByMonth += 0.5
		}
		priceOld -= priceOld * PercentLossByMonth / 100.0
		priceNew -= priceNew * PercentLossByMonth / 100.0

	}

	result := new(structResult)    // new struct for the result storing

	result.months = months
	result.left = priceOld + float64(months * SavingPerMonth) - priceNew   //use int(priceOld + float64(months * SavingPerMonth) - priceNew + 0.5) instead if you need the nearest integer


	return result
}

