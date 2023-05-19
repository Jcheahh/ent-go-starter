// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/commissionstructureschema"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// CommissionStructureSchema is the model entity for the CommissionStructureSchema schema.
type CommissionStructureSchema struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CommissionValue holds the value of the "commissionValue" field.
	CommissionValue string `json:"commissionValue,omitempty"`
	// CommissionPercentage holds the value of the "commissionPercentage" field.
	CommissionPercentage string `json:"commissionPercentage,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommissionStructureSchemaQuery when eager-loading is set.
	Edges                        CommissionStructureSchemaEdges `json:"edges"`
	product_commission_structure *int
	selectValues                 sql.SelectValues
}

// CommissionStructureSchemaEdges holds the relations/edges for other nodes in the graph.
type CommissionStructureSchemaEdges struct {
	// ProductSeller holds the value of the productSeller edge.
	ProductSeller []*UserSeller `json:"productSeller,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedProductSeller map[string][]*UserSeller
}

// ProductSellerOrErr returns the ProductSeller value or an error if the edge
// was not loaded in eager-loading.
func (e CommissionStructureSchemaEdges) ProductSellerOrErr() ([]*UserSeller, error) {
	if e.loadedTypes[0] {
		return e.ProductSeller, nil
	}
	return nil, &NotLoadedError{edge: "productSeller"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CommissionStructureSchema) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case commissionstructureschema.FieldID:
			values[i] = new(sql.NullInt64)
		case commissionstructureschema.FieldName, commissionstructureschema.FieldDescription, commissionstructureschema.FieldCommissionValue, commissionstructureschema.FieldCommissionPercentage:
			values[i] = new(sql.NullString)
		case commissionstructureschema.ForeignKeys[0]: // product_commission_structure
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CommissionStructureSchema fields.
func (css *CommissionStructureSchema) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case commissionstructureschema.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			css.ID = int(value.Int64)
		case commissionstructureschema.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				css.Name = value.String
			}
		case commissionstructureschema.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				css.Description = value.String
			}
		case commissionstructureschema.FieldCommissionValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commissionValue", values[i])
			} else if value.Valid {
				css.CommissionValue = value.String
			}
		case commissionstructureschema.FieldCommissionPercentage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commissionPercentage", values[i])
			} else if value.Valid {
				css.CommissionPercentage = value.String
			}
		case commissionstructureschema.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field product_commission_structure", value)
			} else if value.Valid {
				css.product_commission_structure = new(int)
				*css.product_commission_structure = int(value.Int64)
			}
		default:
			css.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CommissionStructureSchema.
// This includes values selected through modifiers, order, etc.
func (css *CommissionStructureSchema) Value(name string) (ent.Value, error) {
	return css.selectValues.Get(name)
}

// QueryProductSeller queries the "productSeller" edge of the CommissionStructureSchema entity.
func (css *CommissionStructureSchema) QueryProductSeller() *UserSellerQuery {
	return NewCommissionStructureSchemaClient(css.config).QueryProductSeller(css)
}

// Update returns a builder for updating this CommissionStructureSchema.
// Note that you need to call CommissionStructureSchema.Unwrap() before calling this method if this CommissionStructureSchema
// was returned from a transaction, and the transaction was committed or rolled back.
func (css *CommissionStructureSchema) Update() *CommissionStructureSchemaUpdateOne {
	return NewCommissionStructureSchemaClient(css.config).UpdateOne(css)
}

// Unwrap unwraps the CommissionStructureSchema entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (css *CommissionStructureSchema) Unwrap() *CommissionStructureSchema {
	_tx, ok := css.config.driver.(*txDriver)
	if !ok {
		panic("ent: CommissionStructureSchema is not a transactional entity")
	}
	css.config.driver = _tx.drv
	return css
}

// String implements the fmt.Stringer.
func (css *CommissionStructureSchema) String() string {
	var builder strings.Builder
	builder.WriteString("CommissionStructureSchema(")
	builder.WriteString(fmt.Sprintf("id=%v, ", css.ID))
	builder.WriteString("name=")
	builder.WriteString(css.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(css.Description)
	builder.WriteString(", ")
	builder.WriteString("commissionValue=")
	builder.WriteString(css.CommissionValue)
	builder.WriteString(", ")
	builder.WriteString("commissionPercentage=")
	builder.WriteString(css.CommissionPercentage)
	builder.WriteByte(')')
	return builder.String()
}

// NamedProductSeller returns the ProductSeller named value or an error if the edge was not
// loaded in eager-loading with this name.
func (css *CommissionStructureSchema) NamedProductSeller(name string) ([]*UserSeller, error) {
	if css.Edges.namedProductSeller == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := css.Edges.namedProductSeller[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (css *CommissionStructureSchema) appendNamedProductSeller(name string, edges ...*UserSeller) {
	if css.Edges.namedProductSeller == nil {
		css.Edges.namedProductSeller = make(map[string][]*UserSeller)
	}
	if len(edges) == 0 {
		css.Edges.namedProductSeller[name] = []*UserSeller{}
	} else {
		css.Edges.namedProductSeller[name] = append(css.Edges.namedProductSeller[name], edges...)
	}
}

// CommissionStructureSchemas is a parsable slice of CommissionStructureSchema.
type CommissionStructureSchemas []*CommissionStructureSchema