import { calcOptimizedRange, isIntraday } from './dashboard';
import { Range, Unit } from './timerange';

describe('Dashboard Utilities', () => {
  it('checks if a timeframe is intraday', () => {
    expect(isIntraday({ n: 1, unit: Unit.min })).toBe(true);
    expect(isIntraday({ n: 1, unit: Unit.day })).toBe(false);
  });

  it('calculates optimized range for minute unit', () => {
    expect(calcOptimizedRange({ n: 1, unit: Unit.min })).toBe(Range.oneMonth);
    expect(calcOptimizedRange({ n: 15, unit: Unit.min })).toBe(Range.sixMonths);
  });

  it('calculates optimized range for hour unit', () => {
    expect(calcOptimizedRange({ n: 1, unit: Unit.hour })).toBe(Range.sixMonths);
    expect(calcOptimizedRange({ n: 5, unit: Unit.hour })).toBe(Range.oneYear);
  });

  it('calculates optimized range for day unit', () => {
    expect(calcOptimizedRange({ n: 1, unit: Unit.day })).toBe(Range.fiveYear);
  });

  it('calculates optimized range for unsupported units', () => {
    expect(calcOptimizedRange({ n: 1, unit: 'unsupported' as Unit })).toBe(Range.oneDay);
  });
});
