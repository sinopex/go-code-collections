// Package memory_leak 内存泄露范例
//
// - 获取长字符串中的一段导致长字符串未释放
// - 同样，获取长slice中的一段导致长slice未释放
// - 在长slice新建slice导致泄漏
// - time.Ticker未关闭导致泄漏
// - Finalizer导致泄漏
// - Deferring Function Call导致泄漏
// - http请求response未关闭
// - goroutine泄漏
package memory_leak
