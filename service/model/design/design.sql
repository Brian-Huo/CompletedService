CREATE TABLE `b_design` (
    design_id int unsigned NOT NUll AUTO_INCREMENT,
    company_id int unsigned NOT NULL,
    service_id int unsigned NOT NUll,
    price float unsigned NOT NUll,
    comments longtext default "",
    PRIMARY KEY(design_id)
);