package main

func IsValidSubsequence(array []int, sequence []int) bool {
	seq :=0
	arr:=0
	for arr < len(array) && seq < len(sequence){
		if array[arr] == sequence[seq]{
			seq +=1
		}
		arr +=1
	}
	return seq == len(sequence)
}

