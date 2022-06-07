package models

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
	UserName string
	Score    int32
}

type Users []User

func (u *Users) ConvertToString() string {
	var builder strings.Builder
	builder.WriteString("Топ игроки\n\n")
	for i, user := range *u {
		name := fmt.Sprintf("%s", user.FullName)
		if len(name) > 17 {
			name = name[:14] + "..."
		}
		text := fmt.Sprintf("%d| %s — %d\n", i+1, name, user.Score)
		builder.WriteString(text)
	}
	return builder.String()
}
