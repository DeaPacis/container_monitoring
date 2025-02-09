CREATE TABLE Containers(
    container_id     VARCHAR(12) PRIMARY KEY,
    ip_address       VARCHAR(15) NOT NULL,
    response_time    INTEGER NOT NULL,
    last_checked     VARCHAR(20) NOT NULL
);