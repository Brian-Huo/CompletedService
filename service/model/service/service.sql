-- Base data table: service --
CREATE TABLE `b_service` (
    service_id int unsigned NOT NULL AUTO_INCREMENT,
    service_type int unsigned NOT NULL,
    service_scope varchar(255) NOT NULL,
    service_name varchar(100) NOT NULL,
    service_photo varchar(255),
    service_description longtext NOT NULL,
    service_price float unsigned NOT NULL,
    PRIMARY KEY(service_id)
);
ALTER TABLE `b_service` ADD FOREIGN KEY (service_type) REFERENCES b_category(category_id);