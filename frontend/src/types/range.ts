export enum RangeEnum {
  oneDay = '1day',
  threeDays = '3days',
  oneWeek = '1week',
  oneMonth = '1month',
  threeMonths = '3months',
  sixMonths = '6months',
  oneYear = '1year',
  fiveYear = '5year',
}
export const RangeOptions: RangeEnum[] = [
  RangeEnum.oneDay,
  RangeEnum.threeDays,
  RangeEnum.oneWeek,
  RangeEnum.oneMonth,
  RangeEnum.threeMonths,
  RangeEnum.sixMonths,
  RangeEnum.oneYear,
  RangeEnum.fiveYear,
];

export const rangeToDates = (range: RangeEnum): { start: Date; end: Date } => {
  const end = new Date();
  const start = new Date();
  switch (range) {
    case RangeEnum.oneDay:
      start.setDate(end.getDate() - 1); // 1 jour en arrière par rapport à la date du jour
      break;
    case RangeEnum.threeDays:
      start.setDate(end.getDate() - 3); // 3 jours en arrière par rapport à la date du jour
      break;
    case RangeEnum.oneWeek:
      start.setDate(end.getDate() - 7); // 1 semaine en arrière par rapport à la date du jour
      break;
    case RangeEnum.oneMonth:
      start.setMonth(end.getMonth() - 1); // 1 mois en arrière par rapport à la date du jour
      break;
    case RangeEnum.threeMonths:
      start.setMonth(end.getMonth() - 3); // 3 mois en arrière par rapport à la date du jour
      break;
    case RangeEnum.sixMonths:
      start.setMonth(end.getMonth() - 6); // 6 mois en arrière par rapport à la date du jour
      break;
    case RangeEnum.oneYear:
      start.setFullYear(end.getFullYear() - 1); // 1 an en arrière par rapport à la date du jour
      break;
    case RangeEnum.fiveYear:
      start.setFullYear(end.getFullYear() - 5); // 5 ans en arrière par rapport à la date du jour
      break;
    default:
      break;
  }

  return { start, end };
};
