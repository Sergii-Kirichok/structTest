package module

type Module struct {
	vendor    string //vendor name. Ex.Optex, etc
	statusCur string //Current Status. Ex. Started, Processing, Paused, Stopped, Restarted
	statusAsk string //Start, Stop, Restart
}

func NewModule() *Module{
	return &Module{}
}

func (m *Module) SetVendorName(s string){ m.vendor = s}
func (m *Module) SetStatusCur(s string){ m.statusCur = s}
func (m *Module) SetStatusAsk(s string){ m.statusAsk = s}

func (m Module) VendorName(s string) string { return m.vendor }
func (m Module) StatCur(s string)  string { return m.statusCur }
func (m Module) StatAsk(s string)  string { return m.statusAsk }