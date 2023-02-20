-- Relation data table: subscribe record --
CREATE TABLE `r_broadcast` (
    group_id int unsigned NOT NULL,
    order_id int unsigned NOT NULL,
    PRIMARY KEY(group_id, order_id)
);
ALTER TABLE `r_broadcast` ADD FOREIGN KEY (group_id) REFERENCES b_subscribe_group(group_id);
ALTER TABLE `r_broadcast` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id);