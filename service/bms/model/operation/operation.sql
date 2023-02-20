-- Relation data table: operation (contractor-order) --
CREATE TABLE `r_operation` (
    operation_id int unsigned NOT NULL AUTO_INCREMENT,
    contractor_id int unsigned NOT NULL,
    order_id int unsigned NOT NULL,
    operation tinyint(3) NOT NULL,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE `b_operation` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(id); 
ALTER TABLE `b_operation` ADD FOREIGN KEY (order_id) REFERENCES b_order(id); 