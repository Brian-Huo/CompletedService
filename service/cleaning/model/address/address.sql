-- Base data table: address --
CREATE TABLE `b_address` (
    address_id int unsigned NOT NULL AUTO_INCREMENT,
    street varchar(255) NOT NULL,
    suburb varchar(50) NOT NULL,
    postcode char(4) NOT NULL,
    property varchar(255) NOT NULL,
    city varchar(50) NOT NULL,
    lat DECIMAL(10, 8) NOT NULL,
    lng DECIMAL(11, 8) NOT NULL,
    formatted varchar(255) NOT NULL,
    PRIMARY KEY(address_id)
);
ALTER TABLE `b_address` ADD FOREIGN KEY (postcode) REFERENCES b_region(postcode); 
ALTER TABLE `b_address` ADD FOREIGN KEY (property) REFERENCES b_property(property_name);