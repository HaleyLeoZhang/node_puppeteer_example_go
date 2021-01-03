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

func easyjson55af6af2DecodeNodePuppeteerExampleGoCommonModelPo(in *jlexer.Lexer, out *SupplierImage) {
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
		case "src_origin":
			out.SrcOrigin = string(in.String())
		case "src_own":
			out.SrcOwn = string(in.String())
		case "progress":
			out.Progress = int(in.Int())
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
func easyjson55af6af2EncodeNodePuppeteerExampleGoCommonModelPo(out *jwriter.Writer, in SupplierImage) {
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
		const prefix string = ",\"src_origin\":"
		out.RawString(prefix)
		out.String(string(in.SrcOrigin))
	}
	{
		const prefix string = ",\"src_own\":"
		out.RawString(prefix)
		out.String(string(in.SrcOwn))
	}
	{
		const prefix string = ",\"progress\":"
		out.RawString(prefix)
		out.Int(int(in.Progress))
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
func (v SupplierImage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson55af6af2EncodeNodePuppeteerExampleGoCommonModelPo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SupplierImage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson55af6af2EncodeNodePuppeteerExampleGoCommonModelPo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SupplierImage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson55af6af2DecodeNodePuppeteerExampleGoCommonModelPo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SupplierImage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson55af6af2DecodeNodePuppeteerExampleGoCommonModelPo(l, v)
}
