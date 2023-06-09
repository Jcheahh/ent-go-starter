// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/product"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Price holds the value of the "price" field.
	Price string `json:"price,omitempty"`
	// DateCreated holds the value of the "dateCreated" field.
	DateCreated string `json:"dateCreated,omitempty"`
	// DateUpdated holds the value of the "dateUpdated" field.
	DateUpdated string `json:"dateUpdated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges                      ProductEdges `json:"edges"`
	category_products          *int
	group_buy_product          *int
	marketing_campaign_product *int
	review_product             *int
	shop_products              *int
	transaction_product        *int
	user_influencer_products   *int
	view_analytics_product     *int
	selectValues               sql.SelectValues
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// ProductSeller holds the value of the productSeller edge.
	ProductSeller []*UserSeller `json:"productSeller,omitempty"`
	// Reviews holds the value of the reviews edge.
	Reviews []*Review `json:"reviews,omitempty"`
	// Images holds the value of the images edge.
	Images []*Image `json:"images,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// ProductAttributes holds the value of the productAttributes edge.
	ProductAttributes []*ProductAttribute `json:"productAttributes,omitempty"`
	// Variations holds the value of the variations edge.
	Variations []*ProductVariation `json:"variations,omitempty"`
	// CommissionStructure holds the value of the commissionStructure edge.
	CommissionStructure []*CommissionStructureSchema `json:"commissionStructure,omitempty"`
	// Shop holds the value of the shop edge.
	Shop []*Shop `json:"shop,omitempty"`
	// GroupBuys holds the value of the groupBuys edge.
	GroupBuys []*GroupBuy `json:"groupBuys,omitempty"`
	// ProductPageViews holds the value of the productPageViews edge.
	ProductPageViews []*ProductPageView `json:"productPageViews,omitempty"`
	// BlogPosts holds the value of the blogPosts edge.
	BlogPosts []*BlogPost `json:"blogPosts,omitempty"`
	// MarketingCampaigns holds the value of the marketingCampaigns edge.
	MarketingCampaigns []*MarketingCampaign `json:"marketingCampaigns,omitempty"`
	// Chats holds the value of the chats edge.
	Chats []*Chat `json:"chats,omitempty"`
	// EmailCampaign holds the value of the emailCampaign edge.
	EmailCampaign []*EmailCampaign `json:"emailCampaign,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [15]bool
	// totalCount holds the count of the edges above.
	totalCount [15]map[string]int

	namedProductSeller       map[string][]*UserSeller
	namedReviews             map[string][]*Review
	namedImages              map[string][]*Image
	namedCategories          map[string][]*Category
	namedTags                map[string][]*Tag
	namedProductAttributes   map[string][]*ProductAttribute
	namedVariations          map[string][]*ProductVariation
	namedCommissionStructure map[string][]*CommissionStructureSchema
	namedShop                map[string][]*Shop
	namedGroupBuys           map[string][]*GroupBuy
	namedProductPageViews    map[string][]*ProductPageView
	namedBlogPosts           map[string][]*BlogPost
	namedMarketingCampaigns  map[string][]*MarketingCampaign
	namedChats               map[string][]*Chat
	namedEmailCampaign       map[string][]*EmailCampaign
}

// ProductSellerOrErr returns the ProductSeller value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductSellerOrErr() ([]*UserSeller, error) {
	if e.loadedTypes[0] {
		return e.ProductSeller, nil
	}
	return nil, &NotLoadedError{edge: "productSeller"}
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ReviewsOrErr() ([]*Review, error) {
	if e.loadedTypes[1] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ImagesOrErr() ([]*Image, error) {
	if e.loadedTypes[2] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[3] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[4] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// ProductAttributesOrErr returns the ProductAttributes value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductAttributesOrErr() ([]*ProductAttribute, error) {
	if e.loadedTypes[5] {
		return e.ProductAttributes, nil
	}
	return nil, &NotLoadedError{edge: "productAttributes"}
}

// VariationsOrErr returns the Variations value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) VariationsOrErr() ([]*ProductVariation, error) {
	if e.loadedTypes[6] {
		return e.Variations, nil
	}
	return nil, &NotLoadedError{edge: "variations"}
}

// CommissionStructureOrErr returns the CommissionStructure value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) CommissionStructureOrErr() ([]*CommissionStructureSchema, error) {
	if e.loadedTypes[7] {
		return e.CommissionStructure, nil
	}
	return nil, &NotLoadedError{edge: "commissionStructure"}
}

// ShopOrErr returns the Shop value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ShopOrErr() ([]*Shop, error) {
	if e.loadedTypes[8] {
		return e.Shop, nil
	}
	return nil, &NotLoadedError{edge: "shop"}
}

// GroupBuysOrErr returns the GroupBuys value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) GroupBuysOrErr() ([]*GroupBuy, error) {
	if e.loadedTypes[9] {
		return e.GroupBuys, nil
	}
	return nil, &NotLoadedError{edge: "groupBuys"}
}

// ProductPageViewsOrErr returns the ProductPageViews value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductPageViewsOrErr() ([]*ProductPageView, error) {
	if e.loadedTypes[10] {
		return e.ProductPageViews, nil
	}
	return nil, &NotLoadedError{edge: "productPageViews"}
}

// BlogPostsOrErr returns the BlogPosts value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) BlogPostsOrErr() ([]*BlogPost, error) {
	if e.loadedTypes[11] {
		return e.BlogPosts, nil
	}
	return nil, &NotLoadedError{edge: "blogPosts"}
}

// MarketingCampaignsOrErr returns the MarketingCampaigns value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) MarketingCampaignsOrErr() ([]*MarketingCampaign, error) {
	if e.loadedTypes[12] {
		return e.MarketingCampaigns, nil
	}
	return nil, &NotLoadedError{edge: "marketingCampaigns"}
}

// ChatsOrErr returns the Chats value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ChatsOrErr() ([]*Chat, error) {
	if e.loadedTypes[13] {
		return e.Chats, nil
	}
	return nil, &NotLoadedError{edge: "chats"}
}

// EmailCampaignOrErr returns the EmailCampaign value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) EmailCampaignOrErr() ([]*EmailCampaign, error) {
	if e.loadedTypes[14] {
		return e.EmailCampaign, nil
	}
	return nil, &NotLoadedError{edge: "emailCampaign"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			values[i] = new(sql.NullInt64)
		case product.FieldName, product.FieldDescription, product.FieldPrice, product.FieldDateCreated, product.FieldDateUpdated:
			values[i] = new(sql.NullString)
		case product.ForeignKeys[0]: // category_products
			values[i] = new(sql.NullInt64)
		case product.ForeignKeys[1]: // group_buy_product
			values[i] = new(sql.NullInt64)
		case product.ForeignKeys[2]: // marketing_campaign_product
			values[i] = new(sql.NullInt64)
		case product.ForeignKeys[3]: // review_product
			values[i] = new(sql.NullInt64)
		case product.ForeignKeys[4]: // shop_products
			values[i] = new(sql.NullInt64)
		case product.ForeignKeys[5]: // transaction_product
			values[i] = new(sql.NullInt64)
		case product.ForeignKeys[6]: // user_influencer_products
			values[i] = new(sql.NullInt64)
		case product.ForeignKeys[7]: // view_analytics_product
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case product.FieldPrice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = value.String
			}
		case product.FieldDateCreated:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dateCreated", values[i])
			} else if value.Valid {
				pr.DateCreated = value.String
			}
		case product.FieldDateUpdated:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dateUpdated", values[i])
			} else if value.Valid {
				pr.DateUpdated = value.String
			}
		case product.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category_products", value)
			} else if value.Valid {
				pr.category_products = new(int)
				*pr.category_products = int(value.Int64)
			}
		case product.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_buy_product", value)
			} else if value.Valid {
				pr.group_buy_product = new(int)
				*pr.group_buy_product = int(value.Int64)
			}
		case product.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field marketing_campaign_product", value)
			} else if value.Valid {
				pr.marketing_campaign_product = new(int)
				*pr.marketing_campaign_product = int(value.Int64)
			}
		case product.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field review_product", value)
			} else if value.Valid {
				pr.review_product = new(int)
				*pr.review_product = int(value.Int64)
			}
		case product.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field shop_products", value)
			} else if value.Valid {
				pr.shop_products = new(int)
				*pr.shop_products = int(value.Int64)
			}
		case product.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field transaction_product", value)
			} else if value.Valid {
				pr.transaction_product = new(int)
				*pr.transaction_product = int(value.Int64)
			}
		case product.ForeignKeys[6]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_influencer_products", value)
			} else if value.Valid {
				pr.user_influencer_products = new(int)
				*pr.user_influencer_products = int(value.Int64)
			}
		case product.ForeignKeys[7]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field view_analytics_product", value)
			} else if value.Valid {
				pr.view_analytics_product = new(int)
				*pr.view_analytics_product = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Product.
// This includes values selected through modifiers, order, etc.
func (pr *Product) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryProductSeller queries the "productSeller" edge of the Product entity.
func (pr *Product) QueryProductSeller() *UserSellerQuery {
	return NewProductClient(pr.config).QueryProductSeller(pr)
}

// QueryReviews queries the "reviews" edge of the Product entity.
func (pr *Product) QueryReviews() *ReviewQuery {
	return NewProductClient(pr.config).QueryReviews(pr)
}

// QueryImages queries the "images" edge of the Product entity.
func (pr *Product) QueryImages() *ImageQuery {
	return NewProductClient(pr.config).QueryImages(pr)
}

// QueryCategories queries the "categories" edge of the Product entity.
func (pr *Product) QueryCategories() *CategoryQuery {
	return NewProductClient(pr.config).QueryCategories(pr)
}

// QueryTags queries the "tags" edge of the Product entity.
func (pr *Product) QueryTags() *TagQuery {
	return NewProductClient(pr.config).QueryTags(pr)
}

// QueryProductAttributes queries the "productAttributes" edge of the Product entity.
func (pr *Product) QueryProductAttributes() *ProductAttributeQuery {
	return NewProductClient(pr.config).QueryProductAttributes(pr)
}

// QueryVariations queries the "variations" edge of the Product entity.
func (pr *Product) QueryVariations() *ProductVariationQuery {
	return NewProductClient(pr.config).QueryVariations(pr)
}

// QueryCommissionStructure queries the "commissionStructure" edge of the Product entity.
func (pr *Product) QueryCommissionStructure() *CommissionStructureSchemaQuery {
	return NewProductClient(pr.config).QueryCommissionStructure(pr)
}

// QueryShop queries the "shop" edge of the Product entity.
func (pr *Product) QueryShop() *ShopQuery {
	return NewProductClient(pr.config).QueryShop(pr)
}

// QueryGroupBuys queries the "groupBuys" edge of the Product entity.
func (pr *Product) QueryGroupBuys() *GroupBuyQuery {
	return NewProductClient(pr.config).QueryGroupBuys(pr)
}

// QueryProductPageViews queries the "productPageViews" edge of the Product entity.
func (pr *Product) QueryProductPageViews() *ProductPageViewQuery {
	return NewProductClient(pr.config).QueryProductPageViews(pr)
}

// QueryBlogPosts queries the "blogPosts" edge of the Product entity.
func (pr *Product) QueryBlogPosts() *BlogPostQuery {
	return NewProductClient(pr.config).QueryBlogPosts(pr)
}

// QueryMarketingCampaigns queries the "marketingCampaigns" edge of the Product entity.
func (pr *Product) QueryMarketingCampaigns() *MarketingCampaignQuery {
	return NewProductClient(pr.config).QueryMarketingCampaigns(pr)
}

// QueryChats queries the "chats" edge of the Product entity.
func (pr *Product) QueryChats() *ChatQuery {
	return NewProductClient(pr.config).QueryChats(pr)
}

// QueryEmailCampaign queries the "emailCampaign" edge of the Product entity.
func (pr *Product) QueryEmailCampaign() *EmailCampaignQuery {
	return NewProductClient(pr.config).QueryEmailCampaign(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(pr.Price)
	builder.WriteString(", ")
	builder.WriteString("dateCreated=")
	builder.WriteString(pr.DateCreated)
	builder.WriteString(", ")
	builder.WriteString("dateUpdated=")
	builder.WriteString(pr.DateUpdated)
	builder.WriteByte(')')
	return builder.String()
}

// NamedProductSeller returns the ProductSeller named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedProductSeller(name string) ([]*UserSeller, error) {
	if pr.Edges.namedProductSeller == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedProductSeller[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedProductSeller(name string, edges ...*UserSeller) {
	if pr.Edges.namedProductSeller == nil {
		pr.Edges.namedProductSeller = make(map[string][]*UserSeller)
	}
	if len(edges) == 0 {
		pr.Edges.namedProductSeller[name] = []*UserSeller{}
	} else {
		pr.Edges.namedProductSeller[name] = append(pr.Edges.namedProductSeller[name], edges...)
	}
}

// NamedReviews returns the Reviews named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedReviews(name string) ([]*Review, error) {
	if pr.Edges.namedReviews == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedReviews[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedReviews(name string, edges ...*Review) {
	if pr.Edges.namedReviews == nil {
		pr.Edges.namedReviews = make(map[string][]*Review)
	}
	if len(edges) == 0 {
		pr.Edges.namedReviews[name] = []*Review{}
	} else {
		pr.Edges.namedReviews[name] = append(pr.Edges.namedReviews[name], edges...)
	}
}

// NamedImages returns the Images named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedImages(name string) ([]*Image, error) {
	if pr.Edges.namedImages == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedImages[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedImages(name string, edges ...*Image) {
	if pr.Edges.namedImages == nil {
		pr.Edges.namedImages = make(map[string][]*Image)
	}
	if len(edges) == 0 {
		pr.Edges.namedImages[name] = []*Image{}
	} else {
		pr.Edges.namedImages[name] = append(pr.Edges.namedImages[name], edges...)
	}
}

// NamedCategories returns the Categories named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedCategories(name string) ([]*Category, error) {
	if pr.Edges.namedCategories == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedCategories[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedCategories(name string, edges ...*Category) {
	if pr.Edges.namedCategories == nil {
		pr.Edges.namedCategories = make(map[string][]*Category)
	}
	if len(edges) == 0 {
		pr.Edges.namedCategories[name] = []*Category{}
	} else {
		pr.Edges.namedCategories[name] = append(pr.Edges.namedCategories[name], edges...)
	}
}

// NamedTags returns the Tags named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedTags(name string) ([]*Tag, error) {
	if pr.Edges.namedTags == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedTags[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedTags(name string, edges ...*Tag) {
	if pr.Edges.namedTags == nil {
		pr.Edges.namedTags = make(map[string][]*Tag)
	}
	if len(edges) == 0 {
		pr.Edges.namedTags[name] = []*Tag{}
	} else {
		pr.Edges.namedTags[name] = append(pr.Edges.namedTags[name], edges...)
	}
}

// NamedProductAttributes returns the ProductAttributes named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedProductAttributes(name string) ([]*ProductAttribute, error) {
	if pr.Edges.namedProductAttributes == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedProductAttributes[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedProductAttributes(name string, edges ...*ProductAttribute) {
	if pr.Edges.namedProductAttributes == nil {
		pr.Edges.namedProductAttributes = make(map[string][]*ProductAttribute)
	}
	if len(edges) == 0 {
		pr.Edges.namedProductAttributes[name] = []*ProductAttribute{}
	} else {
		pr.Edges.namedProductAttributes[name] = append(pr.Edges.namedProductAttributes[name], edges...)
	}
}

// NamedVariations returns the Variations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedVariations(name string) ([]*ProductVariation, error) {
	if pr.Edges.namedVariations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedVariations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedVariations(name string, edges ...*ProductVariation) {
	if pr.Edges.namedVariations == nil {
		pr.Edges.namedVariations = make(map[string][]*ProductVariation)
	}
	if len(edges) == 0 {
		pr.Edges.namedVariations[name] = []*ProductVariation{}
	} else {
		pr.Edges.namedVariations[name] = append(pr.Edges.namedVariations[name], edges...)
	}
}

// NamedCommissionStructure returns the CommissionStructure named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedCommissionStructure(name string) ([]*CommissionStructureSchema, error) {
	if pr.Edges.namedCommissionStructure == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedCommissionStructure[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedCommissionStructure(name string, edges ...*CommissionStructureSchema) {
	if pr.Edges.namedCommissionStructure == nil {
		pr.Edges.namedCommissionStructure = make(map[string][]*CommissionStructureSchema)
	}
	if len(edges) == 0 {
		pr.Edges.namedCommissionStructure[name] = []*CommissionStructureSchema{}
	} else {
		pr.Edges.namedCommissionStructure[name] = append(pr.Edges.namedCommissionStructure[name], edges...)
	}
}

// NamedShop returns the Shop named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedShop(name string) ([]*Shop, error) {
	if pr.Edges.namedShop == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedShop[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedShop(name string, edges ...*Shop) {
	if pr.Edges.namedShop == nil {
		pr.Edges.namedShop = make(map[string][]*Shop)
	}
	if len(edges) == 0 {
		pr.Edges.namedShop[name] = []*Shop{}
	} else {
		pr.Edges.namedShop[name] = append(pr.Edges.namedShop[name], edges...)
	}
}

// NamedGroupBuys returns the GroupBuys named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedGroupBuys(name string) ([]*GroupBuy, error) {
	if pr.Edges.namedGroupBuys == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedGroupBuys[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedGroupBuys(name string, edges ...*GroupBuy) {
	if pr.Edges.namedGroupBuys == nil {
		pr.Edges.namedGroupBuys = make(map[string][]*GroupBuy)
	}
	if len(edges) == 0 {
		pr.Edges.namedGroupBuys[name] = []*GroupBuy{}
	} else {
		pr.Edges.namedGroupBuys[name] = append(pr.Edges.namedGroupBuys[name], edges...)
	}
}

// NamedProductPageViews returns the ProductPageViews named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedProductPageViews(name string) ([]*ProductPageView, error) {
	if pr.Edges.namedProductPageViews == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedProductPageViews[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedProductPageViews(name string, edges ...*ProductPageView) {
	if pr.Edges.namedProductPageViews == nil {
		pr.Edges.namedProductPageViews = make(map[string][]*ProductPageView)
	}
	if len(edges) == 0 {
		pr.Edges.namedProductPageViews[name] = []*ProductPageView{}
	} else {
		pr.Edges.namedProductPageViews[name] = append(pr.Edges.namedProductPageViews[name], edges...)
	}
}

// NamedBlogPosts returns the BlogPosts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedBlogPosts(name string) ([]*BlogPost, error) {
	if pr.Edges.namedBlogPosts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedBlogPosts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedBlogPosts(name string, edges ...*BlogPost) {
	if pr.Edges.namedBlogPosts == nil {
		pr.Edges.namedBlogPosts = make(map[string][]*BlogPost)
	}
	if len(edges) == 0 {
		pr.Edges.namedBlogPosts[name] = []*BlogPost{}
	} else {
		pr.Edges.namedBlogPosts[name] = append(pr.Edges.namedBlogPosts[name], edges...)
	}
}

// NamedMarketingCampaigns returns the MarketingCampaigns named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedMarketingCampaigns(name string) ([]*MarketingCampaign, error) {
	if pr.Edges.namedMarketingCampaigns == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedMarketingCampaigns[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedMarketingCampaigns(name string, edges ...*MarketingCampaign) {
	if pr.Edges.namedMarketingCampaigns == nil {
		pr.Edges.namedMarketingCampaigns = make(map[string][]*MarketingCampaign)
	}
	if len(edges) == 0 {
		pr.Edges.namedMarketingCampaigns[name] = []*MarketingCampaign{}
	} else {
		pr.Edges.namedMarketingCampaigns[name] = append(pr.Edges.namedMarketingCampaigns[name], edges...)
	}
}

// NamedChats returns the Chats named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedChats(name string) ([]*Chat, error) {
	if pr.Edges.namedChats == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedChats[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedChats(name string, edges ...*Chat) {
	if pr.Edges.namedChats == nil {
		pr.Edges.namedChats = make(map[string][]*Chat)
	}
	if len(edges) == 0 {
		pr.Edges.namedChats[name] = []*Chat{}
	} else {
		pr.Edges.namedChats[name] = append(pr.Edges.namedChats[name], edges...)
	}
}

// NamedEmailCampaign returns the EmailCampaign named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedEmailCampaign(name string) ([]*EmailCampaign, error) {
	if pr.Edges.namedEmailCampaign == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedEmailCampaign[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedEmailCampaign(name string, edges ...*EmailCampaign) {
	if pr.Edges.namedEmailCampaign == nil {
		pr.Edges.namedEmailCampaign = make(map[string][]*EmailCampaign)
	}
	if len(edges) == 0 {
		pr.Edges.namedEmailCampaign[name] = []*EmailCampaign{}
	} else {
		pr.Edges.namedEmailCampaign[name] = append(pr.Edges.namedEmailCampaign[name], edges...)
	}
}

// Products is a parsable slice of Product.
type Products []*Product
