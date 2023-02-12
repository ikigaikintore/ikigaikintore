import react from '@vitejs/plugin-react'
import * as path from 'path'
import { defineConfig, loadEnv } from 'vite'
import tsconfigPaths from 'vite-tsconfig-paths'

// https://vitejs.dev/config/
export default defineConfig((config) => {
    const env = loadEnv(config.mode, process.cwd())
    const envWithProcessPrefix = Object.entries(env).reduce(
        (prev, [key, val]) => {
            return {
                ...prev,
                ['process.env' + key]: `'${val}`,
            }
        },
        {}
    )
    return {
        root: './',
        plugins: [react(), tsconfigPaths()],
        resolve: {
            alias: {
                '@/': path.join(__dirname, './src/'),
            },
        },
        define: envWithProcessPrefix,
        server: {
            host: true,
            port: 3000,
        },
    }
})
