package model

type Email struct {
	To           string
	Subject      string
	TemplateName string
	Data         map[string]any
}
