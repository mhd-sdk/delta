import { AuthResponse, AuthState, User } from '@/features/auth/types';
import axios from 'axios';
import { create } from 'zustand';
import { persist } from 'zustand/middleware';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

interface WebAuthnStore extends AuthState {
  register: (username: string, email: string) => Promise<void>;
  login: (username: string) => Promise<void>;
  logout: () => void;
  checkAuth: () => Promise<boolean>;
  clearErrors: () => void;
}

export const useWebAuthnStore = create<WebAuthnStore>()(
  persist(
    (set, get) => ({
      user: null,
      token: null,
      expiresAt: null,
      isAuthenticated: false,
      isLoading: false,
      error: null,

      register: async (username: string, email: string) => {
        try {
          set({ isLoading: true, error: null });

          // Step 1: Begin registration
          const beginResponse = await axios.post(`${API_URL}/auth/register/begin`, {
            username,
            email,
          });

          const options = beginResponse.data;

          // Step 2: Create credentials
          const credential = (await navigator.credentials.create({
            publicKey: {
              ...options,
              challenge: base64URLToBuffer(options.challenge),
              user: {
                ...options.user,
                id: base64URLToBuffer(options.user.id),
              },
              excludeCredentials: options.excludeCredentials?.map((credential: any) => ({
                ...credential,
                id: base64URLToBuffer(credential.id),
              })),
            },
          })) as PublicKeyCredential;

          // Step 3: Finish registration
          const attestationResponse = credential.response as AuthenticatorAttestationResponse;

          const finishResponse = await axios.post(`${API_URL}/auth/register/finish`, {
            id: credential.id,
            rawId: bufferToBase64URL(Buffer.from(credential.rawId)),
            type: credential.type,
            response: {
              clientDataJSON: bufferToBase64URL(Buffer.from(attestationResponse.clientDataJSON)),
              attestationObject: bufferToBase64URL(Buffer.from(attestationResponse.attestationObject)),
            },
          });

          const authData: AuthResponse = finishResponse.data;

          const user: User = {
            username: authData.username,
            email: authData.email,
          };

          set({
            isLoading: false,
            user,
            token: authData.token,
            expiresAt: new Date(authData.expires_at),
            isAuthenticated: true,
          });
        } catch (error) {
          set({
            isLoading: false,
            error: error instanceof Error ? error.message : 'Registration failed',
          });
          throw error;
        }
      },

      login: async (username: string) => {
        try {
          set({ isLoading: true, error: null });

          // Step 1: Begin login
          const beginResponse = await axios.post(`${API_URL}/auth/login/begin`, {
            username,
          });

          const options = beginResponse.data;

          // Step 2: Get credentials
          const credential = (await navigator.credentials.get({
            publicKey: {
              ...options,
              challenge: base64URLToBuffer(options.challenge),
              allowCredentials: options.allowCredentials?.map((credential: any) => ({
                ...credential,
                id: base64URLToBuffer(credential.id),
              })),
            },
          })) as PublicKeyCredential;

          // Step 3: Finish login
          const assertionResponse = credential.response as AuthenticatorAssertionResponse;

          const finishResponse = await axios.post(`${API_URL}/auth/login/finish`, {
            id: credential.id,
            rawId: bufferToBase64URL(Buffer.from(credential.rawId)),
            type: credential.type,
            response: {
              clientDataJSON: bufferToBase64URL(Buffer.from(assertionResponse.clientDataJSON)),
              authenticatorData: bufferToBase64URL(Buffer.from(assertionResponse.authenticatorData)),
              signature: bufferToBase64URL(Buffer.from(assertionResponse.signature)),
              userHandle: assertionResponse.userHandle ? bufferToBase64URL(Buffer.from(assertionResponse.userHandle)) : null,
            },
          });

          const authData: AuthResponse = finishResponse.data;

          const user: User = {
            username: authData.username,
            email: authData.email,
          };

          set({
            isLoading: false,
            user,
            token: authData.token,
            expiresAt: new Date(authData.expires_at),
            isAuthenticated: true,
          });
        } catch (error) {
          set({
            isLoading: false,
            error: error instanceof Error ? error.message : 'Login failed',
          });
          throw error;
        }
      },

      logout: () => {
        set({
          user: null,
          token: null,
          expiresAt: null,
          isAuthenticated: false,
        });
      },

      checkAuth: async () => {
        const { token, expiresAt } = get();

        // If no token or expired, return false
        if (!token || !expiresAt || new Date() > expiresAt) {
          set({ isAuthenticated: false, user: null, token: null, expiresAt: null });
          return false;
        }

        try {
          // Verify token with backend
          const response = await axios.get(`${API_URL}/auth/verify`, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          });

          const authData: AuthResponse = response.data;

          const user: User = {
            username: authData.username,
            email: authData.email,
          };

          set({
            user,
            token: authData.token,
            expiresAt: new Date(authData.expires_at),
            isAuthenticated: true,
          });

          return true;
        } catch (error) {
          set({ isAuthenticated: false, user: null, token: null, expiresAt: null });
          return false;
        }
      },

      clearErrors: () => {
        set({ error: null });
      },
    }),
    {
      name: 'webauthn-auth-storage',
    }
  )
);

// Utility functions for WebAuthn operations
function bufferToBase64URL(buffer: ArrayBuffer): string {
  const bytes = new Uint8Array(buffer);
  let str = '';

  for (const byte of bytes) {
    str += String.fromCharCode(byte);
  }

  return btoa(str).replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/, '');
}

function base64URLToBuffer(base64URL: string): ArrayBuffer {
  const base64 = base64URL.replace(/-/g, '+').replace(/_/g, '/');
  const paddingLength = (4 - (base64.length % 4)) % 4;
  const padded = base64 + '='.repeat(paddingLength);

  const binary = atob(padded);
  const buffer = new ArrayBuffer(binary.length);
  const bytes = new Uint8Array(buffer);

  for (let i = 0; i < binary.length; i++) {
    bytes[i] = binary.charCodeAt(i);
  }

  return buffer;
}
