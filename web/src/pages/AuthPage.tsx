import { useState, type FormEvent } from "react";
import { Link, useNavigate } from "react-router-dom";
import { ApiError, login, register } from "../api/client";
import { useAuth } from "../auth";

type Mode = "login" | "register";

export function AuthPage({ mode }: { mode: Mode }) {
  const navigate = useNavigate();
  const { setUser } = useAuth();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [busy, setBusy] = useState(false);

  const isLogin = mode === "login";

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    setBusy(true);
    setError(null);
    try {
      const user = isLogin
        ? await login({ username, password })
        : await register({ username, password });
      setUser(user);
      navigate("/");
    } catch (err) {
      setError(err instanceof ApiError ? err.message : "Что-то пошло не так");
    } finally {
      setBusy(false);
    }
  }

  return (
    <div className="auth-page">
      <form className="auth-card" onSubmit={onSubmit}>
        <Link to="/" className="brand">
          Blip<span>w</span>
        </Link>
        <h1>{isLogin ? "С возвращением" : "Создай аккаунт"}</h1>
        <p>
          {isLogin
            ? "Войди тем же username/password, что через API."
            : "Регистрация идёт в твой Go backend."}
        </p>

        <div className="field">
          <label htmlFor="username">Username</label>
          <input
            id="username"
            autoComplete="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>

        <div className="field">
          <label htmlFor="password">Password</label>
          <input
            id="password"
            type="password"
            autoComplete={isLogin ? "current-password" : "new-password"}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>

        {error ? <p className="error">{error}</p> : null}

        <button className="btn btn-primary" type="submit" disabled={busy} style={{ width: "100%" }}>
          {busy ? "…" : isLogin ? "Войти" : "Зарегистрироваться"}
        </button>

        <p className="auth-footer">
          {isLogin ? (
            <>
              Нет аккаунта? <Link to="/register">Регистрация</Link>
            </>
          ) : (
            <>
              Уже есть? <Link to="/login">Войти</Link>
            </>
          )}
        </p>
      </form>
    </div>
  );
}
