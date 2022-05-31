-- Payment Initialize --
-- QME TECH Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("1111111111111111", "QME Tech. Ltd.", "2030-05-21 00:00:00", "111");
-- WIX Corperate Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("0000000000000000", "WIX Corp.", "2030-05-21 00:00:00", "000");
-- AK Knowall Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("2222222222222222", "AK Knowall", "2050-05-21 00:00:00", "222");

-- Category Initialize --
INSERT INTO b_category (category_name, category_description) values ("Departure Cleaning", "Departure Cleaning includes: ...");
INSERT INTO b_category (category_name, category_description) values ("General House Cleaning", "General House Cleaning includes: ...");
INSERT INTO b_category (category_name, category_description) values ("Commerical Cleaning", "Commerical Cleaning includes: ...");
INSERT INTO b_category (category_name, category_description) values ("Office Cleaning", "Office Cleaning includes: ...");

-- Service Initialize --
--  Departure Cleaning Standards --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Bedroom", "Departure cleaning - Bedroom", 45);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Livingroom", "Departure cleaning - Livingroom", 45);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Bathroom", "Departure cleaning - Bathroom", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Toilet", "Departure cleaning - Toilet", 30);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Balcony", "Departure cleaning - Balcony", 30);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Garage", "Departure cleaning - Garage", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Garden", "Departure cleaning - Garden", 50);
-- Departure Cleaning Additions --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Microwave", "Departure cleaning - Microwave", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Refrigerator", "Departure cleaning - Refrigerator", 25);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Dishwasher", "Departure cleaning - Dishwasher", 40);

-- General House Cleaning Standards --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Bedroom", "General house cleaning - Bedroom", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Livingroom", "General house cleaning  - Livingroom", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Bathroom", "General house cleaning  - Bathroom", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Toilet", "General house cleaning  - Toilet", 25);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Balcony", "General house cleaning  - Balcony", 25);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Garage", "General house cleaning  - Garage", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Garden", "General house cleaning  - Garden", 45);

-- Company Initialize --
-- QME TECH Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, deposite_rate, finance_status) values ("QME Tech. Ltd.", 1, "", "0123456789", 10, 1);
-- AK Knowall Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, deposite_rate, finance_status) values ("AK Knowall", 3, "Shengzhe Zhang", "0420830301", 10, 1);