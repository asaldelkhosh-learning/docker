package main

func main() {
	inp := Input{}
	inp = inp.Init()
	Init()

	for true {
		cmd := inp.Get()

		switch {
		case cmd == "new":
			// New process function
		case cmd == "kill":
			// Kill process
		case cmd == "monitor":
			// Monitoring
		case cmd == "terminate":
			break
		}
	}
}
