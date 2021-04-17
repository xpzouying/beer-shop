// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/biz"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/data/ent/beer"
)

// Beer is the model entity for the Beer schema.
type Beer struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Count holds the value of the "count" field.
	Count int64 `json:"count,omitempty"`
	// Price holds the value of the "price" field.
	Price int64 `json:"price,omitempty"`
	// Images holds the value of the "images" field.
	Images []biz.Image `json:"images,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Beer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case beer.FieldImages:
			values[i] = &[]byte{}
		case beer.FieldID, beer.FieldCount, beer.FieldPrice:
			values[i] = &sql.NullInt64{}
		case beer.FieldName, beer.FieldDescription:
			values[i] = &sql.NullString{}
		case beer.FieldCreatedAt, beer.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Beer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Beer fields.
func (b *Beer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case beer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int64(value.Int64)
		case beer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case beer.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = value.String
			}
		case beer.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				b.Count = value.Int64
			}
		case beer.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				b.Price = value.Int64
			}
		case beer.FieldImages:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field images", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &b.Images); err != nil {
					return fmt.Errorf("unmarshal field images: %w", err)
				}
			}
		case beer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case beer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Beer.
// Note that you need to call Beer.Unwrap() before calling this method if this Beer
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Beer) Update() *BeerUpdateOne {
	return (&BeerClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Beer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Beer) Unwrap() *Beer {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Beer is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Beer) String() string {
	var builder strings.Builder
	builder.WriteString("Beer(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", name=")
	builder.WriteString(b.Name)
	builder.WriteString(", description=")
	builder.WriteString(b.Description)
	builder.WriteString(", count=")
	builder.WriteString(fmt.Sprintf("%v", b.Count))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", b.Price))
	builder.WriteString(", images=")
	builder.WriteString(fmt.Sprintf("%v", b.Images))
	builder.WriteString(", created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Beers is a parsable slice of Beer.
type Beers []*Beer

func (b Beers) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
