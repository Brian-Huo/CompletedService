-- Relation data table: subscription --
CREATE TABLE `r_subscription` (
    subscription_id int unsigned NOT NULL AUTO_INCREMENT,
    category_id int unsigned NOT NULL,
    contractor_id int unsigned NOT NULL,
    UNIQUE UNI_subscribe (category_id, contractor_id),
    PRIMARY KEY(subscription_id)
);
ALTER TABLE `r_subscription` ADD FOREIGN KEY (category_id) REFERENCES b_category(category_id);
ALTER TABLE `r_subscription` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id);