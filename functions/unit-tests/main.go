package main

func getMonthlyPrice(tier string) int {
	
	switch tier{
		case "basic":
			return 100.00 * 100
		case "premium":
			return 150.00 * 100
		case "enterprise":
			return 500.00 * 100
		
		default:
			return 0
	}
}
