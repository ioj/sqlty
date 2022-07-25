{{ define "struct_declaration" }}
  type {{ .Name }} struct {
    {{ range .Params }}
      {{- $jsonTag := lowerFirstLetter .Name -}}
      {{ .Name }} {{ .Type.Name }} `json:"{{ $jsonTag }}"`
    {{ end }}
  }
{{ end }}

{{ define "struct_functions" }}
  func (dst *{{ .Name }}) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
    if src == nil {
      return errors.New("NULL values can't be decoded. Scan into a &*{{ .Name }} to handle NULLs")
    }

    if err := (pgtype.CompositeFields{ {{- range .Params -}}
      &dst.{{- .Name -}},
      {{ end -}} }).DecodeBinary(ci, src); err != nil {
      return err
    }

    return nil
  }
{{ end }}