package main

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "short name",
			user: User{
				Name:  "u",
				Email: "test@example.com",
			},
			wantErr: true, // Name < 3 chars
		},
		{
			name: "long name",
			user: User{
				Name:  "I have at least 32 characters, but I'm not sure.",
				Email: "test@example.com",
			},
			wantErr: true, // Name > 32 chars
		},
		{
			name: "missing email",
			user: User{
				Name:  "rafael",
				Email: "",
			},
			wantErr: true, // Email is required
		},
		{
			name: "short name and missing email",
			user: User{
				Name:  "a",
				Email: "",
			},
			wantErr: true, // Both errors
		},
		{
			name: "valid user",
			user: User{
				Name:  "rafael",
				Email: "test@example.com",
			},
			wantErr: false, // No errors expected
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
