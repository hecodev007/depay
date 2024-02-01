CREATE
DATABASE  IF NOT EXISTS `TronService`;
USE
`TronService`;

--
-- Table structure for table `Address`
--

CREATE TABLE `Address`
(
    `ID`             bigint      NOT NULL AUTO_INCREMENT,
    `BASE58_ADDRESS` varchar(40) NOT NULL,
    `WALLET_GROUP`   varchar(256) DEFAULT NULL,
    `WALLET_LABEL`   varchar(256) DEFAULT NULL,
    `GAS_ADDRESS`    boolean      DEFAULT false,
    `WALLET_ADDRESS` boolean      DEFAULT false,
    PRIMARY KEY (`ID`),
    UNIQUE KEY `BASE58_ADDRESS_UNIQUE` (`BASE58_ADDRESS`),
    KEY              `WALLET_GROUP_INDEX` (`WALLET_GROUP`,`WALLET_LABEL`)
);


--
-- Table structure for table `Block_Processing`

CREATE TABLE `Block_Processing`
(
    `NODE_NAME`                   varchar(20) NOT NULL,
    `LAST_PROCESSED_BLOCK_NUMBER` bigint      NOT NULL,
    `LAST_PROCESSED_BLOCK_HASH`   varchar(128) DEFAULT NULL,
    `BASE_BLOCK_NUMBER`           bigint      NOT NULL,
    `LAST_UPDATED`                timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`NODE_NAME`)
);


--
-- Table structure for table `Block_Processing_History`
--

CREATE TABLE `Block_Processing_History`
(
    `ID`              bigint    NOT NULL AUTO_INCREMENT,
    `VERSION`         bigint    NOT NULL,
    `BLOCK_NUMBER`    bigint    NOT NULL,
    `CREATION_DATE`   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `PROCESSING_DATE` timestamp NULL DEFAULT NULL,
    `PROCESSED`       boolean            DEFAULT false,
    `NODE_NAME`       varchar(20)        DEFAULT NULL,
    PRIMARY KEY (`ID`),
    UNIQUE KEY `BLOCK_NUMBER_UNIQUE` (`BLOCK_NUMBER`),
    KEY               `BPH_NODE_NAME_FK` (`NODE_NAME`),
    CONSTRAINT `Block_Processing_History_FK_1` FOREIGN KEY (`NODE_NAME`) REFERENCES `Block_Processing` (`NODE_NAME`) ON DELETE CASCADE ON UPDATE CASCADE
);



CREATE TABLE `Trc20_Tokens`
(
    `ID`               int          NOT NULL AUTO_INCREMENT,
    `CCY`              varchar(64)  NOT NULL,
    `CREATION_DATE`     timestamp    NOT NULL,
    `UPDATE_DATE`     timestamp    NOT NULL,
    `TOKEN_ADDRESS` varchar(255) NOT NULL,
    `DECIMAL_PLACES`   int(11) NOT NULL DEFAULT 6,
    PRIMARY KEY (`ID`),
    UNIQUE KEY `UK_Tron_Trc20_Tokens_Ccy` (`CCY`)
);
--
-- Table structure for table `Block_Rescanning`

CREATE TABLE `Block_Rescanning`
(
    `NODE_NAME`                   varchar(20) NOT NULL,
    `LAST_PROCESSED_BLOCK_NUMBER` bigint      NOT NULL,
    `LAST_PROCESSED_BLOCK_HASH`   varchar(128) DEFAULT NULL,
    `BASE_BLOCK_NUMBER`           bigint      NOT NULL,
    `LAST_UPDATED`                timestamp NULL DEFAULT NULL,
    `STATUS`                      int DEFAULT NULL,
    `NONDE_ID`                    int DEFAULT NULL,
    PRIMARY KEY (`NODE_NAME`)
);