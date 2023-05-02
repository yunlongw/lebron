package main

import (
	"fmt"
	"math/rand"
)

/**
在 Go 中，chan（通道）是一种特殊的类型，用于在不同的 Goroutine 之间传递数据。通道有以下几个特点：

1.类型必须指定：在创建通道时，必须指定通道所能传递的数据类型。通道只能传递这种类型的数据，否则会在编译时产生错误。
2.可以阻塞：当从通道中读取数据时，如果通道为空，则会被阻塞。同样，当向通道中写入数据时，如果通道已满，则也会被阻塞。
3.支持双向和单向操作：通道可以用于双向传递数据，也可以将其用作只读或只写的通道。
4.作用类似于管道：通道可以被用来传递数据，类似于 Unix 系统中的管道（pipe）。

通道在 Go 语言中被广泛应用于多个 Goroutine 之间的通信和协作。通常情况下，通道被用来解决并发编程中的同步和数据共享
问题，可以保证数据的安全性和可靠性。

要更好地理解通道，可以将其类比为管道，数据就像水流一样在通道中传递。通道的缓冲区大小类似于管道的容量，当管道/通道已满
时，就会被阻塞，等待容量增大/数据被消费。

同时，通道也可以用于实现事件的发布和订阅，一个 Goroutine 可以向通道中发送事件，另一个 Goroutine 可以从通道中接收
事件并处理。这种模式类似于观察者模式。

总之，通道是 Go 语言中的一个重要特性，使用通道可以轻松地实现并发编程中的同步和数据共享，提高程序的可靠性和性能。

创建一个 chan ，要指定管道两头的输入和输出 , 通常输出端是一个 for 死循环等待输出
*/

type GenerateFunc func(source chan<- interface{})

func generateRandomNumbers(source chan<- interface{}) {
	defer close(source)
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		source <- num
	}
}

func main() {
	// 创建一个 channel
	ch := make(chan interface{})

	// 定义一个 GenerateFunc 类型的函数变量
	var f GenerateFunc

	// 将 generateRandomNumbers 函数赋值给函数变量 f
	f = generateRandomNumbers

	// 在一个单独的 goroutine 中调用函数变量 f，并将 channel 作为参数传递给它
	go func() {
		f(ch)
	}()

	// 从 channel 中读取数据，并打印输出
	for num := range ch {
		fmt.Println(num)
	}

}
