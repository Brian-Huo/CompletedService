-- Base data table: region --
CREATE TABLE `b_region` (
    region_id int unsigned NOT NULL AUTO_INCREMENT,
    region_name varchar(255) NOT NULL,
    region_type varchar(255) NOT NULL, 
    postcode char(4) NOT NULL UNIQUE,
    state_code varchar(5) NOT NULL,
    state_name varchar(255) NOT NULL,
    charge_type int unsigned NOT NULL DEFAULT 0,
    charge_amount int unsigned NOT NULL DEFAULT 0,
    service_status tinyint unsigned NOT NULL DEFAULT 1,
    PRIMARY KEY(region_id)
);