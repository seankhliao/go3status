package mod

import (
	"time"

	"github.com/seankhliao/go3status/protocol"
)

type Time struct {
	Mod

	// Specific
	Format      string
	FormatShort string
	Color       string
}

func NewTime() Module {
	return &Time{
		Mod: NewMod("time", 1),

		Format: time.ANSIC,
	}
}

func (m *Time) NewBlock(t time.Time) *protocol.Block {
	return &protocol.Block{
		FullText: time.Now().Format(m.Format),
		Name:     m.name,
		Instance: m.Instance,
	}
}

func (m *Time) Start(blocks []*protocol.Block, pos int) {
	m.Mod.Start(blocks, pos, m.NewBlock)
}
