*Channels* are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.

Send values into a channel using the `channel <-` syntax.
The `<-channel` syntax receives a value from the channel.

By default sends and receives block until both the sender and receiver are ready.
