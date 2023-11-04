// Package point is a package that defines the points behavior and handles them.
package point

var instance *Handler

func init() {
	instance = &Handler{
		points: make(map[string]point),
	}
}

// Handler is a struct that holds the points.
type Handler struct {
	points map[string]point
}

// Instance returns the instance of the handler.
func Instance() *Handler {
	return instance
}

// AddPoint adds a point to the handler.
func (h *Handler) AddPoint(name string, p point) {
	h.points[name] = p
}

// RemovePoint removes the point with the given name.
func (h *Handler) RemovePoint(name string) {
	delete(h.points, name)
}

// ExecutePoint executes the point with the given name.
func (h *Handler) ExecutePoint(name string) {
	if p, ok := h.points[name]; ok {
		p.Execute()
	}
}

// ResumePoint resumes the point with the given name.
func (h *Handler) ResumePoint(name string) {
	if p, ok := h.points[name]; ok {
		p.Resume()
	}
}
