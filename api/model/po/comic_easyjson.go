// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package po

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7950832dDecodeNodePuppeteerExampleGoApiModelPo(in *jlexer.Lexer, out *Comic) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	out.Model = new(Model)
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "channel":
			out.Channel = int(in.Int())
		case "source_id":
			out.SourceID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "pic":
			out.Pic = string(in.String())
		case "intro":
			out.Intro = string(in.String())
		case "max_sequence":
			out.MaxSequence = int(in.Int())
		case "weight":
			out.Weight = int(in.Int())
		case "tag":
			out.Tag = int(in.Int())
		case "id":
			out.ID = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7950832dEncodeNodePuppeteerExampleGoApiModelPo(out *jwriter.Writer, in Comic) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"channel\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Channel))
	}
	{
		const prefix string = ",\"source_id\":"
		out.RawString(prefix)
		out.Int(int(in.SourceID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"pic\":"
		out.RawString(prefix)
		out.String(string(in.Pic))
	}
	{
		const prefix string = ",\"intro\":"
		out.RawString(prefix)
		out.String(string(in.Intro))
	}
	{
		const prefix string = ",\"max_sequence\":"
		out.RawString(prefix)
		out.Int(int(in.MaxSequence))
	}
	{
		const prefix string = ",\"weight\":"
		out.RawString(prefix)
		out.Int(int(in.Weight))
	}
	{
		const prefix string = ",\"tag\":"
		out.RawString(prefix)
		out.Int(int(in.Tag))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int(int(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Comic) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7950832dEncodeNodePuppeteerExampleGoApiModelPo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comic) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7950832dEncodeNodePuppeteerExampleGoApiModelPo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comic) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7950832dDecodeNodePuppeteerExampleGoApiModelPo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comic) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7950832dDecodeNodePuppeteerExampleGoApiModelPo(l, v)
}
