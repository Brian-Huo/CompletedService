-- Payment Initialize --
-- QME TECH Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("1111111111111111", "QME Tech. Ltd.", "2030-05-21 00:00:00", "111");
-- WIX Corperate Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("0000000000000000", "WIX Corp.", "2030-05-21 00:00:00", "000");
-- AK Knowall Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("2222222222222222", "AK Knowall", "2050-05-21 00:00:00", "222");

-- Category Initialize --
INSERT INTO b_category (category_addr, category_name, category_description) values ("DC", "Departure Cleaning", "Departure Cleaning is a complete cleaning of a vacation property between short term tenants.");
INSERT INTO b_category (category_addr, category_name, category_description) values ("GHC", "General House Cleaning", "General House Cleaning consists of the basic cleaning tasks that include sweeping, vacuuming, dusting, mopping, etc.");
INSERT INTO b_category (category_addr, category_name, category_description) values ("CC", "Commerical Cleaning", "Commerical Cleaning Commercial office cleaning companies use a wide variety of cleaning methods, chemicals, and equipment to facilitate and expedite the cleaning process.");
INSERT INTO b_category (category_addr, category_name, category_description) values ("OC", "Office Cleaning", "Office Cleaning include all internal, general and routine cleaning - including floors, tiles, partition walls, internal walls, suspended ceilings, lighting, furniture and cleaning, window cleaning, deep cleans of sanitary conveniences and washing facilities, kitchens and dining areas, consumables and feminine hygiene facilities as well as cleaning of telephones, IT, and other periodic cleaning as required.");

-- Service Initialize --
--  Departure Cleaning Standards --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Studio", "Departure cleaning - Studio. In a studio, the other living areas—kitchen, living room, and bedroom—are typically combined into one larger space.", 170);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "Studio (with carpet)", "Departure cleaning - Studio with carpet. In a studio, the other living areas—kitchen, living room, and bedroom—are typically combined into one larger space. (Combined with Carpet Steam Cleaning).", 220);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "1b1h1b", "Departure cleaning - 1b1h1b. 1 Bedroom 1 Hall 1 Bathroom.", 200);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "1b1h1b (with carpet)", "Departure cleaning - 1b1h1b with carpet. 1 Bedroom 1 Hall 1 Bathroom combined with Carpet Steam Cleaning.", 280);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "2b1h1b", "Departure cleaning - 2b1h1b. 2 Bedroom 1 Hall 1 Bathroom.", 240);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "2b1h1b (with carpet)", "Departure cleaning - 2b1h1b with carpet. 2 Bedroom 1 Hall 1 Bathroom combined with Carpet Steam Cleaning.", 360);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "2b1h2b", "Departure cleaning - 2b1h2b. 2 Bedroom 1 Hall 2 Bathroom.", 280);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "2b1h2b (with carpet)", "Departure cleaning - 2b1h2b with carpet. 2 Bedroom 1 Hall 2 Bathroom combined with Carpet Steam Cleaning.", 400);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "3b1h1b", "Departure cleaning - 3b1h1b. 3 Bedroom 1 Hall 1 Bathroom.", 320);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "3b1h1b (with carpet)", "Departure cleaning - 3b1h1b with carpet. 3 Bedroom 1 Hall 1 Bathroom combined with Carpet Steam Cleaning.", 450);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "3b1h2b", "Departure cleaning - 3b1h2b. 3 Bedroom 1 Hall 2 Bathroom.", 360);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "3b1h2b (with carpet)", "Departure cleaning - 3b1h2b with carpet. 3 Bedroom 1 Hall 2 Bathroom combined with Carpet Steam Cleaning.", 490);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b1h1b", "Departure cleaning - 4b1h1b. 4 Bedroom 1 Hall 1 Bathroom.", 380);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b1h1b (with carpet)", "Departure cleaning - 4b1h1b with carpet. 4 Bedroom 1 Hall 1 Bathroom combined with Carpet Steam Cleaning.", 530);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b1h2b", "Departure cleaning - 4b1h2b. 4 Bedroom 1 Hall 2 Bathroom.", 400);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b1h2b (with carpet)", "Departure cleaning - 4b1h2b with carpet. 4 Bedroom 1 Hall 2 Bathroom combined with Carpet Steam Cleaning.", 550);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b2h3b", "Departure cleaning - 4b2h3b. 4 Bedroom 2 Hall 3 Bathroom.", 450);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "4b2h3b (with carpet)", "Departure cleaning - 4b2h3b with carpet. 4 Bedroom 2 Hall 3 Bathroom combined with Carpet Steam Cleaning.", 600);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "5b2h3b", "Departure cleaning - 5b2h3b. 5 Bedroom 2 Hall 3 Bathroom.", 500);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Base Options", "5b2h3b (with carpet)", "Departure cleaning - 5b2h3b with carpet. 5 Bedroom 2 Hall 3 Bathroom combined with Carpet Steam Cleaning.", 650);
-- Departure Cleaning Additions --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Bedroom", "Departure cleaning - Bedroom", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Livingroom", "Departure cleaning - Livingroom", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Bathroom", "Departure cleaning - Bathroom", 60);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Toilet", "Departure cleaning - Toilet", 25);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Balcony", "Departure cleaning - Balcony", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Garage", "Departure cleaning - Garage", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Garden", "Departure cleaning - Garden", 100);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Microwave", "Departure cleaning - One Microwave", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Refrigerator", "Departure cleaning - One Refrigerator", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Dishwasher", "Departure cleaning - One Dishwasher", 40);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Louver", "Departure cleaning - Louver per piece", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Windows Glass", "Departure cleaning - Windows glass per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Windows Screen", "Departure cleaning - Windows screen per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Air Conditioning Filter", "Departure cleaning - Air Conditioning Filter per piece", 15);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Mildew on Wall", "Departure cleaning - Mildew per Wall", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Wall Stain", "Departure cleaning - Wall Stain per face", 20);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Garbage Packing/hr", "Departure cleaning - Garbage Packing per hour", 50);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Waste Transportation/cube", "Departure cleaning - Waste Transportation per cubic meter", 150);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Pet Hair Cleaning", "Departure cleaning - Pet Hair Cleaning", 80);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Carpet Stain", "Departure cleaning - Carpet Stain", 80);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (1, "Additional Options", "Glass Glue Replacement", "Departure cleaning - Glass Glue Replacement", 100);

-- General House Cleaning Standards --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Standard Plan(1p/2hr)", "General house cleaning- Standard Plan(2hr) with one professional cleaner in 2 hours working.", 110);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Base Options", "Standard Plan(2p/2hr)", "General house cleaning- Standard Plan(2hr) with two professional cleaner in 2 hours working.", 220);
-- General House Cleaning Additions --
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Additional Hour for 1p", "General house cleaning - Additional Hour for one professinal cleaner plan.", 55);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Additional Hour for 2p", "General house cleaning - Additional Hour for two professinal cleaner plan.", 110);
INSERT INTO b_service (service_type, service_scope, service_name, service_description, service_price) values (2, "Additional Options", "Carpet Steam (One Room)", "General house cleaning  - Carpet Steam (One Room)", 80);

-- Company Initialize --
-- QME TECH Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, deposite_rate, finance_status) values ("QME Tech. Pty Ltd.", 2, "", "0123456789", 10, 1);
-- AK Knowall Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, deposite_rate, finance_status) values ("AK Knowall", 3, "Shengzhe Zhang", "0420830301", 10, 1);