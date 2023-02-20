-- Base data table: customer --
CREATE TABLE `b_customer` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    customer_name varchar(50) NOT NULL,
    customer_type tinyint(3) NOT NULL,
    country_code char(2) NOT NULL,
    customer_phone varchar(15) NOT NULL UNIQUE,
    customer_email varchar(50) NOT NULL UNIQUE,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);