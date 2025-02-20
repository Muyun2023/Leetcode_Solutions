
func myAtoi(s string) int {
	sign := 1
	sum := 0
	plusFound := false
	NumIntParsed := 0
	for _, v := range s {
		if s == "" {
			return 0
		}
		if v >= '0' && v <= '9' {
			NumIntParsed++
			n, err := strconv.ParseInt(string(v), 10, 32)
			if err == nil {
				sum = sum*10 + int(n)
			}
			if sum > math.MaxInt32 {
				if sign < 0 {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
		} else if sum >= 0 && NumIntParsed > 0 {
			return sign * sum
		} else {
				}
				continue
			} else {
				break
			}
		}
	}
	return sign * sum
}
