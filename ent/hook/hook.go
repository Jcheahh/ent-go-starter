// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"entdemo/ent"
	"fmt"
)

// The BankAccountFunc type is an adapter to allow the use of ordinary
// function as BankAccount mutator.
type BankAccountFunc func(context.Context, *ent.BankAccountMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BankAccountFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BankAccountMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BankAccountMutation", m)
}

// The BlogPostFunc type is an adapter to allow the use of ordinary
// function as BlogPost mutator.
type BlogPostFunc func(context.Context, *ent.BlogPostMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BlogPostFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BlogPostMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BlogPostMutation", m)
}

// The CategoryFunc type is an adapter to allow the use of ordinary
// function as Category mutator.
type CategoryFunc func(context.Context, *ent.CategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CategoryMutation", m)
}

// The ChatFunc type is an adapter to allow the use of ordinary
// function as Chat mutator.
type ChatFunc func(context.Context, *ent.ChatMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ChatFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ChatMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ChatMutation", m)
}

// The CommissionStructureSchemaFunc type is an adapter to allow the use of ordinary
// function as CommissionStructureSchema mutator.
type CommissionStructureSchemaFunc func(context.Context, *ent.CommissionStructureSchemaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CommissionStructureSchemaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CommissionStructureSchemaMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CommissionStructureSchemaMutation", m)
}

// The ContentBlockFunc type is an adapter to allow the use of ordinary
// function as ContentBlock mutator.
type ContentBlockFunc func(context.Context, *ent.ContentBlockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ContentBlockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ContentBlockMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ContentBlockMutation", m)
}

// The EmailCampaignFunc type is an adapter to allow the use of ordinary
// function as EmailCampaign mutator.
type EmailCampaignFunc func(context.Context, *ent.EmailCampaignMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmailCampaignFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EmailCampaignMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmailCampaignMutation", m)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *ent.GroupMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GroupMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupMutation", m)
}

// The GroupBuyFunc type is an adapter to allow the use of ordinary
// function as GroupBuy mutator.
type GroupBuyFunc func(context.Context, *ent.GroupBuyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupBuyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GroupBuyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupBuyMutation", m)
}

// The HeroContentFunc type is an adapter to allow the use of ordinary
// function as HeroContent mutator.
type HeroContentFunc func(context.Context, *ent.HeroContentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HeroContentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.HeroContentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HeroContentMutation", m)
}

// The ImageFunc type is an adapter to allow the use of ordinary
// function as Image mutator.
type ImageFunc func(context.Context, *ent.ImageMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ImageFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ImageMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ImageMutation", m)
}

// The LinkVisitFunc type is an adapter to allow the use of ordinary
// function as LinkVisit mutator.
type LinkVisitFunc func(context.Context, *ent.LinkVisitMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LinkVisitFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LinkVisitMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LinkVisitMutation", m)
}

// The MarketingCampaignFunc type is an adapter to allow the use of ordinary
// function as MarketingCampaign mutator.
type MarketingCampaignFunc func(context.Context, *ent.MarketingCampaignMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MarketingCampaignFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MarketingCampaignMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MarketingCampaignMutation", m)
}

// The NotificationFunc type is an adapter to allow the use of ordinary
// function as Notification mutator.
type NotificationFunc func(context.Context, *ent.NotificationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NotificationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NotificationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NotificationMutation", m)
}

// The PaymentMethodFunc type is an adapter to allow the use of ordinary
// function as PaymentMethod mutator.
type PaymentMethodFunc func(context.Context, *ent.PaymentMethodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentMethodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PaymentMethodMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentMethodMutation", m)
}

// The PrimaryContentFunc type is an adapter to allow the use of ordinary
// function as PrimaryContent mutator.
type PrimaryContentFunc func(context.Context, *ent.PrimaryContentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PrimaryContentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PrimaryContentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PrimaryContentMutation", m)
}

// The ProductFunc type is an adapter to allow the use of ordinary
// function as Product mutator.
type ProductFunc func(context.Context, *ent.ProductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProductMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductMutation", m)
}

// The ProductAttributeFunc type is an adapter to allow the use of ordinary
// function as ProductAttribute mutator.
type ProductAttributeFunc func(context.Context, *ent.ProductAttributeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductAttributeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProductAttributeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductAttributeMutation", m)
}

// The ProductPageViewFunc type is an adapter to allow the use of ordinary
// function as ProductPageView mutator.
type ProductPageViewFunc func(context.Context, *ent.ProductPageViewMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductPageViewFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProductPageViewMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductPageViewMutation", m)
}

// The ProductVariationFunc type is an adapter to allow the use of ordinary
// function as ProductVariation mutator.
type ProductVariationFunc func(context.Context, *ent.ProductVariationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductVariationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProductVariationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductVariationMutation", m)
}

// The ReferralLinkFunc type is an adapter to allow the use of ordinary
// function as ReferralLink mutator.
type ReferralLinkFunc func(context.Context, *ent.ReferralLinkMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReferralLinkFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReferralLinkMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReferralLinkMutation", m)
}

// The RefundTransactionsFunc type is an adapter to allow the use of ordinary
// function as RefundTransactions mutator.
type RefundTransactionsFunc func(context.Context, *ent.RefundTransactionsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RefundTransactionsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RefundTransactionsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RefundTransactionsMutation", m)
}

// The ReviewFunc type is an adapter to allow the use of ordinary
// function as Review mutator.
type ReviewFunc func(context.Context, *ent.ReviewMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReviewFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReviewMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReviewMutation", m)
}

// The RewardTypeFunc type is an adapter to allow the use of ordinary
// function as RewardType mutator.
type RewardTypeFunc func(context.Context, *ent.RewardTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RewardTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RewardTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RewardTypeMutation", m)
}

// The ShippingAddressFunc type is an adapter to allow the use of ordinary
// function as ShippingAddress mutator.
type ShippingAddressFunc func(context.Context, *ent.ShippingAddressMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ShippingAddressFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ShippingAddressMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ShippingAddressMutation", m)
}

// The ShopFunc type is an adapter to allow the use of ordinary
// function as Shop mutator.
type ShopFunc func(context.Context, *ent.ShopMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ShopFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ShopMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ShopMutation", m)
}

// The TagFunc type is an adapter to allow the use of ordinary
// function as Tag mutator.
type TagFunc func(context.Context, *ent.TagMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TagFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TagMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TagMutation", m)
}

// The TransactionFunc type is an adapter to allow the use of ordinary
// function as Transaction mutator.
type TransactionFunc func(context.Context, *ent.TransactionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TransactionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TransactionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TransactionMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The UserBuyerFunc type is an adapter to allow the use of ordinary
// function as UserBuyer mutator.
type UserBuyerFunc func(context.Context, *ent.UserBuyerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserBuyerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserBuyerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserBuyerMutation", m)
}

// The UserInfluencerFunc type is an adapter to allow the use of ordinary
// function as UserInfluencer mutator.
type UserInfluencerFunc func(context.Context, *ent.UserInfluencerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserInfluencerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserInfluencerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserInfluencerMutation", m)
}

// The UserSellerFunc type is an adapter to allow the use of ordinary
// function as UserSeller mutator.
type UserSellerFunc func(context.Context, *ent.UserSellerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserSellerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserSellerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserSellerMutation", m)
}

// The ViewAnalyticsFunc type is an adapter to allow the use of ordinary
// function as ViewAnalytics mutator.
type ViewAnalyticsFunc func(context.Context, *ent.ViewAnalyticsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ViewAnalyticsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ViewAnalyticsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ViewAnalyticsMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
