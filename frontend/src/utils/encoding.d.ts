declare module "../utils/encoding" {
  export function bufferToBase64URL(
    buffer: ArrayBuffer | ArrayBufferView,
  ): string;
  export function base64URLToBuffer(base64: string): Uint8Array;
}
