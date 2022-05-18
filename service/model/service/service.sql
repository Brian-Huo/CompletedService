-- Base data table: service --
CREATE TABLE `b_service` (
    service_id int unsigned NOT NULL AUTO_INCREMENT,
    service_type varchar(255) NOT NULL,
    service_scope varchar(255) NOT NULL,
    service_name varchar(100) NOT NUll UNIQUE,
    service_description longtext NOT NUll,
    service_price float unsigned NOT NULL,
    PRIMARY KEY(service_id)
);