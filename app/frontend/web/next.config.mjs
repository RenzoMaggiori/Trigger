/** @type {import('next').NextConfig} */
const nextConfig = {
  env: {
    NEXT_PUBLIC_AUTH_SERVICE_URL: process.env.NEXT_PUBLIC_AUTH_SERVICE_URL,
    NEXT_PUBLIC_ACTION_SERVICE_URL: process.env.NEXT_PUBLIC_ACTION_SERVICE_URL,
    NEXT_PUBLIC_SYNC_SERVICE_URL: process.env.NEXT_PUBLIC_SYNC_SERVICE_URL,
    NEXT_PUBLIC_SETTINGS_SERVICE_URL: process.env.NEXT_PUBLIC_SETTINGS_SERVICE_URL,
    NEXT_PUBLIC_WEB_URL: process.env.NEXT_PUBLIC_WEB_URL,
  },
  output: "standalone",
  images: {
    domains: ['fakeimg.pl'],
  },

};

export default nextConfig;
