package dsa_test

import (
	"testing"
     //"reflect"

	"github.com/Mar1O9/dsa"
)


func TestPush(t *testing.T) {
    var testStack = dsa.NewStack[int]()
    
    testStack.Push(1)
    if testStack.Len() != 1 {
        t.Errorf("len(testStack) == 1, but got: %v", testStack.Len())
    }
}

func TestPop(t *testing.T) {
    var testStack = dsa.NewStack[int]()
    
    testStack.Push(1)
    testStack.Push(19)

    testStack.Pop()
    if testStack.Len() != 1 {
        t.Errorf("len(testStack) == 1, but got: %v", testStack.Len())
    }

}
