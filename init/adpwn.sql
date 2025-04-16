CREATE TABLE adpwn_modules (
                               module_id INT GENERATED ALWAYS AS IDENTITY,
                               key VARCHAR(100) NOT NULL,
                               name VARCHAR(100) NOT NULL,
                               version VARCHAR(10) NOT NULL,
                               author VARCHAR(255) NOT NULL,
                               description VARCHAR(255) NOT NULL,
                               attack_id VARCHAR(255) NOT NULL,
                               loot_path VARCHAR(255) NOT NULL,
                               module_type VARCHAR(100) NOT NULL,
                               execution_metric VARCHAR(100) NOT NULL,
                               PRIMARY KEY(module_id),
                               UNIQUE(key)
);

CREATE TABLE adpwn_modules_dependencies (
                                     previous_module VARCHAR(100) NOT NULL,
                                     next_module VARCHAR(100) NOT NULL,
                                     PRIMARY KEY (previous_module, next_module),
                                     FOREIGN KEY (previous_module) REFERENCES adpwn_modules(key),
                                     FOREIGN KEY (next_module) REFERENCES adpwn_modules(key)
);


CREATE TABLE adpwn_modules_metadata (
    project_uid VARCHAR(255),
    module_id INT,
    CONSTRAINT fk_module
     FOREIGN KEY (module_id)
        REFERENCES adpwn_modules(module_id)
);


CREATE TABLE adpwn_users (
    user_id INT GENERATED ALWAYS AS IDENTITY,
    hash varchar,
    PRIMARY KEY (user_id)
);

CREATE TABLE adpwn_collections
(
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR,
    description VARCHAR,
    PRIMARY KEY (id)
);

CREATE TABLE adpwn_collection_modules
(
    module_key    VARCHAR,
    collection_id INT,
    PRIMARY KEY (module_key, collection_id)
);

CREATE TABLE adpwn_modules_options
(
    module_key VARCHAR,
    option_key VARCHAR,
    label VARCHAR,
    placeholder VARCHAR,
    type VARCHAR,
    required bool,
    PRIMARY KEY (module_key, option_key)
);

CREATE TABLE adpwn_module_runs
(
    module_key VARCHAR,
    run_uid VARCHAR,
    ran_at TIMESTAMP,
    project_uid VARCHAR
);

CREATE TABLE adpwn_module_last_runs
(
    module_key VARCHAR,
    run_uid VARCHAR,
    ran_at TIMESTAMP,
    project_uid VARCHAR,
    parameter jsonb
);

CREATE TABLE adpwn_module_logs (
    project_uid VARCHAR(255),
    module_key INT,
    created_at TIMESTAMP,
    message VARCHAR
);
