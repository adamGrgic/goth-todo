import { readdir, writeFile, mkdir, unlink } from "fs/promises";
import { existsSync } from "fs";
import { join, extname, basename } from "path";
import { createHash } from 'crypto';
import { mkdirSync, writeFileSync, readFileSync, watch} from 'fs';
import path, {  resolve } from 'path';
import { argv } from 'process';
import { rm } from "node:fs/promises";
import type { Dirent } from "node:fs";


const INPUT_DIR = "./src/scripts/js";
const OUTPUT_DIR = "./public/scripts/bun";
const MANIFEST_PATH = "./public/manifest.json";

const manifest: Record<string, string> = {};

async function ensureOutputDir() {
  if (!existsSync(OUTPUT_DIR)) {
    await mkdir(OUTPUT_DIR, { recursive: true });
  }
}

function getHash(content: string): string {
  return createHash("md5").update(content).digest("hex").slice(0, 8);
}

async function deleteHashedFiles(outputDir: string, logicalName: string) {
    const files = await readdir(outputDir);
  
    const pattern = new RegExp(`^${logicalName}\\.[a-f0-9]{8}\\.js$`);
  
    for (const file of files) {
      if (pattern.test(file)) {
        await unlink(join(outputDir, file));
        console.log(`🗑️  Deleted old build: ${file}`);
      }
    }
  }

async function compileFile(entry : Dirent, fullPath: string) {
  const relativePath = fullPath.replace(INPUT_DIR + "/", "");
  const logicalName = basename(entry.name, ".ts");

  const result = await Bun.build({
    entrypoints: [fullPath],
    minify: true,
    target: "bun",
  });

  if (!result.success || result.outputs.length === 0) {
    console.error(`❌ Failed to build: ${relativePath}`);
    return;
  }

  const outputContent = await result!.outputs[0]!.text();
      const hash = getHash(outputContent);
      const hashedName = `${logicalName}.${hash}.js`;
      const outputPath = join(OUTPUT_DIR, hashedName);
      const manifest = JSON.parse(readFileSync(MANIFEST_PATH, 'utf8'));
    //   console.log(`DEBUG: logical name: ${logicalName}`);
    //   console.log(`DEBUG: hashedName: ${hashedName}`);
    //   console.log(`DEBUG: manifest[logicalName]: ${manifest[logicalName]}`);
      
      if (manifest[logicalName].includes(hashedName)) {
            // console.log(`DEBUG: compiled file ${hashedName} already exists, continuing...`)
            return;
      }
      await deleteHashedFiles(OUTPUT_DIR, logicalName);

      await writeFile(outputPath, outputContent);
      
      manifest[logicalName] = outputPath; 
      writeFileSync(MANIFEST_PATH, JSON.stringify(manifest, null, 2));

      console.log(`✅ Compiled ${relativePath} → ${hashedName}`);
}

async function walkAndCompile(dir: string) {
  const entries = await readdir(dir, { withFileTypes: true });
  const compileJobs: Promise<void>[] = [];

  for (const entry of entries) {
    const fullPath = join(dir, entry.name);

    if (entry.isDirectory()) {
      compileJobs.push(walkAndCompile(fullPath));
    } else if (entry.isFile() && extname(entry.name) === ".ts") {
      
      compileJobs.push(compileFile(entry,fullPath))
    }
  }
}

async function run() {
  console.clear();
  console.log("🚀 Starting TS compiler...\n");

  await ensureOutputDir();
  await walkAndCompile(INPUT_DIR);

  console.log("\n✅ All done!");
}


if (argv.includes('--watch')) {
    // console.log(' watch flag given');
    run();
    // console.log(path.dirname(INPUT_DIR));
    watch(path.dirname(INPUT_DIR), { recursive: true }, (event, filename) => {
    if (filename && filename.endsWith('.ts')) {
      console.log(`🔄 Change detected in ${filename}, recompiling...`);
      run();
    }
  });

} else {
    // console.log('no watch flag given')
    run();
}