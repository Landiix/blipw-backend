import { useEffect, useState, type FormEvent } from "react";
import { Link } from "react-router-dom";
import { ApiError, createTweet, getTweets } from "../api/client";
import type { Tweet } from "../api/types";
import { useAuth } from "../auth";

function formatWhen(iso: string) {
  try {
    return new Intl.DateTimeFormat("ru-RU", {
      dateStyle: "medium",
      timeStyle: "short",
    }).format(new Date(iso));
  } catch {
    return iso;
  }
}

export function HomePage() {
  const { user, logout } = useAuth();
  const [tweets, setTweets] = useState<Tweet[]>([]);
  const [content, setContent] = useState("");
  const [loading, setLoading] = useState(true);
  const [sending, setSending] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [formError, setFormError] = useState<string | null>(null);

  async function loadFeed() {
    setError(null);
    try {
      const data = await getTweets();
      setTweets(Array.isArray(data) ? data : []);
    } catch (err) {
      const message =
        err instanceof ApiError
          ? err.message
          : "Не удалось загрузить ленту. API на :8081 запущен?";
      setError(message);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    void loadFeed();
  }, []);

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    if (!user) return;
    const text = content.trim();
    if (!text) {
      setFormError("Напиши хотя бы что-то");
      return;
    }
    if ([...text].length > 280) {
      setFormError("Максимум 280 символов");
      return;
    }

    setSending(true);
    setFormError(null);
    try {
      const created = await createTweet(user.id, text);
      setTweets((prev) => [created, ...prev]);
      setContent("");
    } catch (err) {
      setFormError(err instanceof ApiError ? err.message : "Не отправилось");
    } finally {
      setSending(false);
    }
  }

  const remaining = 280 - [...content].length;

  return (
    <div className="app-shell">
      <header className="topbar">
        <Link to="/" className="brand">
          Blip<span>w</span>
        </Link>
        <div className="nav-actions">
          {user ? (
            <>
              <span className="muted">@{user.username}</span>
              <button type="button" className="btn btn-ghost" onClick={logout}>
                Выйти
              </button>
            </>
          ) : (
            <>
              <Link className="btn btn-ghost" to="/login">
                Войти
              </Link>
              <Link className="btn btn-primary" to="/register">
                Регистрация
              </Link>
            </>
          )}
        </div>
      </header>

      <section className="hero">
        <h1>Короткие сигналы. Живая лента.</h1>
        <p>Короткие посты. Одна лента. Без лишнего шума.</p>
      </section>

      {user ? (
        <form className="panel compose" onSubmit={onSubmit}>
          <textarea
            value={content}
            onChange={(e) => setContent(e.target.value)}
            placeholder="Что нового?"
            maxLength={400}
          />
          <div className="compose-meta">
            <span className="muted">{remaining} символов</span>
            <button className="btn btn-primary" type="submit" disabled={sending}>
              {sending ? "Отправка…" : "Опубликовать"}
            </button>
          </div>
          {formError ? <p className="error">{formError}</p> : null}
        </form>
      ) : (
        <div className="panel">
          <p className="muted" style={{ margin: 0 }}>
            Войди, чтобы писать твиты. Ленту можно смотреть и без входа.
          </p>
        </div>
      )}

      <section className="panel">
        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            gap: "1rem",
          }}
        >
          <strong>Лента</strong>
          <button type="button" className="btn btn-ghost" onClick={() => void loadFeed()}>
            Обновить
          </button>
        </div>

        {loading ? <p className="muted">Загрузка…</p> : null}
        {error ? <p className="error">{error}</p> : null}

        {!loading && !error && tweets.length === 0 ? (
          <p className="empty">Пока пусто — напиши первый blip.</p>
        ) : null}

        <div className="feed">
          {tweets.map((t, i) => (
            <article
              key={t.id}
              className="tweet"
              style={{ animationDelay: `${Math.min(i, 8) * 0.04}s` }}
            >
              <div className="tweet-meta">
                <span className="tweet-author">user #{t.user_id}</span>
                <span>·</span>
                <time dateTime={t.created_at}>{formatWhen(t.created_at)}</time>
              </div>
              <p className="tweet-body">{t.content}</p>
            </article>
          ))}
        </div>
      </section>
    </div>
  );
}
