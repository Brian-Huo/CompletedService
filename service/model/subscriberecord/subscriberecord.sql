-- Relation data table: subscribe record --
CREATE TABLE `r_subscribe_record` (
    group_id int unsigned NOT NULL,
    contractor_id int unsigned NOT NULL,
    PRIMARY KEY(group_id, contractor_id)
);
ALTER TABLE `r_subscribe_record` ADD FOREIGN KEY (group_id) REFERENCES b_subscribe_group(group_id);
ALTER TABLE `r_subscribe_record` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id);