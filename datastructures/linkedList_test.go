package datastructures

import (
	"testing"
)

func testList(t *testing.T) {
	ll := &LinkedList{}
	ll.Add("apple")
	ll.Add("banana")
	ll.Add("cherry")
	ll.Display() // Output: apple -> banana -> cherry -> nil

	ll.Remove("banana")
	ll.Display() // Output: apple -> cherry -> nil

	ll.Remove("apple")
	ll.Display() // Output: cherry -> nil

	ll.Remove("cherry")
	ll.Display() // Output: nil
}
