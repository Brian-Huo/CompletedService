-- Base data table: subscribe group --
CREATE TABLE `b_subscribe_group` (
    group_id int unsigned NOT NULL AUTO_INCREMENT,
    category int unsigned NOT NULL,
    location varchar(5) NOT NULL,
    PRIMARY KEY(group_id),
    UNIQUE KEY(category, location)
);
ALTER TABLE `b_subscribe_group` ADD FOREIGN KEY (category) REFERENCES b_category(category_id);