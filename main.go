package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SmartEquipment struct {
	Type       string `json:"type"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Connection bool   `json:"connection"`
	Status     bool   `json:"status"`
}

type SmartHub struct {
	Equipment []SmartEquipment `json:"equipment"`
}

type SmartControl interface {
	Connect()
	SwitchOn()
	SwitchOff()
	Disconnect()
}

func (s *SmartEquipment) Connect() {
	fmt.Println("Connecting to:", s.Name)
	s.Connection = true
}

func (s *SmartEquipment) SwitchOn() {
	fmt.Println("Switching on:", s.Name)
	s.Status = true
}

func (s *SmartEquipment) SwitchOff() {
	fmt.Println("Switching off:", s.Name)
	s.Status = false
}

func (s *SmartEquipment) Disconnect() {
	fmt.Println("Disconnecting:", s.Name)
	s.Connection = false
}

func (sh *SmartHub) listeq() ([]byte, error) {
	return json.Marshal(sh.Equipment)
}

func (sh *SmartHub) add(e SmartEquipment) {
	sh.Equipment = append(sh.Equipment, e)
	fmt.Println("Added:", e.Name)
}
func listobject(name string) ([]byte, error) {
	fmt.Println("calling list-------------------")
	for _, eq := range hub.Equipment {
		if eq.Name == name {
			fmt.Println("calling list-------------------found : ", name)

			equipmentJSON, err := json.Marshal(eq)
			if err != nil {
				//http.Error(w, "Failed to marshal equipment data", http.StatusInternalServerError)
				return nil, err
			}
			fmt.Println("calling list-------------------listing return")

			return equipmentJSON, nil
		}

	}
	fmt.Println("calling list-------------------last")

	return nil, nil
}
func (sh *SmartHub) DeleteEquipment(name string) {
	fmt.Println("Deleting:", name)
	for i, eq := range sh.Equipment {
		if eq.Name == name {
			sh.Equipment = append(sh.Equipment[:i], sh.Equipment[i+1:]...)
			fmt.Println("Deleted:", name)
			return
		}
	}
	fmt.Println("Equipment not found:", name)
}

func createEquipment(w http.ResponseWriter, r *http.Request) {
	var equipment SmartEquipment
	err := json.NewDecoder(r.Body).Decode(&equipment)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	hub.add(equipment)
	saveData()
	equipmentJSON, _ := listobject(equipment.Name)
	w.Header().Set("Content-Type", "application/json")
	w.Write(equipmentJSON)
	fmt.Println("Equipment Lists:")
}

func deleteEquipment(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is missing", http.StatusBadRequest)
		return
	}

	hub.DeleteEquipment(name)
	saveData()
	listobject(name)

}

func listEquipment(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var equipmentJSON []byte
	var err error
	if name != "" {
		equipmentJSON, err = listobject(name)

	} else {
		equipmentJSON, err = hub.listeq()
	}
	if err != nil {
		http.Error(w, "Failed to list equipment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(equipmentJSON)
	fmt.Println("Equipment List:")
	/*for _, eq := range hub.Equipment {
		fmt.Printf("Name: %s, Type: %s, Location: %s, Connection: %t, Status: %t\n", eq.Name, eq.Type, eq.Location, eq.Connection, eq.Status)
	}
	*/

}

/*
	func listEquipment(w http.ResponseWriter, r *http.Request) {
		// Check if the request is coming from an API request
		if r != nil {
			// API request
			equipmentJSON, err := hub.listeq()
			if err != nil {
				http.Error(w, "Failed to list equipment", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(equipmentJSON)
		} else {
			// Non-API request, print JSON data to terminal
			fmt.Println("Equipment List:")
			for _, eq := range hub.Equipment {
				fmt.Printf("Name: %s, Type: %s, Location: %s, Connection: %t, Status: %t\n", eq.Name, eq.Type, eq.Location, eq.Connection, eq.Status)
			}
		}
	}
*/
var hub SmartHub

func saveData() {
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Error creating data.json:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&hub)
	if err != nil {
		fmt.Println("Error encoding data.json:", err)
		return
	}
}

//here

func main() {
	_, err := os.Stat("data.json")
	if os.IsNotExist(err) {
		hub = SmartHub{Equipment: []SmartEquipment{}}
	} else {
		file, err := os.Open("data.json")
		if err != nil {
			fmt.Println("Error opening data.json:", err)
			return
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&hub)
		if err != nil {
			fmt.Println("Error decoding data.json:", err)
			return
		}
	}

	http.HandleFunc("/equipment/create", createEquipment)
	http.HandleFunc("/equipment/delete", deleteEquipment)
	http.HandleFunc("/equipment/list", listEquipment)

	fmt.Println("Server is running...")
	http.ListenAndServe(":8077", nil)
}
