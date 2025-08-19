import { AuthResponse, AuthState, User } from '@/features/auth/types';
import axios from 'axios';
import { Buffer } from 'buffer';
import { create } from 'zustand';
import { persist } from 'zustand/middleware';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

interface WebAuthnStore extends AuthState {
  register: (username: string) => Promise<void>;
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

      register: async (username: string) => {
        set({ isLoading: true, error: null });
        try {
          // Step 1: Begin registration
          const beginResponse = await axios.post(
            `${API_URL}/auth/register/begin`,
            { username },
            {
              withCredentials: true,
              headers: {
                'Content-Type': 'application/json',
              },
            }
          );

          const options = beginResponse.data.publicKey;

          // Step 2: Create credentials
          const credential = (await navigator.credentials.create({
            publicKey: {
              ...options,
              challenge: base64URLToBuffer(options.challenge),
              user: {
                ...options.user,
                id: base64URLToBuffer(options.user.id),
              },
              excludeCredentials: options.excludeCredentials?.map((credential: { id: string }) => ({
                ...credential,
                id: base64URLToBuffer(credential.id),
              })),
            },
          })) as PublicKeyCredential;

          // Step 3: Finish registration
          const attestationResponse = credential.response as AuthenticatorAttestationResponse;
          const finishResponse = await axios.post(
            `${API_URL}/auth/register/finish?username=${username}`,
            {
              id: credential.id,
              rawId: bufferToBase64URL(Buffer.from(credential.rawId)),
              type: credential.type,
              response: {
                clientDataJSON: bufferToBase64URL(Buffer.from(attestationResponse.clientDataJSON)),
                attestationObject: bufferToBase64URL(Buffer.from(attestationResponse.attestationObject)),
              },
            },
            {
              withCredentials: true,
              headers: {
                'Content-Type': 'application/json',
              },
            }
          );

          const authData: AuthResponse = finishResponse.data;

          const user: User = {
            username: authData.username,
          };

          set({
            isLoading: false,
            user,
            token: authData.token,
            expiresAt: new Date(authData.expires_at),
            isAuthenticated: true,
          });
        } catch (error) {
          let errorMessage = 'Registration failed';

          // Récupérer le corps de l'erreur de l'API si disponible
          if (axios.isAxiosError(error) && error.response) {
            // Utiliser le message d'erreur du serveur s'il existe
            errorMessage = error.response.data?.message || error.response.data || errorMessage;

            // Vous pouvez aussi logger l'erreur complète pour le débogage
            console.error('API Error:', error.response.data);
          } else if (error instanceof Error) {
            errorMessage = error.message;
          }

          set({
            isLoading: false,
            error: errorMessage,
          });
          throw error;
        }
      },

      login: async (username: string) => {
        try {
          set({ isLoading: true, error: null });

          // Step 1: Begin login
          const beginResponse = await axios.post(
            `${API_URL}/auth/login/begin`,
            { username },
            {
              withCredentials: true,
              headers: {
                'Content-Type': 'application/json',
              },
            }
          );

          const options = beginResponse.data;

          // Step 2: Get credentials
          const credential = (await navigator.credentials.get({
            publicKey: {
              ...options,
              challenge: base64URLToBuffer(options.challenge),
              allowCredentials: options.allowCredentials?.map((credential: { id: string }) => ({
                ...credential,
                id: base64URLToBuffer(credential.id),
              })),
            },
          })) as PublicKeyCredential;
          console.log(credential);
          // Step 3: Finish login
          const assertionResponse = credential.response as AuthenticatorAssertionResponse;

          const finishResponse = await axios.post(
            `${API_URL}/auth/login/finish?username=${username}`,
            {
              id: credential.id,
              rawId: bufferToBase64URL(Buffer.from(credential.rawId)),
              type: credential.type,
              response: {
                clientDataJSON: bufferToBase64URL(Buffer.from(assertionResponse.clientDataJSON)),
                authenticatorData: bufferToBase64URL(Buffer.from(assertionResponse.authenticatorData)),
                signature: bufferToBase64URL(Buffer.from(assertionResponse.signature)),
                userHandle: assertionResponse.userHandle ? bufferToBase64URL(Buffer.from(assertionResponse.userHandle)) : null,
              },
            },
            {
              withCredentials: true,
              headers: {
                'Content-Type': 'application/json',
              },
            }
          );

          const authData: AuthResponse = finishResponse.data;

          const user: User = {
            username: authData.username,
          };

          set({
            isLoading: false,
            user,
            token: authData.token,
            expiresAt: new Date(authData.expires_at),
            isAuthenticated: true,
          });
        } catch (error) {
          console.error(error);
          set({
            isLoading: false,
          });
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
          set({
            isAuthenticated: false,
            user: null,
            token: null,
            expiresAt: null,
          });
          return false;
        }

        try {
          // Verify token with backend
          const response = await axios.get(`${API_URL}/auth/verify`, {
            withCredentials: true,
            headers: {
              Authorization: `Bearer ${token}`,
              'Content-Type': 'application/json',
            },
          });

          const authData: AuthResponse = response.data;

          const user: User = {
            username: authData.username,
          };

          set({
            user,
            token: authData.token,
            expiresAt: new Date(authData.expires_at),
            isAuthenticated: true,
          });

          return true;
        } catch {
          set({
            isAuthenticated: false,
            user: null,
            token: null,
            expiresAt: null,
          });
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
  if (!base64URL) {
    console.error('base64URL is undefined or null');
    return new ArrayBuffer(0); // Return empty buffer instead of throwing error
  }

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
