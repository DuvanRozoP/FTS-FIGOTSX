const path = require('path');
const scriptWebpackPlugin = require('./modules/scripts/plugins/exec');
require('dotenv').config();
const isDevelopment = process.env.isDev === 'true';
module.exports = {
  entry: './modules/client/_main.server.tsx',
  mode: isDevelopment ? 'development' : 'production',
  output: {
    filename: 'main.bundle.js',
    path: isDevelopment
      ? path.resolve(__dirname, 'src/cache/')
      : path.resolve(__dirname, 'dist/client/'),
  },
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        exclude: /node_modules/,
        use: 'swc-loader',
      },
    ],
  },
  resolve: {
    extensions: ['.ts', '.tsx', '.js'],
    alias: {
      '@modules': path.resolve(__dirname, 'modules/'),
      '@pages': path.resolve(__dirname, 'src/pages/'),
    },
  },
  plugins: [
    new scriptWebpackPlugin({
      scripts: [
        './modules/scripts/scripts',
        `npx tailwindcss -i src/index.css -o ${
          isDevelopment
            ? path.resolve(
                __dirname,
                'public/index.bundle.css'
              )
            : path.resolve(
                __dirname,
                'dist/static/index.bundle.css'
              )
        }`,
        isDevelopment
          ? ''
          : 'go build -o dist/server main.go',
      ],
      catchMessage: 'Error execute Scripts ⛔',
      doneMessage: 'Complete execute Scripts 🟢',
      doneCompilationMessage: 'All Scripts Complete 🟢',
    }),
  ],
};
