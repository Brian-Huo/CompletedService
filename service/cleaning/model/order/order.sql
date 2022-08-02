-- Base data table: order --
CREATE TABLE `b_order` (
    order_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_id int unsigned NOT NULL,
    address_id int unsigned NOT NULL,
    contractor_id int unsigned,
    finance_id int unsigned,
    category_id int unsigned NOT NULL,
    service_list mediumtext NOT NULL,
    deposite_payment int unsigned,
    deposite_amount float unsigned NOT NULL,
    deposite_date datetime,
    final_payment int unsigned,
    final_amount float unsigned NOT NULL,
    final_payment_date datetime,
    current_deposite_rate int(2) unsigned NOT NULL,
    item_amount float unsigned NOT NULL,
    gst_amount float unsigned NOT NULL,
    total_amount float unsigned NOT NULL,
    order_description longtext,
    post_date datetime NOT NULL,
    reserve_date datetime NOT NULL,
    finish_date datetime,
    status int(3) unsigned NOT NULL,
    urgant_flag tinyint(1) unsigned NOT NULL,
    PRIMARY KEY(order_id)
);
ALTER TABLE `b_order` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(customer_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (address_id) REFERENCES b_address(address_id);
ALTER TABLE `b_order` ADD FOREIGN KEY (finance_id) REFERENCES b_company(company_id);
ALTER TABLE `b_order` ADD FOREIGN KEY (category_id) REFERENCES b_category(category_id);
ALTER TABLE `b_order` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (deposite_payment) REFERENCES b_payment(payment_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (final_payment) REFERENCES b_payment(payment_id); 
-- Base data table indexes order-finance & order-contractor --
CREATE INDEX IDX_Finance ON `b_order` (finance_id);
CREATE INDEX IDX_Contractor ON `b_order` (contractor_id);