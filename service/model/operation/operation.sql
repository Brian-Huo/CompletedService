-- Base data table: operation (contractor-order) --
CREATE TABLE `b_operation` (
    operation_id int unsigned NOT NUll AUTO_INCREMENT,
    contractor_id int unsigned NOT NULL,
    order_id int unsigned NOT NUll,
    operation tinyint(1) NOT NUll,
    issue_date datetime NOT NULL,
    PRIMARY KEY(operation_id)
);
ALTER TABLE `b_operation` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id); 
ALTER TABLE `b_operation` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 