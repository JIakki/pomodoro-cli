package controllers

func (ctrl *Controllers) Start() chan int {
	c := make(chan int)

	go ctrl.pomoInstance.Start(c)

	return c
}
