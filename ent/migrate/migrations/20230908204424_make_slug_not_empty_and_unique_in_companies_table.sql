-- Modify "companies" table
ALTER TABLE "companies" ALTER COLUMN "slug" SET NOT NULL;
-- Create index "companies_slug_key" to table: "companies"
CREATE UNIQUE INDEX "companies_slug_key" ON "companies" ("slug");
