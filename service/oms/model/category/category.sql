-- Base data table: category (service type) --
CREATE TABLE `b_category` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    category_addr varchar(10) NOT NULL,
    category_name varchar(100) NOT NULL UNIQUE,
    category_description mediumtext NOT NULL,
    serve_range float unsigned NOT NULL DEFAULT 50.0,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);