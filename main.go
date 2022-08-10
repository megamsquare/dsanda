package main

import (
	"fmt"
)

// type Node struct {
// 	Val int
// 	Neighbors []*Node
// }

func main() {
	s := "ABAB"
	k := 2
	fmt.Println(characterReplacement(s, k))
	// setNums := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	// fmt.Println(setZeroes(setNums))
	// fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	// fmt.Println(isToeplitzMatrix([][]int{{1,2,3,4},{5,1,2,3},{9,5,1,2}}))
	// fmt.Println(cloneGraph(&Node{Val: 1, Neighbors: []*Node{{Val: 2}, {Val: 3}}}))
	// fmt.Println(canJump([]int{2,3,1,1,4}))
	// fmt.Println(hammingWeight(uint32(00000000000000000000000000001011)))
	// fmt.Println(getSum(3, 2))
	// fmt.Println(maxArea([]int{1, 2, 1})) // [1,2,1]

	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(cleanRoom([][]int{{1,1,1,1,1,0,1,1},{1,1,1,1,1,0,1,1},{1,0,1,1,1,1,1,1},{0,0,0,1,0,0,0,0},{1,1,1,1,1,1,1,1}}))
	// fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	// fmt.Println(search([]int{1, 3}, 3))
	// fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	// fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	// fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	// fmt.Println(productExceptSelf([]int{1,2,3,4}))
	// fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

}

func characterReplacement(s string, k int) int {
    // count := (map[byte]int)
    
    res,l,maxf:=0,0,0
    for r:=0;r<len(s);r++{
        count[s[r]]+=1
        maxf=maxVal(maxf,count[s[r]])
        if (r-l+1)-maxf>k{
            count[s[l]]-=1
            l++
        }
        res=maxVal(res,r-l+1)
    }
    return res
    
}

func maxVal(a, b int) {
    
    if a > b {
        return a
    }
    
    return b
}

// func setZeroes(matrix [][]int) [][]int {
// 	m, n := len(matrix), len(matrix[0])
// 	rowZero := false

// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] == 0 {
// 				matrix[0][j] = 0
// 				if i > 0 {
// 					matrix[i][0] = 0
// 				} else {
// 					rowZero = true
// 				}
// 			}
// 		}
// 	}
	
// 	for i := 1; i <= m; i++ {
// 		for j := 1; i <= n; j++ {
// 			fmt.Println(matrix[i])
// 			if matrix[i][0] == 0 || matrix[0][j] == 0 {
// 				fmt.Println(matrix[i][j])
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}

// 	if matrix[0][0] == 0 {
// 		for i := 0; i < m; i++ {
// 			matrix[i][0] = 0
// 		}
// 	}

// 	if rowZero {
// 		for j := 0; j < n; j++ {
// 			matrix[0][j] = 0
// 		}
// 	}

// 	return matrix
// }

// func lengthOfLIS(nums []int) int {
// 	lis := make([]int, len(nums))

// 	for i := 0; i < len(nums); i++ {
// 		lis[i] = 1
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] {
// 				lis[i] = int(math.Max(float64(lis[i]), float64(lis[j]+1)))
// 			}
// 		}
// 	}

// 	return findMaxNumInArr(lis)
// }

// func findMaxNumInArr(nums []int) int {
// 	max := nums[0]

// 	for _, v := range nums {
// 		if v > max {
// 			max = v
// 		}
// 	}

// 	return max
// }

// func isToeplitzMatrix(matrix [][]int) bool {

//     for i := 0; i < len(matrix)-1; i++ {
//         for j := 0; j < len(matrix[i])-1; j++ {
//             if(matrix[i][j] != matrix[i+1][j+1]) {
//                 return false
//             }
//         }
//     }
//     return true

// }

// func cloneGraph(node *Node) *Node {

// 	m := make(map[*Node]*Node)
// 	var dfs func(node *Node)
// 	dfs = func(node *Node) {
// 		if node == nil {
// 			return
// 		}
// 		newNode := new(Node)
// 		newNode.Val = node.Val
// 		m[node] = newNode
// 		for _, n := range node.Neighbors {
// 			if _, ok := m[n]; !ok {
// 				dfs(n)
// 			}
// 			newNode.Neighbors = append(newNode.Neighbors, m[n])
// 		}
// 	}

// 	dfs(node)
// 	return m[node]
// }

// func dfs(node *Node, m map[*Node]*Node) *Node {
// 	if node == nil {
// 		return nil
// 	}
// 	if m[node] != nil {
// 		return m[node]
// 	}
// 	newNode := &Node{Val: node.Val}
// 	m[node] = newNode
// 	for _, n := range node.Neighbors {
// 		newNode.Neighbors = append(newNode.Neighbors, dfs(n, m))
// 	}
// 	return newNode
// }

// func canJump(nums []int) bool {
//     goal := len(nums) - 1

//     for i := len(nums) - 1; i >= 0; i-- {
//         if i + nums[i] >= goal {
//             goal = i
//         }
//     }

//     if goal == 0 {
//         return true
//     }

//     return false

// }

// func hammingWeight(num uint32) int {
// 	count := 0
// 	fmt.Println(num)
// 	for num != 0 {
// 		// fmt.Println(num)
// 		num &= num - 1
// 		count++
// 	}
// 	return count
// }

// func getSum(a int, b int) int {
//     for (b != 0) {
//         temp := (a & b) << 1
//         a = a ^ b
//         b = temp
//     }
//     return a

// }

// func maxArea(height []int) int {
//     res := 0
//     l,r := 0, len(height)-1

//     for l < r {
//         area := (r - l) * int(math.Min(float64(height[l]), float64(height[r])))
//         res = int(math.Max(float64(res), float64(area)))

//         if height[l] < height[r] {
//             l += 1
//         } else {
//             r -= 1
//         }
//     }

//     return res

// }

// func threeSum(nums []int) [][]int {
//     result := [][]int{}
//     sort.Ints(nums)

//     a := nums[0]

//     for i := 0; i < len(nums); i++ {
//         if i > 0 {
//             a = nums[i]
//         }

//         l, r := i+1, len(nums) - 1
//         for l < r {
//             sumNum := a + nums[l] + nums[r]
//             if sumNum > 0 {
//                 r -= 1
//             } else if sumNum < 0 {
//                 l += 1
//             } else {
//                 result = append(result, []int{a, nums[l], nums[r]})

//             }

//             if nums[l] == nums[l - 1] && l < r {
//                 l += 1
//             }
//         }
//     }

//     return result
// }

// func numUniqueEmails(emails []string) int {
// 	hashMap := make(map[string]int)
// 	for i , v := range emails {
// 		email := strings.Split(v, "@")
// 		local := strings.Split(email[0], "+")[0]
// 		local = strings.Replace(local, ".", "", -1)
// 		hashMap[local+"@"+email[1]] = i
// 	}
// 	return len(hashMap)
// }

// func search(nums []int, target int) int {
//     numLen := len(nums)

//     if target == nums[0] {
//         return 0
//     }

//     hashMap := make(map[int]int, numLen)

//     for i := 0; i < numLen; i++ {
//         hashMap[nums[i]] = i
//         if _, ok := hashMap[target]; ok {
//             return hashMap[target]
//         }
//     }

//     return -1

// }

// func findMin(nums []int) int {
//     result := nums[0]
//     l, r := 0, len(nums) - 1

//     for l <= r {
//         if nums[l] < nums[r] {
//             result = int(math.Min(float64(result), float64(nums[l])))
//         }
//         m := (l+r)/2
//         result = int(math.Min(float64(result), float64(nums[m])))

//         if nums[m] >= nums[l] {
//             l = m + 1
//         } else {
//             r = m - 1
//         }
//     }
//     return result
// }

// func maxProduct(nums []int) int {
// 	res := findMaxNumInArr(nums)
// 	minNum, maxNum := 1, 1
// 	for _, n := range nums {
// 		// fmt.Printf("max: %d, min: %d\n", maxNum, minNum)
// 		minNum, maxNum = findMinAndMaxNum(n*maxNum, n*minNum, n)
// 		maxRes := math.Max(float64(maxNum), float64(minNum))
// 		fmt.Printf("max responce: %d\n", int(maxRes))
// 		res = int(maxRes)
// 	}
// 	return res
// }

// func findMaxNumInArr(nums []int) int {
// 	max := nums[0]

// 	for _, v := range nums {
// 		if v > max {
// 			max = v
// 		}
// 	}

// 	return max
// }

// func findMinAndMaxNum(a, b, c int) (int, int) {
// 	min := 0
// 	max := 0

// 	if a >= b && a >= c {
// 		max = a
// 	} else if b >= a && b >= c {
// 		max = b
// 	} else {
// 		max = c
// 	}

// 	if a <= b && a <= c {
// 		min = a
// 	} else if b <= a && b <= c {
// 		min = b
// 	} else {
// 		min = c
// 	}

// 	return min, max
// }

// func maxSubArray(nums []int) int {
//     maxNum := nums[0]
//     currNum := 0

//     for i := 0; i < len(nums); i++ {
//         if currNum < 0 {
//             currNum = 0
//         }
//         currNum += nums[i]
//         if maxNum < currNum {
//             maxNum = currNum
//         }
//     }

//     return maxNum
// }

// func productExceptSelf(nums []int) []int {
//     numLen := len(nums)

//     result := make([]int, numLen)

//     prefix := 1
//     postFix := 1

//     for i := 0; i < numLen; i++ {
//         result[i] = prefix
//         prefix *= nums[i]
//     }

// 	for i := numLen - 1; i >= 0; i-- {
// 		result[i] *= postFix
// 		postFix *= nums[i]
// 	}

//     return result

// }

// func maxProfit(prices []int) string {
//     l, r := 0, 0
//     maxP := 0
// 	b := 0
// 	s := 0
//     priceLen := len(prices)

//     for i := 0; i <= priceLen-1; i++ {
//         if prices[l] < prices[r] {
//             profit := prices[r] - prices[l]
//             if maxP < profit {
//                 maxP = profit
// 				b = l
// 				s = r
//             }
//         } else {
//             l = r
//         }
//         r += 1
//     }
//     return fmt.Sprintf("Buy index is %d with value of %d, Sell index %d with value of %d", b, prices[b], s, prices[s])

// }

// func twoSum(nums []int, target int) []int {
//     initialNum := map[int]int{}

//     for i, num := range nums {
//         numMatch := target - num
//         if j, isValid := initialNum[numMatch]; isValid {
//             return []int{j, i}
//         }
//         initialNum[num] = i
//     }
//     return []int{}
// }
