create table "admin"(
	id serial primary key,
	name varchar(100),
	email varchar(200),
	password varchar(100),
	phone varchar(20)
);
-- delete table
drop table admin;
-- insert into "nama tabel" (nama field) values(isinya apa)
insert into "admin" (name, email, password, phone) values ('bambang', 'bambang@gmail.com', 'bam123', '082212223445');
-- select (field yang ingin di pilih) dari "nama tabel"
select * from "admin";
-- update "nama tabel" set "field yang ingin di ganti" where id nya 2
update "admin" set name = 'ali' where id = 2;
-- delete from "nama tabel" where id nya 2
delete from "admin" where id = 2;


create table "users"(
	id serial primary key,
	name varchar(100),
	email varchar(100),
	password varchar(100),
	address varchar(100),
	phone varchar(13),
	birth_date varchar(100)
);

-- delete table
drop table users;
-- insert
insert into users (name, email, password, address, phone, birth_date) values
('sumanto', 'manto@gmail.com','sum123', 'Jln. Supratman', '08123444322', '23 agustus 1999'),
('herman', 'herman@gmail.com','her123', 'Jln. Hermanman', '08123333223', '24 agustus 1999');
-- select 
select name, email,password from users;
-- update
update users set name = 'deddy' where id = 1;
-- delete
delete from users where id = 1;

create table "genre"(
	id serial primary key,
	genre_name varchar(100)
);

-- delete table
drop table genre;
-- insert
insert into genre(genre_name) values('horor'),('action'),('slice of life'), ('comedy');
-- update
update genre set genre_name = 'horor' where  id = 4;
-- delete
delete from genre where id == 3;

create table "books"(
	id serial primary key,
	genre_id int,
	title varchar(100),
	author varchar(5),
	publisher varchar(50),
	publish_year varchar(100),
	foreign key(genre_id) references genre(id)
	
);
-- delete table
drop table books;
-- insert
insert into books(genre_id, title, author, publisher, publish_year) 
values(1, 'the conjuring', 'James', 'the sarfan company', '19 july 2023'),
	  (1, 'ritual', 'Adam', 'the sarfan company', '20 Juni 2011');
-- update
update books set title = 'kkn desa penari', author= 'bamli',publisher= '10 Maret 2022' where id = 2;
-- delete
delete from books where  id = 2;

create table "book_request"(
	id serial primary key,
	user_id int,
	approved_admin_id int,
	title varchar(100),
	author varchar(100),
	publisher varchar(100),
	publish_year varchar(100),
	status_request varchar(100),
	foreign key(user_id) references users(id),
	foreign key(approved_admin_id) references admin(id)
);

-- delete table 
drop table book_request;
-- insert table
insert into book_request(user_id, approved_admin_id, title, author, publisher, publish_year, status_request)
	        values (1, 1, 'the conjuring', 'adam', 'the company book', '20 aug 2024', 'waiting list');
-- update
update book_request set title = 'kkn desa penari' where  id = 1;
-- delete
delete from book_request where id = 1;


create table "loans"(
	id serial primary key,
	user_id int,
	book_id int,
	deadline_date varchar(100),
	date_loan varchar(100),
	date_return varchar(100),
	status_loan varchar(50),
	foreign key(user_id) references users(id),
	foreign key(book_id) references books(id)
);

-- delete table
drop table loans;
-- insert data
insert into loans(user_id, book_id, deadline_date, date_loan, date_return, status_loan) 
		    values(1, 1, '20 jan 2025', '11 jan 2023', '-', 'success');
-- update
update loans set book_id = 2 where id = 1;
-- delete
delete from loans where  id = 1;