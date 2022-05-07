CREATE TABLE `r_employee_service` (
    employee_id int unsigned NOT NULL,
    service_id int unsigned NOT NUll,
    PRIMARY KEY(employee_id)
);
ALTER TABLE `r_employee_service` ADD FOREIGN KEY (employee_id) REFERENCES b_employee(employee_id); 
ALTER TABLE `r_employee_service` ADD FOREIGN KEY (service_id) REFERENCES b_service(service_id); 