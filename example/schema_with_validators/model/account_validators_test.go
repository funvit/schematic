package model

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestAccount_Validate(t *testing.T) {
	type fields struct {
		Id        int64
		Login     string
		Name      string
		Password  string
		IsDeleted bool
		Rev       uuid.UUID
		CreatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"valid",
			fields{
				Id:        1,
				Login:     "login",
				Name:      "name",
				Password:  "00000000",
				IsDeleted: false,
				Rev:       uuid.New(),
				CreatedAt: time.Now(),
			},
			false,
		},
		{
			"empty Rev must cause validation error",
			fields{
				Id:        0,
				Login:     "login",
				Name:      "name",
				Password:  "00000000",
				IsDeleted: false,
				Rev:       uuid.Nil,
				CreatedAt: time.Now(),
			},
			true,
		},
		{
			"empty Name must cause validation error",
			fields{
				Id:        0,
				Login:     "login",
				Name:      "",
				Password:  "00000000",
				IsDeleted: false,
				Rev:       uuid.New(),
				CreatedAt: time.Now(),
			},
			true,
		},
		{
			"empty CreatedAt must cause validation error",
			fields{
				Id:        0,
				Login:     "login",
				Name:      "name",
				Password:  "00000000",
				IsDeleted: false,
				Rev:       uuid.New(),
				CreatedAt: time.Time{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Account{
				id:        tt.fields.Id,
				Login:     tt.fields.Login,
				Name:      tt.fields.Name,
				Password:  tt.fields.Password,
				IsDeleted: tt.fields.IsDeleted,
				Rev:       tt.fields.Rev,
				CreatedAt: tt.fields.CreatedAt,
			}
			err := m.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				t.Log("got error:", err)
			}

			t.Log("model stringer:", m)
		})
	}
}
