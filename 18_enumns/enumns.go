package main

import "fmt"

// Enum-like constants for status in int
type Status int

const (
	pending Status = iota
	approved
	rejected
	reviewing
	recived
)

func changedStatus(status Status) {
	fmt.Println("Status changed to:", status)

}

// Enum-like constants for status in string
type StatusString string

const (
	pendingString   StatusString = "pending"
	approvedString  StatusString = "approved"
	rejectedString  StatusString = "rejected"
	reviewingString StatusString = "reviewing"
	recivedString   StatusString = "recived"
)

func changedStatusString(status StatusString) {
	fmt.Println("Status changed to:", status)
}

func main() {
	changedStatus(approved) // changed status to: 1

	changedStatusString(recivedString)

}
