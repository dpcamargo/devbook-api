INSERT INTO usuarios (nome, nick, email, senha) VALUES
("usuário 1", "nick 1", "email@1", "$2a$10$7ztB4Ij4O7gO53ApGB.7YueZvnenLVR6atPu2rOuaeG8GzVxj/Lki"),
("usuário 2", "nick 2", "email@2", "$2a$10$7ztB4Ij4O7gO53ApGB.7YueZvnenLVR6atPu2rOuaeG8GzVxj/Lki"),
("usuário 3", "nick 3", "email@3", "$2a$10$7ztB4Ij4O7gO53ApGB.7YueZvnenLVR6atPu2rOuaeG8GzVxj/Lki"),
("usuário 4", "nick 4", "email@4", "$2a$10$7ztB4Ij4O7gO53ApGB.7YueZvnenLVR6atPu2rOuaeG8GzVxj/Lki"),
("usuário 5", "nick 5", "email@5", "$2a$10$7ztB4Ij4O7gO53ApGB.7YueZvnenLVR6atPu2rOuaeG8GzVxj/Lki");

INSERT INTO seguidores (usuario_id, seguidor_id) VALUES
(1, 2),
(1, 3),
(2, 4),
(2, 3),
(3, 5),
(4, 3),
(5, 2),
(5, 1);

INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES
("Publicacao usuario 1", "Essa é a publicacao do usuario 1", 1),
("Publicacao usuario 2", "Essa é a publicacao do usuario 2", 2),
("Publicacao usuario 3", "Essa é a publicacao do usuario 3", 3);