import { RangeEnum } from './range';

export type Timeframe = {
  n: number;
  unit: Unit;
};

export enum Unit {
  min = 'Min',
  hour = 'Hour',
  day = 'Day',
  week = 'Week',
  month = 'Month',
}

export const timeframeToString = (timeframe: Timeframe): string => {
  return `${timeframe.n} ${timeframe.unit}`;
};

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

export const calcOptimizedRange = (timeframe: Timeframe): RangeEnum => {
  const { n, unit } = timeframe;

  switch (unit) {
    case Unit.min:
      if (n <= 1) return RangeEnum.oneMonth;
      if (n <= 5) return RangeEnum.oneMonth;
      if (n <= 15) return RangeEnum.sixMonths;
      if (n <= 30) return RangeEnum.sixMonths;
      return RangeEnum.sixMonths;

    case Unit.hour:
      if (n <= 1) return RangeEnum.sixMonths;
      if (n <= 4) return RangeEnum.oneYear;
      return RangeEnum.oneYear;

    case Unit.day:
      if (n <= 1) return RangeEnum.fiveYear;
      return RangeEnum.fiveYear;

    case Unit.week:
      return RangeEnum.fiveYear;

    case Unit.month:
      return RangeEnum.fiveYear;

    default:
      return RangeEnum.oneDay;
  }
};
