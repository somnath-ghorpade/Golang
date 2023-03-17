const path = require("path");
module.exports = {
  publicPath: process.env.NODE_ENV === "production" ? "/my-project/" : "/",

  chainWebpack: config => {
    config.resolve.alias.set(
      "vue$",
      path.resolve("node_modules/vue/dist/vue.esm-bundler.js")
    );
  },

  devServer: {
    proxy: "http://localhost:8000"
  },

  pluginOptions: {
    i18n: {
      locale: "en",
      fallbackLocale: "en",
      localeDir: "locales",
      enableLegacy: false,
      runtimeOnly: false,
      compositionOnly: false,
      fullInstall: true
    }
  }
};
