module.exports = {
    chainWebpack: config => {
        config
            .plugin('friendly-errors')
            .tap(args => {
                args[0].clearConsole = false;
                return args;
            });
    },
};
