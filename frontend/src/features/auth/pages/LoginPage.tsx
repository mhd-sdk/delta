import WebAuthnForm from '../components/WebAuthnForm';

const LoginPage = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-background p-4">
      <WebAuthnForm mode="login" />
    </div>
  );
};

export default LoginPage;
