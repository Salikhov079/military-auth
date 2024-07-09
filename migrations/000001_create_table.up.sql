create table admins(
    id uuid primary key default gen_random_uuid(),
    user_name varchar(30) unique,
    password varchar(30),
    soldier_id uuid,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    delated_at bigint default 0
);


INSERT INTO admins (user_name, password)
VALUES
('user1', 'password1'),
('user2', 'password2'),
('user3', 'password3'),
('user4', 'password4'),
('user5', 'password5');



INSERT INTO admins (user_name, password)
VALUES
('dilshod', '1234'),
('string', 'string');