import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import rehypeHighlight from 'rehype-highlight'

export default function Markdown({ children, className = '' }) {
  return (
    <div
      className={`prose prose-slate max-w-none dark:prose-invert prose-headings:scroll-mt-24 prose-h2:text-xl prose-h2:font-bold prose-a:text-brand-600 dark:prose-a:text-brand-400 ${className}`}
    >
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        rehypePlugins={[[rehypeHighlight, { detect: true, ignoreMissing: true }]]}
      >
        {children || ''}
      </ReactMarkdown>
    </div>
  )
}
