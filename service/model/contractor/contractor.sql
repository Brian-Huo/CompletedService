-- Base data table: contractor --
CREATE TABLE `b_contractor` (
    contractor_id int unsigned NOT NULL AUTO_INCREMENT,
    contractor_photo varchar(255),
    contractor_name varchar(255) NOT NULL,
    contractor_type tinyint(3) NOT NUll,
    contact_details char(10) NOT NULL UNIQUE,
    finance_id int unsigned NOT NULL,
    address_id int unsigned,
    link_code char(64) NOT NUll,
    work_status tinyint(3) NOT NUll,
    order_id int unsigned,
    PRIMARY KEY(contractor_id)
);
ALTER TABLE `b_contractor` ADD FOREIGN KEY (finance_id) REFERENCES b_company(company_id);
ALTER TABLE `b_contractor` ADD FOREIGN KEY (address_id) REFERENCES b_address(address_id);