const rules = [
  {
    test: /undefined: (\w+)/,
    hint: (m) => `“${m[1]}” belum dideklarasikan. Cek ejaan atau tambahkan deklarasinya.`,
  },
  {
    test: /declared and not used:? (\w+)?/,
    hint: () => 'Ada variabel yang dideklarasikan tetapi tidak dipakai. Go menolak kode seperti ini.',
  },
  {
    test: /imported and not used:? "?([\w./]+)"?/,
    hint: (m) => `Package “${m[1]}” di-import tetapi tidak dipakai. Hapus import-nya.`,
  },
  {
    test: /missing return/,
    hint: () => 'Fungsi yang punya nilai kembali harus diakhiri dengan return.',
  },
  {
    test: /cannot use .* as .* in/,
    hint: () => 'Tipe data tidak cocok. Periksa tipe variabel atau konversi dengan tipe yang benar.',
  },
  {
    test: /expected .* found/,
    hint: () => 'Ada tanda baca atau kata kunci yang salah posisi. Periksa kurung kurawal dan titik koma.',
  },
  {
    test: /syntax error/,
    hint: () => 'Ada kesalahan sintaks. Periksa kurung, tanda kutip, dan struktur blok.',
  },
  {
    test: /index out of range|nil pointer dereference/,
    hint: () => 'Program panic saat berjalan. Periksa akses index slice atau pointer nil.',
  },
  {
    test: /cannot find package/,
    hint: () => 'Package tidak ditemukan. Pastikan import benar.',
  },
]

export function friendlyError(stderr) {
  if (!stderr) return ''
  for (const rule of rules) {
    const m = stderr.match(rule.test)
    if (m) return rule.hint(m)
  }
  return ''
}
