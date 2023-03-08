package bridge

type Computer interface {
	Show()
	SetDisplay(Display)
}

type Mac struct {
	display Display
}

func NewMac() Computer                    { return &Mac{} }
func (m *Mac) Show()                      { m.display.Display() }
func (m *Mac) SetDisplay(display Display) { m.display = display }

type Linux struct {
	display Display
}

func NewLinux() Computer                    { return &Linux{} }
func (l *Linux) Show()                      { l.display.Display() }
func (l *Linux) SetDisplay(display Display) { l.display = display }

type Display interface {
	Display()
}

type Benq struct{}

func NewBenq() Display   { return &Benq{} }
func (b *Benq) Display() {}

type Aoc struct{}

func NewAoc() Display   { return &Aoc{} }
func (a *Aoc) Display() {}
