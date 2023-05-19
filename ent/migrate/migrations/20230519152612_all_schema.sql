-- Create "bank_accounts" table
CREATE TABLE "bank_accounts" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "xid" bigint NOT NULL, "shop_bank_accounts" bigint NULL, "user_bank_accounts" bigint NULL, PRIMARY KEY ("id"));
-- Create "blog_posts" table
CREATE TABLE "blog_posts" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "title" character varying NOT NULL, "content" character varying NOT NULL, "date_created" character varying NOT NULL, "date_updated" character varying NOT NULL, "product_blog_posts" bigint NULL, PRIMARY KEY ("id"));
-- Create "categories" table
CREATE TABLE "categories" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "description" character varying NOT NULL, "product_categories" bigint NULL, PRIMARY KEY ("id"));
-- Create "chats" table
CREATE TABLE "chats" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "xid" bigint NOT NULL, "product_chats" bigint NULL, PRIMARY KEY ("id"));
-- Create "commission_structure_schemas" table
CREATE TABLE "commission_structure_schemas" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "description" character varying NOT NULL, "commission_value" character varying NOT NULL, "commission_percentage" character varying NOT NULL, "product_commission_structure" bigint NULL, PRIMARY KEY ("id"));
-- Create "content_blocks" table
CREATE TABLE "content_blocks" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "primary_message" character varying NOT NULL, "secondary_message" character varying NOT NULL, "primary_content_content_block" bigint NULL, PRIMARY KEY ("id"));
-- Create "email_campaigns" table
CREATE TABLE "email_campaigns" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "xid" bigint NOT NULL, "product_email_campaign" bigint NULL, PRIMARY KEY ("id"));
-- Create "group_buys" table
CREATE TABLE "group_buys" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "date_created" character varying NOT NULL, "product_price" bigint NOT NULL, "moq" bigint NOT NULL, "start_date" character varying NOT NULL, "end_date" character varying NOT NULL, "product_group_buys" bigint NULL, PRIMARY KEY ("id"));
-- Create "groups" table
CREATE TABLE "groups" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "hero_contents" table
CREATE TABLE "hero_contents" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "primary_message" character varying NOT NULL, "secondary_message" character varying NOT NULL, "product_page_view_hero_content" bigint NULL, PRIMARY KEY ("id"));
-- Create "images" table
CREATE TABLE "images" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "description" character varying NOT NULL, "url" character varying NOT NULL, "content_block_image" bigint NULL, "hero_content_image" bigint NULL, "product_images" bigint NULL, PRIMARY KEY ("id"));
-- Create "link_visits" table
CREATE TABLE "link_visits" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "date_created" character varying NOT NULL, "ip_address" character varying NOT NULL, "sale_value" bigint NOT NULL, "commission_earned" bigint NOT NULL, "referral_link_visits" bigint NULL, "transaction_origin_link" bigint NULL, "user_buyer_links_clicked" bigint NULL, PRIMARY KEY ("id"));
-- Create "marketing_campaigns" table
CREATE TABLE "marketing_campaigns" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "description" character varying NOT NULL, "consumer_purchase_value" character varying NOT NULL, "customer_application_logic" character varying NOT NULL, "initialisation_logic" character varying NOT NULL, "start_date" character varying NOT NULL, "end_date" character varying NOT NULL, "date_created" character varying NOT NULL, "date_updated" character varying NOT NULL, "product_marketing_campaigns" bigint NULL, PRIMARY KEY ("id"));
-- Create "notifications" table
CREATE TABLE "notifications" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "title" character varying NOT NULL, "content" character varying NOT NULL, "date_created" character varying NOT NULL DEFAULT '2023-05-19T23:26:11+08:00', "date_updated" character varying NOT NULL DEFAULT '2023-05-19T23:26:11+08:00', "read" boolean NOT NULL, "user_notifications" bigint NULL, PRIMARY KEY ("id"));
-- Create "payment_methods" table
CREATE TABLE "payment_methods" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "xid" bigint NOT NULL, "user_payment_methods" bigint NULL, PRIMARY KEY ("id"));
-- Create "primary_contents" table
CREATE TABLE "primary_contents" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "placeholder" bigint NULL, "product_page_view_primary_content" bigint NULL, PRIMARY KEY ("id"));
-- Create "product_attributes" table
CREATE TABLE "product_attributes" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" bigint NOT NULL, "description" bigint NOT NULL, "value" bigint NOT NULL, "product_product_attributes" bigint NULL, "product_variation_product_attributes" bigint NULL, PRIMARY KEY ("id"));
-- Create "product_page_views" table
CREATE TABLE "product_page_views" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "version" bigint NOT NULL, "product_product_page_views" bigint NULL, PRIMARY KEY ("id"));
-- Create "product_variations" table
CREATE TABLE "product_variations" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "description" character varying NOT NULL, "price" character varying NOT NULL, "product_variations" bigint NULL, PRIMARY KEY ("id"));
-- Create "products" table
CREATE TABLE "products" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "description" character varying NOT NULL, "price" character varying NOT NULL, "date_created" character varying NOT NULL, "date_updated" character varying NOT NULL, "category_products" bigint NULL, "group_buy_product" bigint NULL, "marketing_campaign_product" bigint NULL, "review_product" bigint NULL, "shop_products" bigint NULL, "transaction_product" bigint NULL, "user_influencer_products" bigint NULL, "view_analytics_product" bigint NULL, PRIMARY KEY ("id"));
-- Create "referral_links" table
CREATE TABLE "referral_links" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "xid" bigint NOT NULL, "name" character varying NOT NULL, "description" character varying NOT NULL, "link" character varying NOT NULL, "user_influencer_referral_links" bigint NULL, PRIMARY KEY ("id"));
-- Create "refund_transactions" table
CREATE TABLE "refund_transactions" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "refund_amount" character varying NOT NULL, "refund_currency" character varying NOT NULL, "refund_reason" character varying NOT NULL, "refund_status" character varying NOT NULL, "date_created" character varying NOT NULL, "date_updated" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "reviews" table
CREATE TABLE "reviews" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "type" character varying NOT NULL, "content" character varying NOT NULL, "rating" character varying NOT NULL, "date_created" character varying NOT NULL, "product_reviews" bigint NULL, "user_buyer_reviews" bigint NULL, "user_influencer_reviews" bigint NULL, PRIMARY KEY ("id"));
-- Create "reward_types" table
CREATE TABLE "reward_types" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "type" character varying NOT NULL, "val" bigint NOT NULL, "marketing_campaign_consumer_reward" bigint NULL, PRIMARY KEY ("id"));
-- Create "shipping_addresses" table
CREATE TABLE "shipping_addresses" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "address" character varying NOT NULL, "city" character varying NOT NULL, "state" character varying NOT NULL, "zip" character varying NOT NULL, "country" character varying NOT NULL, "date_created" character varying NOT NULL DEFAULT '2023-05-19T23:26:11+08:00', "date_updated" character varying NOT NULL DEFAULT '2023-05-19T23:26:11+08:00', "user_shipping_addresses" bigint NULL, PRIMARY KEY ("id"));
-- Create "shops" table
CREATE TABLE "shops" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "description" character varying NOT NULL, "product_shop" bigint NULL, "transaction_shop" bigint NULL, "user_seller_shops" bigint NULL, PRIMARY KEY ("id"));
-- Create "tags" table
CREATE TABLE "tags" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "description" character varying NOT NULL, "product_tags" bigint NULL, "user_influencer_tags" bigint NULL, PRIMARY KEY ("id"));
-- Create "transactions" table
CREATE TABLE "transactions" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "date_created" character varying NOT NULL, "date_updated" character varying NOT NULL, "status" character varying NOT NULL DEFAULT 'PENDING', "payment_method" character varying NOT NULL, "payment_status" character varying NOT NULL, "payment_id" character varying NOT NULL, "payment_amount" character varying NOT NULL, "payment_currency" character varying NOT NULL, "payment_date" character varying NOT NULL, "payment_fee" character varying NOT NULL, "payment_net" character varying NOT NULL, "payment_payer_email" character varying NOT NULL, "payment_payer_first_name" character varying NOT NULL, "payment_payer_last_name" character varying NOT NULL, "payment_payer_id" character varying NOT NULL, "payment_payer_status" character varying NOT NULL, "payment_receiver_email" character varying NOT NULL, "payment_receiver_id" character varying NOT NULL, "payment_tax" character varying NOT NULL, "payment_transaction_id" character varying NOT NULL, "payment_transaction_type" character varying NOT NULL, "payment_pending_reason" character varying NOT NULL, "payment_reason_code" character varying NOT NULL, "group_buy_transaction" bigint NULL, "refund_transactions_transaction" bigint NULL, "shop_transactions" bigint NULL, "user_buyer_transactions" bigint NULL, PRIMARY KEY ("id"));
-- Create "user_buyers" table
CREATE TABLE "user_buyers" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "placeholder" bigint NULL, "review_product_customer" bigint NULL, "transaction_product_customer" bigint NULL, PRIMARY KEY ("id"));
-- Create "user_influencers" table
CREATE TABLE "user_influencers" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "placeholder" bigint NULL, "transaction_product_influencer" bigint NULL, PRIMARY KEY ("id"));
-- Create "user_sellers" table
CREATE TABLE "user_sellers" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "brand_name" character varying NOT NULL, "blog_post_author" bigint NULL, "commission_structure_schema_product_seller" bigint NULL, "product_product_seller" bigint NULL, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "email" character varying NOT NULL, "phone" character varying NOT NULL, "address" character varying NOT NULL, "city" character varying NOT NULL, "state" character varying NOT NULL, "zip" character varying NOT NULL, "country" character varying NOT NULL, "date_created" character varying NOT NULL DEFAULT '2023-05-19T23:26:11+08:00', "date_updated" character varying NOT NULL DEFAULT '2023-05-19T23:26:11+08:00', "notification_recipient" bigint NULL, "user_buyer_user_profile" bigint NULL, "user_influencer_user_profile" bigint NULL, "user_seller_user_profile" bigint NULL, PRIMARY KEY ("id"));
-- Create "view_analytics" table
CREATE TABLE "view_analytics" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "views" bigint NOT NULL, "scrolls" bigint NOT NULL, "exits" bigint NOT NULL, "date_created" character varying NOT NULL, "product_page_view_view_analytics" bigint NULL, PRIMARY KEY ("id"));
-- Modify "bank_accounts" table
ALTER TABLE "bank_accounts" ADD CONSTRAINT "bank_accounts_shops_bankAccounts" FOREIGN KEY ("shop_bank_accounts") REFERENCES "shops" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "bank_accounts_users_bankAccounts" FOREIGN KEY ("user_bank_accounts") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "blog_posts" table
ALTER TABLE "blog_posts" ADD CONSTRAINT "blog_posts_products_blogPosts" FOREIGN KEY ("product_blog_posts") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "categories" table
ALTER TABLE "categories" ADD CONSTRAINT "categories_products_categories" FOREIGN KEY ("product_categories") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "chats" table
ALTER TABLE "chats" ADD CONSTRAINT "chats_products_chats" FOREIGN KEY ("product_chats") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "commission_structure_schemas" table
ALTER TABLE "commission_structure_schemas" ADD CONSTRAINT "commission_structure_schemas_products_commissionStructure" FOREIGN KEY ("product_commission_structure") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "content_blocks" table
ALTER TABLE "content_blocks" ADD CONSTRAINT "content_blocks_primary_contents_contentBlock" FOREIGN KEY ("primary_content_content_block") REFERENCES "primary_contents" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "email_campaigns" table
ALTER TABLE "email_campaigns" ADD CONSTRAINT "email_campaigns_products_emailCampaign" FOREIGN KEY ("product_email_campaign") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "group_buys" table
ALTER TABLE "group_buys" ADD CONSTRAINT "group_buys_products_groupBuys" FOREIGN KEY ("product_group_buys") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "hero_contents" table
ALTER TABLE "hero_contents" ADD CONSTRAINT "hero_contents_product_page_views_heroContent" FOREIGN KEY ("product_page_view_hero_content") REFERENCES "product_page_views" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "images" table
ALTER TABLE "images" ADD CONSTRAINT "images_content_blocks_image" FOREIGN KEY ("content_block_image") REFERENCES "content_blocks" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "images_hero_contents_image" FOREIGN KEY ("hero_content_image") REFERENCES "hero_contents" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "images_products_images" FOREIGN KEY ("product_images") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "link_visits" table
ALTER TABLE "link_visits" ADD CONSTRAINT "link_visits_referral_links_visits" FOREIGN KEY ("referral_link_visits") REFERENCES "referral_links" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "link_visits_transactions_originLink" FOREIGN KEY ("transaction_origin_link") REFERENCES "transactions" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "link_visits_user_buyers_linksClicked" FOREIGN KEY ("user_buyer_links_clicked") REFERENCES "user_buyers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "marketing_campaigns" table
ALTER TABLE "marketing_campaigns" ADD CONSTRAINT "marketing_campaigns_products_marketingCampaigns" FOREIGN KEY ("product_marketing_campaigns") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "notifications" table
ALTER TABLE "notifications" ADD CONSTRAINT "notifications_users_notifications" FOREIGN KEY ("user_notifications") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "payment_methods" table
ALTER TABLE "payment_methods" ADD CONSTRAINT "payment_methods_users_paymentMethods" FOREIGN KEY ("user_payment_methods") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "primary_contents" table
ALTER TABLE "primary_contents" ADD CONSTRAINT "primary_contents_product_page_views_primaryContent" FOREIGN KEY ("product_page_view_primary_content") REFERENCES "product_page_views" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "product_attributes" table
ALTER TABLE "product_attributes" ADD CONSTRAINT "product_attributes_product_variations_productAttributes" FOREIGN KEY ("product_variation_product_attributes") REFERENCES "product_variations" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "product_attributes_products_productAttributes" FOREIGN KEY ("product_product_attributes") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "product_page_views" table
ALTER TABLE "product_page_views" ADD CONSTRAINT "product_page_views_products_productPageViews" FOREIGN KEY ("product_product_page_views") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "product_variations" table
ALTER TABLE "product_variations" ADD CONSTRAINT "product_variations_products_variations" FOREIGN KEY ("product_variations") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "products" table
ALTER TABLE "products" ADD CONSTRAINT "products_categories_products" FOREIGN KEY ("category_products") REFERENCES "categories" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "products_group_buys_product" FOREIGN KEY ("group_buy_product") REFERENCES "group_buys" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "products_marketing_campaigns_product" FOREIGN KEY ("marketing_campaign_product") REFERENCES "marketing_campaigns" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "products_reviews_product" FOREIGN KEY ("review_product") REFERENCES "reviews" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "products_shops_products" FOREIGN KEY ("shop_products") REFERENCES "shops" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "products_transactions_product" FOREIGN KEY ("transaction_product") REFERENCES "transactions" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "products_user_influencers_products" FOREIGN KEY ("user_influencer_products") REFERENCES "user_influencers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "products_view_analytics_product" FOREIGN KEY ("view_analytics_product") REFERENCES "view_analytics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "referral_links" table
ALTER TABLE "referral_links" ADD CONSTRAINT "referral_links_user_influencers_referralLinks" FOREIGN KEY ("user_influencer_referral_links") REFERENCES "user_influencers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "reviews" table
ALTER TABLE "reviews" ADD CONSTRAINT "reviews_products_reviews" FOREIGN KEY ("product_reviews") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "reviews_user_buyers_reviews" FOREIGN KEY ("user_buyer_reviews") REFERENCES "user_buyers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "reviews_user_influencers_reviews" FOREIGN KEY ("user_influencer_reviews") REFERENCES "user_influencers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "reward_types" table
ALTER TABLE "reward_types" ADD CONSTRAINT "reward_types_marketing_campaigns_consumerReward" FOREIGN KEY ("marketing_campaign_consumer_reward") REFERENCES "marketing_campaigns" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "shipping_addresses" table
ALTER TABLE "shipping_addresses" ADD CONSTRAINT "shipping_addresses_users_shippingAddresses" FOREIGN KEY ("user_shipping_addresses") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "shops" table
ALTER TABLE "shops" ADD CONSTRAINT "shops_products_shop" FOREIGN KEY ("product_shop") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "shops_transactions_shop" FOREIGN KEY ("transaction_shop") REFERENCES "transactions" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "shops_user_sellers_shops" FOREIGN KEY ("user_seller_shops") REFERENCES "user_sellers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "tags" table
ALTER TABLE "tags" ADD CONSTRAINT "tags_products_tags" FOREIGN KEY ("product_tags") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "tags_user_influencers_tags" FOREIGN KEY ("user_influencer_tags") REFERENCES "user_influencers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "transactions" table
ALTER TABLE "transactions" ADD CONSTRAINT "transactions_group_buys_transaction" FOREIGN KEY ("group_buy_transaction") REFERENCES "group_buys" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "transactions_refund_transactions_transaction" FOREIGN KEY ("refund_transactions_transaction") REFERENCES "refund_transactions" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "transactions_shops_transactions" FOREIGN KEY ("shop_transactions") REFERENCES "shops" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "transactions_user_buyers_transactions" FOREIGN KEY ("user_buyer_transactions") REFERENCES "user_buyers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "user_buyers" table
ALTER TABLE "user_buyers" ADD CONSTRAINT "user_buyers_reviews_productCustomer" FOREIGN KEY ("review_product_customer") REFERENCES "reviews" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "user_buyers_transactions_productCustomer" FOREIGN KEY ("transaction_product_customer") REFERENCES "transactions" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "user_influencers" table
ALTER TABLE "user_influencers" ADD CONSTRAINT "user_influencers_transactions_productInfluencer" FOREIGN KEY ("transaction_product_influencer") REFERENCES "transactions" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "user_sellers" table
ALTER TABLE "user_sellers" ADD CONSTRAINT "user_sellers_blog_posts_author" FOREIGN KEY ("blog_post_author") REFERENCES "blog_posts" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "user_sellers_commission_structure_schemas_productSeller" FOREIGN KEY ("commission_structure_schema_product_seller") REFERENCES "commission_structure_schemas" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "user_sellers_products_productSeller" FOREIGN KEY ("product_product_seller") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "users" table
ALTER TABLE "users" ADD CONSTRAINT "users_notifications_recipient" FOREIGN KEY ("notification_recipient") REFERENCES "notifications" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "users_user_buyers_userProfile" FOREIGN KEY ("user_buyer_user_profile") REFERENCES "user_buyers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "users_user_influencers_userProfile" FOREIGN KEY ("user_influencer_user_profile") REFERENCES "user_influencers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, ADD CONSTRAINT "users_user_sellers_userProfile" FOREIGN KEY ("user_seller_user_profile") REFERENCES "user_sellers" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Modify "view_analytics" table
ALTER TABLE "view_analytics" ADD CONSTRAINT "view_analytics_product_page_views_viewAnalytics" FOREIGN KEY ("product_page_view_view_analytics") REFERENCES "product_page_views" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;