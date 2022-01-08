module.exports = {
  chainWebpack(config) {
    config.plugin('html').tap((args) => {
      args[0].title = 'Mnemonic Ninja';
      return args;
    });
  },
  pwa: {
    title: 'Mnemonic Ninja',
    themeColor: '#607D8B',
    msTileColor: '#607D8B',
    appleMobileWebAppCapable: 'yes',
    appleMobileWebAppStatusBarStyle: 'black',

    workboxPluginMode: 'InjectManifest',
    workboxOptions: {
      swSrc: './src/service-worker.js',
      exclude: [
        /^manifest.*\.js?$/,
      ],
    },
  },
  css: {
    loaderOptions: {
      scss: {
        additionalData: `
          @use "sass:selector";
        `,
      },
    },
  },
  transpileDependencies: [
    'vuetify',
  ],
};
