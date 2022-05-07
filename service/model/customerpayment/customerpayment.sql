-- Relationship data table: customer-payment --
CREATE TABLE `r_customer_payment` (
    customer_id int unsigned NOT NULL,
    payment_id int unsigned NOT NUll,
    update_date datetime NOT NULL,
    PRIMARY KEY(customer_id)
);
ALTER TABLE `r_customer_payment` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(customer_id); 
ALTER TABLE `r_customer_payment` ADD FOREIGN KEY (payment_id) REFERENCES b_payment(payment_id); 