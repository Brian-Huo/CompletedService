-- Base data table: region --
CREATE TABLE `b_region` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    region_name varchar(255) NOT NULL,
    region_type varchar(255) NOT NULL DEFAULT 'suburb', 
    postcode char(4) NOT NULL UNIQUE,
    charge_type int unsigned NOT NULL DEFAULT 0,
    charge_amount int unsigned NOT NULL DEFAULT 0,
    serve_status tinyint unsigned NOT NULL DEFAULT 1,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);