-- Base data table: schedule (order message) --
CREATE TABLE `b_schedule` (
    contractor_id int unsigned NOT NULL,
    order_id int unsigned NOT NUll,
    PRIMARY KEY(contractor_id, order_id)
);
ALTER TABLE `b_schedule` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id); 
ALTER TABLE `b_schedule` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 