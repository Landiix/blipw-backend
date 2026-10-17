/** Контракт с blipw-backend (json-теги Go). Меняешь API — обнови здесь. */

export type User = {
  id: number;
  username: string;
  created_at: string;
};

export type Tweet = {
  id: number;
  user_id: number;
  content: string;
  created_at: string;
};

export type Credentials = {
  username: string;
  password: string;
};
