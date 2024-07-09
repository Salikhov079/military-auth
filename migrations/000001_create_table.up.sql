create table admins(
    id uuid primary key default gen_random_uuid(),
    user_name varchar(30) unique,
    password varchar(30),
    soldier_id uuid,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delated_at bigint default 0
);


INSERT INTO admins (user_name, password, soldier_id)
VALUES
('user1', 'password1', '4643a4a1-b71a-4fa4-bde8-27e80707a1ed'),
('user2', 'password2', '4643a4a1-b71a-4fa4-bde8-27e80707a1ed'),
('user3', 'password3', '4643a4a1-b71a-4fa4-bde8-27e80707a1ed'),
('user4', 'password4', '4643a4a1-b71a-4fa4-bde8-27e80707a1ed'),
('user5', 'password5', '4643a4a1-b71a-4fa4-bde8-27e80707a1ed');



INSERT INTO admins (user_name, password, soldier_id)
VALUES
('dilshod', '1234', '4643a4a1-b71a-4fa4-bde8-27e80707a1ed'),