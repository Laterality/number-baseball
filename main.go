package main

func main() {
	controller := NewGameSession()
	for !controller.isFinished() {
		controller.proceed()
	}
}
