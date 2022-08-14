package model

import (
	"bufio"

	"io"
	"os"
	"strings"
)

func get_ground_truth() [3000]int {
	ans := [3000]int{}
	// ground_truth, err := os.Open("model/ground_truth.txt")
	ground_truth, err := os.Open("ground_truth.txt")
	defer ground_truth.Close()
	counter := 0
	if err == nil {
		reader := bufio.NewReader(ground_truth)
		for {
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			sliced_line := strings.Split(str, ",")
			// fmt.Println(sliced_line)
			for j := 0; j < 3; j++ {
				if sliced_line[j+1][0] == 'T' {
					ans[counter*3+j] = 1
				}
			}
			if str[0] != 'i' {
				counter++
			}
		}
	} else {
		panic(err)
	}
	return ans
}

func calculate_score(ground_truth *[3000]int, submitted *[]int) (int, int, int, int) { // main, sub1, sub2, sub3
	l := len(*ground_truth)
	if len(*submitted) != l {
		return -1, 0, 0, 0
	}
	var count [3]int

	for i := 0; i < l; i++ {
		if (*ground_truth)[i] == (*submitted)[i] {
			count[i%3]++
		}
	}

	total_score := count[0]*count[0] + count[1]*count[1] + count[2]*count[2]

	return total_score, count[0], count[1], count[2]

}

func possess_submitted_content_into_array(submitted *string) []int {
	var ans []int
	for _, ch := range *submitted {
		if ch == 48 {
			ans = append(ans, 0)
		}
		if ch == 49 {
			ans = append(ans, 1)
		}
	}
	return ans
}

func Score_calculator(submitted *string) (int, int, int, int) {
	gt := get_ground_truth()
	sub := possess_submitted_content_into_array(submitted)
	return calculate_score(&gt, &sub)
}
