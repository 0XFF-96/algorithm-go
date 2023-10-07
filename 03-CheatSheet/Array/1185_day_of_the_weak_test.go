package array


// Sum up the number of days for the years before the given year.
// Handle the case of a leap year.
// Find the number of days for each month of the given year.
// % 7 的方法
//

func dayOfTheWeek(day int, month int, year int) string {
	t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	if month < 3 {
		year -= 1
	}
	date := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return date[(year+year/4+year/400-year/100+t[month-1]+day)%7]
}

