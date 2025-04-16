import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* ルートディレクトリのpagesを使用するように設定 */
  experimental: {
    optimizeCss: false
  },
  pageExtensions: ['tsx', 'ts', 'jsx', 'js']
}

export default nextConfig;
