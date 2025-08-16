import { Range, Timeframe, Unit } from "./timerange";

export const defaultTimeframes: Timeframe[] = [
  { n: 1, unit: Unit.min },
  { n: 5, unit: Unit.min },
  { n: 15, unit: Unit.min },
  { n: 30, unit: Unit.min },
  { n: 1, unit: Unit.hour },
  { n: 4, unit: Unit.hour },
  { n: 1, unit: Unit.day },
  { n: 1, unit: Unit.week },
  { n: 1, unit: Unit.month },
];

export const isIntraday = (timeframe: Timeframe): boolean => {
  return timeframe.unit === Unit.min || timeframe.unit === Unit.hour;
};

export const calcOptimizedRange = (timeframe: Timeframe): Range => {
  const { n, unit } = timeframe;

  switch (unit) {
    case Unit.min:
      if (n <= 1) return Range.oneMonth;
      if (n <= 5) return Range.oneMonth;
      if (n <= 15) return Range.sixMonths;
      if (n <= 30) return Range.sixMonths;
      return Range.sixMonths;

    case Unit.hour:
      if (n <= 1) return Range.sixMonths;
      if (n <= 4) return Range.oneYear;
      return Range.oneYear;

    case Unit.day:
      if (n <= 1) return Range.fiveYear;
      return Range.fiveYear;

    case Unit.week:
      return Range.fiveYear;

    case Unit.month:
      return Range.fiveYear;

    default:
      return Range.oneDay;
  }
};
