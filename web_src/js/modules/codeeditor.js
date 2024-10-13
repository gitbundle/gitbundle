import { colord } from 'colord'
import { basename, extname, isDarkTheme, isObject } from '../utils/index.js'

const languagesByFilename = {}
const languagesByExt = {}

const baseOptions = {
  // fontFamily: 'var(--fonts-monospace)',//NOTE: use `font-mono` class instead
  fontSize: 14, // https://github.com/microsoft/monaco-editor/issues/2242
  guides: { bracketPairs: false, indentation: false },
  links: false,
  minimap: { enabled: false },
  occurrencesHighlight: false,
  overviewRulerLanes: 0,
  renderLineHighlight: 'all',
  renderLineHighlightOnlyWhenFocus: true,
  renderWhitespace: 'none',
  rulers: false,
  scrollbar: { horizontalScrollbarSize: 6, verticalScrollbarSize: 6 },
  scrollBeyondLastLine: false
}

function getEditorconfig(input) {
  try {
    return JSON.parse(input.getAttribute('data-editorconfig'))
  } catch {
    return null
  }
}

function initLanguages(monaco) {
  for (const { filenames, extensions, id } of monaco.languages.getLanguages()) {
    for (const filename of filenames || []) {
      languagesByFilename[filename] = id
    }
    for (const extension of extensions || []) {
      languagesByExt[extension] = id
    }
  }
}

function getLanguage(filename) {
  return (
    languagesByFilename[filename] ||
    languagesByExt[extname(filename)] ||
    'plaintext'
  )
}

function updateEditor(monaco, editor, filename, lineWrapExts) {
  editor.updateOptions(getFileBasedOptions(filename, lineWrapExts))
  const model = editor.getModel()
  const language = model.getLanguageId()
  const newLanguage = getLanguage(filename)
  if (language !== newLanguage)
    monaco.editor.setModelLanguage(model, newLanguage)
}

// export editor for customization - https://github.com/go-gitea/gitea/issues/10409
function exportEditor(editor) {
  if (!window.codeEditors) window.codeEditors = []
  if (!window.codeEditors.includes(editor)) window.codeEditors.push(editor)
}

export async function createMonaco(
  textarea,
  filename,
  editorOpts,
  className = undefined,
  automaticLayout = undefined // this has a perf implication cost
) {
  const monaco = await import(/* webpackChunkName: "monaco" */ 'monaco-editor')

  initLanguages(monaco)
  let { language, ...other } = editorOpts
  if (!language) language = getLanguage(filename)

  const container = document.createElement('div')
  container.className = 'monaco-editor-container'
  if (className) {
    container.classList.add(className)
  }
  textarea.parentNode.appendChild(container)

  // https://github.com/microsoft/monaco-editor/issues/2427
  const styles = window.getComputedStyle(document.documentElement)
  const getProp = (name) => styles.getPropertyValue(name).trim()

  monaco.editor.defineTheme('gitbundle', {
    base: isDarkTheme() ? 'vs-dark' : 'vs',
    inherit: true,
    rules: [
      {
        background: colord(`hsl(${getProp('--b3')})`).toHex()
      }
    ],
    colors: {
      'editor.background': colord(`hsl(${getProp('--b3')})`).toHex(),
      'editor.foreground': colord(`hsl(${getProp('--bc')})`).toHex(),
      'editor.inactiveSelectionBackground': colord(
        `hsla(${getProp('--pf')} / 0.4)`
      ).toHex(),
      'editor.lineHighlightBackground': colord(
        `hsla(${getProp('--pf')} / 0.5)`
      ).toHex(),
      'editor.selectionBackground': colord(
        `hsla(${getProp('--pf')} / 0.3)`
      ).toHex(),
      'editor.selectionForeground': colord(
        `hsla(${getProp('--pf')} / 0.3)`
      ).toHex(),
      'editorLineNumber.background': colord(`hsl(${getProp('--b3')})`).toHex(),
      'editorLineNumber.foreground': colord(`hsl(${getProp('--bc')})`).toHex(),
      'editorWidget.background': colord(`hsl(${getProp('--b1')})`).toHex(),
      'editorWidget.border': colord(`hsl(${getProp('--b2')})`).toHex(),
      'input.background': colord(`hsl(${getProp('--b1')})`).toHex(),
      // 'input.border': getProp('--color-input-border'),
      // 'input.foreground': getProp('--color-input-text'),
      // 'scrollbar.shadow': getProp('--color-shadow'),
      'progressBar.background': colord(`hsl(${getProp('--p')})`).toHex()
    }
  })

  // monaco.languages.register({ id: 'vs.editor.nullLanguage' })
  // monaco.languages.setLanguageConfiguration('vs.editor.nullLanguage', {})

  const options = {
    value: textarea.value,
    theme: 'gitbundle',
    language,
    ...other
  }
  if (automaticLayout) options.automaticLayout = true
  const editor = monaco.editor.create(container, options)

  const model = editor.getModel()
  model.onDidChangeContent(() => {
    textarea.value = editor.getValue()
    textarea.dispatchEvent(new Event('change')) // seems to be needed for jquery-are-you-sure
  })

  window.addEventListener('resize', () => {
    editor.layout()
  })

  exportEditor(editor)

  const loading = document.querySelector('[jq-editor-loading]')
  if (loading) loading.remove()

  return { monaco, editor }
}

function getFileBasedOptions(filename, lineWrapExts) {
  return {
    wordWrap: (lineWrapExts || []).includes(extname(filename)) ? 'on' : 'off'
  }
}

export async function createCodeEditor(
  textarea,
  filenameInput,
  previewFileModes,
  className = undefined,
  automaticLayout = undefined // this has a perf implication cost
) {
  const filename = basename(filenameInput.value)
  const previewLink = document.querySelector('a[data-tab=preview]')
  const markdownExts = (
    textarea.getAttribute('data-markdown-file-exts') || ''
  ).split(',')
  const lineWrapExts = (
    textarea.getAttribute('data-line-wrap-extensions') || ''
  ).split(',')
  const isMarkdown = markdownExts.includes(extname(filename))
  const editorConfig = getEditorconfig(filenameInput)

  if (previewLink) {
    if (isMarkdown && (previewFileModes || []).includes('markdown')) {
      const newUrl = (previewLink.getAttribute('data-url') || '').replace(
        /(.*)\/.*/i,
        `$1/markdown`
      )
      previewLink.setAttribute('data-url', newUrl)
      previewLink.style.display = ''
    } else {
      previewLink.style.display = 'none'
    }
  }

  const { monaco, editor } = await createMonaco(
    textarea,
    filename,
    {
      ...baseOptions,
      ...getFileBasedOptions(filenameInput.value, lineWrapExts),
      ...getEditorConfigOptions(editorConfig)
    },
    className,
    automaticLayout
  )

  filenameInput.addEventListener('keyup', () => {
    const filename = filenameInput.value
    updateEditor(monaco, editor, filename, lineWrapExts)
  })

  return editor
}

function getEditorConfigOptions(ec) {
  if (!isObject(ec)) return {}

  const opts = {}
  opts.detectIndentation = !('indent_style' in ec) || !('indent_size' in ec)
  if ('indent_size' in ec) opts.indentSize = Number(ec.indent_size)
  if ('tab_width' in ec) opts.tabSize = Number(ec.tab_width) || opts.indentSize
  if ('max_line_length' in ec) opts.rulers = [Number(ec.max_line_length)]
  opts.trimAutoWhitespace = ec.trim_trailing_whitespace === true
  opts.insertSpaces = ec.indent_style === 'space'
  opts.useTabStops = ec.indent_style === 'tab'
  return opts
}
