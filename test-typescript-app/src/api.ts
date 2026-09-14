import { ApiResponse } from "./types";

const BASE_URL = "https://localhost:5000/";

async function fetchJson<T>(path: string): Promise<ApiResponse<T>> {
  const response = await fetch(`${BASE_URL}${path}`);

  if (!response.ok) {
    throw new Error(`HTTP ${response.status}: ${response.statusText}`);
  }

  const data = (await response.json()) as T;
  return { data, status: response.status, ok: response.ok };
}

export const api = {
  get: <T>(path: string) => fetchJson<T>(path),
};
