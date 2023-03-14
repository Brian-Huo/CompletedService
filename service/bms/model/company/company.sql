-- Base data table: company --
CREATE TABLE `b_company` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    contactor varchar(255),
    company_name varchar(255) NOT NULL,
    company_phone varchar(15) NOT NULL UNIQUE,
    company_address int unsigned NOT NULL,
    company_deposit int(2) unsigned NOT NULL,
    company_status tinyint(1) NOT NULL DEFAULT 1,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE `b_company` ADD FOREIGN KEY (company_address) REFERENCES b_address(id);