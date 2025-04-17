import { FlatCompat } from '@eslint/eslintrc';
import { dirname } from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const compat = new FlatCompat({
  baseDirectory: __dirname,
});

const eslintConfig = [
  ...compat.extends('next/typescript','plugin:prettier/recommended'),
  {
    rules: {
      'prettier/prettier': [
        'error',
        {
          printWidth: 150,
          tabWidth: 2,
          semi: true,
          singleQuote: true,
          trailingComma: 'es5',
          bracketSpacing: true,
          jsxBracketSameLine: false,
          endOfLine: 'auto',
        },
      ],
    },
  },
];

export default eslintConfig;