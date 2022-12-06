package kata

/*
Count the number of divisors of a positive integer n.

Random tests go up to n = 500000.

Examples (input --> output)
4 --> 3 (1, 2, 4)
5 --> 2 (1, 5)
12 --> 6 (1, 2, 3, 4, 6, 12)
30 --> 8 (1, 2, 3, 5, 6, 10, 15, 30)
Note you should only return a number, the count of divisors.
The numbers between parentheses are shown only for you to see
which numbers are counted in each case.
*/

func Divisors(n int) int {
	//Put your code here
	ns := make([]int, 0, 5000000)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ns = append(ns, i)
		} else {
			continue
		}
		//fmt.Println(ns)
	}
	return len(ns)
}

//other's solutions:

/*
func Divisors(n int)int{
	count:=0
	for i:=1; i <= n; i++{
		if n%i == 0{
			count +=1
		}
	}
	return count
}
*/

/*
func Divisors(n int)int{
  nDiv:=1
  for i:=1; i <= n/2 ; i++ {
    if n % i == 0 {
      nDiv++
    }
  }
  return nDiv
}
*/

/*
func Divisors(n int) int {
  numbers, res := make(map[int]int), 1
  for i := 2; n > 1; {
    if n % i != 0 {
      i++
      continue
    }
    numbers[i] += 1
    n /= i
  }

  for _, number := range numbers {
    res *= (number + 1)
  }
  return res
}
*/

/*
func Divisors(n int) (count int) {
  for i := 1; i < n + 1; i++ {
    if n % i == 0 {
      count++
    }
  }
  return
}
*/

/*
func Divisors(n int)int{
    count := 0
    found := make(map[int]bool)
     for i:=1 ; i <= n ; i++{
        if n%(i*1) == 0 && !found[i*1]{
          count++
          found[i*1] = true
        }
        if n%(i*2) == 0 && !found[i*2]{
          count++
          found[i*2] = true
        }
       if n%(i*3) == 0 && !found[i*3] {
         count++
         found[i*3] = true
       }
       if n%(i*5) == 0 && !found[i*5]{
         count++
         found[i*5] = true
       }
       if n%(i*7) == 0 && !found[i*7]{
         count++
         found[i*7] = true
       }

     }
  return count
}
*/
