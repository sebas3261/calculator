const BASE_URL = import.meta.env.VITE_API_URL ?? "";

export class HttpError extends Error {
  public readonly status: number;
  constructor(
    status: number,
    message: string
  ) {
    super(message);
    this.name = "HttpError";
    this.status = status;
  }
}

export class NetworkError extends Error {
  constructor() {
    super("Unable to connect to the server. Please try again.");
    this.name = "NetworkError";
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
  }).catch(() => { throw new NetworkError(); });

  if (!response.ok) {
    const error = await response.json().catch(() => null);

    throw new HttpError(
      response.status,
      error?.error ?? "An unexpected error occurred"
    );
  }

  try {
    return await response.json() as T;
  } catch {
    throw new Error("The server returned an invalid response.");
  }
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