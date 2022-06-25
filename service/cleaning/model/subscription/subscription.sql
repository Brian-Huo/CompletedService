-- Relation data table: subscribe record --
CREATE TABLE `b_subscription` (
    group_id int unsigned NOT NULL,
    contractor_id int unsigned NOT NULL,
    PRIMARY KEY(group_id, contractor_id)
);
ALTER TABLE `b_subscription` ADD FOREIGN KEY (group_id) REFERENCES b_subscribe_group(group_id);
ALTER TABLE `b_subscription` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id);