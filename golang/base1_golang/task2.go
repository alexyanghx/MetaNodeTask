package base1_golang

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ✅指针,
// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。,
// ,

func Task2_1() {

	method := func(numPoint *int) {
		*numPoint++
	}

	num := 10
	method(&num)
	fmt.Println("指针修改后的值:", num) // 输出 11

}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。,

func Task2_2() {
	method := func(slicePoint *[]int) {
		for i, v := range *slicePoint {
			(*slicePoint)[i] = v * 2
		}
	}

	slice := []int{1, 2, 3, 4, 5}
	method(&slice)
	fmt.Println("指针修改后的切片:", slice) // 输出 [2, 4, 6, 8, 10]
}

// ,
// ✅Goroutine,
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。,

func Task2_3() {
	method := func() {
		go func() {
			for i := 1; i <= 10; i++ {
				if i%2 == 1 {
					fmt.Println("奇数打印:", i)
				}
			}
		}()

		go func() {
			for i := 1; i <= 10; i++ {
				if i%2 == 0 {
					fmt.Println("偶数打印:", i)
				}
			}
		}()
	}

	method()                    // 启动协程
	time.Sleep(2 * time.Second) // 确保协程有时间执行
}

// ,
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。,

func Task2_4() {
	type Task struct {
		Name string
		Func func()
	}

	var wg sync.WaitGroup

	method := func(tasks []Task) {
		wg.Add(len(tasks))

		for _, task := range tasks {
			go func(t Task) {
				defer wg.Done() // 确保任务完成时调用 Done

				start := time.Now()
				fmt.Printf("开始执行任务: %s\n", t.Name)
				t.Func()
				duration := time.Since(start)
				fmt.Printf("任务 %s 执行时间: %v\n", t.Name, duration)
			}(task)
		}
	}
	tasks := []Task{
		{
			Name: "任务1",
			Func: func() {
				time.Sleep(time.Second * 2)
			},
		},
		{
			Name: "任务2",
			Func: func() {
				time.Sleep(time.Second * 3)
			},
		},
	}
	method(tasks) // 启动任务调度器
	wg.Wait()
}

// ,
// ✅面向对象,
// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。,
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	width  float64
	height float64
}

func (rt *Rectangle) Area() float64 {
	return rt.width * rt.height
}

func (rt *Rectangle) Perimeter() float64 {
	return 2 * (rt.width + rt.height)
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius // 使用近似值 3 代替 π
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius // 使用近似值 3 代替 π
}

func Task2_5() {
	var shape Shape
	shape = &Rectangle{width: 5, height: 10}
	fmt.Println("长方形面积:", shape.Area()) // 调用 Area 方法
	fmt.Println("长方形周长:", shape.Perimeter())

	shape = &Circle{radius: 7}
	fmt.Println("圆形面积:", shape.Area())      // 调用 Area 方法
	fmt.Println("圆形周长:", shape.Perimeter()) // 调用 Perimeter 方法
}

// ,
// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。,
// ,

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     // 组合 Person 结构体
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("员工信息: 姓名: %s, 年龄: %d, 员工ID: %s\n", e.Name, e.Age, e.EmployeeID)
}

func Task2_6() {
	employee := &Employee{
		Person:     Person{Name: "张三", Age: 30},
		EmployeeID: "E12345",
	}
	employee.PrintInfo() // 输出员工信息
}

// ✅Channel,
// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。,

func Task2_7() {
	method := func() {
		var ch = make(chan int)
		go func() {
			for i := 1; i <= 10; i++ {
				ch <- i // 将整数发送到通道
			}
			fmt.Println("发送完毕，关闭通道")
			close(ch) // 关闭通道
		}()

		go func() {
			for {
				select {
				case num, ok := <-ch:
					if ok {
						fmt.Println("接收到的整数:", num) // 从通道接收整数并打印
					} else {
						fmt.Println("通道已关闭")
						return
					}
				default:
					fmt.Println("等待接收数据...")
					time.Sleep(100 * time.Millisecond) // 避免忙等待
				}
			}
		}()
	}

	method()
	time.Sleep(2 * time.Second) // 确保协程有时间执行
}

// ,
// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。,
func Task2_8() {
	method := func() {
		ch := make(chan int, 10) // 创建一个缓冲通道，容量为10

		go func() {
			for i := 1; i <= 100; i++ {
				ch <- i // 生产者将整数发送到通道
				fmt.Println("生产者发送:", i)
			}
			close(ch) // 关闭通道
		}()

		go func() {
			for num := range ch { // 从通道接收整数
				fmt.Println("消费者接收到:", num)
			}
			fmt.Println("消费者完成接收")
		}()
	}

	method()
	time.Sleep(5 * time.Second) // 确保协程有时间执行
}

// ,
// ✅锁机制,
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。,
type Counter struct {
	mu    sync.Mutex // 互斥锁
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()         // 上锁
	defer c.mu.Unlock() // 确保在函数结束时解锁
	c.count++           // 递增计数器
}

func Task2_9() {
	counter := &Counter{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment() // 每个协程递增计数器
			}
		}()
	}

	time.Sleep(2 * time.Second)            // 等待所有协程完成
	fmt.Println("计数器的最终值:", counter.count) // 输出计数器的值
}

// ,
// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func Task2_10() {
	var counter uint64

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				// 使用原子操作递增计数器
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	time.Sleep(2 * time.Second)        // 等待所有协程完成
	fmt.Println("无锁计数器的最终值:", counter) // 输出计数器的值
}
