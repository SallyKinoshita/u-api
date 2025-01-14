schema "main" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}

table "company" {
  schema = schema.main
  column "id" {
    type = int
    auto_increment = true
  }
  column "company_name" {
    type = varchar(255)
    null = false
  }
  column "representative" {
    type = varchar(255)
    null = false
  }
  column "phone" {
    type = varchar(20)
    null = false
  }
  column "postal_code" {
    type = varchar(8)
    null = false
  }
  column "address" {
    type = text
    null = false
  }
  column "created_at" {
    type = timestamp
    null = false
  }
  column "updated_at" {
    type = timestamp
    null = false
  }

  primary_key {
    columns = [column.id]
  }
}

table "users" {
  schema = schema.main
  column "id" {
    type = int
    auto_increment = true
  }
  column "company_id" {
    type = int
    null = false
  }
  column "name" {
    type = varchar(255)
    null = false
  }
  column "email" {
    type = varchar(255)
    null = false
  }
  column "password" {
    type = varchar(255)
    null = false
  }
  column "created_at" {
    type = timestamp
    null = false
  }
  column "updated_at" {
    type = timestamp
    null = false
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "company_user_fk" {
    columns = [column.company_id]
    ref_columns = [table.company.column.id]
    on_delete = CASCADE
  }

  index "email_idx" {
    columns = [column.email]
    unique = true
  }
}

table "business_partner" {
  schema = schema.main
  column "id" {
    type = int
    auto_increment = true
  }
  column "company_id" {
    type = int
    null = false
  }
  column "company_name" {
    type = varchar(255)
    null = false
  }
  column "representative" {
    type = varchar(255)
    null = false
  }
  column "phone" {
    type = varchar(20)
    null = false
  }
  column "postal_code" {
    type = varchar(8)
    null = false
  }
  column "address" {
    type = text
    null = false
  }
  column "created_at" {
    type = timestamp
    null = false
  }
  column "updated_at" {
    type = timestamp
    null = false
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "company_business_partner_fk" {
    columns = [column.company_id]
    ref_columns = [table.company.column.id]
    on_delete = CASCADE
  }
}

table "bank_account" {
  schema = schema.main
  column "id" {
    type = int
    auto_increment = true
  }
  column "business_partner_id" {
    type = int
    null = false
  }
  column "bank_name" {
    type = varchar(255)
    null = false
  }
  column "branch_name" {
    type = varchar(255)
    null = false
  }
  column "account_number" {
    type = varchar(20)
    null = false
  }
  column "account_name" {
    type = varchar(255)
    null = false
  }
  column "created_at" {
    type = timestamp
    null = false
  }
  column "updated_at" {
    type = timestamp
    null = false
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "business_partner_bank_account_fk" {
    columns = [column.business_partner_id]
    ref_columns = [table.business_partner.column.id]
    on_delete = CASCADE
  }
}

table "invoice" {
  schema = schema.main
  column "id" {
    type = int
    auto_increment = true
  }
  column "company_id" {
    type = int
    null = false
  }
  column "business_partner_id" {
    type = int
    null = false
  }
  column "issue_date" {
    type = date
    null = false
  }
  column "amount" {
    type = decimal(10, 2)
    null = false
  }
  column "fee" {
    type = decimal(10, 2)
    null = false
  }
  column "fee_rate" {
    type = decimal(5, 2)
    null = false
  }
  column "tax" {
    type = decimal(10, 2)
    null = false
  }
  column "tax_rate" {
    type = decimal(5, 2)
    null = false
  }
  column "total_amount" {
    type = decimal(10, 2)
    null = false
  }
  column "due_date" {
    type = date
    null = false
  }
  column "status" {
    type = varchar(20)
    null = false
  }
  column "created_at" {
    type = timestamp
    null = false
  }
  column "updated_at" {
    type = timestamp
    null = false
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "company_invoice_fk" {
    columns = [column.company_id]
    ref_columns = [table.company.column.id]
    on_delete = CASCADE
  }

  foreign_key "business_partner_invoice_fk" {
    columns = [column.business_partner_id]
    ref_columns = [table.business_partner.column.id]
    on_delete = CASCADE
  }

  index "issue_date_idx" {
    columns = [column.issue_date]
  }

  index "due_date_idx" {
    columns = [column.due_date]
  }
}
