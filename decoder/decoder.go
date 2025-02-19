package decoder

import (
	"fmt"

	"example.com/types"
	"github.com/BertoldVdb/go-ais"
	"github.com/adrianmo/go-nmea"
)

type Decoder struct {
	pendingMessages map[int64][]byte
}

func New() *Decoder {
	return &Decoder{
		pendingMessages: make(map[int64][]byte),
	}
}

func (d *Decoder) toPacket(raw string) (ais.Packet, error) {
	sentence, err := nmea.Parse(raw)
	if err != nil {
		return nil, err
	}

	c := ais.CodecNew(false, true)

	if sentence, ok := sentence.(nmea.VDMVDO); ok {
		if sentence.NumFragments == 1 {
			packet := c.DecodePacket(sentence.Payload)
			return packet, nil
		}

		// Multi-fragment message
		firstPart := d.pendingMessages[sentence.MessageID]
		if firstPart == nil && sentence.FragmentNumber == 1 {
			// Store as pending message
			d.pendingMessages[sentence.MessageID] = sentence.Payload
		} else if firstPart != nil && sentence.FragmentNumber == 2 {
			// Combine fragments and return
			packet := c.DecodePacket(append(firstPart, sentence.Payload...))
			return packet, nil
		} else {
			// Something is wrong, just drop the message
			return nil, fmt.Errorf("received an unaccompanied fragment")
		}

		return nil, nil
	} else {
		return nil, fmt.Errorf("cannot handle non-VDM sentences")
	}
}

type LabeledPacket struct {
	Packet *ais.Packet
	Label  types.AISMessageType
}

func (d *Decoder) Decode(raw string) (*LabeledPacket, error) {
	result, err := d.toPacket(raw)

	if result != nil {
		switch r := result.(type) {
		case ais.PositionReport:
			return &LabeledPacket{
				Packet: &result,
				Label:  types.PositionReportLabel,
			}, nil
		case ais.StaticDataReport:
			return &LabeledPacket{
				Packet: &result,
				Label:  types.StaticDataReportLabel,
			}, nil
		case ais.ShipStaticData:
			return &LabeledPacket{
				Packet: &result,
				Label:  types.ShipStaticDataLabel,
			}, nil
		default:
			return nil, fmt.Errorf("cannot handle %T messages", r)
		}
	}

	return nil, err
}
