package main

//https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go

// simple
type EBase int

const (
	EA EBase = iota
	EC
	ET
	EG
)

// complex
type EnumItem struct {
	index int
	name  string
}

type Enum struct {
	items []EnumItem
}

func (enum Enum) Name(findIndex int) string {
	for _, item := range enum.items {
		if item.index == findIndex {
			return item.name
		}
	}
	return "ID not found"
}

func (enum Enum) Index(findName string) int {
	for idx, item := range enum.items {
		if findName == item.name {
			return idx
		}
	}
	return -1
}

func (enum Enum) Last() (int, string) {
	n := len(enum.items)
	return n - 1, enum.items[n-1].name
}

var AgentTypes = Enum{[]EnumItem{{0, "StaffMember"}, {1, "Organization"}, {1, "Automated"}}}
var AccountTypes = Enum{[]EnumItem{{0, "Basic"}, {1, "Advanced"}}}
var FlagTypes = Enum{[]EnumItem{{0, "Custom"}, {1, "System"}}}

func main() {
	s := AgentTypes.Name(1)
	print(s)
}
