CREATE TABLE tasks (
  id UUID PRIMARY KEY,
  title VARCHAR(240) NOT NULL,
  description VARCHAR(240) NOT NULL,
  creator UUID NOT NULL,
  asigned UUID NOT NULL,
  status_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_tasks_statuses FOREIGN KEY (status_id) REFERENCES statuses (id),
  CONSTRAINT fk_tasks_users FOREIGN KEY (creator) REFERENCES users (id),
  CONSTRAINT fk_tasks_users FOREIGN KEY (asigned) REFERENCES users (id)
);

CREATE TABLE boards (
  id UUID PRIMARY KEY,
  title VARCHAR(240) NOT NULL,
  description VARCHAR(240) NOT NULL,
  owner UUID NOT NULL,
  team_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_boards_teams FOREIGN KEY (team_id) REFERENCES teams (id),
  CONSTRAINT fk_boards_users FOREIGN KEY (owner) REFERENCES users (id)
);

CREATE TABLE teams (
  id UUID PRIMARY KEY,
  title VARCHAR(240) NOT NULL,
  description VARCHAR(240) NOT NULL,
  owner UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_teams_users FOREIGN KEY (owner) REFERENCES users (id)
);

CREATE TABLE teams_users (
  team_id UUID NOT NULL,
  user_id UUID NOT NULL,
  PRIMARY KEY (team_id, user_id),
  CONSTRAINT fk_teams_users_teams FOREIGN KEY (team_id) REFERENCES teams (id),
  CONSTRAINT fk_teams_users_users FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(240) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE statuses (
  id UUID PRIMARY KEY,
  title VARCHAR(240) NOT NULL,
  board_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE statuses_next_statuses (
  status_id UUID NOT NULL,
  next_status UUID NOT NULL,
  PRIMARY KEY (status_id, next_status),
  CONSTRAINT fk_statuses_next_statuses_statuses FOREIGN KEY (status_id) REFERENCES statuses (id),
  CONSTRAINT fk_statuses_next_statuses_statuses FOREIGN KEY (next_status) REFERENCES statuses (id)
);

CREATE TABLE statuses_previous_statuses (
  status_id UUID NOT NULL,
  previous_status UUID NOT NULL,
  PRIMARY KEY (status_id, previous_status),
  CONSTRAINT fk_statuses_previous_statuses_statuses FOREIGN KEY (status_id) REFERENCES statuses (id),
  CONSTRAINT fk_statuses_previous_statuses_statuses FOREIGN KEY (previous_status) REFERENCES statuses (id)
);