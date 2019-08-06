package storage

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/blokje5/rdbms/pkg/test_util"
)

func Test_parseHeader(t *testing.T) {
	tests := []struct {
		name   string
		args   []byte
		want   header
	}{
		{ "parseHeader with directory header page set", []byte{0, 0, 12, 1}, header { DirectoryHeaderPageType, 12, 1 }, },
		{ "parseHeader with default page set", []byte{1, 0, 12, 1}, header { DefaultPageType, 12, 1 }, },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseHeader(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("header.parseHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_header_writeBytes(t *testing.T) {
	type fields struct {
		pageType PageType
		pageSize uint16
		nextPage uint8
	}
	tests := []struct {
		name    string
		fields  fields
		wantB   []byte
		wantErr bool
	}{
		{ "header_writeBytes with directory header page", fields { DirectoryHeaderPageType, 12, 1 }, []byte{0, 0, 12, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := header{
				pageType: tt.fields.pageType,
				pageSize: tt.fields.pageSize,
				nextPage: tt.fields.nextPage,
			}
			w := &bytes.Buffer{}
			if err := h.writeBytes(w); (err != nil) != tt.wantErr {
				t.Errorf("header.writeBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.Equals(t, w.Bytes(), tt.wantB)
		})
	}
}
