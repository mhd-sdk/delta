import { Range, rangeToTimeframe, Unit } from './timerange';

describe('rangeToTimeframe', () => {
  it('returns the correct timeframe for each range', () => {
    expect(rangeToTimeframe(Range.oneDay)).toEqual({ n: 1, unit: Unit.day });
    expect(rangeToTimeframe(Range.threeDays)).toEqual({ n: 3, unit: Unit.day });
    expect(rangeToTimeframe(Range.oneWeek)).toEqual({ n: 1, unit: Unit.week });
    expect(rangeToTimeframe(Range.oneMonth)).toEqual({ n: 1, unit: Unit.month });
    expect(rangeToTimeframe(Range.threeMonths)).toEqual({ n: 3, unit: Unit.month });
    expect(rangeToTimeframe(Range.sixMonths)).toEqual({ n: 6, unit: Unit.month });
    expect(rangeToTimeframe(Range.oneYear)).toEqual({ n: 12, unit: Unit.month });
    expect(rangeToTimeframe(Range.fiveYear)).toEqual({ n: 60, unit: Unit.month });
  });

  it('returns a default timeframe for unsupported ranges', () => {
    expect(rangeToTimeframe('unsupported' as Range)).toEqual({ n: 1, unit: Unit.day });
  });
});
