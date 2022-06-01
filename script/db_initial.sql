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
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Bedroom", "Departure cleaning - Bedroom", 45);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Livingroom", "Departure cleaning - Livingroom", 45);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Bathroom", "Departure cleaning - Bathroom", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Toilet", "Departure cleaning - Toilet", 30);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Balcony", "Departure cleaning - Balcony", 30);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Garage", "Departure cleaning - Garage", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Garden", "Departure cleaning - Garden", 50);
-- Departure Cleaning Additions --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Microwave", "Departure cleaning - Microwave", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Refrigerator", "Departure cleaning - Refrigerator", 25);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Dishwasher", "Departure cleaning - Dishwasher", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Louver", "Departure cleaning - Louver per piece", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Windows Glass", "Departure cleaning - Windows glass per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Windows Screen", "Departure cleaning - Windows screen per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Air Conditioning Filter", "Departure cleaning - Air Conditioning Filter per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Mildew on Wall", "Departure cleaning - Mildew on Wall", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Wall Stain", "Departure cleaning - Wall Stain per face", 10);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Garbage Packing", "Departure cleaning - Garbage Packing per hour", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Waste Transportation", "Departure cleaning - Waste Transportation per cubic meter", 150);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Pet Hair Cleaning", "Departure cleaning - Pet Hair Cleaning", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Carpet Stain", "Departure cleaning - Carpet Stain", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Glass Glue Replacement", "Departure cleaning - Glass Glue Replacement", 110);

-- General House Cleaning Standards --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Bedroom", "General house cleaning - Bedroom", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Livingroom", "General house cleaning  - Livingroom", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Bathroom", "General house cleaning  - Bathroom", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Toilet", "General house cleaning  - Toilet", 25);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Balcony", "General house cleaning  - Balcony", 25);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Garage", "General house cleaning  - Garage", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Garden", "General house cleaning  - Garden", 45);
-- General House Cleaning Additions --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Microwave", "General House cleaning - Microwave", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Refrigerator", "General House cleaning - Refrigerator", 25);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Dishwasher", "General House cleaning - Dishwasher", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Louver", "General House cleaning - Louver per piece", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Windows Glass", "General House cleaning - Windows glass per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Windows Screen", "General House cleaning - Windows screen per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Air Conditioning Filter", "General Housee cleaning - Air Conditioning Filter per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Mildew on Wall", "General House cleaning - Mildew on Wall", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Wall Stain", "General House cleaning - Wall Stain per face", 10);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Garbage Packing", "General House cleaning - Garbage Packing per hour", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Waste Transportation", "General House cleaning - Waste Transportation per cubic meter", 150);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Pet Hair Cleaning", "General House cleaning - Pet Hair Cleaning", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Carpet Stain", "General House cleaning - Carpet Stain", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Glass Glue Replacement", "General House cleaning - Glass Glue Replacement", 110);

-- Company Initialize --
-- QME TECH Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, deposite_rate, finance_status) values ("QME Tech. Ltd.", 2, "", "0123456789", 10, 2);
-- AK Knowall Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, deposite_rate, finance_status) values ("AK Knowall", 3, "Shengzhe Zhang", "0420830301", 10, 2);