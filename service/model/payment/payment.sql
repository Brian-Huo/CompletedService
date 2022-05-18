-- Base data table: payment --
CREATE TABLE `b_payment` (
    payment_id int unsigned NOT NULL AUTO_INCREMENT,
    card_number char(16) NOT NULL UNIQUE,
    holder_name varchar(255) NOT NULL,
    expiry_time datetime NOT NUll,
    security_code char(3) NOT NUll,
    PRIMARY KEY(payment_id)
);