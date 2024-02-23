CREATE TABLE Customer (
    CustomerID BIGINT AUTO_INCREMENT PRIMARY KEY,
    ClientCustomerID BIGINT,
    InsertDate TIMESTAMP
);

CREATE TABLE ChannelType (
    ChannelTypeID SMALLINT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(30)
);

CREATE TABLE EventType (
    EventTypeID SMALLINT PRIMARY KEY,
    Name VARCHAR(30)
);


CREATE TABLE Content (
    ContentID INT AUTO_INCREMENT PRIMARY KEY,
    ClientContentID BIGINT,
    InsertDate TIMESTAMP
);

CREATE TABLE CustomerData (
    CustomerChannelID BIGINT AUTO_INCREMENT PRIMARY KEY,
    CustomerID BIGINT,
    ChannelTypeID SMALLINT,
    ChannelValue VARCHAR(600),
    InsertDate TIMESTAMP,
    FOREIGN KEY (CustomerID) REFERENCES Customer(CustomerID),
    FOREIGN KEY (ChannelTypeID) REFERENCES ChannelType(ChannelTypeID)
);

CREATE TABLE CustomerEvent (
    EventID BIGINT AUTO_INCREMENT PRIMARY KEY,
    ClientEventID BIGINT,
    InsertDate TIMESTAMP
);

CREATE TABLE CustomerEventData (
    EventDataID BIGINT AUTO_INCREMENT PRIMARY KEY,
    EventID BIGINT,
    ContentID INT,
    CustomerID BIGINT,
    EventTypeID SMALLINT,
    EventDate TIMESTAMP,
    Quantity SMALLINT,
    InsertDate TIMESTAMP,
    FOREIGN KEY (EventID) REFERENCES CustomerEvent(EventID),
    FOREIGN KEY (ContentID) REFERENCES Content(ContentID),
    FOREIGN KEY (CustomerID) REFERENCES Customer(CustomerID),
    FOREIGN KEY (EventTypeID) REFERENCES EventType(EventTypeID)
);


CREATE TABLE ContentPrice (
    ContentPriceID MEDIUMINT AUTO_INCREMENT PRIMARY KEY,
    ContentID INT,
    Price DECIMAL(8,2),
    Currency CHAR(3),
    InsertDate TIMESTAMP,
    FOREIGN KEY (ContentID) REFERENCES Content(ContentID)
);

INSERT INTO ChannelType (ChannelTypeID, Name) VALUES
(1, 'Email'),
(2, 'PhoneNumber'),
(3, 'Postal'),
(4, 'MobileID'),
(5, 'Cookie');


INSERT INTO EventType (EventTypeID, Name) VALUES
(1, 'sent'),
(2, 'view'),
(3, 'click'),
(4, 'visit'),
(5, 'cart'),
(6, 'purchase');