package filter

import (
	"bufio"
	"io"

	"github.com/felixge/traceutils/pkg/encoding"
)

// FilterTrace reads a runtime/trace file from r and writes a filtered
// version of it to w. The filtration is done by slipping all the events
// apart from the ones that are explicitly mentioned.
// On success FilterTrace returns nil.
// If an error occurs, FilterTrace returns the error.
func FilterTrace(r io.Reader, w io.Writer, include []encoding.EventType) error {
	// Initialize encoder and decoder
	buf := bufio.NewWriter(w)
	enc := encoding.NewEncoder(buf)
	dec := encoding.NewDecoder(r)

	// Obfuscate all string events
	var ev encoding.Event
	for {
		// Decode event
		if err := dec.Decode(&ev); err != nil {
			if err == io.EOF {
				// We're done
				return buf.Flush()
			}
			return err
		}

		// Check if the event is in the filter
		if err := checkEvent(enc, ev, include); err != nil {
			return err
		}
	}
}

// checkEvent checks if the event is in the filter and writes it to the encoder if it is.
func checkEvent(enc *encoding.Encoder, ev encoding.Event, include []encoding.EventType) error {
	for _, eventType := range include {
		if ev.Type == eventType {
			// Encode the filtered event
			return enc.Encode(&ev)
		}
	}

	return nil
}
