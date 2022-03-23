package pkg

import (
	"reflect"
	"testing"
	"time"
)

func TestGetService(t *testing.T) {
	t.Run("Panic on not initialize", func(t *testing.T) {
		defer func() { recover() }()

		GetJwtService()
		t.Errorf("Did not panic")
	})
}

func TestCheckTokenCreation(t *testing.T) {
	InitJwt("asdasdasd", time.Hour * 72)
	jwt := GetJwtService()
	
	t.Run("create token", func(t *testing.T) {
		got := jwt.CreateToken("test@co.jp")
		isType := reflect.TypeOf(got).Kind() != reflect.String
		if isType == true {
			t.Errorf("createToken() = %v, want %v", got, false)
		}
	})
}