-- Base data table: service item --
CREATE TABLE `b_item` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    item_category int unsigned NOT NULL,
    item_scope varchar(255) NOT NULL,
    item_name varchar(255) NOT NULL,
    item_image varchar(255),
    item_description mediumtext NOT NULL,
    item_price float unsigned NOT NULL,
    item_discount int(2) NOT NULL DEFAULT 0,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE `b_item` ADD FOREIGN KEY (item_category) REFERENCES b_category(id);