import { useState } from 'react'
import { useMutation } from '@tanstack/react-query'
import { api } from '../api'
import CodeEditor from './CodeEditor'
import OutputPanel from './OutputPanel'

const DEFAULT_CODE = `package main

import "fmt"

func main() {
	fmt.Println("Halo, Go!")
}
`

export default function Playground({ initialCode }) {
  const [code, setCode] = useState(initialCode || DEFAULT_CODE)
  const [result, setResult] = useState(null)
  const [error, setError] = useState('')

  const run = useMutation({
    mutationFn: () => api.run(code),
    onMutate: () => {
      setError('')
      setResult(null)
    },
    onSuccess: (data) => setResult(data),
    onError: (err) => setError(err.message || 'Gagal menjalankan kode.'),
  })

  return (
    <div className="space-y-3">
      <CodeEditor value={code} onChange={setCode} height={300} />
      <div className="flex items-center gap-2">
        <button className="btn-primary" onClick={() => run.mutate()} disabled={run.isPending}>
          {run.isPending ? 'Menjalankan…' : '▶ Run'}
        </button>
        <button className="btn-ghost text-xs" onClick={() => { setCode(DEFAULT_CODE); setResult(null); setError('') }}>
          Reset
        </button>
      </div>
      <OutputPanel result={result} running={run.isPending} error={error} />
    </div>
  )
}
