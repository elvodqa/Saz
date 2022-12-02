package ui

type UI interface {
	Update()
	Draw()
	OnResize()
}

type Container struct {
	Elements []UI
}

type Panel struct {
}
