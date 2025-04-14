import { safeParseNumber } from './type-converter'

describe('型変換ユーティリティ', () => {
  describe('safeParseNumber', () => {
    it('有効な数値文字列を正しく数値型に変換すること', () => {
      expect(safeParseNumber('123')).toBe(123)
      expect(safeParseNumber('0')).toBe(0)
      expect(safeParseNumber('-456')).toBe(-456)
      expect(safeParseNumber('123.45')).toBe(123.45)
    })

    it('空文字列に対して指定されたデフォルト値を返すこと', () => {
      expect(safeParseNumber('')).toBe(0)
      expect(safeParseNumber(null, -1)).toBe(-1)
      expect(safeParseNumber(undefined, 999)).toBe(999)
    })

    it('無効な数値文字列に対してデフォルト値を返すこと', () => {
      expect(safeParseNumber('abc')).toBe(0)
      expect(safeParseNumber('123abc', -1)).toBe(-1)
      expect(safeParseNumber('abc123', 999)).toBe(999)
    })
  })
})
