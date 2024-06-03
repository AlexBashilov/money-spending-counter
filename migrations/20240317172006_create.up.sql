CREATE TABLE book_cost_items (
                       id BIGSERIAl NOT NULL PRIMARY KEY,
                       item_name VARCHAR (30) NOT NULL unique ,
                       code INT unique,
                       description VARCHAR
);


CREATE TABLE book_daily_expense (
                                 id BIGSERIAl NOT NULL PRIMARY KEY,
                                 amount VARCHAR (30) NOT NULL,
                                 item int,
                                 date timestamp
);
