import esbuild from 'esbuild';

esbuild.build({
  entryPoints: ['index.ts'],
  bundle: true,
  platform: 'node',
  target: 'node18',
  outfile: 'dist/index.js',
  format: 'esm',
  external: [
    'playwright',
    'playwright-core',
    '@playwright/*',
    'chromium-bidi',
    'amqplib'
  ],
  sourcemap: true,
  minify: false
}).catch(() => process.exit(1));
