package main

import (
	"fmt"
)

type Square struct {
	Volume                Word
	SawEnvelopeDisabled   bool
	LengthCounterDisabled bool
	DutyCycle             Word
	LowPeriod             Word
	HighPeriod            Word
	LengthCounter         Word
}

type Triangle struct {
	Value                    Word
	InternalCountersDisabled bool
	LowPeriod                Word
	HighPeriod               Word
	LengthCounter            Word
}

type Apu struct {
	Square1Enabled  bool
	Square2Enabled  bool
	TriangleEnabled bool
	NoiseEnabled    bool
	DmcEnabled      bool
	Square1         Square
	Square2         Square
	Triangle
}

func (a *Apu) Init() {
}

func (a *Apu) RegRead(addr int) (Word, error) {
	switch addr {
	case 0x4015:
		return a.ReadStatus(), nil
	}

	return 0, nil
}

func (a *Apu) RegWrite(v Word, addr int) {
	fmt.Printf("APU RegWrite: 0x%X\n", addr)
	switch addr & 0xFF {
	case 0x0:
		a.WriteSquare1Control(v)
	case 0x1:
		a.WriteSquare1Sweeps(v)
	case 0x2:
		a.WriteSquare1Low(v)
	case 0x3:
		a.WriteSquare1High(v)
	case 0x4:
		a.WriteSquare2Control(v)
	case 0x5:
		a.WriteSquare2Sweeps(v)
	case 0x6:
		a.WriteSquare2Low(v)
	case 0x7:
		a.WriteSquare2High(v)
	case 0x8:
		a.WriteTriangleControl(v)
	case 0xA:
		a.WriteTriangleLow(v)
	case 0xB:
		a.WriteTriangleHigh(v)
	case 0x15:
		a.WriteControlFlags1(v)
	case 0x17:
		a.WriteControlFlags2(v)
	}
}

// $4015 (w)
func (a *Apu) WriteControlFlags1(v Word) {
	// 76543210
	//    |||||
	//    ||||+- Square 1 (0: disable; 1: enable)
	//    |||+-- Square 2
	//    ||+--- Triangle
	//    |+---- Noise
	//    +----- DMC
	a.Square1Enabled = (v & 0x1) == 0x1
	a.Square2Enabled = ((v >> 1) & 0x1) == 0x1
	a.TriangleEnabled = ((v >> 2) & 0x1) == 0x1
	a.NoiseEnabled = ((v >> 3) & 0x1) == 0x1
	a.DmcEnabled = ((v >> 4) & 0x1) == 0x1
}

// $4015 (r)
func (a *Apu) ReadStatus() Word {
	// if-d nt21   DMC IRQ, frame IRQ, length counter statuses
	return 0
}

// $4017
func (a *Apu) WriteControlFlags2(v Word) {
	// fd-- ----   5-frame cycle, disable frame interrupt
	fmt.Println("WriteControl2!")
}

// $4000
func (a *Apu) WriteSquare1Control(v Word) {
	// 76543210
	// ||||||||
	// ||||++++- Volume
	// |||+----- Saw Envelope Disable (0: use internal counter for volume; 1: use Volume for volume)
	// ||+------ Length Counter Disable (0: use Length Counter; 1: disable Length Counter)
	// ++------- Duty Cycle
	a.Square1.Volume = v & 0xF
	a.Square1.SawEnvelopeDisabled = (v>>4)&0x1 == 1
	a.Square1.LengthCounterDisabled = (v>>5)&0x1 == 1
	a.Square1.DutyCycle = (v >> 6) & 0x3
}

// $4001
func (a *Apu) WriteSquare1Sweeps(v Word) {
}

// $4002
func (a *Apu) WriteSquare1Low(v Word) {
	a.Square1.LowPeriod = v
}

// $4003
func (a *Apu) WriteSquare1High(v Word) {
	a.Square1.HighPeriod = v & 0xF
	a.Square1.LengthCounter = v >> 3
}

// $4004
func (a *Apu) WriteSquare2Control(v Word) {
	// 76543210
	// ||||||||
	// ||||++++- Volume
	// |||+----- Saw Envelope Disable (0: use internal counter for volume; 1: use Volume for volume)
	// ||+------ Length Counter Disable (0: use Length Counter; 1: disable Length Counter)
	// ++------- Duty Cycle
	a.Square2.Volume = v & 0xF
	a.Square2.SawEnvelopeDisabled = (v>>4)&0x1 == 1
	a.Square2.LengthCounterDisabled = (v>>5)&0x1 == 1
	a.Square2.DutyCycle = (v >> 6) & 0x3
}

// $4005
func (a *Apu) WriteSquare2Sweeps(v Word) {
}

// $4006
func (a *Apu) WriteSquare2Low(v Word) {
	a.Square2.LowPeriod = v
}

// $4007
func (a *Apu) WriteSquare2High(v Word) {
	a.Square2.HighPeriod = v & 0xF
	a.Square2.LengthCounter = v >> 3
}

// $4008
func (a *Apu) WriteTriangleControl(v Word) {
	// 76543210
	// ||||||||
	// |+++++++- Value
	// +-------- Control Flag (0: use internal counters; 1: disable internal counters)
	a.Triangle.Value = v & 0x7F
	a.Triangle.InternalCountersDisabled = (v>>7)&0x1 == 1
}

// $400A
func (a *Apu) WriteTriangleLow(v Word) {
	a.Triangle.LowPeriod = v
}

// $400B
func (a *Apu) WriteTriangleHigh(v Word) {
	a.Triangle.HighPeriod = v & 0xF
	a.Triangle.LengthCounter = v >> 3
}