ALTER DATABASE CleanService CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

/*Drop table commands*/
Drop table r_employee_service;
Drop table r_employee_order;
Drop table b_employee;

Drop table r_order_design;
Drop table b_order;

Drop table r_company_payment;
Drop table r_company_service;
Drop table b_company;

Drop table r_customer_payment;
Drop table r_customer_address;
Drop table b_customer;

Drop table b_service;
Drop table b_payment;
Drop table b_address;


/*Create table command*/
-- Base data table: address --
CREATE TABLE `b_address` (
    address_id int unsigned NOT NULL AUTO_INCREMENT,
    address_details varchar(255) NOT NULL,
    suburb varchar(50) NOT NULL,
    postcode char(4) NOT NULL,
    state_code char(3) NOT NULL,
    country varchar(255),
    PRIMARY KEY(address_id)
);

-- Base data table: payment --
CREATE TABLE `b_payment` (
    payment_id int unsigned NOT NULL AUTO_INCREMENT,
    PRIMARY KEY(payment_id)
);

-- Base data table: service --
CREATE TABLE `b_service` (
    service_id int unsigned NOT NULL AUTO_INCREMENT,
    service_type varchar(255) NOT NULL,
    service_description longtext,
    PRIMARY KEY(service_id)
);

-- Base data table: customer --
CREATE TABLE `b_customer` (
    customer_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_name varchar(50) NOT NULL,
    country_code char(2) NOT NULL,
    contact_details varchar(15) NOT NULL,
    PRIMARY KEY(customer_id)
);

-- Relationship data table: customer-address --
CREATE TABLE `r_customer_address` (
    customer_id int unsigned NOT NULL,
    address_id int unsigned NOT NUll,
    update_date datetime,
    PRIMARY KEY(customer_id, address_id)
);
ALTER TABLE `r_customer_address` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(customer_id); 
ALTER TABLE `r_customer_address` ADD FOREIGN KEY (address_id) REFERENCES b_address(address_id); 

-- Relationship data table: customer-payment --
CREATE TABLE `r_customer_payment` (
    customer_id int unsigned NOT NULL,
    payment_id int unsigned NOT NUll,
    update_date datetime,
    PRIMARY KEY(customer_id, payment_id)
);
ALTER TABLE `r_customer_payment` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(customer_id); 
ALTER TABLE `r_customer_payment` ADD FOREIGN KEY (payment_id) REFERENCES b_payment(payment_id); 

-- Base data table: company --
CREATE TABLE `b_company` (
    company_id int unsigned NOT NULL AUTO_INCREMENT,
    company_name varchar(255) NOT NULL,
    payment_id int unsigned NOT NUll,
    director_name varchar(255),
    contact_details char(10) NOT NULL,
    registered_address int unsigned,
    deposite_rate int(2) unsigned NOT NULL,
    PRIMARY KEY(company_id)
);
ALTER TABLE `b_company` ADD FOREIGN KEY (payment_id) REFERENCES b_payment(payment_id); 
ALTER TABLE `b_company` ADD FOREIGN KEY (registered_address) REFERENCES b_address(address_id); 

-- Relationship data table: company-service --
CREATE TABLE `r_company_service` (
    design_id int unsigned NOT NUll AUTO_INCREMENT,
    company_id int unsigned NOT NULL,
    service_id int unsigned NOT NUll,
    price int unsigned NOT NUll,
    comments longtext,
    PRIMARY KEY(design_id)
);
ALTER TABLE `r_company_service` ADD FOREIGN KEY (company_id) REFERENCES b_company(company_id); 
ALTER TABLE `r_company_service` ADD FOREIGN KEY (service_id) REFERENCES b_service(service_id); 

-- Relationship data table: customer-payment --
CREATE TABLE `r_company_payment` (
    company_id int unsigned NOT NULL,
    payment_id int unsigned NOT NUll,
    update_date datetime,
    PRIMARY KEY(company_id, payment_id)
);
ALTER TABLE `r_company_payment` ADD FOREIGN KEY (company_id) REFERENCES b_company(company_id);  
ALTER TABLE `r_customer_payment` ADD FOREIGN KEY (payment_id) REFERENCES b_payment(payment_id); 

-- Base data table: order --
CREATE TABLE `b_order` (
    order_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_id int unsigned NOT NULL,
    company_id int unsigned NOT NUll,
    address_id int unsigned NOT NULL,
    deposite_payment int unsigned,
    deposite_amount int unsigned NOT NUll,
    current_deposite_rate int(2) unsigned NOT NUll,
    deposite_date datetime,
    final_payment int unsigned,
    final_amount int unsigned NOT NULL,
    final_payment_date datetime,
    total_fee int unsigned NOT NUll,
    order_description longtext,
    post_date datetime NOT NUll,
    reserve_date datetime NOT NUll,
    finish_date datetime,
    status tinyint(3) NOT NULL,
    PRIMARY KEY(order_id)
);
ALTER TABLE `b_order` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(customer_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (company_id) REFERENCES b_company(company_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (address_id) REFERENCES b_address(address_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (deposite_payment) REFERENCES b_payment(payment_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (final_payment) REFERENCES b_payment(payment_id); 

-- Relationship data table: order-design --
CREATE TABLE `r_order_design` (
    order_id int unsigned NOT NULL,
    design_id int unsigned NOT NUll,
    cur_price int unsigned NOT NUll,
    comments longtext,
    PRIMARY KEY(order_id, design_id)
);
ALTER TABLE `r_order_design` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 
ALTER TABLE `r_order_design` ADD FOREIGN KEY (design_id) REFERENCES r_company_service(design_id); 

-- Base data table: employee --
CREATE TABLE `b_employee` (
    employee_id int unsigned NOT NULL AUTO_INCREMENT,
    employee_photo varchar(255),
    employee_name varchar(255) NOT NULL,
    contact_details char(10) NOT NULL,
    company_id int unsigned NOT NULL,
    link_code char(20) NOT NUll,
    work_status tinyint(2) default 0,
    serve_id int unsigned,
    PRIMARY KEY(employee_id)
);
ALTER TABLE `b_employee` ADD FOREIGN KEY (company_id) REFERENCES b_company(company_id);

-- Relationship data table: employee-service --
CREATE TABLE `r_employee_service` (
    employee_id int unsigned NOT NULL,
    service_id int unsigned NOT NUll,
    PRIMARY KEY(employee_id, serve_id)
);
ALTER TABLE `r_employee_service` ADD FOREIGN KEY (employee_id) REFERENCES b_employee(employee_id); 
ALTER TABLE `r_employee_service` ADD FOREIGN KEY (service_id) REFERENCES b_service(service_id); 

-- Relationship data table: employee-order --
CREATE TABLE `r_employee_order` (
    operation_id int unsigned NOT NUll AUTO_INCREMENT,
    employee_id int unsigned NOT NULL,
    order_id int unsigned NOT NUll,
    operation tinyint(1) NOT NUll,
    issue_date datetime NOT NULL,
    PRIMARY KEY(operation_id)
);
ALTER TABLE `r_employee_order` ADD FOREIGN KEY (employee_id) REFERENCES b_employee(employee_id); 
ALTER TABLE `r_employee_order` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 