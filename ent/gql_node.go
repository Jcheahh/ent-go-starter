// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/bankaccount"
	"entdemo/ent/blogpost"
	"entdemo/ent/category"
	"entdemo/ent/chat"
	"entdemo/ent/commissionstructureschema"
	"entdemo/ent/contentblock"
	"entdemo/ent/emailcampaign"
	"entdemo/ent/group"
	"entdemo/ent/groupbuy"
	"entdemo/ent/herocontent"
	"entdemo/ent/image"
	"entdemo/ent/linkvisit"
	"entdemo/ent/marketingcampaign"
	"entdemo/ent/notification"
	"entdemo/ent/paymentmethod"
	"entdemo/ent/primarycontent"
	"entdemo/ent/product"
	"entdemo/ent/productattribute"
	"entdemo/ent/productpageview"
	"entdemo/ent/productvariation"
	"entdemo/ent/referrallink"
	"entdemo/ent/refundtransactions"
	"entdemo/ent/review"
	"entdemo/ent/rewardtype"
	"entdemo/ent/shippingaddress"
	"entdemo/ent/shop"
	"entdemo/ent/tag"
	"entdemo/ent/transaction"
	"entdemo/ent/user"
	"entdemo/ent/userbuyer"
	"entdemo/ent/userinfluencer"
	"entdemo/ent/userseller"
	"entdemo/ent/viewanalytics"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

var bankaccountImplementors = []string{"BankAccount", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*BankAccount) IsNode() {}

var blogpostImplementors = []string{"BlogPost", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*BlogPost) IsNode() {}

var categoryImplementors = []string{"Category", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Category) IsNode() {}

var chatImplementors = []string{"Chat", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Chat) IsNode() {}

var commissionstructureschemaImplementors = []string{"CommissionStructureSchema", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*CommissionStructureSchema) IsNode() {}

var contentblockImplementors = []string{"ContentBlock", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ContentBlock) IsNode() {}

var emailcampaignImplementors = []string{"EmailCampaign", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*EmailCampaign) IsNode() {}

var groupImplementors = []string{"Group", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Group) IsNode() {}

var groupbuyImplementors = []string{"GroupBuy", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*GroupBuy) IsNode() {}

var herocontentImplementors = []string{"HeroContent", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*HeroContent) IsNode() {}

var imageImplementors = []string{"Image", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Image) IsNode() {}

var linkvisitImplementors = []string{"LinkVisit", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*LinkVisit) IsNode() {}

var marketingcampaignImplementors = []string{"MarketingCampaign", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MarketingCampaign) IsNode() {}

var notificationImplementors = []string{"Notification", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Notification) IsNode() {}

var paymentmethodImplementors = []string{"PaymentMethod", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*PaymentMethod) IsNode() {}

var primarycontentImplementors = []string{"PrimaryContent", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*PrimaryContent) IsNode() {}

var productImplementors = []string{"Product", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Product) IsNode() {}

var productattributeImplementors = []string{"ProductAttribute", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ProductAttribute) IsNode() {}

var productpageviewImplementors = []string{"ProductPageView", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ProductPageView) IsNode() {}

var productvariationImplementors = []string{"ProductVariation", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ProductVariation) IsNode() {}

var referrallinkImplementors = []string{"ReferralLink", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ReferralLink) IsNode() {}

var refundtransactionsImplementors = []string{"RefundTransactions", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*RefundTransactions) IsNode() {}

var reviewImplementors = []string{"Review", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Review) IsNode() {}

var rewardtypeImplementors = []string{"RewardType", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*RewardType) IsNode() {}

var shippingaddressImplementors = []string{"ShippingAddress", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ShippingAddress) IsNode() {}

var shopImplementors = []string{"Shop", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Shop) IsNode() {}

var tagImplementors = []string{"Tag", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Tag) IsNode() {}

var transactionImplementors = []string{"Transaction", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Transaction) IsNode() {}

var userImplementors = []string{"User", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*User) IsNode() {}

var userbuyerImplementors = []string{"UserBuyer", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*UserBuyer) IsNode() {}

var userinfluencerImplementors = []string{"UserInfluencer", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*UserInfluencer) IsNode() {}

var usersellerImplementors = []string{"UserSeller", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*UserSeller) IsNode() {}

var viewanalyticsImplementors = []string{"ViewAnalytics", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*ViewAnalytics) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case bankaccount.Table:
		query := c.BankAccount.Query().
			Where(bankaccount.ID(id))
		query, err := query.CollectFields(ctx, bankaccountImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case blogpost.Table:
		query := c.BlogPost.Query().
			Where(blogpost.ID(id))
		query, err := query.CollectFields(ctx, blogpostImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case category.Table:
		query := c.Category.Query().
			Where(category.ID(id))
		query, err := query.CollectFields(ctx, categoryImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case chat.Table:
		query := c.Chat.Query().
			Where(chat.ID(id))
		query, err := query.CollectFields(ctx, chatImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case commissionstructureschema.Table:
		query := c.CommissionStructureSchema.Query().
			Where(commissionstructureschema.ID(id))
		query, err := query.CollectFields(ctx, commissionstructureschemaImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case contentblock.Table:
		query := c.ContentBlock.Query().
			Where(contentblock.ID(id))
		query, err := query.CollectFields(ctx, contentblockImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case emailcampaign.Table:
		query := c.EmailCampaign.Query().
			Where(emailcampaign.ID(id))
		query, err := query.CollectFields(ctx, emailcampaignImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case group.Table:
		query := c.Group.Query().
			Where(group.ID(id))
		query, err := query.CollectFields(ctx, groupImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case groupbuy.Table:
		query := c.GroupBuy.Query().
			Where(groupbuy.ID(id))
		query, err := query.CollectFields(ctx, groupbuyImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case herocontent.Table:
		query := c.HeroContent.Query().
			Where(herocontent.ID(id))
		query, err := query.CollectFields(ctx, herocontentImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case image.Table:
		query := c.Image.Query().
			Where(image.ID(id))
		query, err := query.CollectFields(ctx, imageImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case linkvisit.Table:
		query := c.LinkVisit.Query().
			Where(linkvisit.ID(id))
		query, err := query.CollectFields(ctx, linkvisitImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case marketingcampaign.Table:
		query := c.MarketingCampaign.Query().
			Where(marketingcampaign.ID(id))
		query, err := query.CollectFields(ctx, marketingcampaignImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case notification.Table:
		query := c.Notification.Query().
			Where(notification.ID(id))
		query, err := query.CollectFields(ctx, notificationImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case paymentmethod.Table:
		query := c.PaymentMethod.Query().
			Where(paymentmethod.ID(id))
		query, err := query.CollectFields(ctx, paymentmethodImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case primarycontent.Table:
		query := c.PrimaryContent.Query().
			Where(primarycontent.ID(id))
		query, err := query.CollectFields(ctx, primarycontentImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case product.Table:
		query := c.Product.Query().
			Where(product.ID(id))
		query, err := query.CollectFields(ctx, productImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case productattribute.Table:
		query := c.ProductAttribute.Query().
			Where(productattribute.ID(id))
		query, err := query.CollectFields(ctx, productattributeImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case productpageview.Table:
		query := c.ProductPageView.Query().
			Where(productpageview.ID(id))
		query, err := query.CollectFields(ctx, productpageviewImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case productvariation.Table:
		query := c.ProductVariation.Query().
			Where(productvariation.ID(id))
		query, err := query.CollectFields(ctx, productvariationImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case referrallink.Table:
		query := c.ReferralLink.Query().
			Where(referrallink.ID(id))
		query, err := query.CollectFields(ctx, referrallinkImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case refundtransactions.Table:
		query := c.RefundTransactions.Query().
			Where(refundtransactions.ID(id))
		query, err := query.CollectFields(ctx, refundtransactionsImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case review.Table:
		query := c.Review.Query().
			Where(review.ID(id))
		query, err := query.CollectFields(ctx, reviewImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case rewardtype.Table:
		query := c.RewardType.Query().
			Where(rewardtype.ID(id))
		query, err := query.CollectFields(ctx, rewardtypeImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case shippingaddress.Table:
		query := c.ShippingAddress.Query().
			Where(shippingaddress.ID(id))
		query, err := query.CollectFields(ctx, shippingaddressImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case shop.Table:
		query := c.Shop.Query().
			Where(shop.ID(id))
		query, err := query.CollectFields(ctx, shopImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case tag.Table:
		query := c.Tag.Query().
			Where(tag.ID(id))
		query, err := query.CollectFields(ctx, tagImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case transaction.Table:
		query := c.Transaction.Query().
			Where(transaction.ID(id))
		query, err := query.CollectFields(ctx, transactionImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		query, err := query.CollectFields(ctx, userImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case userbuyer.Table:
		query := c.UserBuyer.Query().
			Where(userbuyer.ID(id))
		query, err := query.CollectFields(ctx, userbuyerImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case userinfluencer.Table:
		query := c.UserInfluencer.Query().
			Where(userinfluencer.ID(id))
		query, err := query.CollectFields(ctx, userinfluencerImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case userseller.Table:
		query := c.UserSeller.Query().
			Where(userseller.ID(id))
		query, err := query.CollectFields(ctx, usersellerImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case viewanalytics.Table:
		query := c.ViewAnalytics.Query().
			Where(viewanalytics.ID(id))
		query, err := query.CollectFields(ctx, viewanalyticsImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case bankaccount.Table:
		query := c.BankAccount.Query().
			Where(bankaccount.IDIn(ids...))
		query, err := query.CollectFields(ctx, bankaccountImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case blogpost.Table:
		query := c.BlogPost.Query().
			Where(blogpost.IDIn(ids...))
		query, err := query.CollectFields(ctx, blogpostImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case category.Table:
		query := c.Category.Query().
			Where(category.IDIn(ids...))
		query, err := query.CollectFields(ctx, categoryImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case chat.Table:
		query := c.Chat.Query().
			Where(chat.IDIn(ids...))
		query, err := query.CollectFields(ctx, chatImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case commissionstructureschema.Table:
		query := c.CommissionStructureSchema.Query().
			Where(commissionstructureschema.IDIn(ids...))
		query, err := query.CollectFields(ctx, commissionstructureschemaImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case contentblock.Table:
		query := c.ContentBlock.Query().
			Where(contentblock.IDIn(ids...))
		query, err := query.CollectFields(ctx, contentblockImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case emailcampaign.Table:
		query := c.EmailCampaign.Query().
			Where(emailcampaign.IDIn(ids...))
		query, err := query.CollectFields(ctx, emailcampaignImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case group.Table:
		query := c.Group.Query().
			Where(group.IDIn(ids...))
		query, err := query.CollectFields(ctx, groupImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case groupbuy.Table:
		query := c.GroupBuy.Query().
			Where(groupbuy.IDIn(ids...))
		query, err := query.CollectFields(ctx, groupbuyImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case herocontent.Table:
		query := c.HeroContent.Query().
			Where(herocontent.IDIn(ids...))
		query, err := query.CollectFields(ctx, herocontentImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case image.Table:
		query := c.Image.Query().
			Where(image.IDIn(ids...))
		query, err := query.CollectFields(ctx, imageImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case linkvisit.Table:
		query := c.LinkVisit.Query().
			Where(linkvisit.IDIn(ids...))
		query, err := query.CollectFields(ctx, linkvisitImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case marketingcampaign.Table:
		query := c.MarketingCampaign.Query().
			Where(marketingcampaign.IDIn(ids...))
		query, err := query.CollectFields(ctx, marketingcampaignImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case notification.Table:
		query := c.Notification.Query().
			Where(notification.IDIn(ids...))
		query, err := query.CollectFields(ctx, notificationImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case paymentmethod.Table:
		query := c.PaymentMethod.Query().
			Where(paymentmethod.IDIn(ids...))
		query, err := query.CollectFields(ctx, paymentmethodImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case primarycontent.Table:
		query := c.PrimaryContent.Query().
			Where(primarycontent.IDIn(ids...))
		query, err := query.CollectFields(ctx, primarycontentImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case product.Table:
		query := c.Product.Query().
			Where(product.IDIn(ids...))
		query, err := query.CollectFields(ctx, productImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case productattribute.Table:
		query := c.ProductAttribute.Query().
			Where(productattribute.IDIn(ids...))
		query, err := query.CollectFields(ctx, productattributeImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case productpageview.Table:
		query := c.ProductPageView.Query().
			Where(productpageview.IDIn(ids...))
		query, err := query.CollectFields(ctx, productpageviewImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case productvariation.Table:
		query := c.ProductVariation.Query().
			Where(productvariation.IDIn(ids...))
		query, err := query.CollectFields(ctx, productvariationImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case referrallink.Table:
		query := c.ReferralLink.Query().
			Where(referrallink.IDIn(ids...))
		query, err := query.CollectFields(ctx, referrallinkImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case refundtransactions.Table:
		query := c.RefundTransactions.Query().
			Where(refundtransactions.IDIn(ids...))
		query, err := query.CollectFields(ctx, refundtransactionsImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case review.Table:
		query := c.Review.Query().
			Where(review.IDIn(ids...))
		query, err := query.CollectFields(ctx, reviewImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case rewardtype.Table:
		query := c.RewardType.Query().
			Where(rewardtype.IDIn(ids...))
		query, err := query.CollectFields(ctx, rewardtypeImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case shippingaddress.Table:
		query := c.ShippingAddress.Query().
			Where(shippingaddress.IDIn(ids...))
		query, err := query.CollectFields(ctx, shippingaddressImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case shop.Table:
		query := c.Shop.Query().
			Where(shop.IDIn(ids...))
		query, err := query.CollectFields(ctx, shopImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case tag.Table:
		query := c.Tag.Query().
			Where(tag.IDIn(ids...))
		query, err := query.CollectFields(ctx, tagImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case transaction.Table:
		query := c.Transaction.Query().
			Where(transaction.IDIn(ids...))
		query, err := query.CollectFields(ctx, transactionImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, userImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case userbuyer.Table:
		query := c.UserBuyer.Query().
			Where(userbuyer.IDIn(ids...))
		query, err := query.CollectFields(ctx, userbuyerImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case userinfluencer.Table:
		query := c.UserInfluencer.Query().
			Where(userinfluencer.IDIn(ids...))
		query, err := query.CollectFields(ctx, userinfluencerImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case userseller.Table:
		query := c.UserSeller.Query().
			Where(userseller.IDIn(ids...))
		query, err := query.CollectFields(ctx, usersellerImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case viewanalytics.Table:
		query := c.ViewAnalytics.Query().
			Where(viewanalytics.IDIn(ids...))
		query, err := query.CollectFields(ctx, viewanalyticsImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
