package main

import "fmt"

//定义单链表结构体

type Node struct {
    data interface{} //数据域
    next *Node       //指针域
}
type LinkTable struct {
    length   int //储存链表的长度
    headNode *Node
}

/*
type Method interface {
    IsNull() bool                    //1、判断是否为空
    GetLength() int                  //2、获取链表长度
    InsertElem(i int, v interface{}) //3、在指定位置添加元素
    AddElem(v interface{})           //4、在头部插入元素
    AppendElem(v interface{})        //5、在尾部插入元素
    DeleteElem(i int)                //6、删除指定位置元素
    RemoveElem(v interface{})        //7、删除指定值的元素
    ContaineElem(v interface{}) bool //8、是否包含指定值的元素
    LocateElem(i int) interface{}    //9、查找指定位置元素的值
    ShowLinkTable()                       //10、遍历链表所有结点
}
*/
/*单链表的初始化
1、生成新结点作为头结点，用头指针指向头结点
2、头结点的指针域置空
*/
func InitLinkTable() *LinkTable {
    //即构造一个空的单链表L(包含头指针)
    node := new(Node)
    L := new(LinkTable)
    L.headNode = node
    return L
}

/*单链表的取值
1、用指针P指向首元结点，用j做计数器初值赋为1
2、从首元结点开始依次顺着链域(指针域)next向下访问，只要指向当前结点的指针P不为空，
并且没有达到序号为i的结点，则循环执行以下操作：
    2.1、P指向下一个结点
    2.2、计数器j相应加1
3、退出循环时，如果指针P为空，或者计数器j大于i，说明指定的序号i值
不合法(i大于表长n或i小于等于0)，取值失败返回ERROR；否则取值成功，
此时j==i时，P所指的结点就是要找的第i个结点，用参数e保存当前结点的数据域，返回OK
*/
func (linktable *LinkTable) GetElem(index int) int {
    if index <= 0 || index > linktable.length {
        return 0
    } else {
        pre := linktable.headNode
        for j := 0; j < index-1; j++ {
            if pre != nil {
                pre = pre.next
            }
        }
        return pre.data.(int)
    }
}

/*单链表的按值查找
1、用指针p指向首元结点
2、从首元结点开始以此顺着链域next向下查找，只要指向当前结点的
指针p不为空，并且p所指结点的数据域不等于给定值e，则执行以下操作：
    2.1、p指向下一个结点
3、返回p。若查找成功，p此时即为结点的地址值，若查找失败，p的值即为NULL。
*/
func (linktable *LinkTable) LocateElem(v interface{}) bool {
    if linktable.IsNull() {
        fmt.Println("err")
        return false
    } else {
        pre := linktable.headNode
        for pre != nil {
            if pre.data == v {
                return true
            }
            pre = pre.next
        }
        return false
    }
}

/*单链表的插入=>将值为e的新结点插入到表的第i个结点的位置上，即插入到结点a(i-1)与a(i)之间
1、查找结点a(i-1)并由指针p指向该结点
2、生成一个新结点*s
3、将新结点*s的数据域置为e
4、将新结点*s的指针域指向结点a(i)
5、将结点*p的指针域指向新结点*s
*/
func (linktable *LinkTable) InsertElem(index int, v interface{}) {
    if index <= 0 || index > linktable.length {
        fmt.Println("err")
    } else {
        pre := linktable.headNode
        node := &Node{data: v}
        if index == 1 {
            node.next = pre
            linktable.headNode = node
        } else {
            for count := 1; count < index-1; count++ {
                pre = pre.next
            }
            node.next = pre.next
            pre.next = node
        }
        linktable.length--
    }
}

/*单链表的删除
1、查找结点a(i-1)并由指针p指向该结点
2、临时保存待删除结点a(i)的地址在q中，以备释放
3、将结点*p的指针域指向a(i)的直接后继结点
4、释放结点a(i)的空间
*/
func (linktable *LinkTable) DeleteElem(index int) {
    if index <= 0 || index > linktable.length {
        fmt.Println("删除位置不合理")
        return
    } else {
        pre := linktable.headNode
        if index == 1 {
            linktable.headNode = pre.next
        } else {
            pre := linktable.headNode
            for count := 1; count < index-1; count++ {
                pre = pre.next
            }
            pre.next = pre.next.next
        }
        linktable.length--
    }
}

func (linktable *LinkTable) RemoveElem(v interface{}) {
    pre := linktable.headNode
    if pre.data == v {
        linktable.headNode = pre.next
    } else {
        for pre.next != nil {
            if pre.next.data == v {
                pre.next = pre.next.next
                return
            } else {
                pre = pre.next
            }
        }
        fmt.Println("fail")
        return
    }
}

func (linktable *LinkTable) IsNull() bool {
    if linktable.length == 0 {
        return true
    } else {
        return false
    }
}

func (linktable *LinkTable) AddElem(v interface{}) {
    node := &Node{data: v}
    if linktable.IsNull() { //处理空表的插入，否则会导致一个空的头结点后移
        linktable.headNode = node
        linktable.length++
        return
    }
    node.next = linktable.headNode
    linktable.headNode = node
    linktable.length++
    return
}

func (linktable *LinkTable) AppendElem(v interface{}) {
    node := &Node{data: v}
    if linktable.IsNull() {
        linktable.headNode.next = node
    } else {
        cur := linktable.headNode
        for cur.next != nil {
            cur = cur.next
        }
        cur.next = node
    }
    linktable.length++
    return
}

func (linktable *LinkTable) ShowLinkTable() {
    if !linktable.IsNull() {
        cur := linktable.headNode
        for {
            fmt.Printf("\t%v", cur.data)
            if cur.next != nil {
                cur = cur.next
            } else {
                break
            }
        }
        fmt.Printf("\n")
    }
}

func main() {
    L := InitLinkTable()
    msg := []int{12, 5, 3, 8, 55, 13}
    for i := range msg {
        L.AddElem(msg[i])
    }
    fmt.Println("---- 添加元素 ----")
    L.AppendElem(66)
    L.ShowLinkTable()
    fmt.Println("---- 按位删除元素 ----")
    L.DeleteElem(3)
    L.ShowLinkTable()
    fmt.Println("---- 按值删除元素 ----")
    L.RemoveElem(13)
    L.ShowLinkTable()
    fmt.Println("---- 插入元素 ----")
    L.InsertElem(1, 88)
    L.ShowLinkTable()
    fmt.Println("---- 按值寻找元素 ----")
    fmt.Println(L.LocateElem(88))
    fmt.Println("---- 按位寻找元素 ----")
    fmt.Println(L.GetElem(4))
}