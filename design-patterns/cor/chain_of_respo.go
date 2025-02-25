package main

import "fmt"

/*
Chain of Responsibility Real-World Scenarios

Technical Support Escalation:
- Requests are handled by successive support levels (e.g., Level 1 to Level 3) until resolved.

Logging and Event Processing:
- Log messages pass through different handlers (debug, info, error) that decide whether to process or forward them.

Web Middleware:
- HTTP requests traverse middleware layers (authentication, authorization, logging) before reaching the final handler.

Approval Workflows:
- Documents or purchase requests are approved sequentially by managers, directors, etc., with each level deciding to approve or escalate.

Security Filtering:
- Incoming requests pass through a chain of security filters (firewall, antivirus, intrusion detection) that assess or block threats.

***************************************** Discription *****************************************

- So handler only needs to know about the next processor, which will process the requet,
separate of concerns. adding new handler is easy, OC principle.

- If we need to change the next processor, we can change dynamically , as each processor is using same interface....

- So client has acces to next processor, but it doesn't know who it is. Separate of concerns.
*/

type Patient struct {
	name           string
	regisDone      bool
	docCheckUPdone bool
	medicineDone   bool
	paymentDone    bool
}

type Department interface {
	execute(p *Patient)
}

type Receiption struct {
	next Department // interface which all processor will adhere.
}

func (r *Receiption) execute(p *Patient) {
	if p.regisDone {
		fmt.Println("Patient regis is done")
		r.next.execute(p)
		return
	}
	fmt.Println("Patient regis is pedning")
	p.regisDone = true
	r.next.execute(p)
}

type Doctor struct {
	next Department
}

func (r *Doctor) execute(p *Patient) {
	if p.docCheckUPdone {
		fmt.Println("doctor checkup is done")
		r.next.execute(p)
		return
	}
	fmt.Println("doctor checkup is pedning")
	p.docCheckUPdone = true
	r.next.execute(p)
}

type Medical struct {
	next Department
}

func (r *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("medical checkup is done")
		r.next.execute(p)
		return
	}
	fmt.Println("medical checkup is pedning")
	p.medicineDone = true
	r.next.execute(p)
}

type Cashir struct {
	next Department
}

func (r *Cashir) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("payment is done")
		return
	}
	p.paymentDone = true
	fmt.Println("payment is pedning")
}

func main() {
	p := &Patient{}
	d := &Doctor{
		next: &Receiption{
			next: &Medical{
				next: &Cashir{},
			},
		},
	}
	d.execute(p) // main just know about execute method,no concrete implementations of low level modules.

	fmt.Println("Secod time checkup")
	d.execute(p)
}

/*

Summary:

High-Level Module:
The code that sets up and controls the patient processing workflow (in main).

Low-Level Modules:
The individual handlers (Receiption, Doctor, Medical, and Cashir) that implement the concrete actions.

Each handler interacts with the next only through the Department interface, keeping them decoupled and making it easy to modify
 or extend the process without altering the high-level orchestration.

*/
