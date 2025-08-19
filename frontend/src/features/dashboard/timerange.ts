export enum Range {
  oneDay = "1Day",
  threeDays = "3Days",
  oneWeek = "1Week",
  oneMonth = "1Month",
  threeMonths = "3Months",
  sixMonths = "6Months",
  oneYear = "1Year",
  fiveYear = "5Year",
}

export const RangeOptions: Range[] = [
  Range.oneDay,
  Range.threeDays,
  Range.oneWeek,
  Range.oneMonth,
  Range.threeMonths,
  Range.sixMonths,
  Range.oneYear,
  Range.fiveYear,
];

export type Timeframe = {
  n: number;
  unit: Unit;
};

export enum Unit {
  min = "Min",
  hour = "Hour",
  day = "Day",
  week = "Week",
  month = "Month",
}
export const defaultTimeframes: Timeframe[] = [
  { n: 30, unit: Unit.min },
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

export const rangeToTimeframe = (range: Range): Timeframe => {
  switch (range) {
    case Range.oneDay:
      return { n: 1, unit: Unit.day };
    case Range.threeDays:
      return { n: 3, unit: Unit.day };
    case Range.oneWeek:
      return { n: 1, unit: Unit.week };
    case Range.oneMonth:
      return { n: 1, unit: Unit.month };
    case Range.threeMonths:
      return { n: 3, unit: Unit.month };
    case Range.sixMonths:
      return { n: 6, unit: Unit.month };
    case Range.oneYear:
      return { n: 12, unit: Unit.month };
    case Range.fiveYear:
      return { n: 60, unit: Unit.month };
    default:
      return { n: 1, unit: Unit.day };
  }
};
