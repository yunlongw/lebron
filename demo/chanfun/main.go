package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"time"
)

func main() {
	//要处理的数据
	uid := []string{"a", "b", "c", "d", "e", "f"}

	//传递数据的逻辑
	generateFunc := func(source chan<- interface{}) {
		for _, v := range uid {
			source <- v
			fmt.Println("source:", v)
		}
	}

	// 处理数据的逻辑
	mapFunc := func(item interface{}, writer mr.Writer, cancel func(err error)) {
		tmp := item.(string) + ":1"
		writer.Write(tmp)
		fmt.Println("item:", item)
	}

	// 合并的数据逻辑
	reducerFunc := func(pipe <-chan interface{}, writer mr.Writer, cancel func(err error)) {
		var uid []string
		for v := range pipe {
			uid = append(uid, v.(string))
			fmt.Println("pipe:", uid)
		}
		writer.Write(uid)
	}

	// 开始并发处理数据
	// 超时时间
	ctx, cl := context.WithTimeout(context.Background(), time.Second*3)
	res, err := mr.MapReduce(generateFunc, mapFunc, reducerFunc, mr.WithContext(ctx))
	//开启现成控制超时，如果超时则调用cl方法停止所有携程
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("cl")
		cl()
	}()
	fmt.Println(res, err)
}
