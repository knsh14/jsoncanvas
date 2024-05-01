package jsoncanvas

type Canvas struct {
    Nodes []Node `json:"nodes,omitempty"`
    Edges []Edge `json:"edges,omitempty"`
}
