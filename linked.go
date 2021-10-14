// package main

// import "fmt"

// type Node struct {
// 	next *Node // 다음 노드에 대한 포인트
// 	val  int
// }

// func main() {
// 	var root *Node

// 	root = &Node{val: 0} // root는 Node의 주소값 {val: 0}는 필드값

// 	for i := 1; i < 10; i++ {
// 		AddNode(root, i)
// 	}

// 	node := root
// 	for node.next != nil { // 다음 노드가 없을 때까지 반복
// 		fmt.Printf("%d ->", node.val) // 노드 출력
// 		node = node.next              // 다음 노드를 대입
// 	}
// 	fmt.Printf("%d\n", node.val) // 현재의 노드 출력
// }

// func AddNode(root *Node, val int) {
// 	var tail *Node
// 	tail = root            //tail에 root를 대입
// 	for tail.next != nil { // 테일의 마지막이 존재하면 계속 전진
// 		tail = tail.next // 다음 노드를 테일에 대입
// 	}

// 	node := &Node{val: val} // 새로운 필드값 Node{val: 1} 을 주고 새로운  Node 구조체가 만들어지고 그 주소값을 초기화
// 	tail.next = node
// }

package main

import "fmt"

type Node struct {
	next *Node // 다음 노드에 대한 포인트
	val  int
}

func main() {
	var root *Node
	var tail *Node // 맨끝의 노드 변수 추가

	root = &Node{val: 0}
	tail = root // 여기서 추가된 루트는 루트이자 테일

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i) // 마지막 노드를 테일로 갱신
	}
	PrintNodes(root)

	root, tail = RemoveNode(root, root, tail)

	PrintNodes(root)

}

func AddNode(tail *Node, val int) *Node { //반환값을 넣어줌
	node := &Node{val: val} // 새로운 필드값 Node{val: 1} 을 주고 새로운  Node 구조체가 만들어지고 그 주소값을 초기화
	tail.next = node        // 위의 값을 새로운 tail의 next로 들어감
	return node             // 새로 추가된 노드를 반환
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root { // root를 없앨때
		root = root.next
		if root == nil { // 노드가 하나밖에 없을때
			tail = nil // 테일도 없어진다
		}
		return root, tail
	}
	prev := root
	for prev.next != node {
		prev = prev.next // prev의 다음 노드가 현재노드와 같아지면 포문을 빠져나옴
	}
	prev.next = prev.next.next

	if node == tail { // tail을 지우고 싶을때
		prev.next = nil // 다음 노드가 없어야됨
		tail = prev     // 한칸 건너뜀
	} else {
		prev.next = prev.next.next

	}
	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil { // 다음 노드가 없을 때까지 반복
		fmt.Printf("%d ->", node.val) // 노드 출력
		node = node.next              // 다음 노드를 대입
	}
	fmt.Printf("%d\n", node.val) // 현재의 노드 출력
}
