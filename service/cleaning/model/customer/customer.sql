-- Base data table: customer --
CREATE TABLE `b_customer` (
    customer_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_name varchar(50) NOT NULL,
    customer_type tinyint(3) NOT NULL,
    country_code char(2) NOT NULL,
    customer_phone varchar(15) NOT NULL UNIQUE,
    customer_email varchar(255) NOT NULL UNIQUE,
    PRIMARY KEY(customer_id)
);