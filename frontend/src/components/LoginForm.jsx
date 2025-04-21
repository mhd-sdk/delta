import React, { useState } from 'react';
import { bufferToBase64URL, base64URLToBuffer } from '../utils/encoding';

const LoginForm = () => {
  const [username, setUsername] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');

  const handleSignIn = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setErrorMessage('');

    try {
      // Étape 1: Commencer l'authentification
      const beginResponse = await fetch('/api/auth/login/begin', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username }),
        credentials: 'include'
      });

      if (!beginResponse.ok) {
        const error = await beginResponse.json();
        throw new Error(error.error || 'Erreur lors de l'authentification');
      }

      const options = await beginResponse.json();

      // Préparation des options WebAuthn pour utiliser Touch ID
      const publicKeyOptions = {
        ...options,
        challenge: base64URLToBuffer(options.challenge),
        allowCredentials: options.allowCredentials?.map(credential => ({
          ...credential,
          id: base64URLToBuffer(credential.id)
        })),
        userVerification: 'required' // Force la vérification biométrique
      };

      // Étape 2: Utilisation du lecteur d'empreinte pour authentification
      const credential = await navigator.credentials.get({
        publicKey: publicKeyOptions
      });

      // Étape 3: Finaliser l'authentification
      const authResponse = credential.response;
      const finishResponse = await fetch('/api/auth/login/finish', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          id: credential.id,
          rawId: bufferToBase64URL(credential.rawId),
          type: credential.type,
          response: {
            clientDataJSON: bufferToBase64URL(authResponse.clientDataJSON),
            authenticatorData: bufferToBase64URL(authResponse.authenticatorData),
            signature: bufferToBase64URL(authResponse.signature),
            userHandle: authResponse.userHandle ? bufferToBase64URL(authResponse.userHandle) : null
          }
        }),
        credentials: 'include'
      });

      if (!finishResponse.ok) {
        const error = await finishResponse.json();
        throw new Error(error.error || 'Erreur lors de l'authentification');
      }

      const result = await finishResponse.json();
      // Stockage du token JWT et redirection vers le dashboard
      localStorage.setItem('authToken', result.token);
      window.location.href = '/dashboard';

    } catch (error) {
      setErrorMessage(error.message || 'Une erreur est survenue');
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSignUp = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setErrorMessage('');

    try {
      // Étape 1: Commencer l'enregistrement
      const beginResponse = await fetch('/api/auth/register/begin', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username }),
        credentials: 'include'
      });

      if (!beginResponse.ok) {
        const error = await beginResponse.json();
        throw new Error(error.error || 'Erreur lors de l'enregistrement');
      }

      const options = await beginResponse.json();

      // Préparation des options WebAuthn pour utiliser Touch ID
      const publicKeyOptions = {
        ...options,
        challenge: base64URLToBuffer(options.challenge),
        user: {
          ...options.user,
          id: base64URLToBuffer(options.user.id)
        },
        excludeCredentials: options.excludeCredentials?.map(credential => ({
          ...credential,
          id: base64URLToBuffer(credential.id)
        })),
        authenticatorSelection: {
          authenticatorAttachment: 'platform', // Utilise le lecteur d'empreinte intégré
          residentKey: 'preferred',
          userVerification: 'required'
        }
      };

      // Étape 2: Création de la clé avec le lecteur d'empreinte
      const credential = await navigator.credentials.create({
        publicKey: publicKeyOptions
      });

      // Étape 3: Finaliser l'enregistrement
      const attestationResponse = credential.response;
      const finishResponse = await fetch('/api/auth/register/finish', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          id: credential.id,
          rawId: bufferToBase64URL(credential.rawId),
          type: credential.type,
          response: {
            clientDataJSON: bufferToBase64URL(attestationResponse.clientDataJSON),
            attestationObject: bufferToBase64URL(attestationResponse.attestationObject)
          }
        }),
        credentials: 'include'
      });

      if (!finishResponse.ok) {
        const error = await finishResponse.json();
        throw new Error(error.error || 'Erreur lors de l'enregistrement');
      }

      const result = await finishResponse.json();
      // Stockage du token JWT et redirection vers le dashboard
      localStorage.setItem('authToken', result.token);
      window.location.href = '/dashboard';

    } catch (error) {
      setErrorMessage(error.message || 'Une erreur est survenue');
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="auth-form">
      <h2>Authentification Passwordless</h2>
      {errorMessage && <div className="error-message">{errorMessage}</div>}
      
      <form>
        <div className="form-group">
          <label htmlFor="username">Nom d'utilisateur:</label>
          <input
            type="text"
            id="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>

        <div className="button-group">
          <button 
            onClick={handleSignIn}
            disabled={isLoading || !username}
            className="btn-signin"
          >
            {isLoading ? 'Chargement...' : 'Se connecter avec Touch ID'}
          </button>
          
          <button 
            onClick={handleSignUp}
            disabled={isLoading || !username}
            className="btn-signup"
          >
            {isLoading ? 'Chargement...' : 'S\'inscrire avec Touch ID'}
          </button>
        </div>
      </form>

      <div className="info-message">
        <p>Utilisez votre empreinte digitale pour vous authentifier</p>
      </div>
    </div>
  );
};

export default LoginForm; 