/** @type {import('next').NextConfig} */
const nextConfig = {
  env: {
    NEXT_PUBLIC_AUTH_SERVICE_URL: process.env.NEXT_PUBLIC_AUTH_SERVICE_URL,
    NEXT_PUBLIC_ACTION_SERVICE_URL: process.env.NEXT_PUBLIC_ACTION_SERVICE_URL,
  },
  output: "standalone",
};

export default nextConfig;
