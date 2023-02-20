-- Base data table: contractor --
CREATE TABLE `b_contractor` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    contractor_photo varchar(255),
    contractor_name varchar(255) NOT NULL,
    contractor_type tinyint(3) NOT NULL,
    contractor_phone varchar(15) NOT NULL UNIQUE,
    contractor_company int unsigned NOT NULL,
    contractor_address int unsigned,
    link_code char(64) NOT NULL,
    serve_range float unsigned NOT NULL DEFAULT 50.0,
    contractor_status tinyint(3) NOT NULL,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE `b_contractor` ADD FOREIGN KEY (contractor_company) REFERENCES b_company(id);
ALTER TABLE `b_contractor` ADD FOREIGN KEY (contractor_address) REFERENCES b_address(id);
-- Base data table indexes finance-contractor --
CREATE INDEX IDX_Company_con ON `b_contractor` (contractor_company);