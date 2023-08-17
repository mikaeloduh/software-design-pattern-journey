package main

// ConnectTelecomCommand
type ConnectTelecomCommand struct {
	telecom *Telecom
}

func (c ConnectTelecomCommand) Execute() {
	c.telecom.Connect()
}

func (c ConnectTelecomCommand) Undo() {
	c.telecom.Disconnect()
}

// DisconnectTelecomCommand
type DisconnectTelecomCommand struct {
	telecom *Telecom
}

func (c DisconnectTelecomCommand) Execute() {
	c.telecom.Disconnect()
}

func (c DisconnectTelecomCommand) Undo() {
	c.telecom.Connect()
}
