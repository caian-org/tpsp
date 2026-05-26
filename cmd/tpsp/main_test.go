package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetColorForStatus(t *testing.T) {
	tests := []struct {
		name        string
		statusColor string
		want        string
	}{
		{"verde lowercase", "verde", colorGreen},
		{"verde uppercase", "VERDE", colorGreen},
		{"verde mixed case", "Verde", colorGreen},
		{"amarelo lowercase", "amarelo", colorYellow},
		{"amarelo uppercase", "AMARELO", colorYellow},
		{"vermelho lowercase", "vermelho", colorRed},
		{"vermelho uppercase", "VERMELHO", colorRed},
		{"cinza lowercase", "cinza", colorDim},
		{"cinza mixed case", "Cinza", colorDim},
		{"unknown color falls back to reset", "azul", colorReset},
		{"empty string falls back to reset", "", colorReset},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, getColorForStatus(tc.statusColor))
		})
	}
}

func TestFormatLineName(t *testing.T) {
	tests := []struct {
		name string
		line string
		want string
	}{
		{"Linha-Verde becomes Verde", "Linha-Verde", "Verde"},
		{"Linha-AZUL becomes Azul", "Linha-AZUL", "Azul"},
		{"Linha-amarela becomes Amarela", "Linha-amarela", "Amarela"},
		{"single segment lowercase", "vermelha", "Vermelha"},
		{"single segment uppercase", "PRATA", "Prata"},
		{"empty string returns empty", "", ""},
		{"multiple dashes uses last", "Linha-Expresso-Turquesa", "Turquesa"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, formatLineName(tc.line))
		})
	}
}

func TestNormalizeStatus(t *testing.T) {
	tests := []struct {
		name   string
		status string
		want   string
	}{
		{"operações normais lowercase", "operações normais", "Operação Normal"},
		{"operações normais mixed case", "Operações Normais", "Operação Normal"},
		{"operações encerradas lowercase", "operações encerradas", "Operação Encerrada"},
		{"operações encerradas uppercase", "OPERAÇÕES ENCERRADAS", "Operação Encerrada"},
		{"unknown status passes through", "Velocidade Reduzida", "Velocidade Reduzida"},
		{"trims leading and trailing spaces", "   operações normais   ", "Operação Normal"},
		{"passthrough with trim", "  Paralisada  ", "Paralisada"},
		{"empty string remains empty", "", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, normalizeStatus(tc.status))
		})
	}
}

func TestIsValidService(t *testing.T) {
	tests := []struct {
		name    string
		service string
		want    bool
	}{
		{"metro lowercase", "metro", true},
		{"Metro title case", "Metro", true},
		{"METRO uppercase", "METRO", true},
		{"cptm lowercase", "cptm", true},
		{"viamobilidade", "viamobilidade", true},
		{"viaquatro", "viaquatro", true},
		{"ônibus is invalid", "ônibus", false},
		{"empty string is invalid", "", false},
		{"random word is invalid", "trem", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, isValidService(tc.service))
		})
	}
}

func TestFilterByService(t *testing.T) {
	data := []ServiceData{
		{
			Type: "metro",
			ListItem: []LineItem{
				{Line: "Linha-Azul", Status: "Operação Normal"},
				{Line: "Linha-Verde", Status: "Operação Normal"},
			},
		},
		{
			Type: "cptm",
			ListItem: []LineItem{
				{Line: "Linha-Rubi", Status: "Operação Normal"},
			},
		},
	}

	t.Run("empty filter returns all lines", func(t *testing.T) {
		got := filterByService(data, "")
		assert.Len(t, got, 3)
	})

	t.Run("metro filter returns only metro lines", func(t *testing.T) {
		got := filterByService(data, "metro")
		assert.Len(t, got, 2)
		assert.Equal(t, "Linha-Azul", got[0].Line)
		assert.Equal(t, "Linha-Verde", got[1].Line)
	})

	t.Run("cptm filter returns only cptm lines", func(t *testing.T) {
		got := filterByService(data, "cptm")
		assert.Len(t, got, 1)
		assert.Equal(t, "Linha-Rubi", got[0].Line)
	})

	t.Run("filter is case-insensitive", func(t *testing.T) {
		got := filterByService(data, "METRO")
		assert.Len(t, got, 2)
	})

	t.Run("unknown service returns no lines", func(t *testing.T) {
		got := filterByService(data, "unknown")
		assert.Empty(t, got)
	})
}
