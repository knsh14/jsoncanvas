package jsoncanvas

type NodeType string

const (
	NodeText  NodeType = "text"
	NodeFile  NodeType = "file"
	NodeLink  NodeType = "link"
	NodeGroup NodeType = "group"
)

type Node struct {
	ID     string      `json:"id"`
	Type   NodeType    `json:"type"`
	X      int         `json:"x"`
	Y      int         `json:"y"`
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Color  CanvasColor `json:"color,omitempty"`
}

type TextNode struct {
	Node
	Text string `json:"text"`
}
type FileNode struct {
	Node
	File    string `json:"file"`
	Subpath string `json:"subpath,omitempty"` // always prefix is "#" TODO: add validation
}
type LinkNode struct {
	Node
	URL string `json:"url"`
}
type GroupNode struct {
	Node
	Label           string              `json:"label,omitempty"`
	Background      string              `json:"background,omitempty"`
	BackgroundStyle BackgroundStyleType `json:"backgroundStyle,omitempty"`
}

type BackgroundStyleType string

const (
	BackgroundStyleTypeCover  BackgroundStyleType = "cover"
	BackgroundStyleTypeRatio  BackgroundStyleType = "ratio"
	BackgroundStyleTypeRepeat BackgroundStyleType = "repeat"
)
