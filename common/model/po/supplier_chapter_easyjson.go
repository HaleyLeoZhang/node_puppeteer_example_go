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

func easyjson8aebe6f2DecodeNodePuppeteerExampleGoCommonModelPo(in *jlexer.Lexer, out *SupplierChapter) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
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
		case "related_id":
			out.RelatedId = int(in.Int())
		case "sequence":
			out.Sequence = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "id":
			out.Id = int(in.Int())
		case "status":
			out.Status = uint8(in.Uint8())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
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
func easyjson8aebe6f2EncodeNodePuppeteerExampleGoCommonModelPo(out *jwriter.Writer, in SupplierChapter) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"related_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.RelatedId))
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
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Status))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SupplierChapter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8aebe6f2EncodeNodePuppeteerExampleGoCommonModelPo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SupplierChapter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8aebe6f2EncodeNodePuppeteerExampleGoCommonModelPo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SupplierChapter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8aebe6f2DecodeNodePuppeteerExampleGoCommonModelPo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SupplierChapter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8aebe6f2DecodeNodePuppeteerExampleGoCommonModelPo(l, v)
}
