package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

func (h Human) Speak() {
	fmt.Printf("Greetings! I am %s %s, at your service.\n", h.Name, h.Surname)
}

func (h Human) Work() {
	fmt.Printf("%s %s is working!\n", h.Name, h.Surname)
}

type Action struct {
	Human
	ActionType string
}

func (a *Action) Do() {
	switch a.ActionType {
	case "speak":
		a.Speak()
	case "work":
		a.Work()
	default:
		fmt.Printf("Unknown action: %s\n", a.ActionType)
	}
}

func main() {
	sj := Action{
		Human: Human{
			Name:    "Sergey",
			Surname: "Atryakhin",
		},
		ActionType: "speak",
	}

	sj.Do()

	sj.ActionType = "work"
	sj.Do()

	sj.ActionType = "unknown"
	sj.Do()
}
