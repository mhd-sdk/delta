// Conversion d'un buffer en chaîne base64URL
export function bufferToBase64URL(buffer) {
  const bytes = new Uint8Array(buffer);
  let str = '';
  
  for (const charCode of bytes) {
    str += String.fromCharCode(charCode);
  }
  
  const base64 = btoa(str)
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=+$/, '');
    
  return base64;
}

// Conversion d'une chaîne base64URL en buffer
export function base64URLToBuffer(base64) {
  const base64Std = base64
    .replace(/-/g, '+')
    .replace(/_/g, '/');
    
  const padLen = (4 - (base64Std.length % 4)) % 4;
  const padded = base64Std + '='.repeat(padLen);
  
  const binary = atob(padded);
  const buffer = new Uint8Array(binary.length);
  
  for (let i = 0; i < binary.length; i++) {
    buffer[i] = binary.charCodeAt(i);
  }
  
  return buffer;
} 