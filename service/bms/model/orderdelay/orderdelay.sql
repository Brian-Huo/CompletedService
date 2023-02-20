-- Raletion data table: order-delay (order message) --
CREATE TABLE `r_order_delay` (
    contractor_id int unsigned NOT NULL,
    order_id int unsigned NOT NUll,
    PRIMARY KEY(contractor_id, order_id)
);
ALTER TABLE `r_order_delay` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id); 
ALTER TABLE `r_order_delay` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 