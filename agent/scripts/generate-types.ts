#!/usr/bin/env ts-node
/**
 * Generates TypeScript types from JSON Schema definitions
 * Run with: npm run generate:types
 */

import { compileFromFile } from 'json-schema-to-typescript';
import * as fs from 'fs/promises';
import * as path from 'path';
import * as yaml from 'js-yaml';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const SCHEMA_DIR = path.join(__dirname, '../../shared/schemas');
const OUTPUT_DIR = path.join(__dirname, '../src/types/generated');

interface SchemaDocument {
  components?: {
    schemas?: Record<string, any>;
  };
}

async function generateTypes() {
  console.log('🔧 Generating TypeScript types from message schemas...\n');

  // Ensure output directory exists
  await fs.mkdir(OUTPUT_DIR, { recursive: true });

  // Read and parse the messages schema
  const messagesSchemaPath = path.join(SCHEMA_DIR, 'messages.yaml');
  const messagesYaml = await fs.readFile(messagesSchemaPath, 'utf-8');
  const messagesSchema = yaml.load(messagesYaml) as SchemaDocument;

  if (!messagesSchema.components?.schemas) {
    throw new Error('No schemas found in messages.yaml');
  }

  const schemas = messagesSchema.components.schemas;

  // Generate types for each schema
  const typeDefinitions: string[] = [];

  // Header comment
  typeDefinitions.push(`/**
 * Auto-generated TypeScript types from JSON Schema
 * DO NOT EDIT MANUALLY - run 'npm run generate:types' to regenerate
 *
 * Source: /shared/schemas/messages.yaml
 * Generated: ${new Date().toISOString()}
 */

`);

  // Convert each schema to TypeScript
  for (const [schemaName, schema] of Object.entries(schemas)) {
    console.log(`  ✓ Generating ${schemaName}...`);

    // Create a standalone schema document for compilation
    const standaloneSchema = {
      $schema: 'http://json-schema.org/draft-07/schema#',
      $id: schemaName,
      ...schema,
      // Add component references so they can resolve (match the $ref format)
      components: {
        schemas: schemas
      },
      definitions: schemas
    };

    // Write temporary file
    const tempPath = path.join(OUTPUT_DIR, `${schemaName}.temp.json`);
    await fs.writeFile(tempPath, JSON.stringify(standaloneSchema, null, 2));

    // Compile to TypeScript
    try {
      const ts = await compileFromFile(tempPath, {
        bannerComment: '',
        style: {
          semi: true,
          singleQuote: true,
        },
        unreachableDefinitions: true,
      });

      // Extract just the interface/type definition (remove duplicate definitions)
      const lines = ts.split('\n');
      const interfaceStart = lines.findIndex(line =>
        line.includes(`export interface ${schemaName}`) ||
        line.includes(`export type ${schemaName}`)
      );

      if (interfaceStart !== -1) {
        let braceCount = 0;
        let interfaceEnd = interfaceStart;
        let started = false;

        for (let i = interfaceStart; i < lines.length; i++) {
          const line = lines[i];
          if (line.includes('{')) {
            braceCount += (line.match(/{/g) || []).length;
            started = true;
          }
          if (line.includes('}')) {
            braceCount -= (line.match(/}/g) || []).length;
          }

          if (started && braceCount === 0) {
            interfaceEnd = i;
            break;
          }
        }

        const interfaceLines = lines.slice(interfaceStart, interfaceEnd + 1);
        typeDefinitions.push(interfaceLines.join('\n'));
        typeDefinitions.push('\n');
      }

      // Clean up temp file
      await fs.unlink(tempPath);
    } catch (error) {
      console.error(`  ✗ Error generating ${schemaName}:`, error);
      await fs.unlink(tempPath).catch(() => {});
    }
  }

  // Write all types to a single file
  const outputPath = path.join(OUTPUT_DIR, 'messages.ts');
  await fs.writeFile(outputPath, typeDefinitions.join('\n'));

  console.log(`\n✅ Types generated successfully at: ${outputPath}`);
  console.log(`📦 Generated ${Object.keys(schemas).length} type definitions\n`);
}

// Run the generator
generateTypes().catch((error) => {
  console.error('❌ Type generation failed:', error);
  process.exit(1);
});
