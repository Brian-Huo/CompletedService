-- Base data table: employee --
CREATE TABLE `b_employee` (
    employee_id int unsigned NOT NULL AUTO_INCREMENT,
    employee_photo varchar(255),
    employee_name varchar(255) NOT NULL,
    contact_details char(10) NOT NULL,
    company_id int unsigned NOT NULL,
    link_code char(20) NOT NUll,
    work_status tinyint(3) default 0,
    order_id int unsigned,
    PRIMARY KEY(employee_id)
);
ALTER TABLE `b_employee` ADD FOREIGN KEY (company_id) REFERENCES b_company(company_id);