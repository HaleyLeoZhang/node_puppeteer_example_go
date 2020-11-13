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

func easyjson7d177735DecodeNodePuppeteerExampleGoApiModelPo(in *jlexer.Lexer, out *ComicPage) {
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
			out.SourceId = int(in.Int())
		case "sequence":
			out.Sequence = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "progress":
			out.Progress = int(in.Int())
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
func easyjson7d177735EncodeNodePuppeteerExampleGoApiModelPo(out *jwriter.Writer, in ComicPage) {
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
		out.Int(int(in.SourceId))
	}
	{
		const prefix string = ",\"sequence\":"
		out.RawString(prefix)
		out.Int(int(in.Sequence))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"progress\":"
		out.RawString(prefix)
		out.Int(int(in.Progress))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int(int(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ComicPage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d177735EncodeNodePuppeteerExampleGoApiModelPo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ComicPage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d177735EncodeNodePuppeteerExampleGoApiModelPo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ComicPage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d177735DecodeNodePuppeteerExampleGoApiModelPo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ComicPage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d177735DecodeNodePuppeteerExampleGoApiModelPo(l, v)
}
