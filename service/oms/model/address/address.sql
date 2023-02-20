-- Base data table: address --
CREATE TABLE `b_address` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    street varchar(255) NOT NULL,
    suburb varchar(50) NOT NULL,
    postcode char(4) NOT NULL,
    statecode varchar(50) NOT NULL,
    lat DECIMAL(10, 8) NOT NULL,
    lng DECIMAL(11, 8) NOT NULL,
    formatted varchar(255) NOT NULL,
    PRIMARY KEY(id)
);