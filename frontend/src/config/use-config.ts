export const useConfig = () => {

  const isLocal = () => process.env.NEXT_PUBLIC_ENVIRONMENT === "local"
  const emulatorHost = () => process.env.NEXT_PUBLIC_FIREBASE_EMULATOR_HOST || "http://localhost:9099"
  const functionsEmulatorHost = () => {
    return { host: "localhost", port: 5001 }
  }
  const emails = process.env.NEXT_PUBLIC_USER_AUTH?.split(",") || []

  return {
    functionsEmulatorHost,
    emails,
    isLocal,
    emulatorHost,
  }
}