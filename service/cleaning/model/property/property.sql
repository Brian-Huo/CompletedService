-- Base data table: property --
CREATE TABLE `b_property` (
    property_id int unsigned NOT NULL AUTO_INCREMENT,
    property_name varchar(255) NOT NULL UNIQUE,
    property_description mediumtext NOT NULL,
    charge_type int unsigned NOT NULL DEFAULT 0,
    charge_amount int unsigned NOT NULL DEFAULT 0,
    service_status tinyint unsigned NOT NULL,
    PRIMARY KEY(property_id)
);