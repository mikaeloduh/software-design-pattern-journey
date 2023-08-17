package entity

// ConnectTelecomCommand
type ConnectTelecomCommand struct {
	Telecom *Telecom
}

func (c ConnectTelecomCommand) Execute() {
	c.Telecom.Connect()
}

func (c ConnectTelecomCommand) Undo() {
	c.Telecom.Disconnect()
}

// DisconnectTelecomCommand
type DisconnectTelecomCommand struct {
	Telecom *Telecom
}

func (c DisconnectTelecomCommand) Execute() {
	c.Telecom.Disconnect()
}

func (c DisconnectTelecomCommand) Undo() {
	c.Telecom.Connect()
}
