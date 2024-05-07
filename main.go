package main

import "fmt"

type SmartEquipment struct {
	Type       string
	Location   string
	Name       string
	connection bool
	Status     bool
}

type SmartHub struct {
	Equipment []SmartEquipment
}

type SmartControl interface {
	Connect()
	SwitchOn()
	SwitchOff()
	Disconnect()
}

func (s *SmartEquipment) Connect() {
	fmt.Println("Connecting to : ", s.Name)
	s.connection = true
}

func (s *SmartEquipment) SwitchOn() {
	fmt.Println("Switching on : ", s.Name)
	s.Status = true
}

func (s *SmartEquipment) SwitchOff() {
	fmt.Println("Switching off : ", s.Name)
	s.Status = false

}

func (s *SmartEquipment) Disconnect() {
	fmt.Println("Disconnecting : ", s.Name)
	s.connection = false
}

func (sh *SmartHub) listeq() {
	fmt.Println("listing")

	for _, eq := range sh.Equipment {
		fmt.Println(eq.Name, "location:", eq.Location, "Status:", eq.Status)
	}
	fmt.Println("printing length of hub :  ", len(sh.Equipment))

}

func (sh *SmartHub) add(e SmartEquipment) {
	sh.Equipment = append(sh.Equipment, e)
	fmt.Println("Added  : ", e.Name)
}

func (sh *SmartHub) DeleteEquipment(name string) {

	fmt.Println("deleteing :", name)
	for i, eq := range sh.Equipment {
		if eq.Name == name {
			sh.Equipment = append(sh.Equipment[:i], sh.Equipment[i+1:]...)
			fmt.Println("Deleted :: ", name)

		}

	}
	fmt.Println("printing length of hub afetr delete:  ", len(sh.Equipment))

}

func main() {
	hub := SmartHub{}
	hub.listeq()
	hub.add(SmartEquipment{Type: "Light", Location: "bedroom", Name: "eq1", connection: false, Status: false})
	hub.add(SmartEquipment{Type: "Fan", Location: "hall", Name: "eq2", connection: false, Status: true})
	new := SmartEquipment{Type: "Switch", Location: "bathroom", Name: "eq3", connection: true, Status: true}
	hub.add(new)
	hub.listeq()

	hub.DeleteEquipment("eq3")
	hub.Equipment[0].SwitchOn()
	hub.Equipment[0].Connect()
	hub.Equipment[1].Disconnect()

	hub.listeq()

}
