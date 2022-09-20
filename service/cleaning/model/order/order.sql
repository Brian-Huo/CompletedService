-- Base data table: order --
CREATE TABLE `b_order` (
    order_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_id int unsigned NOT NULL,
    address_id int unsigned NOT NULL,
    contractor_id int unsigned,
    finance_id int unsigned,
    category_id int unsigned NOT NULL,
    basic_items JSON NOT NULL,
    additional_items JSON,
    order_description mediumtext,
    order_comments mediumtext,
    current_deposite_rate int(2) unsigned NOT NULL,
    deposite_amount float unsigned NOT NULL,
    final_amount float unsigned NOT NULL,
    item_amount float unsigned NOT NULL,
    gst_amount float unsigned NOT NULL,
    surcharge_item varchar(255) NOT NULL DEFAULT 'None',
    surcharge_rate int unsigned NOT NULL DEFAULT 0,
    surcharge_amount float unsigned NOT NULL DEFAULT 0,
    surcharge_description mediumtext,
    total_amount float unsigned NOT NULL,
    balance_amount float unsigned NOT NULL DEFAULT 0,
    post_date datetime NOT NULL,
    reserve_date datetime NOT NULL,
    finish_date datetime,
    payment_date datetime,
    status int(3) unsigned NOT NULL,
    urgant_flag tinyint(1) unsigned NOT NULL,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    CHECK (JSON_VALID(basic_items)),
    CHECK (JSON_VALID(additional_items)),
    PRIMARY KEY(order_id)
);
ALTER TABLE `b_order` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(customer_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (address_id) REFERENCES b_address(address_id);
ALTER TABLE `b_order` ADD FOREIGN KEY (finance_id) REFERENCES b_company(company_id);
ALTER TABLE `b_order` ADD FOREIGN KEY (category_id) REFERENCES b_category(category_id);
ALTER TABLE `b_order` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id); 
-- Base data table indexes order-finance & order-contractor --
CREATE INDEX IDX_Finance_order ON `b_order` (finance_id);
CREATE INDEX IDX_Contractor_order ON `b_order` (contractor_id);