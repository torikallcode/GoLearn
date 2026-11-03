import Editor from '@monaco-editor/react'
import { useTheme } from '../store/theme'

export default function CodeEditor({ value, onChange, height = 320, readOnly = false }) {
  const theme = useTheme((s) => s.theme)
  return (
    <div className="overflow-hidden rounded-xl border border-slate-200 dark:border-slate-800">
      <Editor
        height={height}
        defaultLanguage="go"
        theme={theme === 'dark' ? 'vs-dark' : 'light'}
        value={value}
        onChange={(v) => onChange?.(v ?? '')}
        loading={<div className="p-4 text-sm text-slate-500">Memuat editor…</div>}
        options={{
          readOnly,
          fontSize: 13,
          minimap: { enabled: false },
          scrollBeyondLastLine: false,
          smoothScrolling: true,
          tabSize: 4,
          automaticLayout: true,
          padding: { top: 12, bottom: 12 },
          renderLineHighlight: 'none',
        }}
      />
    </div>
  )
}
