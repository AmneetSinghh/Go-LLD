package main

import "fmt"

/*
📌 Template Method Pattern
✅ What It Is:
The Template Method Pattern defines a fixed algorithm ( fixed sequence enforce krvana va ) structure in a base class, while allowing subclasses to override specific steps without modifying the overall process.

✅ When to Use:

When a process has a fixed sequence of steps, but some steps need customization.
When you want to enforce a strict execution order while allowing variations.
When you need to avoid code duplication by reusing a structured workflow across multiple implementations.
When you want to follow the Open/Closed Principle (OCP) by adding new behaviors via subclasses without modifying existing code.
🚀 Best for scenarios like report generation

*/

type IOtp interface {
	getRandomOtp(int) string
	saveOtpCache(string)
	getMessage(string) string
	sendNotifications(string) error
}

type Otp struct {
	Iotp IOtp // its child clases.
}

// this is in the base class, so al child will use same, this doest know on which otp ist is performinng
func (o *Otp) getAndSetRandomOtp(OtpLenth int) error {
	otp := o.Iotp.getRandomOtp(OtpLenth)
	o.Iotp.saveOtpCache(otp)
	message := o.Iotp.getMessage(otp)
	err := o.Iotp.sendNotifications(message)
	if err != nil {
		return err
	}
	return nil
}

type Sms struct {
	Otp
}

func (s *Sms) getRandomOtp(len int) string {
	randomOtp := "1234"
	fmt.Println("SMs: generatn romdom otp ", randomOtp)
	return randomOtp
}
func (s *Sms) saveOtpCache(otp string) {
	fmt.Println("SMs: save into cache ")
}

func (s *Sms) getMessage(otp string) string {
	return "SMS:: otp for login is :  " + otp
}

func (s *Sms) sendNotifications(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func (s *Sms) getAndSetRandomOtp(OtpLenth int) error { // we have't exposed this method in inteface, so cliet will not have visibility to this.
	fmt.Printf("internal calling", OtpLenth)
	return nil
}

type Email struct {
	//Otp
}

func (s *Email) getRandomOtp(len int) string {
	randomOtp := "12345"
	fmt.Println("Email: generatn romdom otp ", randomOtp)
	return randomOtp
}
func (s *Email) saveOtpCache(otp string) {
	fmt.Println("Email: save into cache ")
}

func (s *Email) getMessage(otp string) string {
	return "Email:: otp for login is :  " + otp
}

func (s *Email) sendNotifications(message string) error {
	fmt.Printf("Email: sending sms: %s\n", message)
	return nil
}

func main() {
	var smsOtp IOtp = &Sms{}
	o := Otp{
		Iotp: smsOtp,
	}

	o.getAndSetRandomOtp(10)

	var emailOtp IOtp = &Sms{}
	o = Otp{
		Iotp: emailOtp,
	}

	o.getAndSetRandomOtp(10)

}

/*

 Main Takeaway:

The Otp base class is responsible for executing the common flow.
Each OTP type (SMS, Email, WhatsApp) overrides only the steps that differ.
This follows the Template Method Pattern because the execution sequence is fixed, but the implementation of individual steps varies.



- IMPORTANT -

- we can't add getAndSetRandomOtp(int) in this interface, then eac child can enforce change the implementation......, contorl is our the childclass...


✅ Acts like an inherited method.
✅ Template Method Pattern (Fixed execution order).
*/
