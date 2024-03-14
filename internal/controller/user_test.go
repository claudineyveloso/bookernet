package controller

import "testing"

func TestChecksValidation(t *testing.T) {
	u := &UserRequest{Email: ""}
	err := u.Validate()
	if err != nil {
		t.Fatal("Erro de validação:", err)
	}

	if err == nil {
		t.Log("A validação passou com sucesso.")
	} else {
		t.Error("A validação falhou.")
	}
}
