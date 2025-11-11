// csv.go: Handler and helper function(s) for dealing with csv input
package main

import (
	"encoding/csv"
	"mime/multipart"
)

// Verifys that the provided CSV's headers are in accordance with
// databento's format. Does not confirm the underlying data.
func VerifyTradesHeaders(headers *multipart.FileHeader) (bool, int, string) {
	file, err := headers.Open()
	if err != nil {
		return false, 500, "Error opening CSV header file"
	}
	defer file.Close()

	reader := csv.NewReader(file)
	h, err := reader.Read()
	if err != nil {
		return false, 500, "Error reading CSV file"
	}

	expected := []string{"ts_recv", "ts_event", "rtype", "publisher_id", "instrument_id", "action", "side", "depth", "price", "size", "flags", "ts_in_delta", "sequence", "symbol"}
	if len(h) != len(expected) {
		return false, 404, "Not correct number of headers"
	}
	i := 0
	for i < len(h) {
		if expected[i] != h[i] {
			return false, 404, "Given headers do not match databento schema, not they must be ordered properly!"
		}
		i += 1
	}

	return true, 201, "OK"
}
