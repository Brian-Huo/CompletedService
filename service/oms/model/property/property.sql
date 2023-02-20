-- Base data table: property --
CREATE TABLE `b_property` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    property_name varchar(255) NOT NULL UNIQUE,
    property_description mediumtext NOT NULL,
    charge_type int unsigned NOT NULL DEFAULT 0,
    charge_amount int unsigned NOT NULL DEFAULT 0,
    serve_status tinyint unsigned NOT NULL DEFAULT 1,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);