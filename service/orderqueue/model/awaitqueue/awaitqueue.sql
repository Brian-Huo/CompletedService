-- Relation data table: awaiting order queue (order message) --
CREATE TABLE `r_await_queue` (
    order_id int unsigned NOT NUll,
    vacancy  int NOT NULL,
    PRIMARY KEY(order_id)
);
ALTER TABLE `r_await_queue` ADD FOREIGN KEY (order_id) REFERENCES b_order(order_id); 