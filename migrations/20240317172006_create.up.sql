CREATE TABLE book_cost_items (
                       id BIGSERIAl NOT NULL PRIMARY KEY,
                       item_name VARCHAR (30) NOT NULL unique ,
                       code int,
                       description VARCHAR
);


CREATE TABLE book_daily_expense (
                                 id BIGSERIAl NOT NULL PRIMARY KEY,
                                 amount VARCHAR (30) NOT NULL unique ,
                                 item int,
                                 date timestamp
);
