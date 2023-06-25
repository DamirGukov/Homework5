package main

import "fmt"

type Package interface {
	GetSenderAdress() string
	GetRecipientAdress() string
	Send()
}

type Box struct {
	AdressSender    string
	AdressRecipient string
}

type Envelope struct {
	AdressSender    string
	AdressRecipient string
}

func (b Box) GetSenderAdress() string {
	return b.AdressSender
}

func (b Box) GetRecipientAdress() string {
	return b.AdressRecipient
}

func (b Box) Send() {
	fmt.Println("Sending box to:", b.GetRecipientAdress(), "from:", b.GetSenderAdress())
}

func (e Envelope) GetSenderAdress() string {
	return e.AdressSender
}

func (e Envelope) GetRecipientAdress() string {
	return e.AdressRecipient
}

func (e Envelope) Send() {
	fmt.Println("Sending envelope to:", e.GetRecipientAdress(), "from:", e.GetSenderAdress())
}

type SortingDepartment struct{}

func (sd SortingDepartment) SortAndSend(p Package, delivery string, deliverySave map[int]int) {

	fmt.Println("Fast or regular delivery?")
	fmt.Scan(&delivery)

	if delivery == "Fast" || delivery == "fast" {
		fmt.Println("Sorting box, and sending it by fast deliver")
		deliverySave[1]++
		p.Send()
	} else {
		fmt.Println("Sorting box, and sending it by regular deliver")
		deliverySave[2]++
		p.Send()
	}

	fmt.Println("Fast or regular delivery?")
	fmt.Scan(&delivery)

	if delivery == "fast" || delivery == "Fast" {
		fmt.Println("Sorting envelope, and sending it by fast deliver")
		deliverySave[1]++
		p.Send()
	} else {
		fmt.Println("Sorting box, and sending it by regular deliver")
		deliverySave[2]++
		p.Send()
	}
}

func main() {
	b := Box{
		AdressSender:    "12 High Street, London, SW1A 1AA",
		AdressRecipient: "27 Park Avenue, Manchester, M14 5PT",
	}

	e := Envelope{
		AdressSender:    "8 Queen's Road, Birmingham, B1 1RD",
		AdressRecipient: "45 Windsor Gardens, Edinburgh, EH1 2HU",
	}
	var delivery string
	deliverySave := map[int]int{
		1: 0,
		2: 0,
	}

	sd := SortingDepartment{}
	sd.SortAndSend(b, delivery, deliverySave)
	sd.SortAndSend(e, delivery, deliverySave)

	fmt.Println("Fast delivery: ", deliverySave[1])
	fmt.Println("Regular delivery: ", deliverySave[2])
}
