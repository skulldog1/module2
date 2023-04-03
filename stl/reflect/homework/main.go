package main

import (
	"fmt"
	"reflect"
	"time"
)

var reflectedStructs map[string][]Field

type Field struct {
	Name string
	Tags map[string]string
}

type UserDTO struct {
	ID            int       `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null"`
	Name          string    `json:"name" db:"name" db_type:"varchar(55)" db_default:"default null" db_ops:"create,update"`
	Phone         string    `json:"phone" db:"phone" db_type:"varchar(34)" db_default:"default null" db_index:"index,unique" db_ops:"create,update"`
	Email         string    `json:"email" db:"email" db_type:"varchar(89)" db_default:"default null" db_index:"index,unique" db_ops:"create,update"`
	Password      string    `json:"password" db:"password" db_type:"varchar(144)" db_default:"default null" db_ops:"create,update"`
	Status        int       `json:"status" db:"status" db_type:"int" db_default:"default 0" db_ops:"create,update"`
	Role          int       `json:"role" db:"role" db_type:"int" db_default:"not null" db_ops:"create,update"`
	Verified      bool      `json:"verified" db:"verified" db_type:"boolean" db_default:"not null" db_ops:"create,update"`
	EmailVerified bool      `json:"email_verified" db:"email_verified" db_type:"boolean" db_default:"not null" db_ops:"create,update"`
	PhoneVerified bool      `json:"phone_verified" db:"phone_verified" db_type:"boolean" db_default:"not null" db_ops:"create,update"`
	CreatedAt     time.Time `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index"`
	DeletedAt     time.Time `json:"deleted_at" db:"deleted_at" db_type:"timestamp" db_default:"default null" db_index:"index"`
}

type EmailVerifyDTO struct {
	ID        int       `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null"`
	Email     string    `json:"email" db:"email" db_type:"varchar(89)" db_default:"not null" db_index:"index,unique" db_ops:"create,update"`
	UserID    int       `json:"user_id,omitempty" db:"user_id" db_ops:"create" db_type:"int" db_default:"not null" db_index:"index"`
	Hash      string    `json:"hash,omitempty" db:"hash" db_ops:"create" db_type:"char(36)" db_default:"not null" db_index:"index"`
	Verified  bool      `json:"verified" db:"verified" db_type:"boolean" db_default:"not null" db_ops:"create,update"`
	CreatedAt time.Time `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index"`
}

func main() {
	structs := []interface{}{
		UserDTO{},
		EmailVerifyDTO{},
	}

	reflectedStructs = make(map[string][]Field)

	for _, s := range structs {
		t := reflect.TypeOf(s)
		fields := []Field{}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tags := make(map[string]string)

			for _, tag := range []string{"json", "db"} {
				tagVal := field.Tag.Get(tag)
				if tagVal != "" {
					tags[tag] = tagVal
				}
			}

			name := field.Name
			fields = append(fields, Field{Name: name, Tags: tags})
		}

		reflectedStructs[t.Name()] = fields
	}

	fmt.Println(reflectedStructs)
}
