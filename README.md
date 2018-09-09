# Tech Task for DX Exchange
### Requirements: 
• solution should be written on Go; <br>
• code should be commented in English; <br>
• source data should be taken from stdin and results should go to stdout and be informative. <br>
### Task: 
A man has a rather old car being worth $2000. He saw a secondhand car being worth
$8000. He wants to keep his old car until he can buy the secondhand one.
He thinks he can save $1000 each month but the prices of his old car and of the new
one decrease of 1.5 percent per month. Furthermore the percent of loss increases by a
fixed 0.5 percent at the end of every two months.
Can you help him? Our man finds it difficult to make all these calculations.
How many months will it take him to save up enough money to buy the car he wants,
and how much money will he have left over?

## Comments
### Console input example:
<b>2000 8000 1000 1.5</b>
### Console output example:
<b> Money left after purchase : $766.16 <br>
Months required for purchase: 6</b>

### Important info
- <b>Unit test is available. You can run it with a command "go test".</b> 
- <b>The remaining money is not converted to an integer but you can change the precision value inside the settings block.</b> 
- <b>Calculations are processed at the end of each considered month but if the value of the old car is bigger than the value of the new one or equal there is no need to wait.</b>
