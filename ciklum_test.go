package main

import "testing"


type testPair struct {
	values structValues
	results structResult
}


/*  Input and output values for testing  */
var tests = []testPair{
	{structValues{2000, 8000, 1000, 1.5}, structResult{6, 766.15812082500087853987}},
	{structValues{12000, 8000, 1000, 1.5}, structResult{0, 4000}},
	{structValues{8000, 12000, 500, 1.0}, structResult{8, 597.47939434833278937731}},
	{structValues{18000, 32000, 1500, 1.25}, structResult{8, 332.06546169545981683768}},
	{structValues{7500, 32000, 300, 1.55}, structResult{25, 121.72311533138963568490}},
}




/* test function for Calculate() */
func TestCalculate(t *testing.T) {
	for _,pair := range tests {

		res := Calculate(&pair.values)

		if pair.results != *res {
			t.Error(
				"For", pair.values,
				"expected", pair.results,
				"got", res,
			)
		}
	}
}
