import type { Credentials, Tweet, User } from "./types";

export class ApiError extends Error {
  status: number;

  constructor(status: number, message: string) {
    super(message);
    this.status = status;
  }
}

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const headers = new Headers(init?.headers);
  if (init?.body && !headers.has("Content-Type")) {
    headers.set("Content-Type", "application/json");
  }

  // Когда на бэкенде появится JWT — раскомментируй и клади токен в localStorage:
  // const token = localStorage.getItem("blipw_token");
  // if (token) headers.set("Authorization", `Bearer ${token}`);

  const res = await fetch(path, { ...init, headers });
  const text = await res.text();
  const body = text ? safeJson(text) : null;

  if (!res.ok) {
    const message =
      typeof body === "string" && body.trim()
        ? body.trim()
        : res.statusText || "Request failed";
    throw new ApiError(res.status, message);
  }

  return body as T;
}

function safeJson(text: string): unknown {
  try {
    return JSON.parse(text);
  } catch {
    return text;
  }
}

export function getTweets() {
  return request<Tweet[]>("/api/tweets");
}

export function createTweet(userId: number, content: string) {
  // Временно: user_id с клиента. После шага 3 roadmap — только { content }.
  return request<Tweet>("/api/tweets", {
    method: "POST",
    body: JSON.stringify({ user_id: userId, content }),
  });
}

export function register(creds: Credentials) {
  return request<User>("/api/auth/register", {
    method: "POST",
    body: JSON.stringify(creds),
  });
}

export function login(creds: Credentials) {
  return request<User>("/api/auth/login", {
    method: "POST",
    body: JSON.stringify(creds),
  });
}
