package encoding

import (
	"errors"

	"github.com/tierklinik-dobersberg/events-service/internal/automation/common"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// TextEncoder represents an encoder that will generate a byte stream
// with UTF-8 encoding.
type TextEncoder struct {
	// Encoding always holds the `utf-8` value.
	// FIXME: this should be TextEncoder.prototype.encoding instead
	Encoding EncodingName

	encoder encoding.Encoding
}

// NewTextEncoder returns a new TextEncoder object instance that will
// generate a byte stream with UTF-8 encoding.
func NewTextEncoder() *TextEncoder {
	return &TextEncoder{
		encoder:  unicode.UTF8,
		Encoding: UTF8EncodingFormat,
	}
}

// Encode takes a string as input and returns an encoded byte stream.
func (te *TextEncoder) Encode(text string) ([]byte, error) {
	if te.encoder == nil {
		return nil, errors.New("encoding not set")
	}

	enc := te.encoder.NewEncoder()
	encoded, err := enc.Bytes([]byte(text))
	if err != nil {
		return nil, common.NewError(common.TypeError, "unable to encode text; reason: "+err.Error())
	}

	return encoded, nil
}
