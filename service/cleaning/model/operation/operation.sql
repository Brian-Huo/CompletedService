-- Base data table: operation (contractor-order) --
CREATE TABLE `b_operation` (
    operation_id int unsigned NOT NULL AUTO_INCREMENT,
    contractor_id int unsigned NOT NULL,
    order_id int unsigned NOT NULL,
    operation tinyint(3) NOT NULL,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(operation_id)
);
ALTER TABLE `b_operation` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id); 
ALTER TABLE `b_operation` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 