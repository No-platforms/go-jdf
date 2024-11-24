package go-jdf

import (
    "fmt"
)

// JalaliDate represents a date in Jalali calendar.
type JalaliDate struct {
    Year  int
    Month int
    Day   int
}

// GregorianDate represents a date in Gregorian calendar.
type GregorianDate struct {
    Year  int
    Month int
    Day   int
}

// ToGregorian converts a Jalali date to a Gregorian date.
func ToGregorian(jd JalaliDate) (GregorianDate, error) {
    jy := jd.Year + 1595
    days := -355668 + (365 * jy) + ((jy / 33) * 8) + ((jy % 33 + 3) / 4) + jd.Day + ((jd.Month < 7) * (jd.Month - 1) * 31) + ((jd.Month - 7) * 30) + 186

    gy := 400 * (days / 146097)
    days %= 146097

    if days > 36524 {
        gy += 100 * ((days - 1) / 36524)
        days %= 36524
        if days >= 365 {
            days++
        }
    }

    gy += 4 * (days / 1461)
    days %= 1461

    if days > 365 {
        gy += (days - 1) / 365
        days = (days - 1) % 365
    }

    gd := days + 1
    sal_a := []int{0, 31, getDaysInMonth(gy, 2), 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

    gm := 0
    for gm < len(sal_a) && gd > sal_a[gm] {
        gd -= sal_a[gm]
        gm++
    }

    return GregorianDate{Year: gy, Month: gm, Day: gd}, nil
}

// ToJalali converts a Gregorian date to a Jalali date.
func ToJalali(gd GregorianDate) (JalaliDate, error) {
    gYear := gd.Year
    gMonth := gd.Month
    gDay := gd.Day

    g_d_m := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

    gy2 := gYear
    if gMonth > 2 {
        gy2++
    }

    days := (355666 + (365 * gYear)) + (gy2/4 - gy2/100 + gy2/400) + gDay + g_d_m[gMonth-1]

    jy := -1595 + (33 * (days / 12053))
    days %= 12053

    jy += (4 * (days / 1461))
    days %= 1461

    if days > 365 {
        jy += (days - 1) / 365
        days = (days - 1) % 365
    }

    var jm int
    var jd int

    if days < 186 {
        jm = (days / 31) + 1
        jd = (days % 31) + 1
    } else {
        jm = ((days -186)/30) +7
        jd = ((days -186)%30)+1
    }

    return JalaliDate{Year: jy, Month: jm, Day: jd}, nil
}

// getDaysInMonth returns the number of days in a given month of a year.
func getDaysInMonth(year int, month int) int {
	if month == 2 {
		// Check for leap year in Gregorian calendar.
		if (year%4 ==0 && year%100 !=0) || year%400 ==0 {
			return 29 // Leap year February has an extra day.
		}
		return 28 // Regular February.
	}
	
	// For other months.
	monthsWith31Days := []int{1,3,5,7,8,10,12}
	for _, m := range monthsWith31Days {
		if month == m {
			return 31 
		}
	}
	return 30 // Remaining months have only up to thirty days.
}

// String formats the Jalali date as a string.
func (jd JalaliDate) String() string {
	return fmt.Sprintf("%04d/%02d/%02d", jd.Year, jd.Month, jd.Day)
}

// String formats the Gregorian date as a string.
func (gd GregorianDate) String() string {
	return fmt.Sprintf("%04d/%02d/%02d", gd.Year, gd.Month, gd.Day)
}

// jdate_words is a helper function for converting numbers to Persian words.
// This function can be expanded as needed based on requirements.
func jdateWords(array map[string]int) map[string]string {
	result := make(map[string]string)
	for key, num := range array {
		switch key {
		case "mm":
			months := []string{"فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور", "مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"}
			result[key] = months[num-1]
		case "rr":
			days := []string{"یک", "دو", "سه", "چهار", "پنج", "شش", "هفت", "هشت", "نه", "ده"}
			result[key] = days[num-1]
		default:
			result[key] = fmt.Sprintf("%d", num)
		}
	}
	return result
}