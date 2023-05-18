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
		field.String("name"),
		field.String("description"),
		field.String("price"),
		field.String("dateCreated"),
		field.String("dateUpdated"),
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
		edge.To("commissionStructure", CommissionStructure.Type),
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
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
