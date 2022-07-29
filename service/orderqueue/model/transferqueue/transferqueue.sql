-- Relation data table: transfer order queue (order message) --
CREATE TABLE `r_transfer_queue` (
    order_id int unsigned NOT NUll,
    contact varchar(15) NOT NULL,
    PRIMARY KEY(order_id)
);
ALTER TABLE `r_transfer_queue` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 