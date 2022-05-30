ALTER DATABASE cleaningservice_test CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

/*Drop table commands*/
Drop table b_operation;
Drop table b_order;

Drop table r_subscribe_record;
Drop table b_contractor;

Drop table b_service;
Drop table b_category;

Drop table b_company;
Drop table b_customer;
Drop table b_payment;
Drop table b_address;

/*Create table command*/
-- Base data table: address --
CREATE TABLE `b_address` (
    address_id int unsigned NOT NULL AUTO_INCREMENT,
    street varchar(255) NOT NULL,
    suburb varchar(50) NOT NULL,
    postcode char(4) NOT NULL,
    city varchar(5) NOT NULL,
    state_code char(3) NOT NULL,
    country char(2) NOT NULL DEFAULT 'AU',
    lat DECIMAL(10, 8) NOT NULL,
    lng DECIMAL(11, 8) NOT NULL,
    formatted varchar(255) NOT NULL,
    PRIMARY KEY(address_id)
);

-- Base data table: payment --
CREATE TABLE `b_payment` (
    payment_id int unsigned NOT NULL AUTO_INCREMENT,
    card_number char(16) NOT NULL UNIQUE,
    holder_name varchar(255) NOT NULL,
    expiry_time datetime NOT NULL,
    security_code char(3) NOT NULL,
    PRIMARY KEY(payment_id)
);

-- Base data table: customer --
CREATE TABLE `b_customer` (
    customer_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_name varchar(50) NOT NULL,
    customer_type tinyint(3) NOT NULL,
    country_code char(2) NOT NULL,
    contact_details varchar(15) NOT NULL UNIQUE,
    PRIMARY KEY(customer_id)
);

-- Base data table: company/finance --
CREATE TABLE `b_company` (
    company_id int unsigned NOT NULL AUTO_INCREMENT,
    company_name varchar(255) NOT NULL,
    payment_id int unsigned,
    director_name varchar(255),
    contact_details varchar(15) NOT NULL UNIQUE,
    registered_address int unsigned,
    deposite_rate int(2) unsigned NOT NULL,
    finance_status tinyint(1) NOT NULL,
    PRIMARY KEY(company_id)
);
ALTER TABLE `b_company` ADD FOREIGN KEY (payment_id) REFERENCES b_payment(payment_id); 
ALTER TABLE `b_company` ADD FOREIGN KEY (registered_address) REFERENCES b_address(address_id);

-- Base data table: category (service type) --
CREATE TABLE `b_category` (
    category_id int unsigned NOT NULL AUTO_INCREMENT,
    category_name varchar(100) NOT NULL UNIQUE,
    category_description mediumtext NOT NULL,
    serve_range float unsigned NOT NULL,
    PRIMARY KEY(category_id)
);

-- Base data table: service --
CREATE TABLE `b_service` (
    service_id int unsigned NOT NULL AUTO_INCREMENT,
    service_type int unsigned NOT NULL,
    service_scope varchar(255) NOT NULL,
    service_name varchar(255) NOT NULL,
    service_photo varchar(255),
    service_description longtext NOT NULL,
    service_price float unsigned NOT NULL,
    PRIMARY KEY(service_id)
);
ALTER TABLE `b_service` ADD FOREIGN KEY (service_type) REFERENCES b_category(category_id);

-- Base data table: contractor --
CREATE TABLE `b_contractor` (
    contractor_id int unsigned NOT NULL AUTO_INCREMENT,
    contractor_photo varchar(255),
    contractor_name varchar(255) NOT NULL,
    contractor_type tinyint(3) NOT NULL,
    contact_details varchar(15) NOT NULL UNIQUE,
    finance_id int unsigned NOT NULL,
    address_id int unsigned,
    link_code char(64) NOT NULL,
    work_status tinyint(3) NOT NULL,
    order_id int unsigned,
    category_list varchar(255),
    PRIMARY KEY(contractor_id)
);
ALTER TABLE `b_contractor` ADD FOREIGN KEY (finance_id) REFERENCES b_company(company_id);
ALTER TABLE `b_contractor` ADD FOREIGN KEY (address_id) REFERENCES b_address(address_id);

-- Relation data table: subscribe record --
CREATE TABLE `r_subscribe_record` (
    category_id int unsigned NOT NULL,
    contractor_id int unsigned NOT NULL,
    PRIMARY KEY(category_id, contractor_id)
);
ALTER TABLE `r_subscribe_record` ADD FOREIGN KEY (category_id) REFERENCES b_category(category_id);
ALTER TABLE `r_subscribe_record` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id);

-- Base data table: order --
CREATE TABLE `b_order` (
    order_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_id int unsigned NOT NULL,
    address_id int unsigned NOT NULL,
    contractor_id int unsigned,
    finance_id int unsigned,
    category_id int unsigned NOT NULL,
    service_list mediumtext NOT NULL,
    deposite_payment int unsigned NOT NULL,
    deposite_amount float unsigned NOT NULL,
    deposite_date datetime NOT NULL,
    final_payment int unsigned,
    final_amount float unsigned NOT NULL,
    final_payment_date datetime,
    current_deposite_rate int(2) unsigned NOT NULL,
    gst_amount float unsigned NOT NULL,
    total_fee float unsigned NOT NULL,
    order_description longtext,
    post_date datetime NOT NULL,
    reserve_date datetime NOT NULL,
    finish_date datetime,
    status int(3) unsigned NOT NULL,
    urgant_flag tinyint(1) unsigned NOT NULL,
    PRIMARY KEY(order_id)
);
ALTER TABLE `b_order` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(customer_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (address_id) REFERENCES b_address(address_id);
ALTER TABLE `b_order` ADD FOREIGN KEY (finance_id) REFERENCES b_company(company_id);
ALTER TABLE `b_order` ADD FOREIGN KEY (category_id) REFERENCES b_category(category_id);
ALTER TABLE `b_order` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (deposite_payment) REFERENCES b_payment(payment_id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (final_payment) REFERENCES b_payment(payment_id); 

-- Base data table: operation (contractor-order) --
CREATE TABLE `b_operation` (
    operation_id int unsigned NOT NULL AUTO_INCREMENT,
    contractor_id int unsigned NOT NULL,
    order_id int unsigned NOT NULL,
    operation tinyint(3) NOT NULL,
    issue_date datetime NOT NULL,
    PRIMARY KEY(operation_id)
);
ALTER TABLE `b_operation` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id); 
ALTER TABLE `b_operation` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 