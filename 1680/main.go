package main

func main() {
	n := 12
	concatenatedBinary(n)
}

func concatenatedBinary(n int) int {
	res, shift, mod := 0, 0, 1e9+7
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			shift++
		}
		res = (res<<shift + i) % int(mod)
	}
	return res
	//string=>int
	//temp, _ := strconv.Atoi(sum)
	//
	//res, _ := strconv.ParseInt(strconv.Itoa(temp), 2, 64)
	//int%mod
	//int=>int64=>string
	//sum += strconv.FormatInt(int64(i), 2)
}
