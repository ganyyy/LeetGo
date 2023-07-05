package main

func maximumEvenSplit(finalSum int64) []int64 {
	var res []int64
	if finalSum%2 > 0 {
		return res
	}
	for i := int64(2); i <= finalSum; i += 2 {
		res = append(res, i)
		finalSum -= i
	}
	if len(res) == 0 {
		return res
	}
	res[len(res)-1] += finalSum
	return res
}
