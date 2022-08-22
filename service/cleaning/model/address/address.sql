-- Base data table: address --
CREATE TABLE `b_address` (
    address_id int unsigned NOT NULL AUTO_INCREMENT,
    street varchar(255) NOT NULL,
    suburb varchar(50) NOT NULL,
    postcode char(4) NOT NULL,
    city varchar(50) NOT NULL,
    state_code char(3) NOT NULL,
    country char(2) NOT NULL DEFAULT 'AU',
    lat DECIMAL(10, 8) NOT NULL,
    lng DECIMAL(11, 8) NOT NULL,
    formatted varchar(255) NOT NULL,
    PRIMARY KEY(address_id)
);
