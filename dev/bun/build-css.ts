import { compileAsync } from 'sass';
import CleanCSS from 'clean-css';
import { createHash } from 'crypto';
import { mkdirSync, writeFileSync, readFileSync } from 'fs';
import { resolve } from 'path';

// Input and output paths
const inputFile = './src/styles/main.scss';
const outputDir = './public/styles';
const manifestPath = './public/manifest.json';

// Ensure output directory exists
mkdirSync(outputDir, { recursive: true });

async function buildCSS() {
  try {
    // Compile SCSS (suppress warnings)
    const compiled = await compileAsync(inputFile, {
      loadPaths: ['node_modules'],
      style: 'expanded',
      quietDeps: true,
      logger: { warn: () => {} },
    });

    // Minify the compiled CSS
    const minified = new CleanCSS().minify(compiled.css);
    if (minified.errors.length) {
      throw new Error(minified.errors.join('\n'));
    }

    // Generate hash from minified CSS
    const hash = createHash('md5').update(minified.styles).digest('hex').slice(0, 8);
    const filename = `main.${hash}.css`;

    // Write the hashed, minified CSS file
    writeFileSync(resolve(outputDir, filename), minified.styles);

    // Update manifest.json
    const manifest = JSON.parse(readFileSync(manifestPath, 'utf8'));
    manifest['css'] = filename; // Update or add your CSS filename entry here
    writeFileSync(manifestPath, JSON.stringify(manifest, null, 2));

    console.log(`✅ CSS compiled and minified: ${filename}`);
    console.log(`✅ manifest.json updated.`);
  } catch (error) {
    console.error('❌ Build failed:', error);
  }
}

// Execute the build
buildCSS();
