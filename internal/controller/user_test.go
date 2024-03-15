package controller

import "testing"

func TestChecksValidation(t *testing.T) {
	u := &UserRequest{Email: "uai@uai.com.br", Password: "12345678", UserType: "Admin"}
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
