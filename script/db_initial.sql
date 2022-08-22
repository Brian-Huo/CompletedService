-- Payment Initialize --
-- QME TECH Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("1111111111111111", "QME Tech. Ltd.", "2030-05-21 00:00:00", "111");
-- WIX Corperate Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("0000000000000000", "WIX Corp.", "2030-05-21 00:00:00", "000");
-- AK Knowall Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("2222222222222222", "AK Knowall", "2050-05-21 00:00:00", "222");

-- Category Initialize --
INSERT INTO b_category (category_addr, category_name, category_description) values ("DC", "Departure Cleaning", "Departure Cleaning includes: ...");
INSERT INTO b_category (category_addr, category_name, category_description) values ("GHC", "General House Cleaning", "General House Cleaning includes: ...");
INSERT INTO b_category (category_addr, category_name, category_description) values ("CC", "Commerical Cleaning", "Commerical Cleaning includes: ...");
INSERT INTO b_category (category_addr, category_name, category_description) values ("OC", "Office Cleaning", "Office Cleaning includes: ...");

-- Service Initialize --
--  Departure Cleaning Standards --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Studio", "Departure cleaning - Studio", 187);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Studio (with carpet)", "Departure cleaning - Studio with carpet", 242);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "1b1h1b", "Departure cleaning - 1b1h1b", 220);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "1b1h1b (with carpet)", "Departure cleaning - 1b1h1b with carpet", 308);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "2b1h1b", "Departure cleaning - 2b1h1b", 264);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "2b1h1b (with carpet)", "Departure cleaning - 2b1h1b with carpet", 396);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "2b1h2b", "Departure cleaning - 2b1h2b", 308);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "2b1h2b (with carpet)", "Departure cleaning - 2b1h2b with carpet", 440);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "3b1h1b", "Departure cleaning - 3b1h1b", 352);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "3b1h1b (with carpet)", "Departure cleaning - 3b1h1b with carpet", 495);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "3b1h2b", "Departure cleaning - 3b1h2b", 396);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "3b1h2b (with carpet)", "Departure cleaning - 3b1h2b with carpet", 539);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b1h1b", "Departure cleaning - 4b1h1b", 418);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b1h1b (with carpet)", "Departure cleaning - 4b1h1b with carpet", 583);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b1h2b", "Departure cleaning - 4b1h2b", 440);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b1h2b (with carpet)", "Departure cleaning - 4b1h2b with carpet", 605);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b2h3b", "Departure cleaning - 4b2h3b", 495);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b2h3b (with carpet)", "Departure cleaning - 4b2h3b with carpet", 660);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "5b2h3b", "Departure cleaning - 5b2h3b", 550);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "5b2h3b (with carpet)", "Departure cleaning - 5b2h3b with carpet", 715);
-- Departure Cleaning Additions --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Bedroom", "Departure cleaning - Bedroom", 45);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Livingroom", "Departure cleaning - Livingroom", 45);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Bathroom", "Departure cleaning - Bathroom", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Toilet", "Departure cleaning - Toilet", 30);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Balcony", "Departure cleaning - Balcony", 30);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Garage", "Departure cleaning - Garage", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Garden", "Departure cleaning - Garden", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Microwave", "Departure cleaning - Microwave", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Refrigerator", "Departure cleaning - Refrigerator", 25);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Dishwasher", "Departure cleaning - Dishwasher", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Louver", "Departure cleaning - Louver per piece", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Windows Glass", "Departure cleaning - Windows glass per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Windows Screen", "Departure cleaning - Windows screen per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Air Conditioning Filter", "Departure cleaning - Air Conditioning Filter per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Mildew on Wall", "Departure cleaning - Mildew on Wall", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Wall Stain", "Departure cleaning - Wall Stain per face", 10);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Garbage Packing", "Departure cleaning - Garbage Packing per hour", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Waste Transportation", "Departure cleaning - Waste Transportation per cubic meter", 150);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Pet Hair Cleaning", "Departure cleaning - Pet Hair Cleaning", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Carpet Stain", "Departure cleaning - Carpet Stain", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Glass Glue Replacement", "Departure cleaning - Glass Glue Replacement", 110);

-- General House Cleaning Standards --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Standard Plan(2hr)", "General house cleaning- Standard Plan(2hr)", 121);
-- General House Cleaning Additions --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Additional Hour", "General house cleaning - Additional Hour", 60);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Carpet Steam (One Room)", "General house cleaning  - Carpet Steam (One Room)", 66);

-- Company Initialize --
-- QME TECH Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, deposite_rate, finance_status) values ("QME Tech. Pty Ltd.", 2, "", "0123456789", 10, 1);
-- AK Knowall Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, deposite_rate, finance_status) values ("AK Knowall", 3, "Shengzhe Zhang", "0420830301", 10, 1);