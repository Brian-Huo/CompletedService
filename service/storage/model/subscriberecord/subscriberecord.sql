-- Relation data table: subscribe record --
CREATE TABLE `r_subscribe_record` (
    category_id int unsigned NOT NULL,
    contractor_id int unsigned NOT NULL,
    PRIMARY KEY(category_id, contractor_id)
);
ALTER TABLE `r_subscribe_record` ADD FOREIGN KEY (category_id) REFERENCES b_category(category_id);
ALTER TABLE `r_subscribe_record` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id);