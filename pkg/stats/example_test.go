package stats

import (
	"github.com/Ulugbek999/bank5/pkg/bank/types"
	
	
	
	
	"fmt"
	
)






func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   53_00,
			Category: "store",
		},
		{
			ID:       2,
			Amount:   51_00,
			Category: "store",
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "store",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
} 


func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "pharmacy",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "restaurant",
		},
	}

	inCategory := types.Category("auto")
	totalInCategory := TotalInCategory(payments, inCategory)
	fmt.Println(totalInCategory)
	//Output: 1000000

} 