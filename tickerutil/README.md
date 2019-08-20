> Ticker

每隔 duration 时间会把当前的时间点放入到 channel 中，应用可以从 channel 进行读取。应用需要周期性的时间间隔，可以使用此方法。
使用 Ticker 有两种方式，NewTicker 可以获取 Ticker 实例，Stop 可以显示的停止 Tick 运行。Stop 可以释放 timer 资源，但不会关闭 channel，防止应用层报错。
如果 Ticker 一直随应用运行，不会关闭，可以使用 time.Tick 直接获取 time channel。这个没有 Ticker 实例，无法显示关闭。 timer 会一直运行。

> Timer

定时器，和 Tick 类似，经过 duration 时间，Timer 会触发，并且往 channel 写入当前时间点，此时 Timer 不再计时。当应用层重新调用 Reset 函数，才又开始计时，这个是和 Tick 不同的。Tick 是周期性的计时。
Timer 还支持计时结束时，触发自定义函数。AfterFunc 会返回 Timer 实例。AfterFunc 调用后，只会计时结束后触发一次自定义函数调用，如果需要再次触发，需要显示调用 Timer.Reset 函数。
如果结束定时器，调用 Stop 即可。

> 参考链接：
- [go timer 和 ticker 的区别](https://learnku.com/articles/23578/the-difference-between-go-timer-and-ticker)
- [GO Timer 机制探究](http://blog.bruceding.com/485.html)