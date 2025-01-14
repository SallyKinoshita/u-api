-- Create "company" table
CREATE TABLE `company` (`id` int NOT NULL AUTO_INCREMENT, `company_name` varchar(255) NOT NULL, `representative` varchar(255) NOT NULL, `phone` varchar(20) NOT NULL, `postal_code` varchar(8) NOT NULL, `address` text NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "business_partner" table
CREATE TABLE `business_partner` (`id` int NOT NULL AUTO_INCREMENT, `company_id` int NOT NULL, `company_name` varchar(255) NOT NULL, `representative` varchar(255) NOT NULL, `phone` varchar(20) NOT NULL, `postal_code` varchar(8) NOT NULL, `address` text NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), INDEX `company_business_partner_fk` (`company_id`), CONSTRAINT `company_business_partner_fk` FOREIGN KEY (`company_id`) REFERENCES `company` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "bank_account" table
CREATE TABLE `bank_account` (`id` int NOT NULL AUTO_INCREMENT, `business_partner_id` int NOT NULL, `bank_name` varchar(255) NOT NULL, `branch_name` varchar(255) NOT NULL, `account_number` varchar(20) NOT NULL, `account_name` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), INDEX `business_partner_bank_account_fk` (`business_partner_id`), CONSTRAINT `business_partner_bank_account_fk` FOREIGN KEY (`business_partner_id`) REFERENCES `business_partner` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "invoice" table
CREATE TABLE `invoice` (`id` int NOT NULL AUTO_INCREMENT, `company_id` int NOT NULL, `business_partner_id` int NOT NULL, `issue_date` date NOT NULL, `amount` decimal(10,2) NOT NULL, `fee` decimal(10,2) NOT NULL, `fee_rate` decimal(5,2) NOT NULL, `tax` decimal(10,2) NOT NULL, `tax_rate` decimal(5,2) NOT NULL, `total_amount` decimal(10,2) NOT NULL, `due_date` date NOT NULL, `status` varchar(20) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), INDEX `business_partner_invoice_fk` (`business_partner_id`), INDEX `company_invoice_fk` (`company_id`), INDEX `due_date_idx` (`due_date`), INDEX `issue_date_idx` (`issue_date`), CONSTRAINT `business_partner_invoice_fk` FOREIGN KEY (`business_partner_id`) REFERENCES `business_partner` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT `company_invoice_fk` FOREIGN KEY (`company_id`) REFERENCES `company` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "users" table
CREATE TABLE `users` (`id` int NOT NULL AUTO_INCREMENT, `company_id` int NOT NULL, `name` varchar(255) NOT NULL, `email` varchar(255) NOT NULL, `password` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), INDEX `company_user_fk` (`company_id`), UNIQUE INDEX `email_idx` (`email`), CONSTRAINT `company_user_fk` FOREIGN KEY (`company_id`) REFERENCES `company` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- Insert seed data

-- Insert seed data into "company"
INSERT INTO `company` (`id`, `company_name`, `representative`, `phone`, `postal_code`, `address`, `created_at`, `updated_at`)
VALUES
(1, 'Tech Solutions Inc.', 'John Doe', '123-456-7890', '1234567', '123 Tech Street, City, State', NOW(), NOW()),
(2, 'Business Corp.', 'Jane Smith', '987-654-3210', '7654321', '456 Business Avenue, City, State', NOW(), NOW());

-- Insert seed data into "business_partner"
INSERT INTO `business_partner` (`id`, `company_id`, `company_name`, `representative`, `phone`, `postal_code`, `address`, `created_at`, `updated_at`)
VALUES
(1, 1, 'Partner A', 'Alice Johnson', '111-222-3333', '1111111', '123 Partner Road, City, State', NOW(), NOW()),
(2, 1, 'Partner B', 'Bob Brown', '444-555-6666', '2222222', '456 Partner Lane, City, State', NOW(), NOW());

-- Insert seed data into "bank_account"
INSERT INTO `bank_account` (`id`, `business_partner_id`, `bank_name`, `branch_name`, `account_number`, `account_name`, `created_at`, `updated_at`)
VALUES
(1, 1, 'First Bank', 'Main Branch', '1234567890', 'Partner A Account', NOW(), NOW()),
(2, 2, 'Second Bank', 'Central Branch', '0987654321', 'Partner B Account', NOW(), NOW());

-- Insert seed data into "invoice"
INSERT INTO `invoice` (`id`, `company_id`, `business_partner_id`, `issue_date`, `amount`, `fee`, `fee_rate`, `tax`, `tax_rate`, `total_amount`, `due_date`, `status`, `created_at`, `updated_at`)
VALUES
(1, 1, 1, '2024-01-01', 1000.00, 50.00, 5.00, 100.00, 10.00, 1150.00, '2024-01-31', 'unpaid', NOW(), NOW()),
(2, 1, 2, '2024-01-05', 2000.00, 100.00, 5.00, 200.00, 10.00, 2300.00, '2024-02-05', 'paid', NOW(), NOW());

-- Insert seed data into "users"
INSERT INTO `users` (`id`, `company_id`, `name`, `email`, `password`, `created_at`, `updated_at`)
VALUES
(1, 1, 'Admin User', 'admin@techsolutions.com', 'securepassword123', NOW(), NOW()),
(2, 2, 'Manager User', 'manager@businesscorp.com', 'anothersecurepassword456', NOW(), NOW());
