import express from 'express';

export type UserModel = {
    id: string;
    username: string;
    // Other user properties
};
    
  
const app = express();


const port = 3000;


app.get('/auth/register/begin', (req, res) => {
    // (Pseudocode) Retrieve the user from the database
// after they've logged in
const user: UserModel = getUserFromDB(loggedInUserId);
// (Pseudocode) Retrieve any of the user's previously-
// registered authenticators
const userPasskeys: Passkey[] = getUserPasskeys(user);

const options: PublicKeyCredentialCreationOptionsJSON = await generateRegistrationOptions({
  rpName,
  rpID,
  userName: user.username,
  // Don't prompt users for additional information about the authenticator
  // (Recommended for smoother UX)
  attestationType: 'none',
  // Prevent users from re-registering existing authenticators
  excludeCredentials: userPasskeys.map(passkey => ({
    id: passkey.id,
    // Optional
    transports: passkey.transports,
  })),
  // See "Guiding use of authenticators via authenticatorSelection" below
  authenticatorSelection: {
    // Defaults
    residentKey: 'preferred',
    userVerification: 'preferred',
    // Optional
    authenticatorAttachment: 'platform',
  },
});

// (Pseudocode) Remember these options for the user
setCurrentRegistrationOptions(user, options);

return options;
  res.send('Hello, TypeScript + Node.js + Express!');
});

const rpName = 'Delta';
const rpID = 'localhost';
const origin = `https://${rpID}`;

// Start the server and listen on the specified port
app.listen(port, () => {
  // Log a message when the server is successfully running
  console.log(`Server is running on http://localhost:${port}`);
});