ALTER DATABASE cleaningservice_test CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

/*Drop table commands*/
DROP TABLE r_operation;
DROP TABLE b_order;

DROP TABLE r_subscription;
DROP TABLE b_contractor;

DROP TABLE b_item;
DROP TABLE b_category;

DROP TABLE b_address;
DROP TABLE b_company;
DROP TABLE b_customer;
DROP TABLE b_property;
DROP TABLE b_region;
DROP TABLE b_surcharge;

/*Create table command*/
-- Base data table: surcharge --
CREATE TABLE `b_surcharge` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    surcharge_name varchar(255) NOT NULL,
    surcharge_description mediumtext NOT NULL,
    charge_type int unsigned NOT NULL DEFAULT 0,
    charge_amount int unsigned NOT NULL DEFAULT 0,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

-- Base data table: region --
CREATE TABLE `b_region` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    region_name varchar(255) NOT NULL,
    region_type varchar(255) NOT NULL DEFAULT 'suburb', 
    postcode char(4) NOT NULL UNIQUE,
    charge_type int unsigned NOT NULL DEFAULT 0,
    charge_amount int unsigned NOT NULL DEFAULT 0,
    serve_status tinyint unsigned NOT NULL DEFAULT 1,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

-- Base data table: property --
CREATE TABLE `b_property` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    property_name varchar(255) NOT NULL UNIQUE,
    property_description mediumtext NOT NULL,
    charge_type int unsigned NOT NULL DEFAULT 0,
    charge_amount int unsigned NOT NULL DEFAULT 0,
    serve_status tinyint unsigned NOT NULL DEFAULT 1,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

-- Base data table: address --
CREATE TABLE `b_address` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    unit varchar(50),
    street varchar(255) NOT NULL,
    suburb varchar(50) NOT NULL,
    postcode char(4) NOT NULL,
    statecode varchar(50) NOT NULL,
    lat DECIMAL(10, 8) NOT NULL,
    lng DECIMAL(11, 8) NOT NULL,
    formatted varchar(255) NOT NULL,
    PRIMARY KEY(id)
);

-- Base data table: customer --
CREATE TABLE `b_customer` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    customer_name varchar(50) NOT NULL,
    customer_type tinyint(3) NOT NULL,
    country_code char(2) NOT NULL,
    customer_phone varchar(15) NOT NULL UNIQUE,
    customer_email varchar(50) NOT NULL UNIQUE,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

-- Base data table: company --
CREATE TABLE `b_company` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    contractor varchar(255),
    company_name varchar(255) NOT NULL,
    company_phone varchar(15) NOT NULL UNIQUE,
    company_address int unsigned NOT NULL,
    company_deposit int(2) unsigned NOT NULL,
    company_status tinyint(1) NOT NULL DEFAULT 1,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE `b_company` ADD FOREIGN KEY (company_address) REFERENCES b_address(id);

-- Base data table: category (service type) --
CREATE TABLE `b_category` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    category_addr varchar(10) NOT NULL,
    category_name varchar(100) NOT NULL UNIQUE,
    category_description mediumtext NOT NULL,
    serve_range float unsigned NOT NULL DEFAULT 50.0,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

-- Base data table: service item --
CREATE TABLE `b_item` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    item_category int unsigned NOT NULL,
    item_scope varchar(255) NOT NULL,
    item_name varchar(255) NOT NULL,
    item_image varchar(255),
    item_description mediumtext NOT NULL,
    item_price float unsigned NOT NULL,
    item_discount int(2) NOT NULL DEFAULT 0,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE `b_item` ADD FOREIGN KEY (item_category) REFERENCES b_category(id);

-- Base data table: contractor --
CREATE TABLE `b_contractor` (
    id int unsigned NOT NULL AUTO_INCREMENT,
    contractor_photo varchar(255),
    contractor_name varchar(255) NOT NULL,
    contractor_type tinyint(3) NOT NULL,
    contractor_phone varchar(15) NOT NULL UNIQUE,
    contractor_company int unsigned NOT NULL,
    contractor_address int unsigned,
    link_code char(64) NOT NULL,
    serve_range float unsigned NOT NULL DEFAULT 50.0,
    contractor_status tinyint(3) NOT NULL,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE `b_contractor` ADD FOREIGN KEY (contractor_company) REFERENCES b_company(id);
ALTER TABLE `b_contractor` ADD FOREIGN KEY (contractor_address) REFERENCES b_address(id);
-- Base data table indexes company-contractor --
CREATE INDEX IDX_Company_con ON `b_contractor` (contractor_company);

-- Relation data table: subscription --
CREATE TABLE `r_subscription` (
    subscription_id int unsigned NOT NULL AUTO_INCREMENT,
    category_id int unsigned NOT NULL,
    contractor_id int unsigned NOT NULL,
    UNIQUE UNI_subscribe (category_id, contractor_id),
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(subscription_id)
);
ALTER TABLE `r_subscription` ADD FOREIGN KEY (category_id) REFERENCES b_category(id);
ALTER TABLE `r_subscription` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(id);

-- Base data table: order --
CREATE TABLE `b_order` (
    order_id int unsigned NOT NULL AUTO_INCREMENT,
    customer_id int unsigned NOT NULL,
    address_id int unsigned NOT NULL,
    contractor_id int unsigned,
    company_id int unsigned,
    category_id int unsigned NOT NULL,
    service_items JSON NOT NULL,
    surcharge_items JSON,
    item_amount float unsigned NOT NULL,
    gst_amount float unsigned NOT NULL,
    deposit_rate int(2) unsigned NOT NULL,
    deposit_amount float unsigned NOT NULL,
    final_amount float unsigned NOT NULL,
    surcharge_amount float unsigned NOT NULL,
    total_amount float unsigned NOT NULL,
    balance_amount float unsigned NOT NULL DEFAULT 0,
    order_description mediumtext,
    order_comments mediumtext,
    post_date datetime NOT NULL,
    reserve_date datetime NOT NULL,
    finish_date datetime,
    payment_date datetime,
    order_status int(3) unsigned NOT NULL,
    urgant_flag tinyint(1) unsigned NOT NULL,
    update_version int(3) unsigned NOT NULL DEFAULT 1,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    CHECK (JSON_VALID(service_items)),
    CHECK (JSON_VALID(surcharge_items)),
    PRIMARY KEY(id)
);
ALTER TABLE `b_order` ADD FOREIGN KEY (customer_id) REFERENCES b_customer(id); 
ALTER TABLE `b_order` ADD FOREIGN KEY (address_id) REFERENCES b_address(id);
ALTER TABLE `b_order` ADD FOREIGN KEY (company_id) REFERENCES b_company(id);
ALTER TABLE `b_order` ADD FOREIGN KEY (category_id) REFERENCES b_category(id);
ALTER TABLE `b_order` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(id); 

-- Base data table indexes order-company & order-contractor --
CREATE INDEX IDX_Company_order ON `b_order` (company_id);
CREATE INDEX IDX_Contractor_order ON `b_order` (contractor_id);

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