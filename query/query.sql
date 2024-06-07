
create table "admin"(
	id serial primary key,
	name varchar(100),
	email varchar(200),
	password varchar(100),
	phone varchar(20),
	created_at timestamp default now()
);

create table "users"(
	id serial primary key,
	name varchar(100),
	email varchar(100),
	password varchar(100),
	address varchar(100),
	phone varchar(13),
	birth_date varchar(100),
	created_at timestamp default now()
);

create table "genre"(
	id serial primary key,
	genre_name varchar(100),
	created_at timestamp default now()
);

create table "books"(
	id serial primary key,
	genre_id int,
	title varchar(100) not null,
	author varchar(5),
	publisher varchar(50),
	publish_year int,
	created_at timestamp default now(),
	foreign key(genre_id) references genre(id)
	
);

create table "book_request"(
	id serial primary key,
	user_id int,
	approved_admin_id int,
	title varchar(100),
	author varchar(100),
	publisher varchar(100),
	publish_year varchar(100),
	status_request varchar(100),
	created_at timestamp default now(),
	foreign key(user_id) references users(id),
	foreign key(approved_admin_id) references admin(id)
);

create table "loans"(
	id serial primary key,
	user_id int,
	book_id int,
	deadline date not null,
	"return" date,
	status_loan bool default true,
	created_at timestamp default now(),
	foreign key(user_id) references users(id),
	foreign key(book_id) references books(id)
);

-- menghapus tabel database
drop table admin;
drop table users;
drop table genre;
drop table books;
drop table book_request;
drop table loans;


-- INSERT ADMIN
insert into admin("name", email, "password", phone) values
('juan', 'juan@gmial.com', '12345', '08122034433');

-- A. menambah 5 data genre
insert into genre(genre_name) values
('horor'),
('romance'),
('comedy'),
('thriller'),
('action');

-- B. menambah 5 data buku untuk masing2 genre
insert into books(genre_id, title) values
(5, 'jurassic park');

-- C. menambah 5 users 
insert into users("name", email) values  
('hiro', 'hiro@gmail.com');


-- D. tampilkan data buku berdasarkan judul buku;
select title from books; 

-- E.Tampilkan hasil pencarian data buku berdasarkan genre id
select g.genre_name, b.title 
from books b
left join genre g on g.id = b.genre_id;

-- F. Menambahkan data peminjaman buku dengan rincian : ....
insert into loans(user_id, book_id, deadline) values 
(1, 1, '2020-12-12'),
(1, 2, '2020-12-13'),
(1, 6, '2021-01-01'),
(2, 3, '2021-02-02'),
(3, 1, '2021-03-03');

-- G. Mengubah status peminjaman buku yang dipinjam oleh user 1 dan buku dengan id 1 tadi menjadi “dikembalikan”.
update loans set status_loan = false where user_id = 1 and book_id = 1;

-- H. Tampilkan data user beserta buku yang masih dipinjam/belum dikembalikan.
select u."name" , b.title "books title"
from loans l 
left join users u on u.id = l.user_id 
left join books b on b.id = l.book_id
where l.status_loan = true;

-- I. Tampilkan data buku yang belum pernah dipinjam oleh user.
select b.title 
from books b
left join loans l on l.book_id = b.id
where l.book_id is null;










