package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("description").Annotations(
			entgql.OrderField("DESCRIPTION"),
		),
		field.String("price").Annotations(
			entgql.OrderField("PRICE"),
		),
		field.String("dateCreated").Annotations(
			entgql.OrderField("DATECREATED"),
		),
		field.String("dateUpdated").Annotations(
			entgql.OrderField("DATEUPDATED"),
		),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("productSeller", UserSeller.Type),
		edge.To("reviews", Review.Type),
		edge.To("images", Image.Type),
		edge.To("categories", Category.Type),
		edge.To("tags", Tag.Type),
		edge.To("productAttributes", ProductAttribute.Type),
		edge.To("variations", ProductVariation.Type),
		edge.To("commissionStructure", CommissionStructureSchema.Type),
		edge.To("shop", Shop.Type).Required(),
		edge.To("groupBuys", GroupBuy.Type),
		edge.To("productPageViews", ProductPageView.Type),
		edge.To("blogPosts", BlogPost.Type),
		edge.To("marketingCampaigns", MarketingCampaign.Type),
		edge.To("chats", Chat.Type),
		edge.To("emailCampaign", EmailCampaign.Type),
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
