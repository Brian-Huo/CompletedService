-- Base data table: operation (employee-order) --
CREATE TABLE `b_operation` (
    operation_id int unsigned NOT NUll AUTO_INCREMENT,
    employee_id int unsigned NOT NULL,
    order_id int unsigned NOT NUll,
    operation tinyint(1) NOT NUll,
    issue_date datetime NOT NULL,
    PRIMARY KEY(operation_id)
);
ALTER TABLE `b_operation` ADD FOREIGN KEY (employee_id) REFERENCES b_employee(employee_id); 
ALTER TABLE `b_operation` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 