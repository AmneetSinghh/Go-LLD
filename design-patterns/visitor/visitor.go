package main

import "fmt"

/*

✅ Purpose: Separates algorithms from objects, allowing new behaviors to be added without modifying existing object structures.
✅ How It Works: Objects accept a visitor, which then executes the correct method based on the object type (Double Dispatch).


Can We Use Strategy Pattern Instead of Visitor for Adding New Behaviors?
✅ Short Answer:

If each object type has different behaviors, Visitor is better.
If the behavior is the same but needs to be swapped dynamically, Strategy can work. ( multiple strategies)
If you need to add new behaviors later, Visitor is better because you don’t modify existing classes.
But let’s see when Strategy can work instead of Visitor.

*/

/*
strategy :
"Strategy is used when we have multiple interchangeable algorithms that apply the same way to all objects.
 But if the behavior differs based on the object type, then the Visitor Pattern is better. In Visitor,
 we pass the object to a Visitor, and the Visitor uses different strategies for each object."


we pass objet to visitor, he will use streagy.
																		strategy 								visitor
Multiple strategies per object?								❌ No, only one at a time					✅ Yes, multiple at once
Applies differently per object type?						❌ No, same for all objects					✅ Yes, different per object
Can swap dynamically at runtime?							✅ Yes, can change strategies at runtime		❌ No, behavior is fixed per visitor

Best for...	Changing algorithms dynamically	Applying different operations to objects
*/

type Visitor interface {
	VisitInvoice(i *Invoice)
	VisitReceipt(r *Receipt)
}

type BusinessDocument interface {
	Accept(v Visitor) // eah object which want dyamic visitor, must use this
}

type Invoice struct {
	amount int
}

func (i *Invoice) Accept(v Visitor) {
	v.VisitInvoice(i) // delegate to visitor pass concrete type;
}

type Receipt struct {
	amount int
}

func (i *Receipt) Accept(v Visitor) {
	v.VisitReceipt(i) // delegate to visitor pass concrete type;
}

/// visitors
type XMLExportVisitor struct{}

func (x *XMLExportVisitor) VisitInvoice(i *Invoice) {
	fmt.Printf("Invoice amount visitor\n", i.amount)
}
func (x *XMLExportVisitor) VisitReceipt(r *Receipt) {
	fmt.Printf("Receipt amoutn visitor\n", r.amount)
}

/// visitors
type ReportGeneVisitor struct{}

func (x *ReportGeneVisitor) VisitInvoice(i *Invoice) {
	fmt.Printf("Invoice amount visitor:: report generatio\n", i.amount)
}
func (x *ReportGeneVisitor) VisitReceipt(r *Receipt) {
	fmt.Printf("Receipt amoutn visitor:: report generation\n", r.amount)
}

func main() {
	invoice := &Invoice{amount: 1000}
	receipt := &Receipt{amount: 300}

	// visitors
	xmlExporter := &XMLExportVisitor{}
	ReportGeneVisitor := &ReportGeneVisitor{}

	invoice.Accept(xmlExporter)
	receipt.Accept(xmlExporter)

	invoice.Accept(ReportGeneVisitor)
	receipt.Accept(ReportGeneVisitor)
}
