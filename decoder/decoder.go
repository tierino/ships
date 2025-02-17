package decoder

import (
	"fmt"

	"example.com/types"
	"github.com/BertoldVdb/go-ais"
	"github.com/adrianmo/go-nmea"
)

func Decode(raw string) (*types.AISMessage, error) {
	sentence, err := nmea.Parse(raw)
	if err != nil {
		return nil, err
	}

	c := ais.CodecNew(false, true)

	if sentence, ok := sentence.(nmea.VDMVDO); ok {
		result := c.DecodePacket(sentence.Payload)

		if result == nil {
			// handle multipart message
			return nil, nil
		}

		switch r := result.(type) {
		case ais.PositionReport:
			return &types.AISMessage{
				Type:           types.PositionReportType,
				PositionReport: &r,
			}, nil
		case ais.StaticDataReport:
			return &types.AISMessage{
				Type:             types.StaticDataReportType,
				StaticDataReport: &r,
			}, nil
		default:
			return nil, fmt.Errorf("cannot handle %T messages", r)
		}
	}

	return nil, fmt.Errorf("cannot handle %T sentences", sentence)
}
