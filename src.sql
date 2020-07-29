DROP TABLE Reservations;
DROP TABLE Users;

CREATE TABLE Reservations (
    Id serial  NOT NULL,
    BookDate timestamp  NOT NULL,
    SessionDate timestamp  NOT NULL,
	MinuteDuration int  NOT NULL,
    User_Id int  NOT NULL,
    Status varchar(30)  NOT NULL,
    CONSTRAINT Reservation_pk PRIMARY KEY (Id)
);

-- Table: User
CREATE TABLE Users (
    Id serial NOT NULL,
    Name varchar(50)  NOT NULL,
    Surname varchar(50)  NOT NULL,
    Apartment int  NOT NULL,
    Birthday date  NOT NULL,
    Gender varchar(6)  NOT NULL,
    Banned boolean  NOT NULL,
    CONSTRAINT User_pk PRIMARY KEY (Id)
);


CREATE TABLE TEST (
	BookDate timestamp  NOT NULL
);

select * from Reservations;
select * from Users;
select * from TEST;

insert into Users (Name, Surname, Apartment, Birthday, Gender, Banned) 
VALUES ('Marcelo', 'Carbonel', '301','1999-09-29','Male', FALSE)

insert into Reservations (BookDate, SessionDate, MinuteDuration, User_Id, Status) 
VALUES (CURRENT_TIMESTAMP, 

insert into TEST(BookDate) values (current_timestamp)