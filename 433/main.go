package main

func main() {
	start := "AACCGGTT"
	end := "AAACGGTA"
	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	minMutation(start, end, bank)
}

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	bankSet := map[string]bool{}
	for _, parameter := range bank {
		bankSet[parameter] = true
	}
	if _, ok := bankSet[end]; !ok {
		return -1
	}
	queue := []string{start}
	for step := 0; queue != nil; step++ {
		temp := queue
		queue = nil
		for _, str := range temp {
			for i, alphbet := range str {
				for _, set := range "ACGT" {
					if alphbet != set {
						tempStr := str[:i] + string(set) + str[i+1:]
						if _, ok := bankSet[tempStr]; ok {
							if end == tempStr {
								return step + 1
							}
							delete(bankSet, tempStr)
							queue = append(queue, tempStr)
						}
					}
				}
			}
		}
	}
	return -1
}
