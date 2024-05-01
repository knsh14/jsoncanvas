package jsoncanvas

type CanvasColor interface {
	String() string
	implementCanvasColor()
}
type PresetColor string

func (c PresetColor) string() string        { return string(c) }
func (c PresetColor) implementCanvasColor() {}

const (
	PresetColorRed    PresetColor = "1"
	PresetColorOrange PresetColor = "2"
	PresetColorYellow PresetColor = "3"
	PresetColorGreen  PresetColor = "4"
	PresetColorCyan   PresetColor = "5"
	PresetColorPurple PresetColor = "6"
)

type HexColor string

func (c HexColor) string() string        { return string(c) }
func (c HexColor) implementCanvasColor() {}
