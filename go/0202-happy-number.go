func isHappy(n int) bool {    
	total := 0
	alreadySeen := make(map[int]bool)
	for {
		if seen := alreadySeen[n]; seen {
			break
		}
		alreadySeen[n] = true
		strn := fmt.Sprint(n)
		length := len(strn)
		for i := 0; i < length; i++ {
			digit, _ := strconv.Atoi(string(strn[i]))
			total += (digit * digit)
		}
		if total == 1 {
			return true
		}
		n = total
		total = 0
   }   
return false
}
