// DoublyLinked List Project
package main

import "fmt"

//Import fmt for formatting I/O operations

//Utilizing generics: Create node struct for creating new nodes, with type T
type Node[T any] struct{
	//Pointer to next node
	next *Node[T]
	//Pointer to previous node
	prev *Node[T]
	//Value field for storing data / type t / of new nodes
	value T
}

//Utilizing generics: Create linkedlist struct for creating new linked list structures, with type t
type LinkedList[T any] struct{
	//Create head node
	head *Node[T]
	//create length field for storing length of list
	length uint
}

//Define index function which returns value of any given index received, of type t
func (ll *LinkedList[T]) index(idx uint) (T, bool){
	//create variable with empty value, of type T
	var value T
	//Edge Case: If index is out of bounds, return value and false boolean
	if idx >= ll.length{
		return value, false
	}
	//initialize an empty current index
	curIdx := 0
	//Set current node to head to traverse the list
	current := ll.head
	//traverse the list, if it is not the head or the current index is not 0 then perform logic
	for current != ll.head || curIdx == 0{
		if idx == uint(curIdx){
			value = current.value
			break
		}
		curIdx++
		current = current.next
	}
	//increment the current index by one
	curIdx++
	//move to the next node
	current = current.next
	//return the value, and true because it was within the list and true
	return value, true
}

//Define append function for adding new nodes to LinkedList
func (ll *LinkedList[T]) append(value T){
	//Increment because adding a new node, so regardless this logic occurs
	ll.length++
	//Edge Case: Set pointers for new node when length is only one
	if ll.length == 1 {
		//Create new node
		ll.head = &Node[T]{}
		//set pointers for next and previous to new head node
		ll.head.next = ll.head
		ll.head.prev = ll.head
		//set value for new node
		ll.head.value = value
		return
	}
	//Main logic which implements the append function
	//Create new node
	node := &Node[T]{}
	//Set value for new node
	node.value = value
	//Set pointers logic
	//Access the tail node
	tail := ll.head.prev
	//set next pointer in tail to new node
	tail.next = node
	//set new node pointers
	node.prev = tail
	node.next = ll.head
	//set head pointer to the new node
	ll.head.prev = node
}

//Define pop function which removes node from list and returns value
func (ll *LinkedList[T]) pop() T{
	//pop function allows for you to remove and return value
	//edge case for when the length is 0, need to create null variable and decrement counter and assign nodes to null
	//decrement the counter
	ll.length--
	//edge case: if empty list
	if ll.length ==0{
		//create abstract method to create a null variable
		var null T
		//set the head node value to the head
		value := ll.head.value
		//set value of head node to null
		ll.head.value = null
		//return the value
		return value
	}
	//Assign pointers to be removed from tail
	//Access tail , create tail variable to the store the head previous pointer
	tail := ll.head.prev
	//store the previous node's pointer
	tailPrev := tail.prev
	//update head pointer to the tail prev
	ll.head.prev = tailPrev
	//set tail next pointer to head
	tailPrev.next = ll.head
	return tail.value
}

func (ll *LinkedList[T]) printList(){
	//store values in a slice
	values := []T{}

	//Set current to be at the head of the list
	current := ll.head
	//Loop through the list until the counter reaches 0
	for ll.length != 0{
		values = append(values, current.value)
		current = current.next
		if current == ll.head {
			break
		}
	}
	fmt.Println(values)

}



func main(){
	//initialize an empty LinkedList
	ll := LinkedList[int]{}
	ll.append(7)
	ll.append(10)
	ll.append(11)
	ll.append(12)
	ll.pop()
	ll.pop()
	ll.index(7)
	ll.printList()
	fmt.Println(ll.index(1))
}

// go run double.go