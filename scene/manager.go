package scene

// Manager ..
type Manager struct {
	currentScene Scene
	sceneList    []Scene
	renderCtx    *RenderContext
}

// NewManager ..
func NewManager(initialScene Scene, r *RenderContext) *Manager {
	if initialScene != nil {
		initialScene.Init(r)
	}

	return &Manager{
		currentScene: initialScene,
		sceneList:    []Scene{},
		renderCtx:    r,
	}
}

// SetCurrentScene ..
func (sm *Manager) SetCurrentScene(scene Scene) {
	sm.currentScene = scene
}

// AddScene ..
func (sm *Manager) AddScene(scene Scene) {
	if scene != nil {
		scene.Init(sm.renderCtx)
		sm.sceneList = append(sm.sceneList, scene)
	}

}

// UpdateInput ..
func (sm *Manager) UpdateInput() {
	// check for switching scenes

	// update input for current scene
	if sm.currentScene != nil {
		sm.currentScene.UpdateInput(sm.renderCtx)
	}

	// draw updates to shapes
	// this needs to happen before rendering sprites to avoid accidental erase
	r := sm.renderCtx
	r.imd.Draw(r.win)

}

// Render ..
func (sm *Manager) Render() {
	win := sm.renderCtx.win

	// check for switching scenes
	if sm.currentScene != nil {
		sm.currentScene.Render(sm.renderCtx)
	}

	// sm.renderCtx.imd.Draw(win)
	win.Update()
}
