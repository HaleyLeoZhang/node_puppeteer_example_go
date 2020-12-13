// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package vo

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	po "node_puppeteer_example_go/common/model/po"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonEa649aceDecodeNodePuppeteerExampleGoCommonModelVo(in *jlexer.Lexer, out *ImageListResponse) {
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
		case "list":
			if in.IsNull() {
				in.Skip()
				out.List = nil
			} else {
				in.Delim('[')
				if out.List == nil {
					if !in.IsDelim(']') {
						out.List = make([]*po.ComicImage, 0, 8)
					} else {
						out.List = []*po.ComicImage{}
					}
				} else {
					out.List = (out.List)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *po.ComicImage
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(po.ComicImage)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.List = append(out.List, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonEa649aceEncodeNodePuppeteerExampleGoCommonModelVo(out *jwriter.Writer, in ImageListResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"list\":"
		out.RawString(prefix[1:])
		if in.List == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.List {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ImageListResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEa649aceEncodeNodePuppeteerExampleGoCommonModelVo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ImageListResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEa649aceEncodeNodePuppeteerExampleGoCommonModelVo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ImageListResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEa649aceDecodeNodePuppeteerExampleGoCommonModelVo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ImageListResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEa649aceDecodeNodePuppeteerExampleGoCommonModelVo(l, v)
}
func easyjsonEa649aceDecodeNodePuppeteerExampleGoCommonModelVo1(in *jlexer.Lexer, out *ImageListParam) {
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
		case "PageId":
			out.PageId = int(in.Int())
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
func easyjsonEa649aceEncodeNodePuppeteerExampleGoCommonModelVo1(out *jwriter.Writer, in ImageListParam) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"PageId\":"
		out.RawString(prefix[1:])
		out.Int(int(in.PageId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ImageListParam) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEa649aceEncodeNodePuppeteerExampleGoCommonModelVo1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ImageListParam) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEa649aceEncodeNodePuppeteerExampleGoCommonModelVo1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ImageListParam) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEa649aceDecodeNodePuppeteerExampleGoCommonModelVo1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ImageListParam) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEa649aceDecodeNodePuppeteerExampleGoCommonModelVo1(l, v)
}