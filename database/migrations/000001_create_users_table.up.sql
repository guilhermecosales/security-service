CREATE TABLE users
(
    user_id             UUID                                               NOT NULL,
    first_name          VARCHAR(50)                                        NOT NULL,
    last_name           VARCHAR(50)                                        NOT NULL,
    email               VARCHAR(255) UNIQUE                                NOT NULL,
    password            VARCHAR                                            NOT NULL,
    locked              BOOLEAN                                            NOT NULL,
    credentials_expired BOOLEAN                                            NOT NULL,
    enabled             BOOLEAN                                            NOT NULL,
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT pk_users_user_id PRIMARY KEY (user_id)
);
