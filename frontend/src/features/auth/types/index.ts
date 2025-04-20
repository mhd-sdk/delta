export interface User {
  username: string;
  email: string;
}

export interface AuthState {
  user: User | null;
  token: string | null;
  expiresAt: Date | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
}

export interface AuthResponse {
  token: string;
  expires_at: string;
  username: string;
  email: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
}

export interface LoginRequest {
  username: string;
}

export interface ErrorResponse {
  error: string;
}
