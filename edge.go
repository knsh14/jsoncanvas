package jsoncanvas

type EdgeSide string

const (
	EdgeSideTop    EdgeSide = "top"
	EdgeSideRight  EdgeSide = "right"
	EdgeSideBottom EdgeSide = "bottom"
	EdgeSideLeft   EdgeSide = "left"
)

type EdgeEnd string

const (
	EdgeEndNone  EdgeEnd = "none"
	EdgeEndArrow EdgeEnd = "arrow"
)

type Edge struct {
	ID       string      `json:"id"`
	FromNode string      `json:"fromNode"`
	FromSide EdgeSide    `json:"fromSide,omitempty"`
	FromEnd  EdgeEnd     `json:"fromEnd,omitempty"`
	ToNode   string      `json:"toNode"`
	ToSide   EdgeSide    `json:"toSide,omitempty"`
	ToEnd    EdgeEnd     `json:"toEnd,omitempty"`
	Color    CanvasColor `json:"color,omitempty"`
	Label    string      `json:"label,omitempty"`
}
