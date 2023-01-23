package main

import (
	"fmt"
	"os"

	"github.com/felixge/traceutils/pkg/encoding"
	"github.com/felixge/traceutils/pkg/filter"
)

func FilterCommand(args []string) error {
	// Check the number of arguments
	if len(args) != 2 {
		return fmt.Errorf("expected 2 arguments, got %d", len(args))
	}

	// Open the input file
	inFile, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inFile.Close()

	// Open the output file
	outFile, err := os.Create(args[1])
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}
	defer outFile.Close()

	// Filter the trace file
	include := []encoding.EventType{
		encoding.EventFrequency,
		encoding.EventGCStart,
		encoding.EventGCDone,
		encoding.EventStack,
		encoding.EventString,
		encoding.EventUserTaskCreate,
		encoding.EventUserTaskEnd,
		encoding.EventUserRegion,
		encoding.EventUserLog,
		encoding.EventCPUSample,
		encoding.EventCount,
	}
	return filter.FilterTrace(inFile, outFile, include)
}
