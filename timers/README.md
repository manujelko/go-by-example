We often want to execute Go code at some point in the future, or repeatedly at some interval.
Go's built-in *timer* and *ticker* feature make both of these tasks easy.

Timers represent a single event in the future.
You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.

If you just wanted to wait, you could have used time.Sleep.
One reason a timer might be useful is that you can cancel the timer before it fires.
