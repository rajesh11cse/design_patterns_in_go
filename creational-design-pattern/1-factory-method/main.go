package main

import "fmt"

type iPlan interface {
	setName(name string)
	setRate(power int)
	getName() string
	getRate() int
}

type plan struct {
	name string
	rate int
}

func (g *plan) setName(name string) {
	g.name = name
}

func (g *plan) getName() string {
	return g.name
}

func (g *plan) setRate(power int) {
	g.rate = power
}

func (g *plan) getRate() int {
	return g.rate
}

type domestic struct {
	plan
}
type commercial struct {
	plan
}

func domesticPlan() iPlan {
	return &domestic{
		plan: plan{
			name: "Domestic plan",
			rate: 5,
		},
	}
}

func commercialPlan() iPlan {
	return &commercial{
		plan: plan{
			name: "Commercial plan",
			rate: 7,
		},
	}
}

// Factory method
func getPlanFactory(plan string) (iPlan, error) {
	if plan == "domestic plan" {
		return domesticPlan(), nil
	}
	if plan == "commercial plan" {
		return commercialPlan(), nil
	}
	return nil, fmt.Errorf("Wrong plan type passed")
}

func main() {
	plan1, _ := getPlanFactory("domestic plan")   // Client 1
	plan2, _ := getPlanFactory("commercial plan") // Client 2
	printDetails(plan1)
	printDetails(plan2)
}

func printDetails(g iPlan) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getRate())
	fmt.Println()
}
