CREATE TABLE `b_address` (
    address_id int unsigned NOT NULL AUTO_INCREMENT,
    address_details varchar(255) NOT NULL,
    suburb varchar(50) NOT NULL,
    postcode char(4) NOT NULL,
    state_code char(3) NOT NULL,
    country varchar(255) DEFAULT 'Australia',
    PRIMARY KEY(address_id)
);