CREATE TABLE Projects (
    ID int NOT NULL,
    Name varchar(255) NOT NULL,
    CONSTRAINT PK_Projects PRIMARY KEY (ID)
)

INSERT into Projects (ID, Name) VALUES (1, "TestProject");