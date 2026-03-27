// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  ssr: true,

  routeRules: {
    '/wonderful/**': { ssr: false }
  },

  modules: [
    '@nuxt/a11y',
    '@nuxt/eslint',
    '@nuxt/hints',
    '@nuxt/image',
    '@nuxt/scripts',
    '@nuxt/test-utils',
    '@nuxt/ui',
    '@vueuse/nuxt'
  ],

  css: ['~/assets/css/main.css'],

  ui: {
    prefix: 'Nuxt'
  },

  fonts: {
    families: [
      {
        name: 'Inter',
      },
      {
        name: 'Fira Code',
      }
    ]
  },

  runtimeConfig: {
    public: {
      apiDomain: 'http://127.0.0.1:8000',
      wsDomain: 'ws://127.0.0.1:8000'
    }
  }
})
