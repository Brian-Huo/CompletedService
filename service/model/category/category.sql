-- Base data table: category (service type) --
CREATE TABLE `b_category` (
    category_id int unsigned NOT NULL AUTO_INCREMENT,
    category_name varchar(100) NOT NULL UNIQUE,
    category_description mediumtext NOT NULL,
    serve_range float unsigned NOT NULL DEFAULT 50.0,
    PRIMARY KEY(category_id)
);