-- Base data table: category (service type) --
CREATE TABLE `b_category` (
    category_id int unsigned NOT NULL AUTO_INCREMENT,
    category_name varchar(255) NOT NULL,
    PRIMARY KEY(category_id)
);