package scene

// Scene..
type Scene interface {
	Init(*RenderContext)
	UpdateInput(*RenderContext)
	Render(*RenderContext)
}
