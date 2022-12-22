// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package modelsRestApi

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

func easyjson32a60ab0Decode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(in *jlexer.Lexer, out *CategoryInfo) {
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
		case "category_name":
			out.CategoryName = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "subscribers_count":
			out.SubscribersCount = int(in.Int())
		case "subscribed":
			out.Subscribed = bool(in.Bool())
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
func easyjson32a60ab0Encode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(out *jwriter.Writer, in CategoryInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"category_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.CategoryName))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"subscribers_count\":"
		out.RawString(prefix)
		out.Int(int(in.SubscribersCount))
	}
	{
		const prefix string = ",\"subscribed\":"
		out.RawString(prefix)
		out.Bool(bool(in.Subscribed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CategoryInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson32a60ab0Encode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CategoryInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson32a60ab0Encode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CategoryInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson32a60ab0Decode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CategoryInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson32a60ab0Decode20222GoToTeamInternalServerRestAPICategoryComponentDeliveryModelsRestApi(l, v)
}