import { Panel } from "./panel";
import { Range, Timeframe, Unit } from "./timerange";

export const rangeToDates = (range: Range): { start: Date; end: Date } => {
  const end = new Date();
  const start = new Date();
  switch (range) {
    case Range.oneDay:
      start.setDate(end.getDate() - 1);
      break;
    case Range.threeDays:
      start.setDate(end.getDate() - 3);
      break;
    case Range.oneWeek:
      start.setDate(end.getDate() - 7);
      break;
    case Range.oneMonth:
      start.setMonth(end.getMonth() - 1);
      break;
    case Range.threeMonths:
      start.setMonth(end.getMonth() - 3);
      break;
    case Range.sixMonths:
      start.setMonth(end.getMonth() - 6);
      break;
    case Range.oneYear:
      start.setFullYear(end.getFullYear() - 1);
      break;
    case Range.fiveYear:
      start.setFullYear(end.getFullYear() - 5);
      break;
    default:
      break;
  }

  return { start, end };
};

export const timeframeToString = (timeframe: Timeframe): string => {
  return `${timeframe.n} ${timeframe.unit}`;
};

export const formatLayout = (panels: Panel[]): ReactGridLayout.Layout[] => {
  return panels.map(({ id, x, y, width, height }) => ({
    i: id,
    h: height,
    w: width,
    x,
    y,
  }));
};

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
