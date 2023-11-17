package nnkosmos

// 使用 m 个 goroutine 计算自然数求和，
// 也就是 1 + 2 + 3 + ⋯ + n，也就是 n(n+1)/2。
func concurrentSum(n int, m int) int {
	ch := make(chan int)

	segLen := n / m
	for i := 0; i < m; i++ {
		start := i*segLen + 1
		end := start + segLen - 1
		if i == m-1 {
			end = n
		}
		go func() {
			var partialSum int
			for j := start; j <= end; j++ {
				partialSum += j
			}
			ch <- partialSum
		}()
	}

	var result int
	for i := 0; i < m; i++ {
		result += <-ch
	}
	close(ch)
	return result
}

// 问：在这段代码中，最后有必要关闭 channel 吗？
//
// GPT-4：
// 在你提供的 Go 语言代码片段中，虽然关闭 channel 不是绝对必要的，但它是一种良好的编程实践。
// 关闭 channel 可以明确地表示不再有数据发送至该 channel，
// 这对于防止潜在的资源泄漏和提高代码的清晰度是有好处的。
// 在该代码中，所有的 goroutine 都完成了对 channel 的数据发送，
// 并且主 goroutine 已经从 channel 中接收了所有数据，
// 因此，在 return result 之前关闭 channel 是安全的。
// 这样做有助于避免在其他部分的代码中意外地向这个 channel 发送数据，也使得代码更容易理解和维护。
