-- Base data table: order --
CREATE TABLE `b_order` (
    order_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_id int unsigned NOT NULL,
    address_id int unsigned NOT NULL,
    contractor_id int unsigned,
    company_id int unsigned,
    category_id int unsigned NOT NULL,
    service_items JSON NOT NULL,
    surcharge_items JSON,
    item_amount float unsigned NOT NULL,
    gst_amount float unsigned NOT NULL,
    deposit_rate int(2) unsigned NOT NULL,
    deposit_amount float unsigned NOT NULL,
    final_amount float unsigned NOT NULL,
    surcharge_amount float unsigned NOT NULL,
    total_amount float unsigned NOT NULL,
    balance_amount float unsigned NOT NULL DEFAULT 0,
    order_description mediumtext,
    order_comments mediumtext,
    post_date datetime NOT NULL,
    reserve_date datetime NOT NULL,
    finish_date datetime,
    payment_date datetime,
    order_status int(3) unsigned NOT NULL,
    urgant_flag tinyint(1) unsigned NOT NULL,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    CHECK (JSON_VALID(service_items)),
    CHECK (JSON_VALID(surcharge_items)),
    PRIMARY KEY(id)
);
ALTER TABLE `b_order` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (address_id) REFERENCES b_address(id);
ALTER TABLE `b_order` ADD FOREIGN KEY (company_id) REFERENCES b_company(id);
ALTER TABLE `b_order` ADD FOREIGN KEY (category_id) REFERENCES b_category(id);
ALTER TABLE `b_order` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(id); 

-- Base data table indexes order-company & order-contractor --
CREATE INDEX IDX_Company_order ON `b_order` (company_id);
CREATE INDEX IDX_Contractor_order ON `b_order` (contractor_id);