/*
约瑟夫环问题
思路：
一： 数学方式，取模计算结果。代码少，运行效率低？python低……应该和内部实现有关
二： 生成一个1~n的数组，将剔除的对象打印（或存储）并赋值为0，迭代+1前进行判断，跳过0值
三： 构建链式结构？猜测：构建结构体{内存地址：数据值}，单独列表储存内存地址序列（不知道如何实现）
Q1: 时间复杂度和空间复杂度是啥
Q2: 效率高 代码量 关联 Python和Go不同
Q3: 模拟 递归 通顶
*/
package HomeWork

import (
	"fmt"
)

// 递归的思路
func main() {
	var people, circle, s = 0, 0, 0
	fmt.Printf("请输入people和circle")
	_, _ = fmt.Scan(&people, &circle)
	fmt.Printf("there are %d people and %d times a round", people, circle)
	var i int
	for i = 2; i <= people; i++ {
		s = (s + circle) % i
	}
	fmt.Printf("the winner is %d", s+1)
}

/*
模拟算法
func main(){
	var n,find,count int //find=1找到下一个猴子让它出去,count计数是否数到3
	fmt.Scan(&n)
	num := 0        //当num到达N-1时，只剩一只猴子
	flag :=[]int{}  //flag[i],该猴子已经出去了

	for i:=0;i<n;i++ {
		flag = append(flag,0)
	}
	i := 0          //找下一只要出去的猴子从i开始找
	for num != n-1 {
		find = 0
		count=0
		for ;find==0;i++{
			t := i%n
			if flag[t] == 0 {
				count+=1
			}
			if count==3 {
				num ++
				find=1
				flag[t] = 1
			}
		}
	}
	fmt.Println(flag)
	for i=0;i<n;i++ {
		if flag[i] == 0{
			fmt.Println(i)
		}
	}
}
*/

/*
递归算法另一种实现, 虽然看上去更递归，但是想着反而更抽象了
一个环形，每次错位都错的是上轮喊号的余数
func josephus(n int,k int) int{
	if n==1 {
		return 0
	}else {
		return (josephus(n-1,k)+k)%n
	}
}

func main(){
	n:=11
	k:=3
	res:= josephus(n,k)
	fmt.Println(res)
}
*/
