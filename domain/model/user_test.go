package model_test

import (
	"ismaeldf.melo/imersao-fullstack-fullcycle/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewUser(t *testing.T) {
	name := "Ismael"
	email := "meuemail@com"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser("", email)
	require.NotNil(t, err)

	_, err = model.NewUser(name, "")
	require.NotNil(t, err)
}
