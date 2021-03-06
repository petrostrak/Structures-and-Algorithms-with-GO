/*
In a scenario where class responsibilities are removed or added, the decorator pattern is
applied. The decorator pattern helps with subclassing when modifying functionality,
instead of static inheritance. An object can have multiple decorators and run-time
decorators. The single responsibility principle can be achieved using a
decorator. The decorator can be applied to window components and graphical object
modeling. The decorator pattern helps with modifying existing instance attributes and
adding new methods at run-time.
The decorator pattern participants are the component interface, the concrete component
class, and the decorator class:

The concrete component implements the component interface.

The decorator class implements the component interface and provides
additional functionality in the same method or additional methods. The
decorator base can be a participant representing the base class for all decorators.

Let 's say IProcess is an interface with the process method. ProcessClass implements
an interface with the process method. ProcessDecorator implements the process
interface and has an instance of ProcessClass . ProcessDecorator can add more
functionality than ProcessClass ,
*/

package main

import "fmt"

// IProcess Interface
type IProcess interface {
	process()
}

// ProcessClass struct
type ProcessClass struct{}

// ProcessClass method process
func (p ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

// ProcessDecorator struct
type ProcessDecorator struct {
	processInstance *ProcessClass
}

// In the following code, the ProcessDecorator class process method invokes the process
// method on the decorator instance of ProcessClass

// ProcessDecorator class method process
func (pd *ProcessDecorator) process() {
	if pd.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		pd.processInstance.process()
	}
}

func main() {
	// var process = &ProcessClass{}
	var decoretor = &ProcessDecorator{}
	decoretor.process()
	decoretor.processInstance.process()
	decoretor.process()
}
