CREATE TABLE `b_company` (
    company_id int unsigned NOT NULL AUTO_INCREMENT,
    company_name varchar(255) NOT NULL,
    payment_id int unsigned,
    director_name varchar(255),
    contact_details char(10) NOT NULL,
    registered_address int unsigned,
    deposite_rate int(2) unsigned NOT NULL,
    PRIMARY KEY(company_id)
);