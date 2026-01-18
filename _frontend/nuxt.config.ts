// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  ssr: true,
  
  routeRules: {
    'games/**': { ssr: false },
    'rankings': { ssr: true }
  },

  modules: [
    '@nuxt/eslint',
    '@nuxt/image',
    '@nuxt/test-utils',
    '@nuxt/scripts',
    '@nuxt/ui',
    '@vueuse/nuxt',
    'pinia-plugin-persistedstate',
    '@pinia/nuxt'
  ],

  css: [
    '~/assets/css/main.css'
  ],

  ui: {
    prefix: 'Nuxt'
  }
})