-- Base data table: order --
CREATE TABLE `b_order` (
    order_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_id int unsigned NOT NULL,
    address_id int unsigned NOT NULL,
    employee_id int unsigned NOT NULL,
    company_id int unsigned,
    service_list mediumtext NOT NUll,
    deposite_payment int unsigned NOT NUll,
    deposite_amount float unsigned NOT NUll,
    current_deposite_rate int(2) unsigned NOT NUll,
    deposite_date datetime NOT NULL,
    final_payment int unsigned,
    final_amount float unsigned NOT NULL,
    final_payment_date datetime,
    gst_amount float unsigned NOT NULL,
    total_fee float unsigned NOT NUll,
    order_description longtext,
    post_date datetime NOT NUll,
    reserve_date datetime NOT NUll,
    finish_date datetime,
    status int(3) unsigned NOT NULL,
    PRIMARY KEY(order_id)
);
ALTER TABLE `b_order` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(customer_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (company_id) REFERENCES b_company(company_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (employee_id) REFERENCES b_employee(employee_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (address_id) REFERENCES b_address(address_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (deposite_payment) REFERENCES b_payment(payment_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (final_payment) REFERENCES b_payment(payment_id); 