const BASE_URL = import.meta.env.VITE_API_URL ?? "";

export class HttpError extends Error {
  constructor(
    public readonly status: number,
    message: string
  ) {
    super(message);
    this.name = "HttpError";
  }
}

async function request<T>(
  method: string,
  endpoint: string,
  body?: unknown
): Promise<T> {
  const response = await fetch(`${BASE_URL}${endpoint}`, {
    method,
    headers: {
      "Content-Type": "application/json",
    },
    ...(body !== undefined && {
      body: JSON.stringify(body),
    }),
  });

  if (!response.ok) {
    const error = await response.json().catch(() => null);

    throw new HttpError(
      response.status,
      error?.error ?? "An unexpected error occurred"
    );
  }

  return response.json() as Promise<T>;
}

export const http = {
  get: <T>(url: string) =>
    request<T>("GET", url),

  post: <T>(url: string, body: unknown) =>
    request<T>("POST", url, body),

  put: <T>(url: string, body: unknown) =>
    request<T>("PUT", url, body),

  delete: <T>(url: string) =>
    request<T>("DELETE", url),
};