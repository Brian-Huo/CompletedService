-- Payment Initialize --
-- QME TECH Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("1111111111111111", "QME Tech. Ltd.", "2030-05-21 00:00:00", "111");
-- WIX Corperate Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("0000000000000000", "WIX Corp.", "2030-05-21 00:00:00", "000");
-- AK Knowall Payment Details --
INSERT INTO b_payment (card_number, holder_name, expiry_time, security_code) values ("2222222222222222", "AK Knowall", "2050-05-21 00:00:00", "222");

-- Address Initialize --
INSERT INTO b_address (street, suburb, postcode, property, city, lat, lng, formatted) values("6 Dalgety Street", "Oakleigh", "3166", "Melbourne", -37.89205644303595, 145.08993779784805, "Unit 11, 6 Dalgety Street, Oakleigh VIC 3166, AU");

-- Company Initialize --
-- QME TECH Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, registered_address, deposite_rate, finance_status) values ("QME Tech. Pty Ltd.", 2, "", "0123456789", 1, 10, 1);
-- AK Knowall Company Details --
INSERT INTO b_company (company_name, payment_id, director_name, contact_details, registered_address, deposite_rate, finance_status) values ("AK Knowall", 3, "Shengzhe Zhang", "0420830301", 1, 10, 1);