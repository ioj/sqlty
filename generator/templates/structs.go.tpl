{{ define "struct_declaration" }}
  type {{ .Name }} struct {
    {{ range .Params }}
      {{- $jsonTag := lowerFirstLetter .Name -}}
      {{ .Name }} {{ .Type.Name }} `json:"{{ $jsonTag }}"`
    {{ end }}
  }
{{ end }}
