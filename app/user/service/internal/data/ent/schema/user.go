package schema

import (
	v1 "ashop/api/user/service/v1"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("username").NotEmpty().MaxLen(50),
		field.String("password_hash").MaxLen(64),
		field.String("phone").MaxLen(20),
		field.String("email").MaxLen(30).Optional(),
		field.Int8("status").Default(int8(v1.UserStatus_NORMAL)),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("delete_at").
			Optional().Nillable().Comment("Soft deletion time").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("delete_at"),
		index.Fields("phone"),
	}
}
