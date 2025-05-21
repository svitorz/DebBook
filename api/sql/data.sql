INSERT INTO USUARIOS (
  nome, nick, email, senha
) VALUES 
("Usuario 1", "user_1", "usuario1@gmail.com", "$2a$10$R/F1FmKK4BGB8an8n/dcOew5plWtsj6I3.SBn1mlYEkwLZrQgsFL6"),
("Usuario 2", "user_2", "usuario2@gmail.com", "$2a$10$R/F1FmKK4BGB8an8n/dcOew5plWtsj6I3.SBn1mlYEkwLZrQgsFL6"),
("Usuario 3", "user_3", "usuario3@gmail.com", "$2a$10$R/F1FmKK4BGB8an8n/dcOew5plWtsj6I3.SBn1mlYEkwLZrQgsFL6"),
("Usuario 4", "user_4", "usuario4@gmail.com", "$2a$10$R/F1FmKK4BGB8an8n/dcOew5plWtsj6I3.SBn1mlYEkwLZrQgsFL6"),
("Usuario 5", "user_5", "usuario5@gmail.com", "$2a$10$R/F1FmKK4BGB8an8n/dcOew5plWtsj6I3.SBn1mlYEkwLZrQgsFL6"),
("Usuario 6", "user_6", "usuario6@gmail.com", "$2a$10$R/F1FmKK4BGB8an8n/dcOew5plWtsj6I3.SBn1mlYEkwLZrQgsFL6"),
("Usuario 7", "user_7", "usuario7@gmail.com", "$2a$10$R/F1FmKK4BGB8an8n/dcOew5plWtsj6I3.SBn1mlYEkwLZrQgsFL6"),
("Usuario 8", "user_8", "usuario8@gmail.com", "$2a$10$R/F1FmKK4BGB8an8n/dcOew5plWtsj6I3.SBn1mlYEkwLZrQgsFL6");

INSERT INTO SEGUIDORES(usuario_id, seguidor)
VALUES
(1,2),
  (2,3),
  (3,4),
(1,8);
