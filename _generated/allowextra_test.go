package _generated

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestAllowExtraEncodeDecodeNormal(t *testing.T) {
	e := &AllowExtra{Foo: "a", Bar: "b"}

	var buf bytes.Buffer
	w := msgp.NewWriter(&buf)
	if err := e.EncodeMsg(w); err != nil {
		t.Error(err)
	}
	w.Flush()

	var out AllowExtra
	rdr := msgp.NewReader(&buf)
	if err := (&out).DecodeMsg(rdr); err != nil {
		t.Error(err)
	}
	if out != *e {
		t.Error("unexpected output")
	}
}

func TestAllowExtraMarshalUnmarshalNormal(t *testing.T) {
	e := &AllowExtra{Foo: "a", Bar: "b"}

	bts, err := e.MarshalMsg(nil)
	if err != nil {
		t.Error(err)
	}

	var out AllowExtra
	if _, err := (&out).UnmarshalMsg(bts); err != nil {
		t.Error(err)
	}
	if out != *e {
		t.Error("unexpected output")
	}
}

func TestAllowExtraDecodeExtra(t *testing.T) {
	e := &AllowExtra{Foo: "a", Bar: "b"}

	var out AllowExtra
	var buf bytes.Buffer
	buf.Write([]byte{148, 161, 97, 161, 98, 161, 99, 161, 100})
	rdr := msgp.NewReader(&buf)
	if err := (&out).DecodeMsg(rdr); err != nil {
		t.Error(err)
	}
	if out != *e {
		t.Error("unexpected output")
	}
	assertNoUnread(t, rdr)
}

func TestAllowExtraUnmarshalExtra(t *testing.T) {
	e := &AllowExtra{Foo: "a", Bar: "b"}

	var out AllowExtra
	bts := []byte{148, 161, 97, 161, 98, 161, 99, 161, 100, 1}
	rest, err := (&out).UnmarshalMsg(bts)
	if err != nil {
		t.Error(err)
	}
	if out != *e {
		t.Error("unexpected output")
	}
	if !bytes.Equal(rest, []byte{1}) {
		t.Error("trailing bytes missing")
	}
}

func TestAllowExtraDecodeExtraEmpty(t *testing.T) {
	e := &AllowExtraEmpty{}

	var out AllowExtraEmpty
	var buf bytes.Buffer
	buf.Write([]byte{148, 161, 97, 161, 98, 161, 99, 161, 100})
	rdr := msgp.NewReader(&buf)
	if err := (&out).DecodeMsg(rdr); err != nil {
		t.Error(err)
	}
	if out != *e {
		t.Error("unexpected output")
	}
	assertNoUnread(t, rdr)
}

func TestAllowExtraUnmarshalExtraEmpty(t *testing.T) {
	e := &AllowExtraEmpty{}

	var out AllowExtraEmpty
	bts := []byte{148, 161, 97, 161, 98, 161, 99, 161, 100, 1}
	rest, err := (&out).UnmarshalMsg(bts)
	if err != nil {
		t.Error(err)
	}
	if out != *e {
		t.Error("unexpected output")
	}
	if !bytes.Equal(rest, []byte{1}) {
		t.Error("trailing bytes missing")
	}
}

func TestAllowExtraDecodeExtraChild(t *testing.T) {
	e := &AllowExtraChild{
		Before: "1",
		Extra: &AllowExtra{
			Foo: "a",
			Bar: "b",
		},
		After: "1",
	}

	bts := []byte{
		147,     // Object header
		161, 49, // Before
		148,              // Extra header
		161, 97, 161, 98, // Real fields
		161, 99, 161, 100, // Junk fields (ignored)
		161, 49, // After
	}

	var out AllowExtraChild
	rdr := msgp.NewReader(bytes.NewReader(bts))
	if err := (&out).DecodeMsg(rdr); err != nil {
		t.Error(err)
	}
	if out.Before != e.Before {
		t.Error("unexpected before")
	}
	if out.After != e.After {
		t.Error("unexpected before")
	}
	if *(out.Extra) != *(e.Extra) {
		t.Error("unexpected extra")
	}
	assertNoUnread(t, rdr)
}

func TestAllowExtraUnmarshalExtraChild(t *testing.T) {
	e := &AllowExtraChild{
		Before: "1",
		Extra: &AllowExtra{
			Foo: "a",
			Bar: "b",
		},
		After: "1",
	}

	bts := []byte{
		147,     // Object header
		161, 49, // Before
		148,              // Extra header
		161, 97, 161, 98, // Real fields
		161, 99, 161, 100, // Junk fields (ignored)
		161, 49, // After
		1, // Trailing byte
	}

	var out AllowExtraChild
	rest, err := (&out).UnmarshalMsg(bts)
	if err != nil {
		t.Error(err)
	}
	if out.Before != e.Before {
		t.Error("unexpected before")
	}
	if out.After != e.After {
		t.Error("unexpected before")
	}
	if *(out.Extra) != *(e.Extra) {
		t.Error("unexpected extra")
	}
	if !bytes.Equal(rest, []byte{1}) {
		t.Error("trailing bytes missing")
	}
}

func assertNoUnread(t *testing.T, rdr *msgp.Reader) {
	t.Helper()
	rest, err := ioutil.ReadAll(rdr)
	if err != nil {
		t.Error(err)
	}
	if len(rest) > 0 {
		t.Error("unexpected trailing")
	}
}
