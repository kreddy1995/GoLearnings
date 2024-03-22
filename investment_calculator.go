/* sample investement program
package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years = 10

	fmt.Print("Enter investmentAmount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), float64(years))

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
*/
// build a profit calculator
// user input : revenue,expenses & tax rate
// calculate earnings before tax and eranings after tax
// calculate ratio(ebt/profit)
// output ebt,profit and the ratio

package main

import "fmt"

func main() {
	var revenue float64 = 0
	outputText("Enter Revenue: ")
	//fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	var expenses float64 = 0
	outputText("Enter expenses: ")
	//fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	var taxRate float64 = 0
	fmt.Print("Enter TaxRate: ")
	fmt.Scan(&taxRate)

	var ebt = revenue - expenses
	var taxAmount = calculateEbt(revenue, taxRate)
	var profit, ratio float64 = calculateProfitAndRatio(ebt, taxAmount)

	pritnOutput(ebt, profit, ratio)

}

func outputText(text string) {
	fmt.Println(text)
}

func calculateEbt(revenue, taxRate float64) float64 {
	return revenue * taxRate / 100
}

func calculateProfitAndRatio(ebt, taxAmount float64) (float64, float64) {
	profit := ebt - taxAmount
	ratio := ebt / profit

	return profit, ratio
}

func pritnOutput(ebt, profit, ratio float64) {
	fmt.Println("Ebt: ", ebt)
	fmt.Println("profit: ", profit)
	fmt.Println("ratio: ", ratio)
}
