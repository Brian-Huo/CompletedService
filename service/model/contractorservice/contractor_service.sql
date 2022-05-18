-- Relationship data table: contractor-service --
CREATE TABLE `r_contractor_service` (
    contractor_id int unsigned NOT NULL,
    service_id int unsigned NOT NUll,
    PRIMARY KEY(contractor_id, service_id)
);
ALTER TABLE `r_contractor_service` ADD FOREIGN KEY (contractor_id) REFERENCES b_contractor(contractor_id); 
ALTER TABLE `r_contractor_service` ADD FOREIGN KEY (service_id) REFERENCES b_service(service_id); 