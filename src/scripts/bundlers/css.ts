import { compileAsync } from 'sass';
import CleanCSS from 'clean-css';
import { createHash } from 'crypto';
import { mkdirSync, writeFileSync, readFileSync, watch} from 'fs';
import path, {  resolve } from 'path';
import { argv } from 'process';
import { rm, mkdir } from "node:fs/promises";

// Input and output paths
const inputFile = './src/styles/main.scss';
const INPUT_DIR = './src/styles/'
const outputDir = './public/styles';
const manifestPath = './public/manifest.json';

// Ensure output directory exists
mkdirSync(outputDir, { recursive: true });

async function compileCSS() {
  console.clear()
  
  if (argv.includes('--verbose')) {
    console.log('🚀 initializing css compiler...')
  }

  try {
    // Clean current CSS
    if (argv.includes('--verbose')) {
      console.log(`🧹 Cleaning current styles directory ...`)
    }
    
    await rm(outputDir, { recursive: true, force: true });
    
    if (argv.includes('--verbose')) {
      console.log(`✅ Directory "${outputDir}" and its contents have been deleted.`);
    }

    await mkdir(outputDir, { recursive: true });

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

    if (argv.includes('--verbose')) {
      console.log(`✅ CSS compiled and minified: ${filename}`);
      console.log(`✅ manifest.json updated.`);
    }

    if (argv.includes('--watch')) {
      console.log(`⚡ Watching for styling changes in ${INPUT_DIR} and all subdirectories`)
    }
  } catch (error) {
    console.error('❌ Build failed:', error);
  }
}

if (argv.includes('--watch')) {
  compileCSS();
  watch(path.dirname(INPUT_DIR), { recursive: true }, (event, filename) => {
    if (filename && filename.endsWith('.scss')) {
      console.log(`🔄 Change detected in ${filename}, recompiling...`);
      compileCSS();
    }
  });

} else {
  compileCSS();
}
