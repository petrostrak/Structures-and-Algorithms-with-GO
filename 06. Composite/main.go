/*
A composite is a group of similar objects in a single object. Objects are stored in a tree form
to persist the whole hierarchy. The composite pattern is used to change a hierarchical
collection of objects. The composite pattern is modeled on a heterogeneous collection. New
types of objects can be added without changing the interface and the client code. You can
use the composite pattern, for example, for UI layouts on the web, for directory trees, and
for managing employees across departments. The pattern provides a mechanism to access
the individual objects and groups in a similar manner.

The composite pattern comprises the component interface, component class, composite,
and client:

The component interface defines the default behavior of all objects and behaviors
for accessing the components of the composite.

The composite and component classes implement the component interface.
The client interacts with the component interface to invoke methods in the
composite.

Let's say there is an IComposite interface with the perform method and
BranchClass that implements IComposite and has the addLeaf , addBranch , and
perform methods. The Leaflet class implements IComposite with the perform
method. BranchClass has a one-to-many relationship with leafs and branches .
Iterating over the branch recursively, one can traverse the composite tree.
*/

package main

import "fmt"

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet struct
type Leaflet struct {
	name string
}

// Leaflet class method perform
func (l *Leaflet) perform() {
	fmt.Println("Leaflet ", l.name)
}

// Branch struct
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

// The perform method of the Branch class calls the perform method on branch and leafs ,
// as seen in the code:

// Branch class method perform
func (b *Branch) perform() {
	fmt.Println("Branch ", b.name)

	for _, l := range b.leafs {
		l.perform()
	}

	for _, b := range b.branches {
		b.perform()
	}
}

// Branch class method add leaflet
func (b *Branch) add(l Leaflet) {
	b.leafs = append(b.leafs, l)
}

// As shown in the following code, the addBranch method of the Branch class adds a new
// branch :

// Branch clas method addBranch branch
func (b *Branch) addBranch(newBranch Branch) {
	b.branches = append(b.branches, newBranch)
}

// Branch class method getLeaflets
func (b *Branch) getLeaflets() []Leaflet {
	return b.leafs
}

func main() {
	var branch = &Branch{name: "branch 1"}
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}
	var branch2 = Branch{name: "branch 2"}
	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)
	branch.perform()
}
