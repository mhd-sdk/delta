// Déclarations de types pour WebAuthn

interface PublicKeyCredentialCreationOptions {
  challenge: ArrayBuffer;
  rp: {
    name: string;
    id: string;
  };
  user: {
    id: ArrayBuffer;
    name: string;
    displayName: string;
  };
  pubKeyCredParams: {
    type: string;
    alg: number;
  }[];
  timeout?: number;
  excludeCredentials?: {
    id: ArrayBuffer;
    type: string;
    transports?: string[];
  }[];
  authenticatorSelection?: {
    authenticatorAttachment?: string;
    residentKey?: string;
    requireResidentKey?: boolean;
    userVerification?: string;
  };
  attestation?: string;
  extensions?: any;
}

interface PublicKeyCredentialRequestOptions {
  challenge: ArrayBuffer;
  timeout?: number;
  rpId?: string;
  allowCredentials?: {
    id: ArrayBuffer;
    type: string;
    transports?: string[];
  }[];
  userVerification?: string;
  extensions?: any;
}

interface AuthenticatorResponse {
  clientDataJSON: ArrayBuffer;
}

interface AuthenticatorAttestationResponse extends AuthenticatorResponse {
  attestationObject: ArrayBuffer;
}

interface AuthenticatorAssertionResponse extends AuthenticatorResponse {
  authenticatorData: ArrayBuffer;
  signature: ArrayBuffer;
  userHandle: ArrayBuffer | null;
}

interface PublicKeyCredential extends Credential {
  rawId: ArrayBuffer;
  response: AuthenticatorAttestationResponse | AuthenticatorAssertionResponse;
  getClientExtensionResults(): any;
}

interface CredentialCreationOptions {
  publicKey?: PublicKeyCredentialCreationOptions;
  signal?: AbortSignal;
}

interface CredentialRequestOptions {
  publicKey?: PublicKeyCredentialRequestOptions;
  signal?: AbortSignal;
}

interface CredentialsContainer {
  create(options?: CredentialCreationOptions): Promise<Credential | null>;
  get(options?: CredentialRequestOptions): Promise<Credential | null>;
  preventSilentAccess(): Promise<void>;
  store(credential: Credential): Promise<Credential>;
}
