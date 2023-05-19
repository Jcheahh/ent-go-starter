// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/productpageview"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ProductPageView is the model entity for the ProductPageView schema.
type ProductPageView struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductPageViewQuery when eager-loading is set.
	Edges                      ProductPageViewEdges `json:"edges"`
	product_product_page_views *int
	selectValues               sql.SelectValues
}

// ProductPageViewEdges holds the relations/edges for other nodes in the graph.
type ProductPageViewEdges struct {
	// HeroContent holds the value of the heroContent edge.
	HeroContent []*HeroContent `json:"heroContent,omitempty"`
	// PrimaryContent holds the value of the primaryContent edge.
	PrimaryContent []*PrimaryContent `json:"primaryContent,omitempty"`
	// ViewAnalytics holds the value of the viewAnalytics edge.
	ViewAnalytics []*ViewAnalytics `json:"viewAnalytics,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedHeroContent    map[string][]*HeroContent
	namedPrimaryContent map[string][]*PrimaryContent
	namedViewAnalytics  map[string][]*ViewAnalytics
}

// HeroContentOrErr returns the HeroContent value or an error if the edge
// was not loaded in eager-loading.
func (e ProductPageViewEdges) HeroContentOrErr() ([]*HeroContent, error) {
	if e.loadedTypes[0] {
		return e.HeroContent, nil
	}
	return nil, &NotLoadedError{edge: "heroContent"}
}

// PrimaryContentOrErr returns the PrimaryContent value or an error if the edge
// was not loaded in eager-loading.
func (e ProductPageViewEdges) PrimaryContentOrErr() ([]*PrimaryContent, error) {
	if e.loadedTypes[1] {
		return e.PrimaryContent, nil
	}
	return nil, &NotLoadedError{edge: "primaryContent"}
}

// ViewAnalyticsOrErr returns the ViewAnalytics value or an error if the edge
// was not loaded in eager-loading.
func (e ProductPageViewEdges) ViewAnalyticsOrErr() ([]*ViewAnalytics, error) {
	if e.loadedTypes[2] {
		return e.ViewAnalytics, nil
	}
	return nil, &NotLoadedError{edge: "viewAnalytics"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductPageView) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productpageview.FieldID, productpageview.FieldVersion:
			values[i] = new(sql.NullInt64)
		case productpageview.ForeignKeys[0]: // product_product_page_views
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductPageView fields.
func (ppv *ProductPageView) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productpageview.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ppv.ID = int(value.Int64)
		case productpageview.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				ppv.Version = int(value.Int64)
			}
		case productpageview.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field product_product_page_views", value)
			} else if value.Valid {
				ppv.product_product_page_views = new(int)
				*ppv.product_product_page_views = int(value.Int64)
			}
		default:
			ppv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductPageView.
// This includes values selected through modifiers, order, etc.
func (ppv *ProductPageView) Value(name string) (ent.Value, error) {
	return ppv.selectValues.Get(name)
}

// QueryHeroContent queries the "heroContent" edge of the ProductPageView entity.
func (ppv *ProductPageView) QueryHeroContent() *HeroContentQuery {
	return NewProductPageViewClient(ppv.config).QueryHeroContent(ppv)
}

// QueryPrimaryContent queries the "primaryContent" edge of the ProductPageView entity.
func (ppv *ProductPageView) QueryPrimaryContent() *PrimaryContentQuery {
	return NewProductPageViewClient(ppv.config).QueryPrimaryContent(ppv)
}

// QueryViewAnalytics queries the "viewAnalytics" edge of the ProductPageView entity.
func (ppv *ProductPageView) QueryViewAnalytics() *ViewAnalyticsQuery {
	return NewProductPageViewClient(ppv.config).QueryViewAnalytics(ppv)
}

// Update returns a builder for updating this ProductPageView.
// Note that you need to call ProductPageView.Unwrap() before calling this method if this ProductPageView
// was returned from a transaction, and the transaction was committed or rolled back.
func (ppv *ProductPageView) Update() *ProductPageViewUpdateOne {
	return NewProductPageViewClient(ppv.config).UpdateOne(ppv)
}

// Unwrap unwraps the ProductPageView entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ppv *ProductPageView) Unwrap() *ProductPageView {
	_tx, ok := ppv.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductPageView is not a transactional entity")
	}
	ppv.config.driver = _tx.drv
	return ppv
}

// String implements the fmt.Stringer.
func (ppv *ProductPageView) String() string {
	var builder strings.Builder
	builder.WriteString("ProductPageView(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ppv.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", ppv.Version))
	builder.WriteByte(')')
	return builder.String()
}

// NamedHeroContent returns the HeroContent named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ppv *ProductPageView) NamedHeroContent(name string) ([]*HeroContent, error) {
	if ppv.Edges.namedHeroContent == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ppv.Edges.namedHeroContent[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ppv *ProductPageView) appendNamedHeroContent(name string, edges ...*HeroContent) {
	if ppv.Edges.namedHeroContent == nil {
		ppv.Edges.namedHeroContent = make(map[string][]*HeroContent)
	}
	if len(edges) == 0 {
		ppv.Edges.namedHeroContent[name] = []*HeroContent{}
	} else {
		ppv.Edges.namedHeroContent[name] = append(ppv.Edges.namedHeroContent[name], edges...)
	}
}

// NamedPrimaryContent returns the PrimaryContent named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ppv *ProductPageView) NamedPrimaryContent(name string) ([]*PrimaryContent, error) {
	if ppv.Edges.namedPrimaryContent == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ppv.Edges.namedPrimaryContent[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ppv *ProductPageView) appendNamedPrimaryContent(name string, edges ...*PrimaryContent) {
	if ppv.Edges.namedPrimaryContent == nil {
		ppv.Edges.namedPrimaryContent = make(map[string][]*PrimaryContent)
	}
	if len(edges) == 0 {
		ppv.Edges.namedPrimaryContent[name] = []*PrimaryContent{}
	} else {
		ppv.Edges.namedPrimaryContent[name] = append(ppv.Edges.namedPrimaryContent[name], edges...)
	}
}

// NamedViewAnalytics returns the ViewAnalytics named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ppv *ProductPageView) NamedViewAnalytics(name string) ([]*ViewAnalytics, error) {
	if ppv.Edges.namedViewAnalytics == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ppv.Edges.namedViewAnalytics[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ppv *ProductPageView) appendNamedViewAnalytics(name string, edges ...*ViewAnalytics) {
	if ppv.Edges.namedViewAnalytics == nil {
		ppv.Edges.namedViewAnalytics = make(map[string][]*ViewAnalytics)
	}
	if len(edges) == 0 {
		ppv.Edges.namedViewAnalytics[name] = []*ViewAnalytics{}
	} else {
		ppv.Edges.namedViewAnalytics[name] = append(ppv.Edges.namedViewAnalytics[name], edges...)
	}
}

// ProductPageViews is a parsable slice of ProductPageView.
type ProductPageViews []*ProductPageView