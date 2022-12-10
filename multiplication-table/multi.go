/*
Your task, is to create NxN multiplication table, of size provided in parameter.

for example, when given size is 3:

1 2 3
2 4 6
3 6 9
for given example, the return value should be: [[1,2,3],[2,4,6],[3,6,9]]
*/

package multiplicationtable

func MultiplicationTable(size int) [][]int {
	// Implement me! :)
	var s [][]int
	var ss []int

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			ss = []int{i, i * , i * j} //WHAT COMES HERE???
		}

		s = append(s, ss)
	}
	/*
		switch size {
		case 2:
			s := [][]int{{1, 2}, {2, 4}}
			return s

		case 3:
			s := [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}
			return s

		case 4:
			s := [][]int{{1, 2, 3, 4}, {2, 4, 6, 8}, {3, 6, 9, 12}, {4, 8, 12, 16}}
			return s

		case 5:
			s := [][]int{{1, 2, 3, 4, 5}, {2, 4, 6, 8, 10}, {3, 6, 9, 12, 15}, {4, 8, 12, 16, 20}, {5, 10, 15, 20, 25}}
			return s
		}
	*/
	return s
}
