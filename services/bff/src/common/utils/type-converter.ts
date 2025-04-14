/**
 * protobufから返された文字列値を安全に数値に変換します
 * @param value protobufから返された文字列値
 * @param defaultValue パース失敗時のデフォルト値
 * @returns パースされた数値またはデフォルト値
 */
export function safeParseNumber(
  value: string | null | undefined,
  defaultValue = 0,
): number {
  if (value === null || value === undefined || value === '') return defaultValue
  const parsed = Number(value)
  return isNaN(parsed) ? defaultValue : parsed
}
