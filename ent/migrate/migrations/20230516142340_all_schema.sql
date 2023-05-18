-- Modify "notifications" table
ALTER TABLE "notifications" ALTER COLUMN "date_created" SET DEFAULT '2023-05-16T22:23:39+08:00', ALTER COLUMN "date_updated" SET DEFAULT '2023-05-16T22:23:39+08:00';
-- Modify "shipping_addresses" table
ALTER TABLE "shipping_addresses" ALTER COLUMN "date_created" SET DEFAULT '2023-05-16T22:23:39+08:00', ALTER COLUMN "date_updated" SET DEFAULT '2023-05-16T22:23:39+08:00';
-- Modify "users" table
ALTER TABLE "users" ALTER COLUMN "date_created" SET DEFAULT '2023-05-16T22:23:39+08:00', ALTER COLUMN "date_updated" SET DEFAULT '2023-05-16T22:23:39+08:00';
