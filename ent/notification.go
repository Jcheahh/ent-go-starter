// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/notification"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Notification is the model entity for the Notification schema.
type Notification struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// DateCreated holds the value of the "dateCreated" field.
	DateCreated string `json:"dateCreated,omitempty"`
	// DateUpdated holds the value of the "dateUpdated" field.
	DateUpdated string `json:"dateUpdated,omitempty"`
	// Read holds the value of the "read" field.
	Read bool `json:"read,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotificationQuery when eager-loading is set.
	Edges              NotificationEdges `json:"edges"`
	user_notifications *int
	selectValues       sql.SelectValues
}

// NotificationEdges holds the relations/edges for other nodes in the graph.
type NotificationEdges struct {
	// Recipient holds the value of the recipient edge.
	Recipient []*User `json:"recipient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedRecipient map[string][]*User
}

// RecipientOrErr returns the Recipient value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationEdges) RecipientOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Recipient, nil
	}
	return nil, &NotLoadedError{edge: "recipient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Notification) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notification.FieldRead:
			values[i] = new(sql.NullBool)
		case notification.FieldID:
			values[i] = new(sql.NullInt64)
		case notification.FieldTitle, notification.FieldContent, notification.FieldDateCreated, notification.FieldDateUpdated:
			values[i] = new(sql.NullString)
		case notification.ForeignKeys[0]: // user_notifications
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Notification fields.
func (n *Notification) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notification.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case notification.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				n.Title = value.String
			}
		case notification.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				n.Content = value.String
			}
		case notification.FieldDateCreated:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dateCreated", values[i])
			} else if value.Valid {
				n.DateCreated = value.String
			}
		case notification.FieldDateUpdated:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dateUpdated", values[i])
			} else if value.Valid {
				n.DateUpdated = value.String
			}
		case notification.FieldRead:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field read", values[i])
			} else if value.Valid {
				n.Read = value.Bool
			}
		case notification.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_notifications", value)
			} else if value.Valid {
				n.user_notifications = new(int)
				*n.user_notifications = int(value.Int64)
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Notification.
// This includes values selected through modifiers, order, etc.
func (n *Notification) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryRecipient queries the "recipient" edge of the Notification entity.
func (n *Notification) QueryRecipient() *UserQuery {
	return NewNotificationClient(n.config).QueryRecipient(n)
}

// Update returns a builder for updating this Notification.
// Note that you need to call Notification.Unwrap() before calling this method if this Notification
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Notification) Update() *NotificationUpdateOne {
	return NewNotificationClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Notification entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Notification) Unwrap() *Notification {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Notification is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Notification) String() string {
	var builder strings.Builder
	builder.WriteString("Notification(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("title=")
	builder.WriteString(n.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(n.Content)
	builder.WriteString(", ")
	builder.WriteString("dateCreated=")
	builder.WriteString(n.DateCreated)
	builder.WriteString(", ")
	builder.WriteString("dateUpdated=")
	builder.WriteString(n.DateUpdated)
	builder.WriteString(", ")
	builder.WriteString("read=")
	builder.WriteString(fmt.Sprintf("%v", n.Read))
	builder.WriteByte(')')
	return builder.String()
}

// NamedRecipient returns the Recipient named value or an error if the edge was not
// loaded in eager-loading with this name.
func (n *Notification) NamedRecipient(name string) ([]*User, error) {
	if n.Edges.namedRecipient == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := n.Edges.namedRecipient[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (n *Notification) appendNamedRecipient(name string, edges ...*User) {
	if n.Edges.namedRecipient == nil {
		n.Edges.namedRecipient = make(map[string][]*User)
	}
	if len(edges) == 0 {
		n.Edges.namedRecipient[name] = []*User{}
	} else {
		n.Edges.namedRecipient[name] = append(n.Edges.namedRecipient[name], edges...)
	}
}

// Notifications is a parsable slice of Notification.
type Notifications []*Notification
