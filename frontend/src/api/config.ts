export const baseUrl = import.meta.env.VITE_API_BASE_URL ?? '';

export const jsonHeaders = new Headers();
jsonHeaders.append("Content-Type", "application/json");