CREATE TABLE book_cost_items (
                                 id BIGSERIAl NOT NULL PRIMARY KEY,
                                 item_name VARCHAR (30) NOT NULL unique ,
                                 code INT unique,
                                 description VARCHAR
);


CREATE TABLE book_daily_expense (
                                    id BIGSERIAl NOT NULL PRIMARY KEY,
                                    amount FLOAT NOT NULL,
                                    date timestamp,
                                    item VARCHAR REFERENCES book_cost_items (item_name) ON DELETE CASCADE
);

ALTER TABLE book_daily_expense ADD item_id INTEGER;


UPDATE book_daily_expense SET item_id = book_cost_items.id FROM book_cost_items WHERE book_daily_expense.item=book_cost_items.item_name;

