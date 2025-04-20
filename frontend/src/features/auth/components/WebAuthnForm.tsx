import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { AuthResponse } from '@/features/auth/types';
import { zodResolver } from '@hookform/resolvers/zod';
import { useNavigate } from '@tanstack/react-router';
import axios from 'axios';
import { FC, useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

type WebAuthnFormProps = {
  mode: 'register' | 'login';
};

const registerSchema = z.object({
  username: z.string().min(3, {
    message: 'Username must be at least 3 characters long',
  }),
  email: z.string().email({
    message: 'Please enter a valid email address',
  }),
});

const loginSchema = z.object({
  username: z.string().min(1, {
    message: 'Username is required',
  }),
});

type RegisterFormValues = z.infer<typeof registerSchema>;
type LoginFormValues = z.infer<typeof loginSchema>;
type FormValues = RegisterFormValues | LoginFormValues;

// WebAuthn credential type
interface WebAuthnCredential {
  id: ArrayBuffer;
}

const WebAuthnForm: FC<WebAuthnFormProps> = ({ mode }) => {
  const navigate = useNavigate();
  const [isWebAuthnInProgress, setIsWebAuthnInProgress] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const schema = mode === 'register' ? registerSchema : loginSchema;

  const form = useForm<FormValues>({
    resolver: zodResolver(schema),
    defaultValues: {
      username: '',
      ...(mode === 'register' ? { email: '' } : {}),
    },
  });

  const clearErrors = useCallback(() => {
    setError(null);
  }, []);

  const register = useCallback(async (username: string, email: string) => {
    try {
      setIsLoading(true);
      setError(null);

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
          excludeCredentials: options.excludeCredentials?.map((credential: WebAuthnCredential) => ({
            ...credential,
            id: base64URLToBuffer(credential.id),
          })),
        },
      })) as PublicKeyCredential;

      // Step 3: Finish registration
      const attestationResponse = credential.response as AuthenticatorAttestationResponse;

      const finishResponse = await axios.post(`${API_URL}/auth/register/finish`, {
        id: credential.id,
        rawId: arrayBufferToBase64(credential.rawId),
        type: credential.type,
        response: {
          clientDataJSON: arrayBufferToBase64(attestationResponse.clientDataJSON),
          attestationObject: arrayBufferToBase64(attestationResponse.attestationObject),
        },
      });

      const authData: AuthResponse = finishResponse.data;

      setIsLoading(false);
      // Store auth info in localStorage for app-wide access
      localStorage.setItem(
        'auth',
        JSON.stringify({
          username: authData.username,
          email: authData.email,
          token: authData.token,
          expiresAt: authData.expires_at,
          isAuthenticated: true,
        })
      );
    } catch (error) {
      setIsLoading(false);
      setError(error instanceof Error ? error.message : 'Registration failed');
      throw error;
    }
  }, []);

  const login = useCallback(async (username: string) => {
    try {
      setIsLoading(true);
      setError(null);

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
          allowCredentials: options.allowCredentials?.map((credential: WebAuthnCredential) => ({
            ...credential,
            id: base64URLToBuffer(credential.id),
          })),
        },
      })) as PublicKeyCredential;

      // Step 3: Finish login
      const assertionResponse = credential.response as AuthenticatorAssertionResponse;

      const finishResponse = await axios.post(`${API_URL}/auth/login/finish`, {
        id: credential.id,
        rawId: arrayBufferToBase64(credential.rawId),
        type: credential.type,
        response: {
          clientDataJSON: arrayBufferToBase64(assertionResponse.clientDataJSON),
          authenticatorData: arrayBufferToBase64(assertionResponse.authenticatorData),
          signature: arrayBufferToBase64(assertionResponse.signature),
          userHandle: assertionResponse.userHandle ? arrayBufferToBase64(assertionResponse.userHandle) : null,
        },
      });

      const authData: AuthResponse = finishResponse.data;

      setIsLoading(false);
      // Store auth info in localStorage for app-wide access
      localStorage.setItem(
        'auth',
        JSON.stringify({
          username: authData.username,
          email: authData.email,
          token: authData.token,
          expiresAt: authData.expires_at,
          isAuthenticated: true,
        })
      );
    } catch (error) {
      setIsLoading(false);
      setError(error instanceof Error ? error.message : 'Login failed');
      throw error;
    }
  }, []);

  const onSubmit = async (values: FormValues) => {
    try {
      clearErrors();
      setIsWebAuthnInProgress(true);

      if (mode === 'register' && 'email' in values) {
        try {
          await register(values.username, values.email);
          navigate({ to: '/' });
        } catch (error) {
          if (error instanceof Error && error.message.includes('Found no credentials for user')) {
            // This happens when a user exists but doesn't have credentials
            // We can still continue with registration to add credentials for this username
            try {
              // Try to complete registration by creating a credential
              await register(values.username, values.email);
              navigate({ to: '/' });
            } catch (retryError) {
              console.error('Retry registration failed:', retryError);
            }
          } else {
            throw error;
          }
        }
      } else {
        try {
          await login(values.username);
          navigate({ to: '/' });
        } catch (error) {
          if (error instanceof Error && error.message.includes('Found no credentials for user')) {
            setIsWebAuthnInProgress(false);
          } else {
            throw error;
          }
        }
      }
    } catch (error) {
      console.error('Authentication error:', error);
      setIsWebAuthnInProgress(false);
    }
  };

  const getErrorMessage = () => {
    if (!error) return null;

    if (error.includes('Found no credentials for user')) {
      if (mode === 'login') {
        return (
          <div className="text-red-500 text-sm">
            No credentials found for this user. Please register first or use a different account.
            <div className="mt-2">
              <Button type="button" variant="outline" size="sm" onClick={() => navigate({ to: '/register' })}>
                Go to Register
              </Button>
            </div>
          </div>
        );
      } else {
        return (
          <div className="text-amber-500 text-sm">This username exists but doesn't have credentials. Complete registration to add credentials.</div>
        );
      }
    }

    return <div className="text-red-500 text-sm">{error}</div>;
  };

  return (
    <Card className="w-[350px] sm:w-[450px]">
      <CardHeader>
        <CardTitle>{mode === 'register' ? 'Register' : 'Login'}</CardTitle>
        <CardDescription>{mode === 'register' ? 'Create a new account with WebAuthn' : 'Log in to your account with WebAuthn'}</CardDescription>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <FormField
              control={form.control}
              name="username"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input placeholder="Enter your username" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            {mode === 'register' && (
              <FormField
                control={form.control}
                name="email"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Email</FormLabel>
                    <FormControl>
                      <Input placeholder="Enter your email" type="email" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            )}

            {getErrorMessage()}

            <Button type="submit" className="w-full" disabled={isLoading || isWebAuthnInProgress}>
              {isLoading || isWebAuthnInProgress ? 'Authenticating...' : mode === 'register' ? 'Register' : 'Login'}
            </Button>
          </form>
        </Form>
      </CardContent>
      <CardFooter className="flex justify-center">
        {mode === 'register' ? (
          <div className="text-sm text-center">
            Already have an account?{' '}
            <a className="underline cursor-pointer" onClick={() => navigate({ to: '/login' })}>
              Login
            </a>
          </div>
        ) : (
          <div className="text-sm text-center">
            Don't have an account?{' '}
            <a className="underline cursor-pointer" onClick={() => navigate({ to: '/register' })}>
              Register
            </a>
          </div>
        )}
      </CardFooter>
    </Card>
  );
};

// Utility functions for WebAuthn operations
const arrayBufferToBase64 = (buffer: ArrayBuffer): string => {
  const bytes = new Uint8Array(buffer);
  let str = '';

  for (const byte of bytes) {
    str += String.fromCharCode(byte);
  }

  return btoa(str).replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/, '');
};

const base64URLToBuffer = (base64URL: string): ArrayBuffer => {
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
};

export default WebAuthnForm;
